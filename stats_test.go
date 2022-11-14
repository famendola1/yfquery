package yfquery

import "testing"

func TestStatsQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Stats(),
				"/stats",
			},
			{
				Stats().Week(1),
				"/stats;type=week;week=1",
			},
			{
				Stats().LastWeek(),
				"/stats;type=lastweek",
			},
			{
				Stats().LastMonth(),
				"/stats;type=lastmonth",
			},
			{
				Stats().LastWeekAverage(),
				"/stats;type=average_lastweek",
			},
			{
				Stats().LastMonthAverage(),
				"/stats;type=average_lastmonth",
			},
			{
				Stats().Today(),
				"/stats;type=date",
			},
			{
				Stats().Day("2006-10-06"),
				"/stats;type=date;date=2006-10-06",
			},
			{
				Stats().CurrentSeason(),
				"/stats;type=season",
			},
			{
				Stats().Season(2022),
				"/stats;type=season;season=2022",
			},
			{
				Stats().CurrentSeasonAverage(),
				"/stats;type=average_season",
			},
			{
				Stats().SeasonAverage(2022),
				"/stats;type=average_season;season=2022",
			},
		})
}
