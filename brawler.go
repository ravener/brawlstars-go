package brawlstars

// Brawler represents a brawler stats.
type Brawler struct {
  Name string `json:"name"`
  HasSkin bool `json:"hasSkin"`
  Skin string `json:"skin"`
  Trophies int `json:"trophies"`
  HighestTrophies int `json:"highestTrophies"`
  Power int `json:"power"`
  Rank int `json:"rank"`
}

// Check whether this brawler has star power.
func (b *Brawler) HasStarPower() bool {
  return b.Power == 10
}

