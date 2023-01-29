package main

import (
	"github.com/hcdoit/tiktok/cmd/user/dal/db"
	"github.com/hcdoit/tiktok/cmd/user/mw"
	user "github.com/hcdoit/tiktok/kitex_gen/user/userservice"
	"log"
)

func Init() {
	db.Init()
	mw.Init()
}

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
