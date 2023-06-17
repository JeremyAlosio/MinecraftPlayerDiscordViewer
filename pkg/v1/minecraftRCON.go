package v1

import (
	"log"
	"os"
	"strings"

	"github.com/willroberts/minecraft-client"
)

func GetMinecraftPlayerInfo() string {
	// RCON connection details
	hostport := "localhost:25575"
	password := os.Getenv("RCONPASSWORD")

	// Create a new Minecraft RCON client
	client, err := minecraft.NewClient(hostport)
	if err != nil {
		log.Fatalf("Failed to create RCON client: %v", err)
	}
	defer client.Close()

	// Authenticate with the RCON password
	if err := client.Authenticate(password); err != nil {
		log.Fatalf("RCON authentication failed: %v", err)
	}

	// Send an RCON command to retrieve the player count
	response, err := client.SendCommand("/list")
	if err != nil {
		log.Fatalf("Failed to send RCON command: %v", err)
	}

	return formatOutput(response.Body)
}

func formatOutput(body string) string {
	formattedMessage := strings.Builder{}

	splitBody := strings.Split(body, ":")

	formattedMessage.WriteString("Current Players on the Server: \u2028\u2028\u2028\u2028\u2028\u2028" + splitBody[1])

	return formattedMessage.String()
}
