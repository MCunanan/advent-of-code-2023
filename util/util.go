package util

import "regexp"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func RemoveAlphas(str string) string {
	var alphas = regexp.MustCompile(`[a-zA-Z]+`)
	return alphas.ReplaceAllString(str, "")
}
