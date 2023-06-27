package pkg

import "log"

type CloudGuruScraper struct {
	http HttpClient
}


var Guru = CloudGuruScraper{}

func (guru * CloudGuruScraper) Init(token string) {
	if token == "" {
		log.Panic("You must provide a valid token")
	}
	apiUrl := "https://prod-api.acloud.guru/bff/graphql"
	guru.http.Init(apiUrl, map[string]string{
		"Authorization": "Bearer " + token,
	})
	log.Printf("Initiated CloudGuru Scraper, API url ~ %s", apiUrl)

}