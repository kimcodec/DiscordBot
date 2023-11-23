package handlers

import (
	"github.com/bwmarrin/discordgo"

	"log"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "иди нахуй") {
		_, err := s.ChannelMessageSend(m.ChannelID, "сам иди нахуй!!!")
		if err != nil {
			log.Println("Failed to send message", err)
		}
	}
}
