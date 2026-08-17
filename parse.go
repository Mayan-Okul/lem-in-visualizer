package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseColony(r io.Reader) (*Colony, error) {
	scanner := bufio.NewScanner(r)
	colony := &Colony{}
	stage := "ants"
	nextIsStart := false
	nextIsEnd := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if stage == "ants" {
    if strings.HasPrefix(line, "ERROR") {
        return nil, fmt.Errorf("lem-in reported an error: %s", line)
    }
    n, err := strconv.Atoi(line)
    if err != nil {
        return nil, fmt.Errorf("invalid number of ants: %w", err)
    }
    colony.NumAnts = n
    stage = "graph"
    continue
}

		if line == "##start" {
			nextIsStart = true
			continue
		}
		if line == "##end" {
			nextIsEnd = true
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		// Turn line: any line whose first token starts with 'L' is a
		// move line (Lant-room), never a room, since room names can't
		// start with 'L' per the spec.
		if strings.HasPrefix(fields[0], "L") {
			var turn Turn
			for _, tok := range fields {
				tok = strings.TrimPrefix(tok, "L")
				parts := strings.SplitN(tok, "-", 2)
				if len(parts) != 2 {
					continue
				}
				turn.Moves = append(turn.Moves, Move{Ant: parts[0], Room: parts[1]})
			}
			colony.Turns = append(colony.Turns, turn)
			continue
		}

		// Room definition: "name x y"
		if len(fields) == 3 {
			x, errX := strconv.Atoi(fields[1])
			y, errY := strconv.Atoi(fields[2])
			if errX == nil && errY == nil {
				colony.Rooms = append(colony.Rooms, Room{
					Name:    fields[0],
					X:       x,
					Y:       y,
					IsStart: nextIsStart,
					IsEnd:   nextIsEnd,
				})
				nextIsStart = false
				nextIsEnd = false
				continue
			}
		}

		// Link definition: "name1-name2"
		if len(fields) == 1 && strings.Contains(line, "-") {
			parts := strings.SplitN(line, "-", 2)
			if len(parts) == 2 {
				colony.Links = append(colony.Links, Link{
					From: parts[0],
					To:   parts[1],
				})
			}
			continue
		}

		// unknown command: ignored, per spec
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return colony, nil
}