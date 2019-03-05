package brawlstars

// TagID represents the unique ID of a player/club Tag.
// It is what a tag is made of, in other words the ID can be encoded to a tag
// or the tag can be decoded to an ID.
type TagID struct {
  High int `json:"high"`
  Low int `json:"low"`
}
