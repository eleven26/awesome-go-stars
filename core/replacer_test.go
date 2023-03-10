package core

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var newReadmeStr = `## Websites

- [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
- [Awesome Golang Workshops(stars: 20)](https://github.com/amit-davidson/awesome-golang-workshops) - A curated list of awesome golang workshops.

**[⬆ back to top](#contents)**
`

func TestReplace(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(readmeStr))

	mapping := map[string]int{
		"https://github.com/amit-davidson/awesome-golang-workshops": 20,
	}

	res, err := Replace(scanner, mapping)
	assert.Nil(t, err)
	assert.Equal(t, newReadmeStr, res)
}
