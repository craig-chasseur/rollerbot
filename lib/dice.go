package lib

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const maxDice = 100
const maxSides = 1000

// Dice represents a threadsafe set of dice that can be rolled.
type Dice struct {
	rng *rand.Rand
	mtx sync.Mutex
}

// New returns new pre-seeded Dice.
func New() *Dice {
	return &Dice{rng: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

// Roll6Shadowrun rolls n D6es and prints the total hits and any glitches
// according to shadowrun rules.
func (d *Dice) Roll6Shadowrun(n int) string {
	d.mtx.Lock()
	defer d.mtx.Unlock()

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
		switch d.rng.Intn(6) {
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
