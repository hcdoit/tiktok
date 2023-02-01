package dal

import (
	"github.com/hcdoit/tiktok/cmd/video/dal/db"
	"github.com/hcdoit/tiktok/cmd/video/dal/minio"
)

func Init() {
	db.Init()
	minio.Init()
}
