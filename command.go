package erabot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Handle([]string, *discordgo.Session, *discordgo.Message) error
	GetDescription() string
}

func (bot *Bot) generateHelp() string {
	lines := make([]string, 0)
	lines = append(lines, "**!help**: display this help")
	for name, command := range bot.commands {
		lines = append(lines, fmt.Sprintf("**!%s**: %s", name, command.GetDescription()))
	}
	return strings.Join(lines, "\n")
}
