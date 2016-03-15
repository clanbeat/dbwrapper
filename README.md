#dbwrapper

Go package to wrap Clanbeat postgresql database connection and queries

##Usage

Define the connection in main

```go
package main

import (
  "github.com/clanbeat/dbwrapper"
)
var db *dbwrapper.DB


func main() {
  //connect to database
  if len(env.DatabaseURL) > 1 {
    db = database.ConnectWithURL(env.DatabaseURL)
  } else {
    db = database.ConnectDevelopment()
  }
  defer db.Close()
}
```
