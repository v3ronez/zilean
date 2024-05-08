package bot

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var s *discordgo.Session
var err error

func initEnv() {
	path, err := os.Getwd()
	if err != nil {
		slog.Error("Error to load environment variables", "err", err)
		return
	}
	if err := godotenv.Load(path + "/.env"); err != nil {
		slog.Error("Error to load environment variables", "err", err)
		return
	}
}
func init() {
	initEnv()
	var err error
	s, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	s.Identify.Intents = discordgo.IntentMessageContent
	s.Identify.Intents = discordgo.IntentsGuildMessages
}

var (
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer
	commands                       = []*discordgo.ApplicationCommand{
		{
			Name:        "basic-command",
			Description: "Basic command",
		},
	}
	commandHandlers = map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate){
		"basic-command": func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
			s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
	}
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	})
}

func Run() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	if err = s.Open(); err != nil {
		log.Fatalf("Error message: %s", err.Error())
		return
	}
	defer s.Close()
	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for idx, command := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, command.GuildID, command)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", command.Name, err)
		}
		registeredCommands[idx] = cmd
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

}
