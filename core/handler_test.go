package core

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/eleven26/awesome-go-stars/core/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	readmePath, err := filepath.Abs("./readme_test.md")
	assert.Nil(t, err)
	defer os.Remove(readmePath)

	// 返回 mock readme
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(readmeStr))
	}))

	puller := new(mocks.Puller)
	puller.On(
		"Pull",
		"https://api.github.com/repos/amit-davidson/awesome-golang-workshops",
	).Return(NewPullResult(http.StatusOK, []byte(`{"id":1296269,"stargazers_count":20}`)))

	handler := NewHandler(puller, readmePath, ts.URL)
	err = handler.Handle()
	assert.Nil(t, err)

	puller.AssertExpectations(t)

	// assert file content
	bs, err := os.ReadFile(readmePath)
	assert.Nil(t, err)
	assert.Equal(t, newReadmeStr, string(bs))
}
