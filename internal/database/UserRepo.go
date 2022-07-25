package database

import (
	"errors"
	"fmt"
	"time"
)

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	fmt.Println("jo")
	_user := User{}
	db, err := c.readDB()
	if err != nil {
		//fmt.Println("CreateUser", err)
		return _user, err
	}
	user := User{
		age,
		email,
		name,
		password,
		time.Now().UTC(),
	}
	db.Users[email] = user
	err = c.updateDB(db)
	if err != nil {
		fmt.Println("cuser1", err)
		return _user, err
	}

	return user, nil
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	_user := User{}
	db, err := c.readDB()
	if err != nil {
		return _user, err
	}

	user, ok := db.Users[email]
	if !ok {
		return _user, errors.New("user doesn't exist")
	}

	user.Name = name
	user.Password = password
	user.Age = age

	db.Users[email] = user

	err = c.updateDB(db)
	if err != nil {
		return _user, err
	}

	return user, nil
}

func (c Client) GetUser(email string) (User, error) {
	_user := User{}
	db, err := c.readDB()
	if err != nil {
		return _user, err
	}

	user, ok := db.Users[email]
	if !ok {
		return _user, errors.New("user doesn't exist")
	}

	return user, nil
}

func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	_, ok := db.Users[email]
	if !ok {
		return errors.New("user doesn't exist")
	}

	delete(db.Users, email)

	err = c.updateDB(db)
	if err != nil {
		return err
	}

	return nil
}
