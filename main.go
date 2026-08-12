package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	colony, err := ParseColony(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "visualizer error:", err)
		os.Exit(1)
	}

	j, err := colony.ToJSON()
	if err != nil {
		fmt.Fprintln(os.Stderr, "visualizer error:", err)
		os.Exit(1)
	}

	out := "colony_view.html"
	if err := WriteVisualizerHTML(out, j, "stdin"); err != nil {
		fmt.Fprintln(os.Stderr, "visualizer error:", err)
		os.Exit(1)
	}

	openBrowser(out)
}

func openBrowser(path string) {
	var cmd string
	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "start"
	default:
		cmd = "xdg-open"
	}
	exec.Command(cmd, path).Start()
}