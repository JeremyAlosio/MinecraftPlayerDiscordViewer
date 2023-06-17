package v1

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func StartBot() {
	// Discord bot token (replace with your own bot token)
	token := os.Getenv("DISCORDBOTTOKEN")

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Failed to create Discord session: %v", err)
	}

	// Register a callback for the ready event
	dg.AddHandler(onReady)

	// Open a connection to Discord
	err = dg.Open()
	if err != nil {
		log.Fatalf("Failed to open connection to Discord: %v", err)
	}

	log.Println("Bot is running. Press CTRL-C to exit.")

	// Schedule the initial status update
	updateStatus(dg)

	// Start a goroutine to update the status every 30 seconds
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			updateStatus(dg)
		}
	}()

	// Wait until CTRL-C or SIGINT signal is received
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Clean up and close the Discord session
	dg.Close()
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Println("Bot is ready!")
}

func updateStatus(s *discordgo.Session) {
	err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Activities: []*discordgo.Activity{
			{
				Name: GetMinecraftPlayerInfo(),
				Type: discordgo.ActivityTypeGame,
			},
		},
	})
	if err != nil {
		log.Printf("Failed to update bot status: %v", err)
	}
}
