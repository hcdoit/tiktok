package rdb

import (
	"context"
	"fmt"
	"github.com/hcdoit/tiktok/cmd/interact/dal/db"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/redis/go-redis/v9"
	"strconv"
)

// InitCommentCountByVideoID 缓存为空时从数据库读取并更新缓存
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

// GetCommentCountByVideoID 获取缓存，若无缓存则更新
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

// AddCommentCountByVideoID 从已有缓存更新缓存+1
func AddCommentCountByVideoID(ctx context.Context, videoID int64) (err error) {
	key := fmt.Sprintf("video-comment-count:%d", videoID)
	value, err := RDB.Get(ctx, key).Result()
	// 无缓存则直接从数据库更新缓存
	if err == redis.Nil {
		_, err = InitCommentCountByVideoID(ctx, videoID)
		if err != nil {
			return err
		}
		return nil
	}
	// 存在缓存
	if err == nil {
		// 先删除保证数据一致性
		RDB.Del(ctx, key)
		// 更新数据+1
		count, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		count++
		if count <= 0 {
			count, err = InitCommentCountByVideoID(ctx, videoID)
			return err
		}
		// 将更新的数据写回缓存
		return RDB.Set(ctx, fmt.Sprintf("video-comment-count:%d", videoID), count, consts.ForeverExpireDuration).Err()
	}
	return err
}

// MinusCommentCountByVideoID 从已有缓存更新缓存-1
func MinusCommentCountByVideoID(ctx context.Context, videoID int64) (err error) {
	key := fmt.Sprintf("video-comment-count:%d", videoID)
	value, err := RDB.Get(ctx, key).Result()
	// 无缓存则直接从数据库更新缓存
	if err == redis.Nil {
		_, err = InitCommentCountByVideoID(ctx, videoID)
		if err != nil {
			return err
		}
		return nil
	}
	// 存在缓存
	if err == nil {
		// 先删除保证数据一致性
		RDB.Del(ctx, key)
		// 更新数据-1
		count, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		count--
		if count < 0 {
			count, err = InitCommentCountByVideoID(ctx, videoID)
			return err
		}
		// 将更新的数据写回缓存
		return RDB.Set(ctx, fmt.Sprintf("video-comment-count:%d", videoID), count, consts.ForeverExpireDuration).Err()
	}
	return err
}
