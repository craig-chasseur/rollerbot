package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	rollerbot "github.com/craig-chasseur/rollerbot/lib"
)

var dice *rollerbot.Dice

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	res := rollerbot.RunCommand(m.Content, dice)
	if res == nil {
		return
	}
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" "+*res)
}

func main() {
	token := flag.String("t", "", "Discord Bot Token")
	flag.Parse()

	disc, err := discordgo.New("Bot " + *token)
	if err != nil {
		log.Fatalln("Error creating Discord session ", err)
	}

	dice = rollerbot.New()

	disc.AddHandler(messageCreate)

	err = disc.Open()
	if err != nil {
		log.Fatalln("Error opening connection ", err)
	}

	fmt.Println("Bot is running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	disc.Close()
}
