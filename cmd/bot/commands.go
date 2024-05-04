package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	ShowAllMembersOnline = "on"
	ShowALlCommands      = "commands"
)

type Command struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
}

func NewCommand(s *discordgo.Session, m *discordgo.MessageCreate) *Command {
	return &Command{
		session: s,
		message: m,
	}
}

func (c *Command) MembersOnline() (string, error) {
	members, err := c.session.GuildMembers(c.message.GuildID, "", 100)
	if err != nil {
		fmt.Printf("%+s\n", err.Error())
	}
	var ms string
	for _, m := range members {
		ms = ms + m.User.String() + "\n"
	}
	return ms, nil
}

func (c *Command) Validate(m string) bool {
	return strings.HasPrefix(m, "zilean !")
}

func (c Command) ShowCommands() string {
	return "wip"
}
