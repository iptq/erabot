package modcolor

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/iptq/erabot"
)

type ColorCommand struct {
	Description string
}

var command = ColorCommand{
	Description: "change your [color](https://www.w3schools.com/colors/colors_names.asp) (ex: `!color DodgerBlue` or `!color none`)",
}

func successReact(s *discordgo.Session, m *discordgo.Message) {
	s.MessageReactionAdd(m.ChannelID, m.ID, "\xf0\x9f\x91\x8d")
}

func (c ColorCommand) Handle(argv []string, s *discordgo.Session, m *discordgo.Message) error {
	if len(argv) < 2 {
		return errors.New("please choose a color (ex: `!color DodgerBlue`)")
	}

	colorName := strings.ToLower(argv[1])

	value, ok := colorMap[colorName]
	if colorName != "none" && !ok {
		return errors.New("that color doesn't exist. See: https://www.w3schools.com/colors/colors_names.asp")
	}

	// get roles
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return err
	}
	roles, err := s.GuildRoles(channel.GuildID)
	if err != nil {
		return err
	}

	var colorRole *discordgo.Role
	var colorRoleFound = false
	roleMap := make(map[string]string)
	for _, role := range roles {
		roleMap[role.ID] = role.Name
		if role.Name == "Color: "+colorName {
			colorRole = role
			colorRoleFound = true
			break
		}
	}

	member, err := s.GuildMember(channel.GuildID, m.Author.ID)
	if err != nil {
		return err
	}

	// remove existing colors
	// log.Printf("%+v\n", roleMap)
	for _, roleID := range member.Roles {
		role, ok := roleMap[roleID]
		if ok && strings.HasPrefix(role, "Color: ") {
			err := s.GuildMemberRoleRemove(channel.GuildID, m.Author.ID, roleID)
			if err != nil {
				return err
			}
		}
	}

	if colorName != "none" {
		if !colorRoleFound {
			// create the role
			newRole, err := s.GuildRoleCreate(channel.GuildID)
			if err != nil {
				return err
			}
			role, err := s.GuildRoleEdit(channel.GuildID, newRole.ID, "Color: "+colorName, value, false, 0, false)
			colorRole = role
			if err != nil {
				return err
			}
			_, err = s.GuildRoleReorder(channel.GuildID, append([]*discordgo.Role{role}, roles...))
			if err != nil {
				return err
			}
		}

		// add current role
		err = s.GuildMemberRoleAdd(channel.GuildID, m.Author.ID, colorRole.ID)
		if err != nil {
			return err
		}
	}

	successReact(s, m)
	return nil
}

func (c ColorCommand) GetDescription() string {
	return c.Description
}

func Init(bot *erabot.Bot) {
	bot.RegisterCommand("color", command)
}
