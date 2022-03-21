package main

import (
	"github.com/NineSwards/go-job/server"
	"github.com/robfig/cron"
	"log"
)

func main() {
	i := 0
	c := cron.New()
	spec := "*/5 * * * * ?"
	_ = c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})

	c.Start()

	server.RunServer()
	select {}
}