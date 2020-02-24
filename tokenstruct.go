package mtgjson

// Token has information of single token card
// Based on MTGJSON v4.6.2
type Token struct {
  Artist                 string   `json:"artist"`
  BorderColor            string   `json:"borderColor"`
  ColorIdentity          []string `json:"colorIdentity"`
  ColorIndicator         []string `json:"colorIndicator,omitempty"`
  Colors                 []string `json:"colors"`
  IsOnlineOnly           bool     `json:"isOnlineOnly,omitempty"`
  Layout                 string   `json:"layout"`
  Loyalty                string   `json:"loyalty,omitempty"`
  Name                   string   `json:"name,omitempty"`
  Names                  []string `json:"names,omitempty"`
  Number                 string   `json:"number"`
  Power                  string   `json:"power"`
  ReverseRelated         []string `json:"reverseRelated,omitempty"`
	ScryfallID             string   `json:"scryfallId"`
	ScryfallIllustrationID string   `json:"scryfallIllustrationId,omitempty"`
	ScryfallOracleID       string   `json:"scryfallOracleId"`
  Side                   string   `json:"side,omitempty"`
	Subtypes               []string `json:"subtypes"`
	Supertypes             []string `json:"supertypes"`
	Text                   string   `json:"text"`
	Toughness              string   `json:"toughness"`
	Type                   string   `json:"type"`
	Types                  []string `json:"types"`
	UUID                   string   `json:"uuid"`
  Watermark              string   `json:"watermark,omitempty"`
}
