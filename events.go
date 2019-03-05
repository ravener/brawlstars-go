package brawlstars

import "time"

// EventSlot represents a slot in events data.
type EventSlot struct {
  Slot int `json:"slot"`
  SlotName string `json:"slotName"`
  StartTimeInSeconds int `json:"startTimeInSeconds"`
  StartTime time.Time `json:"startTime"`
  EndTimeInSeconds int `json:"endTimeInSeconds"`
  EndTime time.Time `json:"endTime"`
  FreeKeys int `json:"freekeys"`
  MapID int `json:"mapId"`
  MapName string `json:"mapName"`
  GameMode string `json:"gameMode"`
  MapImageURL string `json:"mapImageUrl"`
  HasModifier bool `json:"hasModifier"`
  ModifierID int `json:"modifierId"`
  ModifierName string `json:"modifierName"`
}

// Events represents events data.
type Events struct {
  Upcoming []EventSlot `json:"upcoming"`
  Current []EventSlot `json:"current"`
}
