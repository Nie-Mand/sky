package utils

import (
	"fmt"
	"os"
	"strings"
)

func Mkdir(path string) error {
	folders := strings.Split(path, "/")
	folder := "."
	for i := 0; i < len(folders); i++ {
		if folders[i] == "" {
			continue
		}
		folder = folder + "/" + folders[i]
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			if err != nil {
				err = os.Mkdir(folder, 0755)
				if err != nil {
					return err
				}
			}
		}
	}

	// log.Printf("Created folder: %s", path)
	return nil
}


func Write(path string, content []byte) (error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err  = file.Write(content)
	return err
}

func Exist(path string) (bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func MakeVideoName(section int, title string) string {
	sectionString := fmt.Sprintf("%02d", section)
	return sectionString + " - " + title + ".mp4"
}

func MakeSectionFolder(section int, title string) string {
	sectionString := fmt.Sprintf("%02d", section)
	return sectionString + " - " + title
}