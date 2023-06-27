package cmd

import (
	"log"

	"Nie-Mand/sky-scraper/pkg"
	"Nie-Mand/sky-scraper/utils"

	"github.com/spf13/cobra"
)

const (
	// COURSES, LESSONS, LABS
	COURSES = "courses"
	LESSONS = "lessons"
	LABS = "labs"
)
	

func init() {
  RootCmd.AddCommand(listCmd)
  listCmd.Flags().BoolP("guru", "g", false, "Set the provider as a Cloud Guru")
}

var listCmd = &cobra.Command{
  Use:   "list [courses / lessons / lab] [course ID]",
  Short: "List courses, lessons, or labs from a Cloud Guru",
  Args: cobra.RangeArgs(1, 2),
  Run: func(cmd *cobra.Command, args []string) {
	
	provider, err := utils.GetProvider(cmd.Flags())

	if err != nil {
		log.Println(err)
		return
	}

	what := args[0]
	if what == COURSES && len(args) > 1 {
		log.Println("You don't need to specify a course ID to list courses.")
		return
	} else if (what == LESSONS || what == LABS) && len(args) == 1 {
		log.Println("You need to specify a course ID to list lessons or labs.")
		return
	} else {
		if what == COURSES {
			log.Println("Listing courses...")
		} else if what == LESSONS {
			courseId := args[1]
			_listLessons(courseId, provider)
		} else if what == LABS {
			log.Println("Listing labs...")
		} else {
			log.Printf("Unknown argument: %s", what)
		}
	}
  },
}

func _listLessons(courseId string, provider string) {
	if provider == "guru" {
		pkg.Guru.GetCourseContent(courseId)
	}
}