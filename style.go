package main

import (
	"fmt"
	"net/http"
)

func serveStyle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")

	fmt.Fprint(w, `
* {
	box-sizing: border-box;
}

body {
	margin: 0;
	background: #05090d;
	color: #d8e0e5;
	font-family: monospace;
}

header {
	height: 70px;
	display: flex;
	align-items: center;
	justify-content: space-between;
	padding: 0 32px;
	border-bottom: 1px solid #26343d;
	background: #070d12;
}

.logo {
	color: #f5a623;
	font-weight: bold;
	letter-spacing: 2px;
}

.file {
	border: 1px solid #45515a;
	padding: 10px 20px;
	color: #d8e0e5;
}

.signal {
	color: #f5a623;
}

main {
	padding: 24px;
}

.stats {
	position: absolute;
	z-index: 5;
	margin: 20px;
	background: rgba(10, 17, 22, 0.9);
	border: 1px solid #38454d;
	border-radius: 6px;
	overflow: hidden;
}

.stats div {
	padding: 14px 18px;
	border-bottom: 1px solid #303c43;
}

.stats div:last-child {
	border-bottom: none;
}

.stats strong {
	color: #f5a623;
}

.colony {
	position: relative;
	height: calc(100vh - 120px);
	min-height: 600px;
	border: 1px solid #34434c;
	overflow: hidden;
	background:
		linear-gradient(rgba(0, 180, 220, 0.04) 1px, transparent 1px),
		linear-gradient(90deg, rgba(0, 180, 220, 0.04) 1px, transparent 1px),
		#081017;
	background-size: 40px 40px;
}

.graph {
	position: absolute;
	inset: 0;
	background-image: url('/assets/colony.png');
	background-size: cover;
	background-position: center;
	opacity: 0.65;
}

.terminal {
	position: absolute;
	right: 20px;
	bottom: 20px;
	width: 300px;
	background: rgba(8, 13, 17, 0.94);
	border: 1px solid #4b5961;
	padding: 16px;
	font-size: 13px;
}

.terminal-title {
	color: #f5a623;
	margin-bottom: 12px;
}

.controls {
	position: absolute;
	bottom: 20px;
	left: 50%;
	transform: translateX(-50%);
	display: flex;
	align-items: center;
	gap: 30px;
	padding: 16px 28px;
	background: rgba(15, 19, 22, 0.95);
	border: 1px solid #4b5961;
	border-radius: 8px;
}

button {
	background: transparent;
	border: none;
	color: #d8e0e5;
	cursor: pointer;
	font-family: monospace;
	padding: 12px 18px;
}

button:nth-child(2) {
	background: #f5a623;
	color: #111;
	border-radius: 6px;
}

.speed {
	color: #f5a623;
}
`)
}
