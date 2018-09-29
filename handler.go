package erabot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot *Bot) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		// disregard bot activity
		return
	}

	if strings.HasPrefix(m.Content, "!") {
		line := strings.TrimLeft(m.Content, "!")

		// TODO: some kind of quote parser for this
		argv := strings.Split(line, " ")
		if len(argv) == 0 {
			return
		}
		command, ok := bot.commands[argv[0]]

		// special help command
		if argv[0] == "help" {
			help := bot.generateHelp()
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Description: help,
			})
		}

		// switch on command
		if !ok {
			// just don't do anything
			return
		} else {
			err := command.Handle(argv, s, m.Message)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%s> error: %s", m.Author.ID, err))
			}
		}
	}
}
