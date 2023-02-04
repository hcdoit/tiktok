package dal

import (
	"github.com/hcdoit/tiktok/cmd/social/dal/mdb"
	"github.com/hcdoit/tiktok/cmd/social/dal/rdb"
)

func Init() {
	rdb.Init()
	mdb.Init()
}
