package usecase

import (
	"app/models"
	"app/usecase/auth"

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
