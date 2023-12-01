package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

var (
	logsOtherChannelID string = "692775975175585812"
)

func MemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// recover на случай, если что-то из элементов embed будет nil
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic happened when user joined guild ", r)
		}
	}()
	// формируем вложение
	embed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Nick,
			IconURL: m.AvatarURL(""),
		},
		Title: "Зашел на сервер",
		Type:  discordgo.EmbedTypeArticle,
		Color: 5763719,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: m.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Value:  fmt.Sprintf("<@!%s>", m.User.ID),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Value:  m.User.Username,
				Inline: true,
			},
		},
		Timestamp: m.JoinedAt.Format(time.RFC3339),
	}

	if _, err := s.ChannelMessageSendEmbed(logsOtherChannelID, &embed); err != nil {
		log.Println("cant send embed while user join the guild ", err)
	}
}

func MemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	// recover на случай, если что-то из элементов embed будет nil
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic happened when user joined guild ", r)
		}
	}()
	// формируем вложение
	embed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Nick,
			IconURL: m.AvatarURL(""),
		},
		Title: "Вышел с сервера :(",
		Type:  discordgo.EmbedTypeArticle,
		Color: 15548997,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: m.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Value:  fmt.Sprintf("<@!%s>", m.User.ID),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Value:  m.User.Username,
				Inline: true,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if _, err := s.ChannelMessageSendEmbed(logsOtherChannelID, &embed); err != nil {
		log.Println("cant send embed while user join the guild ", err)
	}
}

func MemberUpdate(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {
	// recover на случай, если что-то из элементов embed будет nil
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic happened when user update ", r)
		}
	}()

	if m.Nick == m.BeforeUpdate.Nick {
		return
	}

	logs, err := s.GuildAuditLog(m.GuildID, "", "", 0, 0)
	if err != nil {
		log.Println("Can't get audit log ", err)
	}

	UserUpdateBy := ""
	for _, entry := range logs.AuditLogEntries {
		if *entry.ActionType == discordgo.AuditLogActionMemberUpdate {
			UserUpdateBy = entry.UserID
			break
		}
	}

	if UserUpdateBy == "" {
		UserUpdateBy = m.User.ID
	}

	// формируем вложение
	embed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Nick,
			IconURL: m.AvatarURL(""),
		},
		Type:  discordgo.EmbedTypeArticle,
		Color: 16705372,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: m.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Value: fmt.Sprintf("**Изменен ник** <@!%s>", m.User.ID),
			},
			&discordgo.MessageEmbedField{
				Value: fmt.Sprintf("Кем: <@!%s>", UserUpdateBy),
			},
			&discordgo.MessageEmbedField{
				Name: "До:",
				Value: func() string {
					if m.BeforeUpdate.Nick == "" {
						return m.User.Username
					}
					return m.BeforeUpdate.Nick
				}(),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "После:",
				Value:  m.Nick,
				Inline: true,
			},
		},
		Timestamp: m.JoinedAt.Format(time.RFC3339),
	}

	if _, err := s.ChannelMessageSendEmbed(logsOtherChannelID, &embed); err != nil {
		log.Println("cant send embed while user join the guild ", err)
	}
}
