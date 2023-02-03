package rdb

import (
	"context"
	"fmt"
	"github.com/hcdoit/tiktok/cmd/interact/dal/db"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func InitCommentCountByVideoID(ctx context.Context, videoID int64) (count int, err error) {
	comments, err := db.QueryCommentByVideoID(ctx, videoID)
	if err != nil {
		return 0, err
	}
	count = len(comments)
	err = RDB.Set(ctx, fmt.Sprintf("video-comment-count:%d", videoID), count, consts.ForeverExpireDuration).Err()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetCommentCountByVideoID(ctx context.Context, videoID int64) (count int, err error) {
	key := fmt.Sprintf("video-comment-count:%d", videoID)
	value, err := RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		count, err = InitCommentCountByVideoID(ctx, videoID)
		if err != nil {
			return 0, err
		}
		return count, nil
	}
	if err != nil {
		return 0, err
	}
	count, err = strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return count, nil

}

func AddCommentCountByVideoID(ctx context.Context, videoID int64) (err error) {
	key := fmt.Sprintf("video-comment-count:%d", videoID)
	value, err := RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		_, err = InitCommentCountByVideoID(ctx, videoID)
		if err != nil {
			return err
		}
		return nil
	}
	count, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	count++
	if count <= 0 {
		count, err = InitCommentCountByVideoID(ctx, videoID)
		return err
	}
	return RDB.Set(ctx, fmt.Sprintf("video-comment-count:%d", videoID), count, consts.ForeverExpireDuration).Err()
}

func MinusCommentCountByVideoID(ctx context.Context, videoID int64) (err error) {
	key := fmt.Sprintf("video-comment-count:%d", videoID)
	value, err := RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		_, err = InitCommentCountByVideoID(ctx, videoID)
		if err != nil {
			return err
		}
		return nil
	}
	count, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	count--
	if count < 0 {
		count, err = InitCommentCountByVideoID(ctx, videoID)
		return err
	}
	return RDB.Set(ctx, fmt.Sprintf("video-comment-count:%d", videoID), count, consts.ForeverExpireDuration).Err()
}
