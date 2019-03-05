package brawlstars

import "time"

// Misc represents some misc data such as shop/season reset times.
type Misc struct {
  TimeUntilSeasonEndInSeconds int `json:"timeUntilSeasonEndInSeconds"`
  TimeUntilSeasonEnd time.Time `json:"timeUntilSeasonEnd"`
  TimeUntilShopResetInSeconds int `json:"timeUntilShopResetInSeconds"`
  TimeUntilShopReset time.Time `json:"timeUntilShopReset"`
  ServerDateYear int `json:"serverDateYear"`
  ServerDateDayOfYear int `json:"serverDateDayOfYear"`
}
