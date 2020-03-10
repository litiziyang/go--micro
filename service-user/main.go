package main

import (
	"github.com/micor-lzy/service-user/user"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcd/v2"
	"time"
)

func main() {
	db, err := CreateCnnection()
	if db != nil {
		defer db.Close()
	}

	if err != nil {
		log.Fatalf("Could not connect to DB:%v", err)
	}

	db.AutoMigrate(&user.User{})
	repo := &UserRepository{db}
	token := &TokenService{repo}

	registerDrive := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"etcd1:2379","etcd2:2379","etcd3:2379",
		}
	})

	srv := micro.NewService(
		micro.Name("service_user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		micro.Registry(registerDrive))
	srv.Init()

	user.RegisterUserServiceHandler(srv.Server(), &service{repo, token})

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
