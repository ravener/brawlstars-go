package brawlstars

// PartialClub represents a partial club that is not full it is found in a Player
// which is enough to get basic club data but you need to fetch the full club if
// you need data like all the members.
type PartialClub struct {
  Client *Client `json:"-"`
  ID TagID `json:"id"`
  Tag string `json:"tag"`
  Name string `json:"name"`
  Role string `json:"role"`
  BadgeID int `json:"badgeId"`
  BadgeURL string `json:"badgeUrl"`
  Members int `json:"members"`
  Trophies int `json:"trophies"`
  RequiredTrophies int `json:"requiredTrophies"`
  OnlineMembers int `json:"onlineMembers"`
}

func (c *PartialClub) FetchFull() (Club, error) {
  return c.Client.GetClub(c.Tag)
}

// ClubMember represents a member in a club, it's kind of a partial player.
type ClubMember struct {
  Client *Client `json:"-"`
  ID TagID `json:"id"`
  Tag string `json:"tag"`
  Role string `json:"role"`
  ExpLevel int `json:"expLevel"`
  Trophies int `json:"trophies"`
  AvatarID int `json:"avatarId"`
  AvatarURL string `json:"avatarUrl"`
}

// Fetches the full player from a member.
func (m *ClubMember) FetchFull() (Player, error) {
  return m.Client.GetPlayer(m.Tag)
}

// Club represents a club in game.
// Clubs can also be partial such as the one in Players, see PartialClub
type Club struct {
  Client *Client `json:"-"`
  ID TagID `json:"id"`
  Tag string `json:"tag"`
  Name string `json:"name"`
  Role string `json:"role"`
  BadgeID int `json:"badgeId"`
  BadgeURL string `json:"badgeUrl"`
  MembersCount int `json:"membersCount"`
  Trophies int `json:"trophies"`
  RequiredTrophies int `json:"requiredTrophies"`
  OnlineMembers int `json:"onlineMembers"`
  Region string `json:"region"`
  Status string `json:"status"`
  Description string `json:"description"`
  Members []ClubMember `json:"members"`
}


