package mtgjson

// Card has information of single card
// Based on MTGJSON v4.6.2
type Card struct {
	Artist                string   `json:"artist"`
	BorderColor           string   `json:"borderColor"`
	ColorIdentity         []string `json:"colorIdentity"`
	ColorIndicator        []string `json:"colorIndicator,omitempty"`
	Colors                []string `json:"colors"`
	ConvertedManaCost     float32  `json:"convertedManaCost"`
	Count                 int      `json:"count,omitempty"`
	DuelDeck              string   `json:"duelDeck,omitempty"`
	EdhrecRank            int      `json:"edhrecRank"`
	FaceConvertedManaCost float32  `json:"faceConvertedManaCost,omitempty"`
	FlavorText            string   `json:"flavorText"`
	ForeignData           []struct {
		FlavorText   string `json:"flavorText"`
		Language     string `json:"language"`
		MultiverseID int    `json:"multiverseId"`
		Name         string `json:"name"`
		Text         string `json:"text"`
		Type         string `json:"type"`
	} `json:"foreignData"`
	FrameVersion     string `json:"frameVersion"`
	Hand             string `json:"hand,omitempty"`
	HasFoil          bool   `json:"hasFoil,omitempty"`
	HasNonFoil       bool   `json:"hasNonFoil,omitempty"`
	IsAlternative    bool   `json:"isAlternative,omitempty"`
	IsArena          bool   `json:"isArena,omitempty"`
	IsFullArt        bool   `json:"isFullArt,omitempty"`
	IsMTGO           bool   `json:"isMtgo,omitempty"`
	IsOnlineOnly     bool   `json:"isOnlineOnly,omitempty"`
	IsOversized      bool   `json:"isOversized,omitempty"`
	IsPaper          bool   `json:"isPaper,omitempty"`
	IsPromo          bool   `json:"isPromo,omitempty"`
	IsReprint        bool   `json:"isReprint,omitempty"`
	IsReserved       bool   `json:"isReserved,omitempty"`
	IsStarter        bool   `json:"isStarter,omitempty"`
	IsStorySpotlight bool   `json:"isStorySpotlight,omitempty"`
	IsTextless       bool   `json:"isTextless,omitempty"`
	IsTimeshifted    bool   `json:"isTimeshifted,omitempty"`
	Layout           string `json:"layout"`
	LeadershipSkills struct {
		Brawl       bool `json:"brawl"`
		Commander   bool `json:"commander"`
		Oathbreaker bool `json:"oathbreaker"`
	} `json:"leadershipSkills,omitempty"`
	Legalities   Legalities `json:"legalities"`
	Life         string     `json:"life,omitempty"`
	Loyalty      string     `json:"loyalty,omitempty"`
	ManaCost     string     `json:"manaCost,omitempty"`
	MCMID        int        `json:"mcmId"`
	MCMMetaID    int        `json:"mcmMetaId"`
	MTGArenaID   int        `json:"mtgArenaId,omitempty"`
	MTGOFoilID   int        `json:"mtgoFoilId,omitempty"`
	MTGOID       int        `json:"mtgoId,omitempty"`
	MTGStocksID  int        `json:"mtgstocksId,omitempty"`
	MultiverseID int        `json:"multiverseId"`
	Name         string     `json:"name,omitempty"`
	Names        []string   `json:"names,omitempty"`
	Number       string     `json:"number"`
	OriginalText string     `json:"originalText"`
	OriginalType string     `json:"originalType"`
	OtherFaceIDs []string   `json:"otherFaceIds,omitempty"`
	Power        string     `json:"power"`
	Prices       struct {
		MTGO      map[string]float32 `json:"mtgo"`
		MTGOFoil  map[string]float32 `json:"mtgoFoil"`
		Paper     map[string]float32 `json:"paper"`
		PaperFoil map[string]float32 `json:"paperFoil"`
	} `json:"prices"`
	Printings    []string `json:"printings"`
	PurchaseURLs struct {
		Cardmarket string `json:"cardmarket"`
		TCGPlayer  string `json:"tcgplayer"`
		MTGStocks  string `json:"mtgstocks"`
	} `json:"purchaseUrls"`
	Rarity         string   `json:"rarity"`
	ReverseRelated []string `json:"reverseRelated,omitempty"`
	Rulings        []struct {
		Date string `json:"date"`
		Text string `json:"text"`
	} `json:"rulings"`
	ScryfallID             string   `json:"scryfallId"`
	ScryfallIllustrationID string   `json:"scryfallIllustrationId,omitempty"`
	ScryfallOracleID       string   `json:"scryfallOracleId"`
	Side                   string   `json:"side,omitempty"`
	Subtypes               []string `json:"subtypes"`
	Supertypes             []string `json:"supertypes"`
	TCGPlayerProductID     int      `json:"tcgplayerProductId"`
	Text                   string   `json:"text"`
	Toughness              string   `json:"toughness"`
	Type                   string   `json:"type"`
	Types                  []string `json:"types"`
	UUID                   string   `json:"uuid"`
	Variations             []string `json:"variations,omitempty"`
	Watermark              string   `json:"watermark,omitempty"`
}

type Legalities struct {
	Brawl     string `json:"brawl,omitempty"`
	Commander string `json:"commander,omitempty"`
	Duel      string `json:"duel,omitempty"`
	Future    string `json:"future,omitempty"`
	Frontier  string `json:"frontier,omitempty"`
	Historic  string `json:"historic,omitempty"`
	Legacy    string `json:"legacy,omitempty"`
	Modern    string `json:"modern,omitempty"`
	Pauper    string `json:"pauper,omitempty"`
	Penny     string `json:"penny,omitempty"`
	Pioneer   string `json:"pioneer,omitempty"`
	Standard  string `json:"standard,omitempty"`
	Vintage   string `json:"vintage,omitempty"`
}
