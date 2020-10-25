package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vpenando/werego"
)

var (
	bot *werego.WereBot
)

func init() {
	var t string
	flag.StringVar(&t, "t", "", "Bot token")
	flag.Parse()
	var err error
	bot, err = werego.NewWereBot(t)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	defer bot.Close()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
