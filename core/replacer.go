package core

import (
	"bufio"
	"strings"
)

func replace(scanner *bufio.Scanner, stars map[string]int) (string, error) {
	var sb strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		links := parse(line)

		for _, link := range links {
			if link.IsRepoUrl() {
				if star, ok := stars[link.Url()]; ok {
					line = strings.Replace(line, link.OldMarkdownLink(), link.NewMarkdownLink(star), 1)
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
