package main

import (
	"Nie-Mand/sky-scraper/pkg"
	"Nie-Mand/sky-scraper/utils"
)

func main() {
	token := utils.GetToken("TOKEN")
	pkg.Guru.Init(token)
}