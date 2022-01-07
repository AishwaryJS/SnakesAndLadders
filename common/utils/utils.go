package utils

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Average(array []int) float64 {
	avg := (float64(Sum(array))) / (float64(len(array)))
	return avg
}

// AsJSON returns the input object in a JSON format
func AsJSON(object interface{}) string {
	prettyJSON, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		log.Warningln("Error in printing as json. Returning empty string.")
		return ""
	}

	return string(prettyJSON)
}
