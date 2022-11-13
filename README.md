# Yahoo Fantasy Query Builders for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/famendola1/yfquery.svg)](https://pkg.go.dev/github.com/famendola1/yfquery)
![License](https://img.shields.io/badge/License-Apache-green)
[![Go Report Card](https://goreportcard.com/badge/github.com/famendola1/yfquery)](https://goreportcard.com/report/github.com/famendola1/yfquery)

## Installation
~~~bash
go get github.com/famendola1/yfquery
~~~

## Yahoo Endpoints
The Yahoo offical documentation for their [Fantasy Sports API](https://developer.yahoo.com/fantasysports/guide) is not comprehensive and incomplete, despite being the offical. For a more complete overview of the supported endpoints, see the [README](https://github.com/edwarddistel/yahoo-fantasy-baseball-reader#yahoo-fantasy-api-docs).

### Before You Start
The functionality of this package require the use of a `*http.Client` that is configured for the Yahoo Fantasy API endpoint. You can use the [github.com/famendola1/yauth](https://pkg.go.dev/github.com/famendola1/yauth) package to configure a `*http.Client` to use.

### Query Builders
The query buiilders were designed to be able to easily generate queries for all the Yahoo Fantasy API endpoints. The builders expose functions that add pieces and parameters to the query. They also expose the following functions:

* `ToString`: Builds the string for the query that the builder represents.
* `Get`: Sends a GET request to the Yahoo Fantasy API for the endpoint that the query represents. A successful query will return a [`FantasyContent`](https://pkg.go.dev/github.com/famendola1/yfantasy/schema#FantasyContent) struct.
* `Post`: Sends a POST request to the Yahoo Fantasy API for the endpoint that the query represents. A successful query will return a [`FantasyContent`](https://pkg.go.dev/github.com/famendola1/yfantasy/schema#FantasyContent) struct.
* `Put`: Sends a PUT request to the Yahoo Fantasy API for the endpoint that the query represents. A successful query will return a [`FantasyContent`](https://pkg.go.dev/github.com/famendola1/yfantasy/schema#FantasyContent) struct.
* `Delete`: Sends a DELETE request to the Yahoo Fantasy API for the endpoint that the query represents. A successful query will return a [`FantasyContent`](https://pkg.go.dev/github.com/famendola1/yfantasy/schema#FantasyContent) struct.
* `Reset`: Clears all query parameters and the payload of the query. If `Reset` is called on a chained query builder (i.e. `Transactions` in `League().Key("key").Transactions()`), only the data on the chained query builder is cleared.

WARNING: The query builders do not validate that the queries they build are valid Yahoo endpoints.

#### Example Usages
```go
// Get all the teams in a league.
Leagues().Key("nba.l.12345").Teams().Get(client)

// Search for a player and return their stats for the past week.
Leagues().Key("nba.l.12345").Players().Search("Donovan Mitchell").Stats().LastWeek().Get(client)

// Get all the leagues a user is in.
Users().Leagues().Get(client)

// Get the rosters for all teams in a league.
League().Key("nba.l.12345").Teams().Roster().Get(client)

// Get all the add and drop transactions in a league.
League().Key("nba.l.12345").Transactions().Types([]string{"add", "drop"}).Get(client)

// Add players to a team.
League().Key("nba.l.12345").Transactions().Add("nba.l.12345.t.1", "nba.p.6030").Post(client)

// Accept a trade.
Transaction().AcceptTrade("248.l.55438.pt.11", "deal!!").Put(client)

// Cancel a waiver claim.
Transaction().CancelWaiverClaim("248.l.55438.w.c.2_6093").Delete(client)

// Reset behavior
q := League().Key("nba.l.12345").Teams()

// /league/nba.l.12345/teams/roster
q.Roster().Get(client)

q.Reset()

// /league/nba.l.12345/teams/stats
q.Stats().Get(client)
```
