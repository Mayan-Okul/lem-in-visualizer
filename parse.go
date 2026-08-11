package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseColony(r *os.File) (*Colony, error) {
	scanner := bufio.NewScanner(r)
	c := &Colony{}

	nextIsStart, nextIsEnd := false, false
	stage := "ants" // ants -> graph -> turns

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			stage = "turns"
			continue
		}

		switch {
		case stage == "ants":
			n, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("expected ant count, got %q", line)
			}
			c.NumAnts = n
			stage = "graph"

		case line == "##start":
			nextIsStart = true

		case line == "##end":
			nextIsEnd = true

		case stage == "graph" && strings.Contains(line, "-") && !strings.HasPrefix(line, "L"):
			parts := strings.SplitN(line, "-", 2)
			c.Links = append(c.Links, Link{From: parts[0], To: parts[1]})

		case stage == "graph":
			fields := strings.Fields(line)
			if len(fields) != 3 {
				continue
			}
			x, err1 := strconv.Atoi(fields[1])
			y, err2 := strconv.Atoi(fields[2])
			if err1 != nil || err2 != nil {
				continue
			}
			c.Rooms = append(c.Rooms, Room{
				Name: fields[0], X: x, Y: y,
				IsStart: nextIsStart, IsEnd: nextIsEnd,
			})
			nextIsStart, nextIsEnd = false, false

		case stage == "turns":
			var turn Turn
			for _, tok := range strings.Fields(line) {
				tok = strings.TrimPrefix(tok, "L")
				parts := strings.SplitN(tok, "-", 2)
				if len(parts) != 2 {
					continue
				}
				turn.Moves = append(turn.Moves, Move{Ant: parts[0], Room: parts[1]})
			}
			c.Turns = append(c.Turns, turn)
		}
	}

	return c, scanner.Err()
}
