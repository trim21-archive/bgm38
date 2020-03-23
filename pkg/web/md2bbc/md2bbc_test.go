package md2bbc_test

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"bgm38/pkg/web/md2bbc"
)

func TestRender(t *testing.T) {
	data, err := ioutil.ReadFile("../../../tests/fixtures/markdown.md")
	assert.Nil(t, err)
	rendered := md2bbc.Render(data)
	expected, err := ioutil.ReadFile("../../../tests/fixtures/bbcode.txt")
	assert.Nil(t, err)
	assert.Equal(t, string(expected), string(rendered), "md2bbc rendered is not eq to expected")
}