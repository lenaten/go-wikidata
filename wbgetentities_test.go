package wikidata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWbGetEntities(t *testing.T) {
	c := client()
	id := "Q95"
	lang := "en"
	in := c.NewWbGetEntitiesInput(id, lang)
	out, err := c.WbGetEntities(in)
	assert.Nil(t, err)
	assert.Equal(t, 1, out.Success)
}
