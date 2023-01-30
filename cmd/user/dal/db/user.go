package db

import (
	"context"
	"github.com/hcdoit/tiktok/pkg/consts"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password string `gorm:"type:varchar(256);not null" json:"password"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
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
