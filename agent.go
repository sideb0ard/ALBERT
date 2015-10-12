package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func newAgent() *Agent {
	return &Agent{
		WurdMap:  make(map[string]int),
		Memories: make(map[time.Time]Event)}
}

func (t *Agent) respond(input string) string {

	input = strings.ToLower(input)
	re := regexp.MustCompile("[^a-z]")
	input = re.ReplaceAllLiteralString(input, " ")
	input = strings.TrimSpace(input)

	t.Memories[time.Now()] = Event{What: input}

	wurds := strings.Split(input, " ")
	for _, w := range wurds {
		nl, _ := regexp.MatchString("^$", w)
		if nl {
			continue
		}
		t.WurdMap[w] += 1
	}

	fmt.Println(t.Memories)
	return input
}
