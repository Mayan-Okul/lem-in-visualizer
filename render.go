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
<link rel="preconnect" href="https://fonts.googleapis.com">
<link href="https://fonts.googleapis.com/css2?family=Hanken+Grotesk:wght@400;500;600;700&family=JetBrains+Mono:wght@400;500;700&display=swap" rel="stylesheet">
<style>{{.CSS}}</style>
</head>
<body>
<div id="bg" style="background-image:url('{{.BG}}')"></div>
<div id="bgScrim"></div>

<div id="topbar">
  <div class="sys">
    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#FFB347" stroke-width="2">
      <circle cx="5" cy="6" r="2.5"/><circle cx="19" cy="6" r="2.5"/>
      <circle cx="12" cy="12" r="2.5"/><circle cx="5" cy="18" r="2.5"/><circle cx="19" cy="18" r="2.5"/>
      <line x1="7" y1="7.5" x2="10" y2="10.5"/><line x1="17" y1="7.5" x2="14" y2="10.5"/>
      <line x1="10" y1="13.5" x2="7" y2="16.5"/><line x1="14" y1="13.5" x2="17" y2="16.5"/>
    </svg>
    <span>LEM-IN_OS_v1.0</span>
  </div>
  <div class="filename">📄 {{.Filename}}</div>
  <div>((•))</div>
</div>

<div id="hud">
  <div>● ANTS: <span class="value" id="antCount">0</span>/<span class="value" id="antTotal"></span></div>
  <div>⇄ MOVES: <span class="value" id="moveCount">0</span></div>
</div>

<canvas id="stage"></canvas>

<div id="term">
  <div class="termHeader">
    <div class="label">TERM_OUTPUT_LOG</div>
    <div class="rec"><div class="recDot"></div>REC</div>
  </div>
  <div id="termLog"></div>
</div>

<div id="controls">
  <button id="resetBtn">⏮ Reset</button>
  <button id="playBtn">▶ Play/Pause</button>
  <button id="stepBtn">› Step</button>
  <input id="speed" type="range" min="150" max="2000" value="700">
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