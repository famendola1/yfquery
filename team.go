package yfquery

import (
	"fmt"
	"strings"
)

// TeamQuery can be used to query the /teams or /team Yahoo Fantasy resource.
type TeamQuery struct {
	query
}

// TeamRosterQuery can be used to query the /roster subresource of Game resource.
type TeamRosterQuery struct {
	query
}

// Team returns a TeamQuery for the /team endpoint.
func Team() *TeamQuery {
	return &TeamQuery{query{resource: "team"}}
}

// Teams returns a TeamQuery for the /teams endpoint.
func Teams() *TeamQuery {
	return &TeamQuery{query{resource: "team", isCollection: true}}
}

// Keys adds the "team_keys" parameter with the given keys to the query.
func (t *TeamQuery) Keys(keys []string) *TeamQuery {
	t.keys = append(t.keys, keys...)
	return t
}

// Key sets the "team_keys" parameter to the the given key. When querying the
// /team endpoint the key will be appended to the query path (i.e. /team/<key>).
func (t *TeamQuery) Key(key string) *TeamQuery {
	t.keys = []string{key}
	return t
}

// Roster adds the "roster" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. team/roster).
func (t *TeamQuery) Roster() *TeamQuery {
	t.outs = append(t.outs, "roster")
	return t
}

// AllMatchups adds the "matchups" subresource to the request to request for
// all matchups. If combined with other subresources, they are all combined into
// the "out" parameter, otherwise it is added to the request path (i.e. team/matchups).
func (t *TeamQuery) AllMatchups() *TeamQuery {
	t.outs = append(t.outs, "matchups")
	return t
}

// Matchups adds the "matchups" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. team/matchups). weeks is expected by
// Yahoo to contain positive integers.
func (t *TeamQuery) Matchups(weeks []int) *TeamQuery {
	t.outs = append(t.outs, "matchups")
	t.params = append(t.params, fmt.Sprintf("weeks=%s", strings.Trim(strings.Replace(fmt.Sprint(weeks), " ", ",", -1), "[]")))
	return t
}

// Standings adds the "standings" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. team/standings).
func (t *TeamQuery) Standings() *TeamQuery {
	t.outs = append(t.outs, "standings")
	return t
}

// DraftResults adds the "draftresults" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. team/draftresults).
func (t *TeamQuery) DraftResults() *TeamQuery {
	t.outs = append(t.outs, "draftresults")
	return t
}

// StatsWithDefaults adds the "stats" subresource to the request. If combined with
// other subresources, they are all combined into the "out" parameter, otherwise
// it is added to the request path (i.e. team/stats).
func (t *TeamQuery) StatsWithDefaults() *TeamQuery {
	t.outs = append(t.outs, "stats")
	return t
}

// Stats returns a StatsQuery for the /stats subresource.
func (t *TeamQuery) Stats() *StatsQuery {
	st := Stats()
	st.base = t.ToString()
	return st
}

// RosterDay returns a TeamRosterQuery for the /roster subresource with the
// "date" parameter set to the provided date. date is expected to be formatted as
// YYYY-MM-DD.
func (t *TeamQuery) RosterDay(date string) *TeamRosterQuery {
	return &TeamRosterQuery{
		query{
			base:     t.ToString(),
			resource: "roster",
			params:   []string{fmt.Sprintf("date=%s", date)},
		},
	}
}

// RosterWeek returns a TeamRosterQuery for the /roster subresource with the
// "week" parameter set to the provided week. week is expected by Yahoo to be a
// positive integer.
func (t *TeamQuery) RosterWeek(week int) *TeamRosterQuery {
	return &TeamRosterQuery{
		query{
			base:     t.ToString(),
			resource: "roster",
			params:   []string{fmt.Sprintf("week=%d", week)},
		},
	}
}
