/*
Copyright © 2024 conneroisu <conneroisu@outlook.com>
*/
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// "fmt"
// "os/exec"

//	func main() {
//	        cmd.Execute()
//	        // x, y, w, h := 154, 389, 179, 374
//	        // geometry := fmt.Sprintf("%d,%d %dx%d", x, y, w, h)
//	        // cmd := exec.Command("/usr/bin/grim", "-g", geometry, "-t", "png", "-l", "6", "/tmp/screenshot.png")
//	        //
//	        // if output, err := cmd.CombinedOutput(); err != nil {
//	        //         fmt.Printf("Error executing grim command: %s\n", err)
//	        //         fmt.Printf("Command output: %s\n", output)
//	        // } else {
//	        //         fmt.Println("Screenshot taken successfully!")
//	        // }
//	}
func main() {
	// Define the options for the screenshot
	opts := WayshotOptions{
		SelectRegion:   true,
		OutputFilename: "screenshot.png",
	}

	// Create and execute the wayshot command
	cmd := NewWayshotCommand(opts)
	output, err := ExecuteCommand(cmd)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Command Output:", output)
}

// WayshotOptions represents the options available for the wayshot command
type WayshotOptions struct {
	SelectRegion   bool
	SelectWindow   bool
	DelaySeconds   int
	Clipboard      bool
	IncludeCursor  bool
	OutputToStdout bool
	OutputFilename string
}

// NewWayshotCommand prepares and returns a wayshot command based on provided options
func NewWayshotCommand(opts WayshotOptions) *exec.Cmd {
	var args []string

	if opts.SelectRegion {
		args = append(args, "-s")
	}
	if opts.SelectWindow {
		args = append(args, "-w")
	}
	if opts.DelaySeconds > 0 {
		args = append(args, fmt.Sprintf("-d %d", opts.DelaySeconds))
	}
	if opts.Clipboard {
		args = append(args, "-c")
	}
	if opts.IncludeCursor {
		args = append(args, "-x")
	}
	if opts.OutputToStdout {
		args = append(args, "-")
	} else if opts.OutputFilename != "" {
		args = append(args, opts.OutputFilename)
	}

	return exec.Command("wayshot", args...)
}

// ExecuteCommand runs the wayshot command and returns the output or an error
func ExecuteCommand(cmd *exec.Cmd) (string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("wayshot failed: %s", stderr.String())
	}
	return out.String(), nil
}
