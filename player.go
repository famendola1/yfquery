package yfquery

import (
	"fmt"
	"net/url"
)

// PlayerQuery can be used to query the /players or /player Yahoo Fantasy endpoints.
type PlayerQuery struct {
	query
}

// PlayerStatus represents a player's ownership status.
type PlayerStatus string

// Enum for a player's ownership status.
const (
	PlayerStatusUnknown   PlayerStatus = ""
	PlayerStatusAvailable              = "A"
	PlayerStatusFreeAgent              = "FA"
	PlayerStatusWaiver                 = "W"
	PlayerStatusTaken                  = "T"
	PlayerStatusKeeper                 = "K"
)

// PlayerSortCriteria represents the sort criteria for players.
type PlayerSortCriteria string

// Enum for player sort criteria.
const (
	PlayerSortCriteriaUnknown       PlayerSortCriteria = ""
	PlayerSortCriteriaName                             = "NAME"
	PlayerSortCriteriaOverallRank                      = "OR"
	PlayerSortCriteriaActualRank                       = "AR"
	PlayerSortCriteriaFantasyPoints                    = "PTS"
)

// PlayerSortType represent the sort type for players.
type PlayerSortType string

// Enum for player sort type.
const (
	PlayerSortTypeUnknown          PlayerSortType = ""
	PlayerSortTypeSeason                          = "season"
	PlayerSortTypeAverageSeason                   = "average_season"
	PlayerSortTypeDate                            = "date"
	PlayerSortTypeWeek                            = "week"
	PlayerSortTypeLastMonth                       = "lastmonth"
	PlayerSortTypeAverageLastMonth                = "average_lastmonth"
	PlayerSortTypeLastWeek                        = "lastweek"
	PlayerSortTypeAverageLastWeek                 = "average_lastweek"
)

// Players returns a PlayerQuery for the /players endpoint.
func Players() *PlayerQuery {
	return &PlayerQuery{query{resource: "player", isCollection: true}}
}

// Player returns a PlayerQuery for the /player endpoint.
func Player() *PlayerQuery {
	return &PlayerQuery{query{resource: "player"}}
}

// Keys adds the "player_keys" parameter with the given keys to the query.
func (p *PlayerQuery) Keys(keys []string) *PlayerQuery {
	p.keys = append(p.keys, keys...)
	return p
}

// Key sets the "player_keys" parameter to the the given key. When querying the
// /game endpoint the key will be appended to the query path (i.e. /player/<key>).
func (p *PlayerQuery) Key(key string) *PlayerQuery {
	p.keys = []string{key}
	return p
}

// Position adds the "position" parameter with the provided position to the query.
// Valid player positions can be provided as input (e.x. "QB", "PG")
func (p *PlayerQuery) Position(pos string) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("position=%s", pos))
	return p
}

// Status adds the "status" parameter with the provided status to the query.
func (p *PlayerQuery) Status(status PlayerStatus) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("status=%s", status))
	return p
}

// Search adds the "search" parameter with the provided name to the query.
func (p *PlayerQuery) Search(name string) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("search=%s", url.QueryEscape(name)))
	return p
}

// Sort adds the "sort" parameter with the provided sort criteria to the query.
func (p *PlayerQuery) Sort(sort PlayerSortCriteria) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort=%s", sort))
	return p
}

// SortByStat adds the "sort" parameter with the provided stat id to the query.
func (p *PlayerQuery) SortByStat(statID int) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort=%d", statID))
	return p
}

// SortType adds the "sort_type" parameter with the provided type to the query.
func (p *PlayerQuery) SortType(sortType PlayerSortType) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort_type=%s", sortType))
	return p
}

// SortSeason adds the "sort_season" parameter with the provided season to the
// query.
func (p *PlayerQuery) SortSeason(season int) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort_season=%d", season))
	return p
}

// SortDate adds the "sort_date" parameter with the provided date to the query.
// date should be formatted as YYYY-MM-DD.
func (p *PlayerQuery) SortDate(date string) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort_date=%s", date))
	return p
}

// SortWeek adds the "sort_week" parameter with the provided week to the query.
// Yahoo only supports this parameter for football. week is expected by Yahoo
// to be a positive integer.
func (p *PlayerQuery) SortWeek(week int) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("sort_week=%d", week))
	return p
}

// Start adds the "start" parameter with the provided start to the query. start
// is expected by Yahoo to be a positive integer.
func (p *PlayerQuery) Start(start int) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("start=%d", start))
	return p
}

// Count adds the "count" parameter with the provided count to the query. count
// is expected by Yahoo to be a positive integer.
func (p *PlayerQuery) Count(count int) *PlayerQuery {
	p.params = append(p.params, fmt.Sprintf("count=%d", count))
	return p
}

// Ownership adds the "ownership" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. player/ownership).
func (p *PlayerQuery) Ownership() *PlayerQuery {
	p.outs = append(p.outs, "ownership")
	return p
}

// Opponent adds the "opponent" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. player/opponent).
func (p *PlayerQuery) Opponent() *PlayerQuery {
	p.outs = append(p.outs, "opponent")
	return p
}

// PercentOwned adds the "percent_owned" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. player/percent_owned).
func (p *PlayerQuery) PercentOwned() *PlayerQuery {
	p.outs = append(p.outs, "percent_owned")
	return p
}

// DraftAnalysis adds the "draft_analysis" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. player/draft_analysis).
func (p *PlayerQuery) DraftAnalysis() *PlayerQuery {
	p.outs = append(p.outs, "draft_analysis")
	return p
}

// StatsWithDefaults adds the "stats" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. player/stats).
func (p *PlayerQuery) StatsWithDefaults() *PlayerQuery {
	p.outs = append(p.outs, "stats")
	return p
}

// Stats returns a StatsQuery for the /stats subresource.
func (p *PlayerQuery) Stats() *StatsQuery {
	st := Stats()
	st.base = p.ToString()
	return st
}
