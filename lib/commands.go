package lib

import (
	"fmt"
	"strconv"
	"strings"
)

func ref(s string) *string {
	return &s
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func consumeInt(s string) (string, int, error) {
	endpos := 0
	for endpos < len(s) && isDigit(s[endpos]) {
		endpos++
	}

	i, err := strconv.Atoi(s[:endpos])
	if err != nil {
		return s, i, err
	}

	return s[endpos:], i, nil
}

// RunCommand attempts to roll dice as specified by cmd, returning a string
// with the results on success, or nil if no valid command was parsed.
func RunCommand(cmd string, d *Dice) *string {
	if !strings.HasPrefix(cmd, "/roll") {
		return nil
	}
	cmd = cmd[5:]

	cmd, sides, err := consumeInt(cmd)
	if err != nil {
		return nil
	}

	srmode := false
	ddmode := false
	if strings.HasPrefix(cmd, "sr") {
		srmode = true
		cmd = cmd[2:]
	} else if strings.HasPrefix(cmd, "dd") {
		ddmode = true
		cmd = cmd[2:]
	} else if strings.HasPrefix(cmd, "dnd") {
		ddmode = true
		cmd = cmd[3:]
	} else if strings.HasPrefix(cmd, "d&d") {
		ddmode = true
		cmd = cmd[3:]
	}

	var numdice int
	if len(cmd) == 0 {
		numdice = 1
	} else {
		matched, _ := fmt.Sscanf(cmd, " %d", &numdice)
		if matched != 1 {
			return nil
		}
	}

	if sides == 6 && !ddmode {
		return ref(d.Roll6Shadowrun(numdice))
	}

	if sides == 20 && !srmode && numdice == 1 {
		return ref(d.Roll20DnD())
	}

	return ref(d.Roll(sides, numdice))
}
