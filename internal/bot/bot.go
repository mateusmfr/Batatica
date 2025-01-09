package bot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

func InitializeBot(token string) error {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		return err
	}

	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!status" {
		s.ChannelMessageSend(m.ChannelID, "ON!")
	}

	if m.Content == "!testapi" {
		resp, err := http.Get("http://localhost:8080/api/test")
		if err != nil {
			log.Printf("Erro ao chamar API: %v", err)
			s.ChannelMessageSend(m.ChannelID, "Erro ao chamar a API.")
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			s.ChannelMessageSend(m.ChannelID, "A API respondeu com sucesso!")
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Erro na API: Status %d", resp.StatusCode))
		}
	}
}
