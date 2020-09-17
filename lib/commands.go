package lib

import "fmt"

func ref(s string) *string {
	return &s
}

// RunCommand attempts to roll dice as specified by cmd, returning a string
// with the results on success, or nil if no valid command was parsed.
func RunCommand(cmd string, d *Dice) *string {
	var numdice int
	matched, _ := fmt.Sscanf(cmd, "/roll6 %d", &numdice)
	if matched == 1 {
		return ref(d.Roll6Shadowrun(numdice))
	}
	return nil
}
