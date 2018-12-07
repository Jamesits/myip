package main

import "regexp"

// get rid of extra characters
func FilterIP(s string) string {
	reg, err := regexp.Compile("[^0-9a-f.:]+")
	if err != nil {
		panic(err)
	}

	return reg.ReplaceAllString(s, "")
}
