package mtgjson

import (
  "fmt"
  "encoding/json"
)

// SetMeta has mtgjson meta data
type SetMeta struct {
  Date string `json:"date"`
  PricesDate string `json:"pricesDate"`
  Version string `json:"version"`
}

const (
  path = "/files/SetList.json"
)

// GetSetList return all card sets as slice of Set
func (c *Client) GetSetList() ([]Set, error) {
  r, err := c.API.doRequest(path)
  if err != nil {
    return nil, fmt.Errorf("c.api.doRequest: %s", err)
  }

  parsedSetList := []Set{}
  if err = json.NewDecoder(r).Decode(&parsedSetList); err != nil {
    return nil, fmt.Errorf("json.NewDecoder(r).Decode: %s", err)
  }
  return parsedSetList, nil
}
