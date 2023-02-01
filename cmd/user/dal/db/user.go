package db

import (
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUserByName query list of user info
func QueryUserByName(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUserByID query list of user info
func QueryUserByID(ctx context.Context, userID int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", userID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
