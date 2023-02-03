package rdb

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func AddFollow(ctx context.Context, followerID int64, followID int64) error {
	followKey := fmt.Sprintf("follow:%d", followerID)
	followerKey := fmt.Sprintf("follower:%d", followID)
	err := RDB.SAdd(ctx, followKey, followID).Err()
	if err != nil {
		return err
	}
	err = RDB.SAdd(ctx, followerKey, followerID).Err()
	if err != nil {
		RDB.SRem(ctx, followKey, followID)
		return err
	}
	return nil
}
func CancelFollow(ctx context.Context, followerID int64, followID int64) error {
	followKey := fmt.Sprintf("follow:%d", followerID)
	followerKey := fmt.Sprintf("follower:%d", followID)
	err := RDB.SRem(ctx, followKey, followID).Err()
	if err != nil {
		return err
	}
	err = RDB.SRem(ctx, followerKey, followerID).Err()
	if err != nil {
		RDB.SAdd(ctx, followKey, followID)
		return err
	}
	return nil
}

func GetFollowList(ctx context.Context, followerID int64) ([]int64, error) {
	followKey := fmt.Sprintf("follow:%d", followerID)
	followList, err := RDB.SMembers(ctx, followKey).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	followIdList := make([]int64, 0)
	for _, str := range followList {
		if id, err := strconv.Atoi(str); err == nil {
			followIdList = append(followIdList, int64(id))
		}

	}
	return followIdList, err
}
func GetFollowCount(ctx context.Context, followerID int64) (int64, error) {
	followKey := fmt.Sprintf("follow:%d", followerID)
	count, err := RDB.SCard(ctx, followKey).Result()
	if err == redis.Nil {
		return 0, nil
	}
	return count, err
}
func CheckFollow(ctx context.Context, followerID int64, followID int64) (bool, error) {
	followKey := fmt.Sprintf("follow:%d", followerID)
	return RDB.SIsMember(ctx, followKey, followID).Result()
}

func GetFollowerList(ctx context.Context, followID int64) ([]int64, error) {
	followerKey := fmt.Sprintf("follower:%d", followID)
	followerList, err := RDB.SMembers(ctx, followerKey).Result()
	if err != nil {
		return nil, err
	}
	followerIdList := make([]int64, 0)
	for _, str := range followerList {
		if id, err := strconv.Atoi(str); err == nil {
			followerIdList = append(followerIdList, int64(id))
		}

	}
	return followerIdList, err
}

func GetFollowerCount(ctx context.Context, followID int64) (int64, error) {
	followerKey := fmt.Sprintf("follower:%d", followID)
	count, err := RDB.SCard(ctx, followerKey).Result()
	if err == redis.Nil {
		return 0, nil
	}
	return count, err
}

func GetFriendList(ctx context.Context, myID int64) ([]int64, error) {
	followKey := fmt.Sprintf("follow:%d", myID)
	followerKey := fmt.Sprintf("follower:%d", myID)
	friendList, err := RDB.SInter(ctx, followKey, followerKey).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	friendIdIdList := make([]int64, 0)
	for _, str := range friendList {
		if id, err := strconv.Atoi(str); err == nil {
			friendIdIdList = append(friendIdIdList, int64(id))
		}

	}
	return friendIdIdList, err

}
