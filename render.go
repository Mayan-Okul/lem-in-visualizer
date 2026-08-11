package main

import (
	"fmt"
	"net/http"
)

func renderPage(w http.ResponseWriter, c *Colony) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

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
				<strong>`+fmt.Sprint(c.NumAnts)+`/`+fmt.Sprint(c.NumAnts)+`</strong>
			</div>

			<div>
				♧ MOVES:
				<strong>`+fmt.Sprint(len(c.Turns))+`</strong>
			</div>
		</div>

		<section class="colony">

			<div class="graph">
				<!-- Colony graph will be rendered here -->
			</div>

			<div class="terminal">
				<div class="terminal-title">
					TERM_OUTPUT_LOG
				</div>

				<div>
					&gt; Waiting for simulation...
				</div>
			</div>

			<div class="controls">

				<button type="button">
					Reset
				</button>

				<button type="button">
					▶ Play/Pause
				</button>

				<button type="button">
					Step
				</button>

				<div class="speed">
					━●━━━━ 1.0x
				</div>

			</div>

		</section>

	</main>

</body>
</html>`)
}
