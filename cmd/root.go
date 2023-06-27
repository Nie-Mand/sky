package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "sky",
	Short: "Sky is a CLI tool for downloading courses from various learning platforms.",
	Run: func(cmd *cobra.Command, args []string) {
	  fmt.Println("Sky ~ Hello World")	
	},
  }
  
  func Execute() {
	if err := RootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }