package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/iptq/erabot"
	"github.com/iptq/erabot/modcolor"
)

func main() {
	var generateConfig bool
	var confFile string

	flag.BoolVar(&generateConfig, "genConf", false, "Generate a default config to use.")
	flag.StringVar(&confFile, "conf", "erabot.conf", "Path to the configuration file.")
	flag.Parse()

	if generateConfig {
		w := toml.NewEncoder(os.Stdout)
		conf := erabot.DefaultConfig
		w.Encode(conf)
		os.Exit(0)
	}

	var conf erabot.Config
	contents, _ := ioutil.ReadFile(confFile)
	toml.Decode(string(contents), &conf)

	bot := erabot.New(&conf)
	modcolor.Init(bot)

	defer bot.Close()
	go bot.Run()

	// wait for signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
