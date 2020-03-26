package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	timeLocal := time.Now()
	fmt.Println("current time:", timeLocal.Round(0))

	host := "0.beevik-ntp.pool.ntp.org"
	timeNtp, err := ntp.Time(host)
	if err != nil {
		log.Fatalf("can't get time from %s, error: %s\n", host, err)
	}
	fmt.Println("exact time:", timeNtp.Round(0))
}
