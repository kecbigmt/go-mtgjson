package mtgjson

import (
  "fmt"
  "encoding/json"
)

// Set has information about the card set
type Set struct {
  Code string `json:"code"`
  Name string `json:"name"`
  Type string `json:"type"`
  ReleaseDate string `json:"releaseDate"`
  TotalSetSize int `json:"totalSetSize"`
  BaseSetSize int `json:"baseSetSize"`
  IsOnlineOnly bool `json:"isOnlineOnly"`
  IsPaperOnly bool `json:"isPaperOnly"`
  Meta SetMeta `json:"meta"`
}

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
