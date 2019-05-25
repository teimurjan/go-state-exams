package utils

import "regexp"

func GetNamedMap(str string, exp *regexp.Regexp) map[string]string {
	match := exp.FindStringSubmatch(str)
	result := make(map[string]string)
	for i, name := range exp.SubexpNames() {

		if i != 0 && name != "" && len(match) > i && match[i] != "" {
			result[name] = match[i]
		}
	}

	return result
}
