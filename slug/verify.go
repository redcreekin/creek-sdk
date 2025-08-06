package slug

import "regexp"

func IsValidSlug(slug string) bool {
	if len(slug) < 1 || len(slug) > 100 {
		return false
	}
	isValid, _ := regexp.MatchString(`^[a-z0-9]+(?:-[a-z0-9]+)*$`, slug)
	return isValid
}
