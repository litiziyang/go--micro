package main

import (
	"errors"
	"github.com/micor-lzy/service-user/user"
	 "golang.org/x/net/context"
)


type service struct {
	repo Repository
	token Authable
}

func (srv service) Create(ctx context.Context, req*user.User, res*user.Response) error {
	err :=srv.repo.Create(ctx,req)
	if err !=nil{
		return err
	}
	return nil
}

func (srv service) Get(ctx context.Context, req*user.User, res*user.Response) error {
	user,err := srv.repo.Get(ctx,req.Id)
	if err !=nil{
		return err
	}
	res.User = user
	return nil
}

func (srv service) GetAll(ctx context.Context, req*user.Request, res*user.Response) error {
	users,err :=srv.repo.GetAll(ctx)
	if err !=nil{
		return err
	}
	res.Users = users
	return nil
}

func (srv service) Auth(ctx context.Context, req *user.User, res*user.Token) error {
	err :=srv.repo.GetByEmailAndPassword(ctx, req)
	if err !=nil{
		return  err
	}

	res.Token , err =srv.token.Encode(req)
	if err !=nil{
		return err
	}
    return nil
}

func (srv service) ValidateToken(ctx context.Context, req*user.Token, res*user.Token) error {
	claims,err:=srv.token.Decode(req.Token)
	if err !=nil{
		return err
	}
	if claims.User.Id ==" "{
		return errors.New("incalid user")
	}
	res.Valid=true
	return nil
}

