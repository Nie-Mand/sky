package cmd

import (
	"fmt"
	"log"

	"Nie-Mand/sky-scraper/pkg"
	"Nie-Mand/sky-scraper/utils"

	"github.com/spf13/cobra"
)
	

func init() {
  RootCmd.AddCommand(getCmd)
  getCmd.Flags().BoolP("guru", "g", false, "Set the provider as a Cloud Guru")
  getCmd.Flags().StringP("quality", "q", "720p", "Set the Quality of the video to download")
}

var getCmd = &cobra.Command{
  Use:   "get [course ID]",
  Short: "Download a course from a learning platform.",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {

	fmt.Println("Sky ~ Hello World", args)

	
	provider, err := utils.GetProvider(cmd.Flags())
	if err != nil {
		log.Println(err)
		return
	}

	quality, err := cmd.Flags().GetString("quality")
	
	if err != nil {
		log.Println(err)
		return
	}

	courseId := args[0]

	if provider == "guru" {
		pkg.Guru.GetCourseContent(courseId, quality)
		return
	}

	log.Printf("Unknown provider %s", provider)
  },
}