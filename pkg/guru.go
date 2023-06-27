package pkg

import (
	"Nie-Mand/sky-scraper/utils"
	"encoding/json"
	"fmt"
	"log"

	"github.com/schollz/progressbar/v3"
)

type ACloudGuruScraper struct {
	http utils.HttpClient
}


var Guru = ACloudGuruScraper{}

func (guru *ACloudGuruScraper) Init() {
	token := utils.GetToken("TOKEN")
	if token == "" {
		log.Println("You must provide a valid token")
		return
	}

	apiUrl := "https://prod-api.acloud.guru/bff/graphql"
	
	guru.http.Init(apiUrl, map[string]string{
		"Authorization": "Bearer " + token,
	})
	log.Printf("Initiated CloudGuru Scraper, API url ~ %s", apiUrl)
}


func (guru *ACloudGuruScraper) GetCourseContent(courseId string, quality string) {
	guru.Init()

	err := utils.Mkdir("downloads/ACloudGuru/" + courseId)

	if err != nil {
		log.Printf("There was an error while creating the folder: %s", err)
	}

	gqlQuery := GuruGetCourseContent	
	
	variables := map[string]string{
		"courseId": courseId,
	}

	body := map[string]interface{}{
		"query":     gqlQuery,
		"variables": variables,
	}

	response, err := guru.http.Post("", body)

	if err != nil {
		log.Panicf("There was an error while fetching the course content: %s", err)
	}

	var result GuruGetCourseContentResponse

	err = json.Unmarshal([]byte(response), &result)
	
	if err != nil {
		log.Printf("There was an error while parsing the response: %s", err)
		return
	}

	if result.Message != "" {
		log.Printf("Message: %s", result.Message)
		return
	}

	readmeFile := "downloads/ACloudGuru/" + courseId + "/README.md"
	log.Println("Successfully fetched course content")
	log.Println("Saving course content to README.md")

	if utils.Exist(readmeFile) {
		log.Printf("README.md already exists, skipping...")
	} else {
		log.Printf("README.md does not exist, creating...")

		readme := makeGuruReadMe(result)

		err = utils.Write("downloads/ACloudGuru/" + courseId + "/README.md", []byte(readme)) 
	
		if err != nil {
			log.Printf("There was an error while writing the README.md file: %s", err)
			return
		}
	}

	videosLength := getHowManyLessons(result.Data.UserCourseOverview.CourseOverview.Sections)
	log.Printf("Downloading course content ~ %d sections, %d videos", len(result.Data.UserCourseOverview.CourseOverview.Sections), videosLength)

	bar := progressbar.DefaultBytes(int64(videosLength))

	for _, section := range result.Data.UserCourseOverview.CourseOverview.Sections {
		// log.Printf("Downloading section: %s", section.Title)
		sectionFolder := utils.MakeSectionFolder(section.Sequence + 1, section.Title)

		err = utils.Mkdir("downloads/ACloudGuru/" + courseId + "/" + sectionFolder)

		if err != nil {
			log.Printf("There was an error while creating the folder: %s", err)
		}


		for _, lesson := range section.Components {
			bar.Describe(fmt.Sprintf("Downloading lesson: %s", lesson.Title))
			
			videoFileName := utils.MakeVideoName(lesson.Sequence + 1, lesson.Title)
			if utils.Exist("downloads/ACloudGuru/" + courseId + "/" + sectionFolder + "/" + videoFileName) {
				// log.Printf("Video of lesson %s already exists, skipping...", lesson.Title)
				bar.Add(1)
				continue
			}


			url, err := getUrl(lesson.Content.Videosources, quality)

			if err != nil {
				log.Printf("There was an error while getting the url for the lesson %s: %s", lesson.Title, err)
				continue
			}
			
			video, err := guru.http.Download(url)
			if err != nil {
				log.Printf("There was an error while downloading the video: %s", err)
				return
				// continue
			}

			err = utils.Write("downloads/ACloudGuru/" + courseId + "/" + sectionFolder + "/" + videoFileName, video)

			if err != nil {
				log.Printf("There was an error while writing the video file: %s", err)
				continue
			}

			bar.Add(1)
		}
	}

}





func getUrl(sources []GuruVideoSourcesType, quality string) (string, error) {
	for _, source := range sources {
		if source.Quality == quality {
			return source.SignedUrl, nil
		}
	}

	return "", fmt.Errorf("Could not find the url for the quality %s", quality)
}

func makeGuruReadMe(result GuruGetCourseContentResponse) string {
	readme := ""
	readme += "# " + result.Data.UserCourseOverview.CourseOverview.Title + "\n\n"

	readme += "## Description\n"
	readme += result.Data.UserCourseOverview.CourseOverview.Description + "\n\n"

	readme += "## Sections\n\n"

	for _, section := range result.Data.UserCourseOverview.CourseOverview.Sections {
		seq := fmt.Sprintf("%d", section.Sequence + 1)
		readme += seq + ". " + section.Title + "\n"

		for _, lesson := range section.Components {
			readme += "   - " + lesson.Title + "\n"
		}
	}

	return readme
}

func getHowManyLessons(sections []GuruSectionType) int {
	lessons := 0
	for _, section := range sections {
		lessons += len(section.Components)
	}

	return lessons
}