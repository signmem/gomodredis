package db

import (
	"fmt"
	"github.com/signmem/gomodredis/g"
	"sync"
)

var (
	m *sync.RWMutex
	hostgroup = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m"}
)

func SyncTest(ExitChan chan bool) {
	max := g.Config().Conncurrency

	var wg sync.WaitGroup
	wg.Add(max)

	done := make(chan struct{})
	i := 1
	for {
		if i < 99 {
			if g.Config().Debug {
				fmt.Println(i)
			}
			i += 1
		} else {
			i = 1
		}
		go work(done,  &wg)
	}
        wg.Add(1)
        go func(wg *sync.WaitGroup) {
                defer wg.Done()
        }(&wg)
        close(done)
        wg.Wait()
	ExitChan <- true
}

func work(done <-chan struct{},  wg *sync.WaitGroup) {
		for _, name := range hostgroup {
			InputHost(name)
		}
        defer func() {
            wg.Done()
        }()
        <-done
}