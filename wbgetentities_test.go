package wikidata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWbGetEntities(t *testing.T) {
	c := client()
	id := "Q95"
	in := c.NewWbGetEntitiesInput(id)
	out, err := c.WbGetEntities(in)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(out.Entities))
	entity := out.Entities[id]
	assert.NotEqual(t, "Google", entity.Labels["en"])
}
