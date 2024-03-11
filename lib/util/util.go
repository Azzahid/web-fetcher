package util

import (
	"os"
	"regexp"
)

func GetUrlName(url string) string {
	regex, _ := regexp.Compile(`^(http|https):\/\/(.+)$`)
	results := regex.FindStringSubmatch(url)
	return results[2]
}

func GetFilePath(url string) string {
	mountedVol := os.Getenv("DOCKER_MOUNTED_VOL")
	filename := GetUrlName(url)
	if mountedVol == "" {
		return filename + ".html"
	} else {
		return mountedVol + "/" + filename + ".html"
	}
}
