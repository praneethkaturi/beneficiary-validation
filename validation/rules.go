package validation

import "regexp"

func validateLength(val string, minLength int, MaxLength int) bool {
	return len(val) >= minLength && len(val) <= MaxLength
}

func validateRegex(val string, regex string) bool {
	matched, _ := regexp.MatchString(regex, val)
	return matched
}

func validateEnum(val string, enumValues []string) bool {
	for _, v := range enumValues {
		if val == v {
			return true
		}
	}
	return false
}
