package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var readmeStr = `## Websites

- [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
- [Awesome Golang Workshops](https://github.com/amit-davidson/awesome-golang-workshops) - A curated list of awesome golang workshops.

**[⬆ back to top](#contents)**`

func TestParse(t *testing.T) {
	links := parse(readmeStr)
	assert.Len(t, links, 2)

	assert.Equal(t, "Awesome Go @LibHunt", links[0].Title())
	assert.Equal(t, "https://go.libhunt.com", links[0].Url())
	assert.Equal(t, "- [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.", links[0].Raw())

	assert.Equal(t, "Awesome Golang Workshops", links[1].Title())
	assert.Equal(t, "https://github.com/amit-davidson/awesome-golang-workshops", links[1].Url())
	assert.Equal(t, "- [Awesome Golang Workshops](https://github.com/amit-davidson/awesome-golang-workshops) - A curated list of awesome golang workshops.", links[1].Raw())
}
