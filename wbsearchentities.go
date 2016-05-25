package wikidata

import "encoding/json"

// WbSearchEntities request output.
type WbSearchEntitiesInput struct {
	Search   string `url:"search"`
	Language string `url:"language"`
}

type Search struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Label string `json:"label"`
}

// WbSearchEntitiesOutput request output.
type WbSearchEntitiesOutput struct {
	Search  []Search `json:"search"`
	Success int      `json:"success"`
}

func (c *Client) NewWbSearchEntitiesInput(search string, language string) *WbSearchEntitiesInput {
	return &WbSearchEntitiesInput{
		Search:   search,
		Language: language,
	}
}

// WbSearchEntities returns a WbSearchEntitiesOutput.
func (c *Client) WbSearchEntities(in *WbSearchEntitiesInput) (out *WbSearchEntitiesOutput, err error) {
	body, err := c.call("wbsearchentities", map[string]string{
		"search":   in.Search,
		"language": in.Language,
	})
	if err != nil {
		return
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&out)
	return
}
