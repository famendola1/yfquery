package yfquery

import "testing"

func TestGameQueryCollection(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Games(),
				"/games",
			},
			{
				Games().IsAvailable(),
				"/games;is_available=1",
			},
			{
				Games().Types([]string{"full", "pickem-team"}),
				"/games;game_types=full,pickem-team",
			},
			{
				Games().Codes([]string{"nba", "nhl", "nfl"}),
				"/games;game_codes=nba,nhl,nfl",
			},
			{
				Games().Seasons([]int{2022}),
				"/games;seasons=2022",
			},
			{
				Games().Keys([]string{"223", "224"}),
				"/games;game_keys=223,224",
			},
			{
				Games().Key("223"),
				"/games;game_keys=223",
			},
			{
				Games().Key("223").Key("224"),
				"/games;game_keys=224",
			},
			{
				Games().Key("nba").RosterPositions(),
				"/games;game_keys=nba/roster_positions",
			},
			{
				Games().Key("nba").GameWeeks(),
				"/games;game_keys=nba/game_weeks",
			},
			{
				Games().Key("nba").PositionTypes(),
				"/games;game_keys=nba/position_types",
			},
			{
				Games().Key("nba").StatCategories(),
				"/games;game_keys=nba/stat_categories",
			},
			{
				Games().Key("nba").StatCategories().RosterPositions(),
				"/games;game_keys=nba;out=stat_categories,roster_positions",
			},
			{
				Games().IsAvailable().Codes([]string{"nba"}).Seasons([]int{2022, 2021}).Types([]string{"full"}).Keys([]string{"223"}),
				"/games;game_keys=223;is_available=1;game_codes=nba;seasons=2022,2021;game_types=full",
			},
		})
}

func TestGameQueryResource(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Game(),
				"/game",
			},
			{
				Game().Key("nba"),
				"/game/nba",
			},
			{
				Game().Key("nba").Leagues(),
				"/game/nba/leagues",
			},
			{
				Game().Key("nba").Teams(),
				"/game/nba/teams",
			},
			{
				Game().Key("nba").RosterPositions(),
				"/game/nba/roster_positions",
			},
			{
				Game().Key("nba").GameWeeks(),
				"/game/nba/game_weeks",
			},
			{
				Game().Key("nba").PositionTypes(),
				"/game/nba/position_types",
			},
			{
				Game().Key("nba").StatCategories(),
				"/game/nba/stat_categories",
			},
			{
				Game().Key("nba").StatCategories().RosterPositions(),
				"/game/nba;out=stat_categories,roster_positions",
			},
		})
}
