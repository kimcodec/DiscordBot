package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

func MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.BeforeUpdate != nil {
		embed := discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    m.Author.Username,
				IconURL: m.Author.AvatarURL(""),
			},
			Type:  discordgo.EmbedTypeArticle,
			Color: 16776960,
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Value: "**Сообщение изменено в <#" + m.ChannelID + ">**" +
						" [Перейти к сообщению]" +
						fmt.Sprintf("(https://discordapp.com/channels/%s/%s/%s)", m.GuildID, m.ChannelID, m.ID),
				},
				&discordgo.MessageEmbedField{
					Name:  "Старое сообщение",
					Value: m.BeforeUpdate.Content,
				},
				&discordgo.MessageEmbedField{
					Name:  "Новое сообщение",
					Value: m.Content,
				}},
			Timestamp: m.Timestamp.Format(time.RFC3339),
		}
		logsChangeChannelID := "692775847572144158"
		_, err := s.ChannelMessageSendEmbed(logsChangeChannelID, &embed)
		if err != nil {
			log.Println("Cant send embed ", err)
		}
	}
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if err := s.State.MessageAdd(m.Message); err != nil {
		log.Println("can't add message to cache ", err)
	}
	if strings.Contains(m.Content, "иди нахуй") {
		_, err := s.ChannelMessageSend(m.ChannelID, "сам иди нахуй!!!")
		if err != nil {
			log.Println("Failed to send message", err)
		}
	}
}
