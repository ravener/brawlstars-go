package brawlstars

import "errors"

// Player represents a player in game.
type Player struct {
  Client *Client `json:"-"`
  Name string `json:"name"`
  Tag string `json:"tag"`
  ID TagID `json:"id"`
  BrawlersUnlocked int `json:"brawlersUnlocked"`
  Victories int `json:"victories"`
  SoloShowdownVictories int `json:"soloShowdownVictories"`
  DuoShowdownVictories int `json:"duoShowdownVictories"`
  TotalExp int `json:"totalExp"`
  ExpFmt string `json:"expFmt"`
  ExpLevel int `json:"expLevel"`
  Trophies int `json:"trophies"`
  HighestTrophies int `json:"highestTrophies"`
  AvatarID int `json:"avatarId"`
  AvatarURL string `json:"avatarUrl"`
  BestTimeAsBigBrawler string `json:"bestTimeAsBigBrawler"`
  BestRoboRumbleTime string `json:"bestRoboRumbleTime"`
  HasSkins bool `json:"hasSkins"`
  Club *PartialClub `json:"club"`
  Brawlers []Brawler `json:"brawlers"`
}

// Fetch the full club the player is in.
// Returns an error if player not in a club.
func (p *Player) GetClub() (Club, error) {
  if p.Club == nil {
    return Club{}, errors.New("Player not in a club.")
  }
  return p.Client.GetClub(p.Club.Tag)
}
