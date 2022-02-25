package db

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

var (
	REDISHOSTKEY []string
	FLUSHKEY int
	AGENTPATH = "/agent.alive"
)


func InputHost(hostname string) {
	timeNow := time.Now().Unix()
	timeNowStr := strconv.FormatInt(timeNow, 10)
	fmt.Println(timeNowStr)
	fullName := AGENTPATH + "/" + hostname
	err := Set(fullName, timeNowStr)
	if err != nil {
		log.Print(err)
	}

}
