package main

import (
    "fmt"
    "github.com/lty5240/blue/service-user/proto"
    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/registry"
    "github.com/micro/go-plugins/registry/consul/v2"
    "log"
    "os"
)

func main() {
    db, err := CreateConnection()
    if db != nil {
        defer db.Close()
    }
    if err != nil {
        log.Fatalf("无法连接数据库，异常：%v", err)
    }

    db.AutoMigrate(&proto.User{})

    repository := &UserRepository{db}
    tokenService := &TokenService{repository}

    registryServer := os.Getenv("REGISTRY_SERVER")

    reg := consul.NewRegistry(func(opt *registry.Options) {
        opt.Addrs = []string{
            registryServer,
        }
    })
    srv := micro.NewService(
        micro.Registry(reg),
        micro.Name("cn.lintyone.user.service"),
    )
    srv.Init()
    proto.RegisterUserServiceHandler(srv.Server(), &service{repository, tokenService})

    if err := srv.Run(); err != nil {
        fmt.Printf("服务运行异常：%s \n", err)
    }
}
