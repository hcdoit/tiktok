package rdb

import (
	"context"
	"fmt"
	"github.com/hcdoit/tiktok/cmd/interact/dal/db"
	"github.com/hcdoit/tiktok/pkg/consts"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func InitFavoriteCountByVideoID(ctx context.Context, videoID int64) (count int, err error) {
	favorites, err := db.QueryFavoriteByVideoID(ctx, videoID)
	if err != nil {
		return 0, err
	}
	count = len(favorites)
	err = RDB.Set(ctx, fmt.Sprintf("video-favorite-count:%d", videoID), count, consts.ForeverExpireDuration).Err()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetFavoriteCountByVideoID(ctx context.Context, videoID int64) (count int, err error) {
	key := fmt.Sprintf("video-favorite-count:%d", videoID)
	value, err := RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		count, err = InitFavoriteCountByVideoID(ctx, videoID)
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

func AddFavoriteCountByVideoID(ctx context.Context, videoID int64) (err error) {
	key := fmt.Sprintf("video-favorite-count:%d", videoID)
	value, err := RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		_, err = InitFavoriteCountByVideoID(ctx, videoID)
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
		count, err = InitFavoriteCountByVideoID(ctx, videoID)
		return err
	}
	return RDB.Set(ctx, fmt.Sprintf("video-favorite-count:%d", videoID), count, consts.ForeverExpireDuration).Err()
}

func MinusFavoriteCountByVideoID(ctx context.Context, videoID int64) (err error) {
	key := fmt.Sprintf("video-favorite-count:%d", videoID)
	value, err := RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		_, err = InitFavoriteCountByVideoID(ctx, videoID)
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
		count, err = InitFavoriteCountByVideoID(ctx, videoID)
		return err
	}
	return RDB.Set(ctx, fmt.Sprintf("video-favorite-count:%d", videoID), count, consts.ForeverExpireDuration).Err()
}
