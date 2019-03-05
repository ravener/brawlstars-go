# Brawlstars Go
Go API wrapper for the unofficial [BrawlAPI](https://brawlapi.cf)

## Authentication
You need a token to use the API, which can be found in the [Discord Server](https://dicord.me/BrawlAPI)

## Install
```sh
$ go get github.com/pollen5/brawlstars-go
```

## Usage
```go
package main

import (
  "fmt"
  "github.com/pollen5/brawlstars-go"
)

func main() {
  client := brawlstars.NewClient("token")
  player := client.GetPlayer("tag")
  fmt.Println(player.Name)
  club := player.GetClub()
  fmt.Println(club.Name)
}
```

## Running Tests
```sh
$ TOKEN=token go test
```

## License
[MIT](LICENSE)
