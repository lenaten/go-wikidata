package wikidata

func client() *Client {
	return New(NewConfig())
}
