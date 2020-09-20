package lib

import (
	"fmt"
	"math/rand"
	"strconv"
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

// Roll20DnD rolls a single D20 and prints the result (possibly with an
// annotation for a critical hit or miss) according to Dungeons & Dragons
// rules.
func (d *Dice) Roll20DnD() string {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	result := d.rng.Intn(20)
	message := ""
	switch result {
	case 0:
		message = " CRITICAL MISS!"
	case 19:
		message = " CRITICAL HIT!"
	}
	return fmt.Sprintf("%d%s", result+1, message)
}

// Roll rolls numdice distinct dice with the specified number of sides and
// prints the results of all the rolls, plus their sum.
func (d *Dice) Roll(sides int, numdice int) string {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	if sides < 1 {
		return "Impossible to roll a die with less than 1 side"
	}
	if sides > maxSides {
		return "That is a frankly silly number of sides for a die to have, and I refuse"
	}

	if numdice < 1 {
		return "Nothing to roll"
	}
	if numdice > maxDice {
		return "I don't have that many dice in my bag"
	}

	resultstr := ""
	sum := 0
	for i := 0; i < numdice; i++ {
		roll := d.rng.Intn(sides) + 1
		if len(resultstr) > 0 {
			resultstr += ", "
		}
		resultstr += strconv.Itoa(roll)
		sum += roll
	}

	if numdice == 1 {
		return resultstr
	}

	return fmt.Sprintf("%s --- TOTAL = %d", resultstr, sum)
}
