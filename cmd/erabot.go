package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/iptq/erabot"
	"github.com/iptq/erabot/modcolor"
)

func main() {
	bot := erabot.New()

	// load plugins
	modcolor.Init(bot)

	defer bot.Close()
	go bot.Run()

	// wait for signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
