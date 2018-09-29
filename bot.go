package erabot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session  *discordgo.Session
	config   *Config
	commands map[string]Command
}

func New(config *Config) (bot *Bot) {
	session, _ := discordgo.New("Bot " + config.Token)
	commands := make(map[string]Command)
	bot = &Bot{
		session:  session,
		commands: commands,
		config:   config,
	}

	// add handlers
	bot.session.AddHandler(bot.messageHandler)

	return
}

func (bot *Bot) Close() {
	bot.session.Close()
	log.Println("disconnected")
}

func (bot *Bot) RegisterCommand(name string, command Command) {
	bot.commands[name] = command
}

func (bot *Bot) Run() {
	bot.session.Open()
	log.Println("connected")
}
