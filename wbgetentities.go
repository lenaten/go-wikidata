package wikidata

import "encoding/json"

// WbGetEntities request output.
type WbGetEntities struct {
	Ids       string `url:"ids"`
	Languages string `url:"languages"`
}

type Label struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}
type Entity struct {
	ID     string           `json:"id"`
	Type   string           `json:"type"`
	Labels map[string]Label `json:"labels"`
}

// WbGetEntitiesOutput request output.
type WbGetEntitiesOutput struct {
	Entities map[string]Entity `json:"entities"`
	Success  int               `json:"success"`
}

func (c *Client) NewWbGetEntitiesInput(ids string) *WbGetEntities {
	return &WbGetEntities{
		Ids: ids,
	}
}

// WbGetEntities returns a WbGetEntitiesOutput.
func (c *Client) WbGetEntities(in *WbGetEntities) (out *WbGetEntitiesOutput, err error) {
	body, err := c.call("wbgetentities", map[string]string{
		"props":     "labels",
		"languages": "en",
		"ids":       in.Ids,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
