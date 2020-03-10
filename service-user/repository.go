package main

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/micor-lzy/service-user/user"

)

type Repository interface {
	GetAll(ctx context.Context)([]*user.User,error)
	Get(ctx context.Context,id string )(*user.User,error)
	Create(ctx context.Context,user *user.User)error
	GetByEmailAndPassword(ctx context.Context,user *user.User)error
}

type UserRepository struct {
	db *gorm.DB
}

func (userRepo UserRepository) GetAll(ctx context.Context) ([]*user.User, error) {
	var users []*user.User
	if err := userRepo.db.Find(&users).Error; err != nil{
		return nil,err
	}
	return users,nil
}

func (userRepo UserRepository) Get(ctx context.Context, id string) (*user.User, error) {
	var user *user.User
	user.Id = id
	if err :=userRepo.db.First(&user).Error;err !=nil{
		return nil,err

	}
	return user,nil
}

func (userRepo UserRepository) Create(ctx context.Context, user *user.User) error {
	if err := userRepo.db.Create(user).Error;err !=nil{
		return err
	}
	return nil
}

func (userRepo UserRepository) GetByEmailAndPassword(ctx context.Context, user *user.User) error {
	if err := userRepo.db.First(&user).Error;err !=nil{
		return err
	}
	return nil
}
