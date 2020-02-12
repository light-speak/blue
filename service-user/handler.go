package main

import (
    "github.com/lty5240/blue/service-user/proto"
    "golang.org/x/net/context"
)

type service struct {
    repository   Repository
    tokenService AuthAble
}

func (s service) Create(ctx context.Context, req *proto.RegisterRequest, res *proto.Response) error {
    //TODO: 验证码匹配

    user := req.User
    if err := s.repository.Create(user); err != nil {
        return err
    }
    res.User = user
    return nil
}

func (s service) Login(ctx context.Context, req *proto.LoginRequest, res *proto.Response) error {
    user, err := s.repository.GetByPhone(req.Phone)
    if err != nil {
        return err
    }
    //TODO: 验证码匹配
    token, err := s.tokenService.Encode(user)
    if err != nil {
        return err
    }
    res.Token.Token = token
    return nil
}

func (s service) Get(ctx context.Context, req *proto.User, res *proto.Response) error {
    user, err := s.repository.Get(req.Id)
    if err != nil {
        return err
    }
    res.User = user
    return nil
}

func (s service) Validate(ctx context.Context, req *proto.Token, res *proto.User) error {
    claim, err := s.tokenService.Decode(req.Token)
    if err != nil {
        return err
    }
    res = claim.User
    return nil
}
