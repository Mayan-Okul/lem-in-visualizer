package main

import (
	"os"
	"text/template"
)

const pageTemplate = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Lem-in Ant Colony Visualizer</title>
<style>{{.CSS}}</style>
</head>
<body>
<div id="bg" style="background-image:url('{{.BG}}')"></div>
<div id="bgOverlay"></div>

<div id="topbar">
  <div>⌗⌗ FORMICARIUM_OS_v1.0</div>
  <div class="filename">📄 {{.Filename}}</div>
  <div>((•))</div>
</div>

<div id="hud">
  <div>● ANTS: <span class="amber" id="antCount"></span>/<span id="antTotal"></span></div>
  <div>⇄ MOVES: <span class="amber" id="moveCount">0</span></div>
</div>

<canvas id="stage"></canvas>

<div id="term">
  <div class="label">TERM_OUTPUT_LOG</div>
  <div id="termLog"></div>
</div>

<div id="controls">
  <button id="resetBtn">⏮ Reset</button>
  <button id="playBtn">▶ Play/Pause</button>
  <button id="stepBtn">› Step</button>
  <input id="speed" type="range" min="150" max="2000" value="800">
</div>

<script>
const COLONY = {{.JSON}};
{{.JS}}
</script>
</body>
</html>`

func WriteVisualizerHTML(path string, colonyJSON string, filename string) error {
	tmpl := template.Must(template.New("page").Parse(pageTemplate))
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, map[string]string{
		"JSON":     colonyJSON,
		"CSS":      visualizerCSS,
		"JS":       visualizerJS,
		"BG":       bgImageDataURI(),
		"Filename": filename,
	})
}