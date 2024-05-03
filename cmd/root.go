/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	grim "github.com/conneroisu/grimgo/pkg/grim"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grimgo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() {
	opts := grim.DefaultGrimOptions()

	// Enable slurp to let the user select a region
	opts.UseSlurp = true
	imageData, err := grim.CaptureImage(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error capturing image: %v\n", err)
		return
	}

	// For demonstration purposes, let's just print the size of the image data
	fmt.Printf("Captured image size: %d bytes\n", len(imageData))

	// Optionally, save the image data to a file
	err = os.WriteFile("captured_image.png", imageData, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save image: %v\n", err)
		return
	}
	fmt.Println("Image saved successfully.")
}
