package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

const maxDice = 100

func roll6(n int) string {
	if n <= 0 {
		return "Nothing to roll"
	}

	if n > maxDice {
		return "I don't have that many dice in my bag"
	}

	hits := 0
	ones := 0
	dice := ""
	for i := 0; i < n; i++ {
		switch rand.Intn(6) {
		case 0:
			dice = dice + "⚀"
			ones++
		case 1:
			dice = dice + "⚁"
		case 2:
			dice = dice + "⚂"
		case 3:
			dice = dice + "⚃"
		case 4:
			dice = dice + "⚄"
			hits++
		case 5:
			dice = dice + "⚅"
			hits++
		}
	}

	glitch := ""
	if ones >= (n+1)/2 {
		if hits == 0 {
			glitch = "CRITICAL GLITCH!!! "
		} else {
			glitch = "GLITCHED! "
		}
	}

	return fmt.Sprintf("%sHits: %d Ones: %d %s", glitch, hits, ones, dice)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages sent by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	var n int
	matched, _ := fmt.Sscanf(m.Content, "/roll6 %d", &n)
	if matched != 1 {
		return
	}
	s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" "+roll6(n))
}

func main() {
	token := flag.String("t", "", "Discord Bot Token")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	disc, err := discordgo.New("Bot " + *token)
	if err != nil {
		log.Fatalln("Error creating Discord session ", err)
	}

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
