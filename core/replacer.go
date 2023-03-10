package core

import (
	"bufio"
	"strings"
)

func replace(scanner *bufio.Scanner, mapping map[string]string) (string, error) {
	var sb strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		links := parse(line)

		for _, link := range links {
			if link.IsRepoUrl() {
				if newTitleUrl, ok := mapping[link.OldTitleUrl()]; ok {
					line = strings.Replace(line, link.OldTitleUrl(), newTitleUrl, 1)
				}
			}
		}

		sb.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return sb.String(), nil
}
