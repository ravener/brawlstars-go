// Brawlstars Go is an easy to use API wrapper for BrawlAPI https://brawlapi.cf
package brawlstars

import (
  "net/http"
  "io/ioutil"
  "errors"
  "encoding/json"
)

// The current version of the wrapper as a string.
const Version string = "0.0.3"

// Client represents a base client used to interact with the API
// To construct one it is recommended to use the NewClient(token) function
// or NewClientWithHttp(token, *http.Client) if you would like to pass a custom
// Http Client to use.
type Client struct {
  HttpClient *http.Client // The http client
  Token string // The authentication token.
}

// NewClient constructs a new client with the given token for authentication.
// 
// This function creates a new http.Client to use for requests use
// NewClientWithHttp if you would like to pass a custom http client.
func NewClient(token string) *Client {
  return &Client{HttpClient:&http.Client{}, Token:token}
}

// NewClientWithHttp constructs a new client with the given token for authentication
// and the given http.Client for requests.
func NewClientWithHttp(token string, client *http.Client) *Client {
  return &Client{HttpClient:client, Token:token}
}

func (c *Client) httpGet(url string) ([]byte, error) {
  req, err := http.NewRequest("GET", "https://brawlapi.cf/api" + url, nil)
  if err != nil {
    return nil, err
  }
  req.Header.Set("Authorization", c.Token)
  req.Header.Set("User-Agent", "BrawlStars-Go")
  res, err := c.HttpClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    return nil, errors.New(res.Status)
  }
  bytes, err := ioutil.ReadAll(res.Body)
  return bytes, err
}

// GetPlayer fetches a player with the given tag and returns a Player
func (c *Client) GetPlayer(tag string) (Player, error) {
  res, err := c.httpGet("/player?tag=" + tag)
  if err != nil {
    return Player{}, err
  }
  p := Player{Client:c}
  err = json.Unmarshal(res, &p)
  if err != nil {
    return Player{}, err
  }
  if p.Club != nil {
    p.Club.Client = c
  }
  return p, nil
}

func (c *Client) GetClub(tag string) (Club, error) {
  res, err := c.httpGet("/club?tag=" + tag)
  if err != nil {
    return Club{}, nil
  }
  club := Club{Client:c}
  err = json.Unmarshal(res, &club)
  if err != nil {
    return Club{}, nil
  }
  for _, cl := range club.Members {
    cl.Client = c
  }
  return club, nil
}

// GetMisc fetches misc data such as shop reset and such.
func (c *Client) GetMisc() (Misc, error) {
  res, err := c.httpGet("/misc")
  if err != nil {
    return Misc{}, err
  }
  var m Misc
  err = json.Unmarshal(res, &m)
  if err != nil {
    return Misc{}, err
  }
  return m, nil
}

// ClubSearch searchs for clubs matching the given query
// and returns an array of Clubs
func (c *Client) ClubSearch(query string) ([]Club, error) {
  res, err := c.httpGet("/clubSearch?query=" + query)
  if err != nil {
    return nil, err
  }
  var clubs []Club
  err = json.Unmarshal(res, &clubs)
  if err != nil {
    return nil, err
  }
  return clubs, nil
}

// GetTopClubs gets top clubs from the leaderboard optionally the results can be
// limited by given count, if passed 0 it won't pass the count to API.
func (c *Client) GetTopClubs(count int) ([]Club, error) {
  var countStr string = ""
  if count != 0 {
    countStr = "?count=" + string(count)
  }
  res, err := c.httpGet("/leaderboards/clubs" + countStr)
  if err != nil {
    return nil, err
  }
  var clubs []Club
  err = json.Unmarshal(res, &clubs)
  if err != nil {
    return nil, err
  }

  for _, cl := range clubs {
    cl.Client = c
  }
  return clubs, nil
}

// GetTopPlayers gets top players from the leaderboard,
// optionally limiting the results with the given count
// if count is 0 it won't be passed to API.
// A brawler name can also be specified to filter the results,
// Brawler can be an empty string if you don't need it.
func (c *Client) GetTopPlayers(count int, brawler string) ([]Player, error) {
  var countStr string = ""
  var brawlerQuery string = ""
  if count != 0 {
    countStr = "?count=" + string(count)
  }
  if brawler != "" {
    var delim string = "?"
    if countStr != "" {
      delim = "&"
    }
    brawlerQuery = delim + "brawler=" + brawler
  }
  res, err := c.httpGet("/leaderboards/players" + countStr + brawlerQuery)
  if err != nil {
    return nil, err
  }
  var p []Player
  err = json.Unmarshal(res, &p)
  if err != nil {
    return nil, err
  }
  for _, pl := range p {
    pl.Client = c
  }
  return p, nil
}

func (c * Client) events(Type string) (Events, error) {
  var typeStr string = ""
  if Type != "" {
    typeStr = "?type=" + Type
  }
  res, err := c.httpGet("/events" + typeStr)
  if err != nil {
    return Events{}, err
  }
  var e Events
  err = json.Unmarshal(res, &e)
  if err != nil {
    return Events{}, err
  }
  return e, nil
}

// GetEvents returns current & upcoming events
// If you need just the current or upcoming use 
// GetCurrentEvents() or GetUpcomingEvents() for smaller responses.
func (c *Client) GetEvents() (Events, error) {
  return c.events("")
}

// GetUpcomingEvents returns an array of EventSlots for the upcoming events.
func (c *Client) GetUpcomingEvents() ([]EventSlot, error) {
  events, err := c.events("upcoming")
  if err != nil {
    return nil, err
  }
  return events.Upcoming, nil
}

// GetCurrentEvents returns an array of EventSlots for the current events.
func (c *Client) GetCurrentEvents() ([]EventSlot, error) {
  events, err := c.events("current")
  if err != nil {
    return nil, err
  }
  return events.Current, nil
}
