package main

import (
	"flag"
	"fmt"
	"github.com/signmem/gomodredis/db"
	"github.com/signmem/gomodredis/g"
	"github.com/signmem/gomodredis/http"
	"log"
	"os"
)

func init () {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		version := g.Version
		fmt.Printf("%s", version)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	g.InitLog()
	log.Println("[INFO] log init success.")
	db.Server = g.Config().Redis.Server
	db.Maxidle = g.Config().Redis.MaxIdle
	db.MaxActive = g.Config().Redis.MaxActive
	db.Idletimeout = g.Config().Redis.IdleTimeout

	if g.Config().Debug {
		log.Println("+++++++++++++++++++++++++++++++++++++")
		log.Println("redis server is:", db.Server)
		log.Printf("redis params: maxidle %d, maxactive %d, maxactive %d \n",
			db.Maxidle, db.MaxActive, db.Idletimeout)
		log.Println("+++++++++++++++++++++++++++++++++++++")
	}
	db.Pool = db.NewPool(db.Maxidle, db.MaxActive, db.Idletimeout, db.Server )
	db.CleanupHook()
}

func main() {

	ExitChan := make(chan bool, 40)

	go http.Start()
	for i:=0; i < g.Config().Conncurrency ; i++ {
		go db.SyncTest(ExitChan)
	}

	close(ExitChan)
	select {}

}