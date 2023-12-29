package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

var (
	logsDeleteChannelID string = "692775888173138010"
	logsChangeChannelID string = "692775847572144158"
	logsPhotoChannelID  string = "692775921878433792"
)

func MessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic happened when delete message ", r)
		}
	}()

	if m.BeforeDelete == nil {
		return
	}

	logs, err := s.GuildAuditLog(m.GuildID, "", "", 0, 0)
	if err != nil {
		log.Println("Can't get audit log ", err)
	}

	UserDeleteBy := ""
	for _, entry := range logs.AuditLogEntries {
		if *entry.ActionType == discordgo.AuditLogActionMessageDelete {
			if entry.TargetID != entry.UserID { // Не ясно, TargetID - это ID сообщения или пользователя
				UserDeleteBy = entry.UserID
			}
			break
		}
	}

	if UserDeleteBy == "" {
		UserDeleteBy = m.BeforeDelete.Author.ID
	}

	// Формируем вложение
	embed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.BeforeDelete.Author.Username,
			IconURL: m.BeforeDelete.Author.AvatarURL(""),
		},
		Type:  discordgo.EmbedTypeArticle,
		Color: 15548997,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Value: fmt.Sprintf("**Удалено** <@!%s>", UserDeleteBy),
			},
			&discordgo.MessageEmbedField{
				Value: fmt.Sprintf("**Сообщение от** <@!%s> **было удалено в** <#%s>",
					m.BeforeDelete.Author.ID, m.BeforeDelete.ChannelID),
			},
			&discordgo.MessageEmbedField{
				Value: m.BeforeDelete.Content,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if _, err := s.ChannelMessageSendEmbed(logsDeleteChannelID, &embed); err != nil {
		log.Println("Cant send embed when Delete ", err)
	}
}

func MessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if m.BeforeUpdate == nil || m.Author == nil || m.Author.Bot {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic happened when update message ", r)
		}
	}()

	if m.Content == m.BeforeUpdate.Content {
		return
	}
	// Формируем вложение
	embed := discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL(""),
		},
		Type:  discordgo.EmbedTypeArticle,
		Color: 16776960,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Value: fmt.Sprintf("**Сообщение изменено в <#%s>**", m.ChannelID) +
					"\n[Перейти к сообщению]" +
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

	_, err := s.ChannelMessageSendEmbed(logsChangeChannelID, &embed)
	if err != nil {
		log.Println("Cant send embed when Update ", err)
	}
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Author.Bot {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic happened when create message ", r)
		}
	}()
	// Добавляем в стек, чтобы видеть сообщения до и после изменения
	if err := s.State.MessageAdd(m.Message); err != nil {
		log.Println("can't add message to cache ", err)
	}

	// Проверяем сообщение на наличие картинок
	for _, v := range m.Attachments {
		if strings.HasPrefix(v.ContentType, "image/") {
			// Формируем вложение
			embed := discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{
					Name:    m.Author.Username,
					IconURL: m.Author.AvatarURL(""),
				},
				Type:  discordgo.EmbedTypeArticle,
				Color: 255,
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Value: fmt.Sprintf("**Сообщение от <@!%s> отправлено в <#%s>**", m.Author.ID,
							m.ChannelID) + "\n[Перейти к сообщению]" +
							fmt.Sprintf("(https://discordapp.com/channels/%s/%s/%s)",
								m.GuildID, m.ChannelID, m.ID),
					},
					&discordgo.MessageEmbedField{
						Value: m.Content,
					},
				},
				Image: &discordgo.MessageEmbedImage{
					URL: v.URL,
				},
				Timestamp: m.Timestamp.Format(time.RFC3339),
			}
			// Отправляем в канал с логами
			if _, err := s.ChannelMessageSendEmbed(logsPhotoChannelID, &embed); err != nil {
				log.Println("cant send embed to logs-photo ", err)
			}
		}
	}

	if strings.Contains(m.Content, "иди нахуй") {
		_, err := s.ChannelMessageSend(m.ChannelID, "сам иди нахуй!!!")
		if err != nil {
			log.Println("Failed to send message", err)
		}
	}
}
