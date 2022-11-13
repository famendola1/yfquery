package yfquery

// UserQuery can be used to query the /users Yahoo Fantasy endpoint.
type UserQuery struct {
	query
}

// Users returns a UserQuery for the /users endpoint.
func Users() *UserQuery {
	return &UserQuery{
		query{
			resource: "users",
			params:   []string{"use_login=1"},
		},
	}
}

// Games returns a GameQuery for the /games subresource.
func (u *UserQuery) Games() *GameQuery {
	return &GameQuery{
		query{
			base:         u.ToString(),
			resource:     "game",
			isCollection: true,
		},
	}
}

// Leagues returns a GameQuery for the /leagues subresource.
func (u *UserQuery) Leagues() *LeagueQuery {
	return &LeagueQuery{
		query{
			base:         u.ToString(),
			resource:     "league",
			isCollection: true,
		},
	}
}

// Teams returns a TeamQuery for the /teams subresource.
func (u *UserQuery) Teams() *TeamQuery {
	return &TeamQuery{
		query{
			base:         u.ToString(),
			resource:     "team",
			isCollection: true,
		},
	}
}
