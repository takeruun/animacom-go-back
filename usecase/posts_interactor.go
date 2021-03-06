package usecase

import (
	models "app/models"
	auth "app/usecase/auth"
	"errors"

	"github.com/google/uuid"
)

type PostsInteractor struct {
	DB         DbRepository
	Posts      PostsRepository
	PostImages PostImagesRepository
	AwsS3      AwsS3Repository
}

func (interactor *PostsInteractor) Create(p models.PostForm, accessToken string) (post models.Post, err error) {
	db := interactor.DB.Connect()

	auth, err := auth.ParseToken(accessToken)
	if err != nil {
		return models.Post{}, error(err)
	}

	file, err := p.Image.Open()
	if err != nil {
		return models.Post{}, error(err)
	}

	uuid, _ := uuid.NewUUID()
	image_url, err := interactor.AwsS3.Upload(file, uuid.String(), "png")
	if err != nil {
		return models.Post{}, error(err)
	}

	post = models.Post{UserId: uint(auth.Uid), CategoryId: p.CategoryId, Title: p.Title, SubTitle: p.SubTitle}
	post, err = interactor.Posts.Add(db, post)
	if err != nil {
		return models.Post{}, error(err)
	}

	interactor.PostImages.Add(db, models.PostImage{PostId: post.ID, Image: image_url})

	return post, nil
}

func (interactor *PostsInteractor) Show(id int, accessToken string) (post models.Post, err error) {
	db := interactor.DB.Connect()

	auth, err := auth.ParseToken(accessToken)
	if err != nil {
		return models.Post{}, error(err)
	}

	post, err = interactor.Posts.FindByID(db, id)
	if err != nil {
		return models.Post{}, error(err)
	}

	if post.UserId != uint(auth.Uid) {
		return models.Post{}, errors.New("Not your posts")
	}

	return post, nil
}

func (interactor *PostsInteractor) Get(accessToken string) (posts []models.Post, err error) {
	db := interactor.DB.Connect()

	auth, err := auth.ParseToken(accessToken)
	if err != nil {
		return []models.Post{}, error(err)
	}

	posts, err = interactor.Posts.FindByUserId(db, auth.Uid)
	if err != nil {
		return []models.Post{}, error(err)
	}

	if posts[0].UserId != uint(auth.Uid) {
		return []models.Post{}, errors.New("Not your posts")
	}

	return posts, nil
}

func (interactor *PostsInteractor) Delete(id int, accessToken string) (err error) {
	db := interactor.DB.Connect()

	auth, err := auth.ParseToken(accessToken)
	if err != nil {
		return error(err)
	}

	post, err := interactor.Posts.FindByID(db, id)

	if post.UserId != uint(auth.Uid) {
		return errors.New("Not your posts")
	}

	err = interactor.Posts.Remove(db, id)
	if err != nil {
		return error(err)
	}

	return nil
}
