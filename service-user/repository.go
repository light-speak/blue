package main

import (
    "github.com/jinzhu/gorm"
    "github.com/lty5240/blue/service-user/proto"
)

type Repository interface {
    Get(id int64) (*proto.User, error)
    Create(user *proto.User) error
    GetByPhone(phone string) (*proto.User, error)
}

type UserRepository struct {
    db *gorm.DB
}

func (repository *UserRepository) GetByPhone(phone string) (*proto.User, error) {
    user := &proto.User{}
    user.Phone = phone
    if err := repository.db.First(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (repository *UserRepository) Get(id int64) (*proto.User, error) {
    user := &proto.User{}
    user.Id = id
    if err := repository.db.First(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (repository *UserRepository) Create(user *proto.User) error {
    if err := repository.db.Create(user).Error; err != nil {
        return err
    }
    return nil
}
