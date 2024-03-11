package util

import "regexp"

func GetUrlName(url string) string {
	regex, _ := regexp.Compile(`^(http|https):\/\/(.+)$`)
	results := regex.FindStringSubmatch(url)
	return results[2]
}
