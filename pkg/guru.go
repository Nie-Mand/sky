package pkg

import (
	"Nie-Mand/sky-scraper/utils"
	"encoding/json"
	"fmt"
	"log"
)

type ACloudGuruScraper struct {
	http HttpClient
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


func (guru *ACloudGuruScraper) GetCourseContent(courseId string) {
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
		return
	} else {
		log.Printf("README.md does not exist, creating...")

		readme := MakeGuruReadMe(result)

		err = utils.Write("downloads/ACloudGuru/" + courseId + "/README.md", []byte(readme)) 
	
		if err != nil {
			log.Printf("There was an error while writing the README.md file: %s", err)
			return
		}
	}
}



func MakeGuruReadMe(result GuruGetCourseContentResponse) string {
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