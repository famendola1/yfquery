package yfquery

import "testing"

func TestUserQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Users(),
				"/users;use_login=1",
			},
			{
				Users().Games(),
				"/users;use_login=1/games",
			},
			{
				Users().Teams(),
				"/users;use_login=1/teams",
			},
			{
				Users().Leagues(),
				"/users;use_login=1/leagues",
			},
		})
}
