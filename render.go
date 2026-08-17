package main

import (
	"os"
	"text/template"
)

const pageTemplate = `<!DOCTYPE html>
<html class="dark" lang="en"><head>
<meta charset="utf-8"/>
<meta content="width=device-width, initial-scale=1.0" name="viewport"/>
<title>LEM-IN_OS_v1.0</title>
{{.Head}}
<style>{{.CSS}}</style>
</head>
<body class="bg-background text-on-background h-screen w-screen overflow-hidden flex flex-col font-body-md">
<div class="fixed inset-0 z-0 pointer-events-none overflow-hidden">
<img alt="Ant Farm Background" class="absolute inset-0 w-full h-full object-cover" src="{{.BG}}"/>
<div class="absolute inset-0 bg-black/30 backdrop-blur-[2px]"></div>
</div>

<header class="bg-surface/80 backdrop-blur-md text-primary border-b border-outline-variant flex justify-between items-center px-gutter w-full h-16 shrink-0 z-10 relative">
<div class="flex items-center gap-4">
<span class="material-symbols-outlined text-primary">account_tree</span>
<h1 class="font-label-caps text-label-caps text-primary tracking-widest">LEM-IN_OS_v1.0</h1>
</div>
<div class="flex items-center gap-2 px-3 py-1 bg-surface-container/80 backdrop-blur-sm rounded-DEFAULT border border-outline-variant font-data-mono text-data-mono text-on-surface-variant">
<span class="material-symbols-outlined text-secondary" style="font-size:16px;">description</span>
<span>{{.Filename}}</span>
</div>
<div class="flex items-center gap-4">
<span class="material-symbols-outlined text-on-surface-variant hover:text-primary transition-colors cursor-pointer">sensors</span>
</div>
</header>

<div class="flex-1 relative w-full h-full flex bg-grid-pattern overflow-hidden">
<main class="flex-1 relative m-canvas-margin border border-outline-variant rounded-lg overflow-hidden flex items-center justify-center bg-transparent">
<div class="relative w-full h-full max-w-4xl max-h-[800px]" id="graphContainer">
<svg class="absolute inset-0 w-full h-full pointer-events-none" id="linkSvg" style="z-index:1;"></svg>
<div class="absolute inset-0 pointer-events-none" id="nodeLayer" style="z-index:2;"></div>
<div class="absolute inset-0 pointer-events-none" id="antLayer" style="z-index:3;"></div>
</div>

<div class="absolute top-4 left-4 glass-panel rounded-lg p-3 min-w-[120px] flex flex-col gap-2">
<div class="flex items-center gap-2">
<div class="w-2 h-2 rounded-full bg-primary ant-glow"></div>
<span class="font-data-mono text-data-mono text-on-surface-variant">ANTS: <span class="text-primary" id="antCount">0/0</span></span>
</div>
<div class="h-px bg-outline-variant w-full"></div>
<div class="flex items-center gap-2">
<span class="material-symbols-outlined text-secondary" style="font-size:16px;">route</span>
<span class="font-data-mono text-data-mono text-on-surface-variant">MOVES: <span class="text-secondary" id="moveCount">0</span></span>
</div>
</div>

<div class="absolute bottom-4 right-4 glass-panel rounded-lg w-64 h-48 flex flex-col overflow-hidden">
<div class="bg-surface-container/80 border-b border-outline-variant px-3 py-1 flex items-center justify-between shrink-0">
<span class="font-label-caps text-label-caps text-on-surface-variant">TERM_OUTPUT_LOG</span>
<span class="material-symbols-outlined text-on-surface-variant" style="font-size:14px;">terminal</span>
</div>
<div class="p-3 font-data-mono text-xs flex-1 overflow-hidden leading-relaxed relative">
<div id="termLog" class="flex flex-col justify-end h-full gap-1"></div>
</div>
</div>
</main>
</div>

<nav class="bg-surface-container-high/80 backdrop-blur-md border-t border-outline-variant shadow-md fixed bottom-0 w-full flex justify-around items-center px-4 py-3 z-50 rounded-t-xl md:flex md:w-[600px] md:left-1/2 md:-translate-x-1/2 md:bottom-4 md:rounded-xl md:border">
<button id="resetBtn" class="flex flex-col items-center justify-center p-2 text-on-surface-variant hover:bg-surface-variant/50 transition-colors rounded-lg group">
<span class="material-symbols-outlined group-hover:text-primary mb-1">skip_previous</span>
<span class="font-data-mono text-[10px]">Reset</span>
</button>
<button id="playBtn" class="flex flex-col items-center justify-center p-2 bg-primary-container/20 text-primary rounded-xl hover:bg-primary-container/30 transition-colors active:scale-90 duration-150">
<span class="material-symbols-outlined mb-1" style="font-variation-settings:'FILL' 1;">play_arrow</span>
<span class="font-data-mono text-[10px]">Play/Pause</span>
</button>
<button id="stepBtn" class="flex flex-col items-center justify-center p-2 text-on-surface-variant hover:bg-surface-variant/50 transition-colors rounded-lg group">
<span class="material-symbols-outlined group-hover:text-primary mb-1">arrow_forward_ios</span>
<span class="font-data-mono text-[10px]">Step</span>
</button>
<div class="flex items-center gap-3 px-4 border-l border-outline-variant ml-2">
<span class="material-symbols-outlined text-on-surface-variant text-sm">speed</span>
<input id="speed" class="w-24 h-1 bg-surface-variant rounded-lg appearance-none cursor-pointer accent-primary" max="2000" min="150" type="range" value="700"/>
<span class="font-data-mono text-[10px] text-on-surface-variant">1.0x</span>
</div>
</nav>

<script>
const COLONY = {{.JSON}};
{{.JS}}
</script>
</body></html>`

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
		"Head":     visualizerHead,
		"BG":       bgImageDataURI(),
		"Filename": filename,
	})
}