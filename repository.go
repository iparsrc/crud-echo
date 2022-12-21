package main

import (
	"errors"
)

var (
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrUserNotFound     = errors.New("user not found")
)

type Repository struct {
	users map[string]User
}

func NewRepository() *Repository {
	return &Repository{
		users: make(map[string]User),
	}
}

func (r Repository) GetUser(email string) (User, error) {
	user, ok := r.users[email]
	if !ok {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

func (r Repository) CreateUser(in User) (User, error) {
	_, ok := r.users[in.Email]
	if ok {
		return User{}, ErrUserAlreadyExist
	}
	r.users[in.Email] = in
	return in, nil
}

func (r Repository) UpdateUser(in User) (User, error) {
	user, ok := r.users[in.Email]
	if !ok {
		return User{}, ErrUserNotFound
	}
	if in.FirstName != "" {
		user.FirstName = in.FirstName
	}
	if in.LastName != "" {
		user.LastName = in.LastName
	}
	if in.Password != "" {
		user.Password = in.Password
	}
	r.users[in.Email] = user
	return user, nil
}

func (r Repository) DeleteUser(email string) (User, error) {
	user, ok := r.users[email]
	if !ok {
		return User{}, ErrUserNotFound
	}
	delete(r.users, email)
	return user, nil
}
