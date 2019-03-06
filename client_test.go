// BrawlStars Go tests.
package brawlstars

import (
  "testing"
  "os"
  "encoding/json"
  "fmt"
)

var token string
var client *Client

func init() {
  token = os.Getenv("TOKEN")
  if token == "" {
    panic("No token found in the variable TOKEN, tests can't be ran.")
  }
  client = NewClient(token)
}

func TestGetPlayer(t *testing.T) {
  p, err := client.GetPlayer("GGJVJLU2");
  if err != nil {
    t.Error(err)
    return
  }
  b, err := json.MarshalIndent(p, "", "  ")
  if err != nil {
    t.Error(err)
    return
  }
  fmt.Println(string(b))
  fmt.Println("")
}


func TestGetClub(t *testing.T) {
  c, err := client.GetClub("QCGVPV8")
  if err != nil {
    t.Error(err)
    return
  }
  b, err := json.MarshalIndent(c, "", "  ")
  if err != nil {
    t.Error(err)
    return
  }
  fmt.Println(string(b))
  fmt.Println("")
}

func TestMisc(t *testing.T) {
  m, err := client.GetMisc()
  if err != nil {
    t.Error(err)
    return
  }
  b, err := json.MarshalIndent(m, "", "  ")
  if err != nil {
    t.Error(err)
    return
  }
  fmt.Println(string(b))
  fmt.Println("")
}

func TestUpcomingEvents(t *testing.T) {
  events, err := client.GetUpcomingEvents()
  if err != nil {
    t.Error(err)
    return
  }
  b, err := json.MarshalIndent(events, "", "  ")
  if err != nil {
    t.Error(err)
    return
  }
  fmt.Println(string(b))
  fmt.Println("")
}

func TestCurrentEvents(t *testing.T) {
  events, err := client.GetCurrentEvents()
  if err != nil {
    t.Error(err)
    return
  }
  b, err := json.MarshalIndent(events, "", "  ")
  if err != nil {
    t.Error(err)
    return
  }
  fmt.Println(string(b))
  fmt.Println("")
}

func TestValidTag(t *testing.T) {
  var res bool
  res = ValidTag("GGJVJLU2")
  if res != true {
    t.Error("ValidTag(GGJVJLU2) returned false but the tag is valid.")
  }
  res = ValidTag("#GGJVJLU2")
  if res != true {
    t.Error("ValidTag(#GGJVJLU2) returned false but the tag is valid.")
  }
  var broke string = "sjsnisnsisnwsjaopqoamzmznsnso"
  res = ValidTag(broke)
  if res != false {
    t.Error(fmt.Sprintf("ValidTag(%s) returned true for an invalid tag.", broke))
  }
}
