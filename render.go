package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func renderPage(w http.ResponseWriter, c *Colony) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var graph strings.Builder

	for _, link := range c.Links {
		graph.WriteString(fmt.Sprintf(
			`<div class="link" data-from="%s" data-to="%s"></div>`,
			link.From,
			link.To,
		))
	}

	for _, room := range c.Rooms {
		class := "room"

		if room.IsStart {
			class += " start"
		}

		if room.IsEnd {
			class += " end"
		}

		graph.WriteString(fmt.Sprintf(
			`<div class="%s" id="room-%s" data-name="%s" data-x="%d" data-y="%d" style="left:%d%%;top:%d%%;">
				<div class="room-node"></div>
				<div class="room-label">%s</div>
			</div>`,
			class,
			room.Name,
			room.Name,
			room.X,
			room.Y,
			room.X,
			room.Y,
			room.Name,
		))
	}

	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>FORMICARIUM_OS_v1.0</title>
	<link rel="stylesheet" href="/style.css">
</head>

<body>

<header>
	<div class="logo">⚊ FORMICARIUM_OS_v1.0</div>
	<div class="file">colony_01.txt</div>
	<div class="signal">◉</div>
</header>

<main>

	<div class="stats">
		<div>
			● ANTS:
			<strong>`+strconv.Itoa(c.NumAnts)+`/`+strconv.Itoa(c.NumAnts)+`</strong>
		</div>

		<div>
			♧ MOVES:
			<strong>`+strconv.Itoa(len(c.Turns))+`</strong>
		</div>
	</div>

	<section class="colony">

		<div class="graph">
			`+graph.String()+`
		</div>

		<div class="terminal">
			<div class="terminal-title">TERM_OUTPUT_LOG</div>
			<div>&gt; Waiting for simulation...</div>
		</div>

		<div class="controls">

			<button type="button">Reset</button>

			<button type="button">▶ Play/Pause</button>

			<button type="button">Step</button>

			<div class="speed">
				━●━━━━ 1.0x
			</div>

		</div>

	</section>

</main>

</body>
</html>`)
}
