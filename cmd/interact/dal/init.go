package dal

import (
	"github.com/hcdoit/tiktok/cmd/interact/dal/db"
	"github.com/hcdoit/tiktok/cmd/interact/dal/rdb"
)

func Init() {
	db.Init()
	rdb.Init()
}
