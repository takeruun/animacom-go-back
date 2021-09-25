package usecase

import (
	"app/models"
	"app/usecase/auth"
)

type PostsInteractor struct {
	DB    DbRepository
	Posts PostsRepository
}

func (interactor *PostsInteractor) Create(p models.PostForm, accessToken string) (post models.Post, err error) {
	db := interactor.DB.Connect()

	auth, err := auth.ParseToken(accessToken)
	if err != nil {
		return models.Post{}, error(err)
	}
	p.UserId = uint(auth.Uid)

	post, err = interactor.Posts.Add(db, p)
	if err != nil {
		return models.Post{}, error(err)
	}

	return post, nil
}
