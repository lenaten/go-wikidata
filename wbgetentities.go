package wikidata

import "encoding/json"

// WbGetEntities request output.
type WbGetEntitiesInput struct {
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

func (c *Client) NewWbGetEntitiesInput(ids string, languages string) *WbGetEntitiesInput {
	return &WbGetEntitiesInput{
		Ids: ids,
		Languages: languages,
	}
}

// WbGetEntities returns a WbGetEntitiesOutput.
func (c *Client) WbGetEntities(in *WbGetEntitiesInput) (out *WbGetEntitiesOutput, err error) {
	body, err := c.call("wbgetentities", map[string]string{
		"props":     "labels",
		"languages": in.Languages,
		"ids":       in.Ids,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
