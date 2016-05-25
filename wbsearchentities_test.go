package wikidata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWbSearchEntities(t *testing.T) {
	c := client()
	search := "ביבי"
	lang := "he"
	in := c.NewWbSearchEntitiesInput(search, lang)
	out, err := c.WbSearchEntities(in)
	assert.Nil(t, err)
	assert.Equal(t, 1, out.Success)
}
