package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	ChannelID = ""
)

func Run() {
	sess, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return
	}
	sess.Identify.Intents = discordgo.IntentMessageContent
	sess.Identify.Intents = discordgo.IntentsGuildMessages
	sess.AddHandler(handlerCommands)
	if err = sess.Open(); err != nil {
		log.Fatalf("Error message: %s", err.Error())
		return
	}
	defer sess.Close()
	fmt.Println("Bot running...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	switch {
	case sig == os.Interrupt:
		_, err := sess.ChannelMessageSend(ChannelID, "Haha até a próxima, 👋🏻🥸")
		if err != nil {
			fmt.Printf("err to shutdown: %s", err)
		}
		sess.Close()
	}

}

func handlerCommands(session *discordgo.Session, message *discordgo.MessageCreate) {
	command := NewCommand(session, message)
	ChannelID = message.ChannelID
	if session.State.User.ID == message.Author.ID {
		return
	}
	isValid := command.Validate(message.Content)
	if !isValid {
		session.ChannelMessageSend(message.ChannelID, "Comando inválido amigão. tente: `zilean !<comando>`")
		return
	}
	switch {
	case strings.Contains(message.Content, "on"):
		members, err := command.MembersOnline()
		if err != nil {
			fmt.Printf("err: %s", err)
		}
		session.ChannelMessageSend(message.ChannelID, members)
		return
	case strings.Contains(message.Content, "commands"):
		commands := command.ShowCommands()
		session.ChannelMessageSend(message.ChannelID, commands)
		return
	default:
		session.ChannelMessageSend(message.ChannelID, "Comando inválido amigão. tente: `zilean !<comando>")
	}
}
