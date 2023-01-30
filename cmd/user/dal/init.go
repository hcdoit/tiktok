package dal

import (
	"github.com/hcdoit/tiktok/cmd/user/dal/db"
	"github.com/hcdoit/tiktok/cmd/user/dal/rdb"
)

func Init() {
	db.Init()
	rdb.Init()
}
