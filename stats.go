package yfquery

import "fmt"

// StatsQuery can be used to query the /stats endpoint.
type StatsQuery struct {
	query
}

// Stats returns a StatsQuery for the /stats subresource endpoint.
func Stats() *StatsQuery {
	return &StatsQuery{query{resource: "stats"}}
}

// Week adds the "type=week" parameter to the query. Only supported by Yahoo for football.
func (s *StatsQuery) Week(week int) *StatsQuery {
	s.params = append(s.params, []string{"type=week", fmt.Sprintf("week=%d", week)}...)
	return s
}

// LastWeek adds the "type=lastweek" parameter to the query.
func (s *StatsQuery) LastWeek() *StatsQuery {
	s.params = append(s.params, "type=lastweek")
	return s
}

// LastMonth adds the "type=lastmonth" parameter to the query.
func (s *StatsQuery) LastMonth() *StatsQuery {
	s.params = append(s.params, "type=lastmonth")
	return s
}

// LastWeekAverage adds the "type=average_lastweek" parameter to the query.
func (s *StatsQuery) LastWeekAverage() *StatsQuery {
	s.params = append(s.params, "type=average_lastweek")
	return s
}

// LastMonthAverage adds the "type=average_lastmonth" parameter to the query.
func (s *StatsQuery) LastMonthAverage() *StatsQuery {
	s.params = append(s.params, "type=average_lastmonth")
	return s
}

// Today adds the "type=date" parameter to the query.
func (s *StatsQuery) Today() *StatsQuery {
	s.params = append(s.params, "type=date")
	return s
}

// Day adds the "type=date" parameter and "date" parameter with the provided
// date to the query. date is expected to be formatted as YYYY-MM-DD.
func (s *StatsQuery) Day(date string) *StatsQuery {
	s.params = append(s.params, []string{"type=date", fmt.Sprintf("date=%s", date)}...)
	return s
}

// CurrentSeason adds the "type=season" parameter to the query.
func (s *StatsQuery) CurrentSeason() *StatsQuery {
	s.params = append(s.params, "type=season")
	return s
}

// Season adds the "type=season" parameter and "season" parameter with the
// provided season to the query.
func (s *StatsQuery) Season(season int) *StatsQuery {
	s.params = append(s.params, []string{"type=season", fmt.Sprintf("season=%d", season)}...)
	return s
}

// CurrentSeasonAverage adds the "type=average_season" parameter to the query.
func (s *StatsQuery) CurrentSeasonAverage() *StatsQuery {
	s.params = append(s.params, "type=average_season")
	return s
}

// SeasonAverage adds the "type=average_season" parameter and "season" parameter
// with the provided season to the query.
func (s *StatsQuery) SeasonAverage(season int) *StatsQuery {
	s.params = append(s.params, []string{"type=average_season", fmt.Sprintf("season=%d", season)}...)
	return s
}
