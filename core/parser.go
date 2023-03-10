package core

import (
	"regexp"

	"github.com/eleven26/awesome-go-stars/contract"
)

func parse(content string) []contract.Link {
	re := regexp.MustCompile(`- (?im)\[([^]]*)]\(([^)]*)\) - .*`)

	var result []contract.Link

	for _, match := range re.FindAllStringSubmatch(content, -1) {
		result = append(result, NewLink(match[1], match[2], match[0]))
	}

	return result
}
