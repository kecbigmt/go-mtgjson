package mtgjson

// Set has information about the card set
// Based on MTGJSON v4.6.2
type Set struct {
  Code             string        `json:"code"`
  Name             string        `json:"name"`
  Type             string        `json:"type"`
  ReleaseDate      string        `json:"releaseDate"`
  TotalSetSize     int           `json:"totalSetSize"`
  BaseSetSize      int           `json:"baseSetSize"`
  Meta             SetMeta       `json:"meta"`
  Cards            []Card        `json:"cards,omitempty"`
  Tokens           []Token       `json:"tokens,omitempty"`
  ParentCode       string        `json:"parentCode,omitempty"`
  IsFoilOnly       bool          `json:"isFoilOnly,omitempty"`
  Block            string        `json:"block,omitempty"`
  BoosterV3        []interface{} `json:"boosterV3"`
  CodeV3           string        `json:"codeV3,omitempty"`
  IsOnlineOnly     bool          `json:"isOnlineOnly,omitempty"`
  IsPaperOnly      bool          `json:"isPaperOnly,omitempty"`
  IsForeignOnly    bool          `json:"isForeignOnly,omitempty"`
  IsPartialPreview bool          `json:"isPartialPreview,omitempty"`
  KeyruneCode      string        `json:"keyruneCode,omitempty"`
  MCMName          string        `json:"mcmName,omitempty"`
  MCMID            string        `json:"mcmId,omitempty"`
  MTGOCode         string `json:"mtgoCode,omitempty"`
  TCGPlayerGroupID string `json:"tcgplayerGroupId,omitempty"`
  Translations     map[string]string `json:"translations,omitempty"`
}

// SetMeta has mtgjson meta data
type SetMeta struct {
  Date string `json:"date"`
  PricesDate string `json:"pricesDate"`
  Version string `json:"version"`
}
