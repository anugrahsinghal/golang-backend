package database

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	_post := Post{}
	db, err := c.readDB()
	if err != nil {
		return _post, err
	}

	_, ok := db.Users[userEmail]
	if !ok {
		return _post, errors.New("user doesn't exist")
	}

	id := uuid.New().String()
	post := Post{
		id,
		time.Now().UTC(),
		text,
		userEmail,
	}

	db.Posts[id] = post

	err = c.updateDB(db)
	if err != nil {
		return _post, err
	}

	return post, nil
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	postArr := make([]Post, 2)
	db, err := c.readDB()
	if err != nil {
		return postArr, err
	}

	_, ok := db.Users[userEmail]
	if !ok {
		return postArr, errors.New("user doesn't exist")
	}

	for _, post := range db.Posts {
		if post.UserEmail == userEmail {
			postArr = append(postArr, post)
		}
	}

	return postArr, nil
}

func (c Client) DeletePost(id string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	_, ok := db.Posts[id]
	if !ok {
		return errors.New("post doesn't exist")
	}

	delete(db.Posts, id)

	err = c.updateDB(db)
	if err != nil {
		return err
	}

	return nil
}
