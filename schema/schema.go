package schema

// FantasyContent models the top-level structure returned by the Yahoo Fantasy API.
type FantasyContent struct {
	Lang        string   `xml:"lang,attr,omitempty"`
	URI         string   `xml:"uri,attr,omitempty"`
	Time        string   `xml:"time,attr,omitempty"`
	Copyright   string   `xml:"copyright,attr,omitempty"`
	RefreshRate string   `xml:"refresh_rate,attr,omitempty"`
	Yahoo       string   `xml:"yahoo,attr,omitempty"`
	Xmlns       string   `xml:"xmlns,attr,omitempty"`
	Games       *Games   `xml:"games,omitempty"`
	Game        *Game    `xml:"game,omitempty"`
	Leagues     *Leagues `xml:"leagues,omitempty"`
	League      *League  `xml:"league,omitempty"`
	Teams       *Teams   `xml:"teams,omitempty"`
	Team        *Team    `xml:"team,omitempty"`
	Users       *Users   `xml:"users,omitempty"`
}

// Users is a list of users.
type Users struct {
	Count int   `xml:"count,attr,omitempty"`
	User  *User `xml:"user,omitempty"`
}

// User models a user in Users.
type User struct {
	GUID    string   `xml:"guid,omitempty"`
	Games   *Games   `xml:"games,omitempty"`
	Leagues *Leagues `xml:"leagues,omitempty"`
	Teams   *Teams   `xml:"teams,omitempty"`
}

// Player models a Yahoo player.
type Player struct {
	PlayerKey                string               `xml:"player_key,omitempty"`
	PlayerID                 int                  `xml:"player_id,omitempty"`
	Name                     *Name                `xml:"name,omitempty"`
	FirstName                string               `xml:"first_name,omitempty"`
	FirstNameInitials        string               `xml:"first_name_initials,omitempty"`
	LastName                 string               `xml:"last_name,omitempty"`
	EditorialPlayerKey       string               `xml:"editorial_player_key,omitempty"`
	EditorialTeamKey         string               `xml:"editorial_team_key,omitempty"`
	EditorialTeamFullName    string               `xml:"editorial_team_full_name,omitempty"`
	EditorialTeamAbbr        string               `xml:"editorial_team_abbr,omitempty"`
	UniformNumber            int                  `xml:"uniform_number,omitempty"`
	DisplayPosition          string               `xml:"display_position,omitempty"`
	Headshot                 *Headshot            `xml:"headshot,omitempty"`
	ImageURL                 string               `xml:"image_url,omitempty"`
	IsUndroppable            bool                 `xml:"is_undroppable,omitempty"`
	PositionType             string               `xml:"position_type,omitempty"`
	PrimaryPosition          string               `xml:"primary_position,omitempty"`
	EligiblePositions        *EligiblePositions   `xml:"eligible_positions,omitempty"`
	HasPlayerNotes           bool                 `xml:"has_player_notes,omitempty"`
	PlayerNotesLastTimestamp uint64               `xml:"player_notes_last_timestamp,omitempty"`
	TransactionData          *TransactionData     `xml:"transaction_data,omitempty"`
	PlayerStats              *PlayerStats         `xml:"player_stats,omitempty"`
	PlayerAdvancedStats      *PlayerAdvancedStats `xml:"player_advanced_stats,omitempty"`
	Status                   string               `xml:"status,omitempty"`
	StatusFull               string               `xml:"status_full,omitempty"`
	InjuryNote               string               `xml:"injury_note,omitempty"`
	HasRecentPlayerNotes     string               `xml:"has_recent_player_notes,omitempty"`
	SelectedPosition         *SelectedPosition    `xml:"selected_position,omitempty"`
	IsKeeper                 *IsKeeper            `xml:"is_keeper,omitempty"`
	PercentOwned             *PercentOwned        `xml:"percent_owned,omitempty"`
	DraftAnalysis            *DraftAnalysis       `xml:"draft_analysis,omitempty"`
	Ownership                *Ownership           `xml:"ownership,omitempty"`
	Opponent                 string               `xml:"opponent,omitempty"`
}

// Name contains information about a Player's name.
type Name struct {
	Full       string `xml:"full,omitempty"`
	First      string `xml:"first,omitempty"`
	Last       string `xml:"last,omitempty"`
	ASCIIFirst string `xml:"ascii_first,omitempty"`
	ASCIILast  string `xml:"ascii_last,omitempty"`
}

// Headshot contains information about a Player's Yahoo photo.
type Headshot struct {
	URL  string `xml:"url,omitempty"`
	Size string `xml:"size,omitempty"`
}

// EligiblePositions contains all the eligible positions a Player can play.
type EligiblePositions struct {
	Position []string `xml:"position,omitempty"`
}

// TransactionData contains transaction information for a Player.
type TransactionData struct {
	Type                string `xml:"type,omitempty"`
	SourceType          string `xml:"source_type,omitempty"`
	DestinationType     string `xml:"destination_type,omitempty"`
	DestinationTeamKey  string `xml:"destination_team_key,omitempty"`
	DestinationTeamName string `xml:"destination_team_name,omitempty"`
	SourceTeamKey       string `xml:"source_team_key,omitempty"`
	SourceTeamName      string `xml:"source_team_name,omitempty"`
}

// PlayerStats contains stat information for a Player.
type PlayerStats struct {
	CoverageType string `xml:"coverage_type,omitempty"`
	Season       string `xml:"season,omitempty"`
	Date         string `xml:"date,omitempty"`
	Stats        *Stats `xml:"stats,omitempty"`
}

// PlayerAdvancedStats contains stat information for a Player.
type PlayerAdvancedStats struct {
	CoverageType string `xml:"coverage_type,omitempty"`
	Season       string `xml:"season,omitempty"`
	Stats        *Stats `xml:"stats,omitempty"`
}

// SelectedPosition contains information on a Player's selected position.
type SelectedPosition struct {
	CoverageType string `xml:"coverage_type,omitempty"`
	Date         string `xml:"date,omitempty"`
	Position     string `xml:"position,omitempty"`
	IsFlex       bool   `xml:"is_flex,omitempty"`
}

// IsKeeper contains keeper information for a Player.
type IsKeeper struct {
	Status string `xml:"status,omitempty"`
	Cost   string `xml:"cost,omitempty"`
	Kept   string `xml:"kept,omitempty"`
}

// Players is a list of players.
type Players struct {
	Count  int      `xml:"count,attr,omitempty"`
	Player []Player `xml:"player,omitempty"`
}

// Transactions is a list of transactions.
type Transactions struct {
	Transaction []Transaction `xml:"transactions,omitempty"`
}

// Transaction represents a Yahoo fantasy transaction.
type Transaction struct {
	TransactionKey    string   `xml:"transaction_key,omitempty"`
	TransactionID     int      `xml:"transaction_id,omitempty"`
	Type              string   `xml:"type,omitempty"`
	Status            string   `xml:"status,omitempty"`
	Timestamp         uint64   `xml:"timestamp,omitempty"`
	WaiverPlayerKey   string   `xml:"waiver_player_key,omitempty"`
	WaiverTeamKey     string   `xml:"waiver_team_key,omitempty"`
	WaiverDate        string   `xml:"waiver_date,omitempty"`
	WaiverPriority    int      `xml:"waiver_priority,omitempty"`
	TraderTeamKey     string   `xml:"trader_team_key,omitempty"`
	TradeeTeamKey     string   `xml:"tradee_team_key,omitempty"`
	TradeProposedTime uint64   `xml:"trade_proposed_time,omitempty"`
	TradeNode         string   `xml:"trade_note,omitempty"`
	VoterTeamKey      string   `xml:"voter_team_key,omitempty"`
	Players           *Players `xml:"players,omitempty"`
}

// Matchups is a list of matchups.
type Matchups struct {
	Matchup []Matchup `xml:"matchup,omitempty"`
}

// Matchup represents a Yahoo matchup.
type Matchup struct {
	Week          int          `xml:"week,omitempty"`
	WeekStart     string       `xml:"week_start,omitempty"`
	WeekEnd       string       `xml:"week_end,omitempty"`
	Status        string       `xml:"status,omitempty"`
	IsPlayoffs    bool         `xml:"is_playoffs,omitempty"`
	IsConsolation bool         `xml:"is_consolation,omitempty"`
	IsTied        bool         `xml:"is_tied,omitempty"`
	WinnerTeamKey string       `xml:"winner_team_key,omitempty"`
	StatWinners   *StatWinners `xml:"stat_winners,omitempty"`
	Teams         *Teams       `xml:"teams,omitempty"`
}

// StatWinners is a list of stat winners.
type StatWinners struct {
	StatWinner []StatWinner `xml:"stat_winner,omitempty"`
}

// StatWinner represents the winner of a stat category.
type StatWinner struct {
	StatID        int    `xml:"stat_id,omitempty"`
	WinnerTeamKey string `xml:"winner_team_key,omitempty"`
	IsTied        bool   `xml:"is_tied,omitempty"`
}

// StatCategories contains the enabled stat categories for a league.
type StatCategories struct {
	Stats *Stats `xml:"stats,omitempty"`
}

// StatModifiers contains the modifiers for stats.
type StatModifiers struct {
	Stats *Stats `xml:"stats,omitempty"`
}

// StatPositionTypes is a list of stat position types.
type StatPositionTypes struct {
	StatPositionType []StatPositionType `xml:"stat_position_type,omitempty"`
}

// StatPositionType contains information about a stat position type.
type StatPositionType struct {
	PositionType      string `xml:"position_type,omitempty"`
	IsOnlyDisplayStat string `xml:"is_only_display_stat,omitempty"`
}

// Stats is a list of stats.
type Stats struct {
	Stat []Stat `xml:"stat,omitempty"`
}

// Stat represents a stat category in Yahoo.
type Stat struct {
	StatID            int                `xml:"stat_id,omitempty"`
	Value             string             `xml:"value,omitempty"`
	Enabled           bool               `xml:"enabled,omitempty"`
	Name              string             `xml:"name,omitempty"`
	DisplayName       string             `xml:"display_name,omitempty"`
	Group             string             `xml:"group,omitempty"`
	Abbr              string             `xml:"abbr,omitempty"`
	SortOrder         string             `xml:"sort_order,omitempty"`
	PositionType      string             `xml:"position_type,omitempty"`
	PositionTypes     *PositionsTypes    `xml:"position_types,omitempty"`
	IsOnlyDisplayStat bool               `xml:"is_only_display_stat,omitempty"`
	StatPositionTypes *StatPositionTypes `xml:"stat_position_types,omitempty"`
	Groups            *Groups            `xml:"groups,omitempty"`
}

// Groups is a list of groups.
type Groups struct {
	Group []Group `xml:"group,omitempty"`
}

// Group contains information on a group.
type Group struct {
	GroupName        string `xml:"group_name,omitempty"`
	GroupDisplayName string `xml:"group_display_name,omitempty"`
	GroupAbbr        string `xml:"group_abbr,omitempty"`
}

// Game represents a Yahoo game
type Game struct {
	GameKey            string           `xml:"game_key,omitempty"`
	GameID             int              `xml:"game_id,omitempty"`
	Name               string           `xml:"name,omitempty"`
	Code               string           `xml:"code,omitempty"`
	Type               string           `xml:"type,omitempty"`
	URL                string           `xml:"url,omitempty"`
	Season             string           `xml:"season,omitempty"`
	IsRegistrationOver bool             `xml:"is_registration_over,omitempty"`
	IsGameOver         bool             `xml:"is_game_over,omitempty"`
	IsOffseason        bool             `xml:"is_offseason,omitempty"`
	RosterPositions    *RosterPositions `xml:"roster_positions,omitempty"`
	PositionsTypes     *PositionsTypes  `xml:"position_types,omitempty"`
	GameWeeks          *GameWeeks       `xml:"game_weeks,omitempty"`
	StatCategories     *StatCategories  `xml:"stat_categories,omitempty"`
	Players            *Players         `xml:"players,omitempty"`
}

// Games is a list Games.
type Games struct {
	Count int    `xml:"count,attr,omitempty"`
	Game  []Game `xml:"game,omitempty"`
}

// League represents a Yahoo league.
type League struct {
	LeagueKey             string        `xml:"league_key,omitempty"`
	LeagueID              int           `xml:"league_id,omitempty"`
	Name                  string        `xml:"name,omitempty"`
	URL                   string        `xml:"url,omitempty"`
	LogoURL               string        `xml:"logo_url,omitempty"`
	DraftStatus           string        `xml:"draft_status,omitempty"`
	NumTeams              int           `xml:"num_teams,omitempty"`
	EditKey               string        `xml:"edit_key,omitempty"`
	WeeklyDeadline        string        `xml:"weekly_deadline,omitempty"`
	LeagueUpdateTimestamp string        `xml:"league_update_timestamp,omitempty"`
	ScoringType           string        `xml:"scoring_type,omitempty"`
	LeagueType            string        `xml:"league_type,omitempty"`
	Renew                 string        `xml:"renew,omitempty"`
	ShortInvitationURL    string        `xml:"short_invitation_url,omitempty"`
	AllowAddToDlExtraPos  string        `xml:"allow_add_to_dl_extra_pos,omitempty"`
	IsProLeague           bool          `xml:"is_pro_league,omitempty"`
	IsCashLeague          bool          `xml:"is_cash_league,omitempty"`
	CurrentWeek           int           `xml:"current_week,omitempty"`
	StartWeek             int           `xml:"start_week,omitempty"`
	StartDate             string        `xml:"start_date,omitempty"`
	EndWeek               int           `xml:"end_week,omitempty"`
	EndDate               string        `xml:"end_date,omitempty"`
	GameCode              string        `xml:"game_code,omitempty"`
	Season                string        `xml:"season,omitempty"`
	IsFinished            bool          `xml:"is_finished,omitempty"`
	Standings             *Standings    `xml:"standings,omitempty"`
	Teams                 *Teams        `xml:"teams,omitempty"`
	Players               *Players      `xml:"players,omitempty"`
	Scoreboard            *Scoreboard   `xml:"scoreboard,omitempty"`
	Settings              *Settings     `xml:"settings,omitempty"`
	DraftResults          *DraftResults `xml:"draft_results,omitempty"`
	Transactions          *Transactions `xml:"transactions,omitempty"`
}

// Leagues is a list Leagues.
type Leagues struct {
	Count  int      `xml:"count,attr,omitempty"`
	League []League `xml:"league,omitempty"`
}

// Scoreboard contains all the matchups for a league.
type Scoreboard struct {
	Count    int       `xml:"count,attr,omitempty"`
	Matchups *Matchups `xml:"matchups,omitempty"`
}

// Standings contains the standings for a league.
type Standings struct {
	Teams *Teams `xml:"teams,omitempty"`
}

// Teams is a list of teams.
type Teams struct {
	Count int    `xml:"count,attr,omitempty"`
	Team  []Team `xml:"team,omitempty"`
}

// Team represents a Yahoo team.
type Team struct {
	TeamKey                  string                    `xml:"team_key,omitempty"`
	TeamID                   int                       `xml:"team_id,omitempty"`
	Name                     string                    `xml:"name,omitempty"`
	IsOwnedByCurrentLogin    bool                      `xml:"is_owned_by_current_login,omitempty"`
	URL                      string                    `xml:"url,omitempty"`
	TeamLogos                *TeamLogos                `xml:"team_logos,omitempty"`
	WaiverPriority           int                       `xml:"waiver_priority,omitempty"`
	NumberOfMoves            int                       `xml:"number_of_moves,omitempty"`
	NumberOfTrades           int                       `xml:"number_of_trades,omitempty"`
	RosterAdds               *RosterAdds               `xml:"roster_adds,omitempty"`
	LeagueScoringType        string                    `xml:"league_scoring_type,omitempty"`
	DraftPosition            int                       `xml:"draft_position,omitempty"`
	HasDraftGrade            bool                      `xml:"has_draft_grade,omitempty"`
	Managers                 *Managers                 `xml:"managers,omitempty"`
	TeamStats                *TeamStats                `xml:"team_stats,omitempty"`
	TeamPoints               *TeamPoints               `xml:"team_points,omitempty"`
	TeamRemainingGames       *TeamRemainingGames       `xml:"team_remaining_games,omitempty"`
	ClinchedPlayoffs         bool                      `xml:"clinched_playoffs,omitempty"`
	TeamStandings            *TeamStandings            `xml:"team_standings,omitempty"`
	Roster                   *Roster                   `xml:"roster,omitempty"`
	Matchups                 *Matchups                 `xml:"matchups,omitempty"`
	RecommendedTradePartners *RecommendedTradePartners `xml:"recommended_trade_partners,omitempty"`
	PositionalRanks          *PositionalRanks          `xml:"positional_ranks,omitempty"`
}

// TeamLogos contains a list of team logos.
type TeamLogos struct {
	TeamLogo []TeamLogo `xml:"team_logo,omitempty"`
}

// TeamLogo contains information on the logo for a team.
type TeamLogo struct {
	Size string `xml:"size,omitempty"`
	URL  string `xml:"url,omitempty"`
}

// RosterAdds contains roster add informatin for a team.
type RosterAdds struct {
	CoverageType  string `xml:"coverage_type,omitempty"`
	CoverageValue int    `xml:"coverage_value,omitempty"`
	Value         int    `xml:"value,omitempty"`
}

// Managers is a list of managers.
type Managers struct {
	Count   int       `xml:"count,attr,omitempty"`
	Manager []Manager `xml:"manager,omitempty"`
}

// Manager contains information about the manager of a team.
type Manager struct {
	ManagerID      int    `xml:"manager_id,omitempty"`
	Nickname       string `xml:"nickname,omitempty"`
	GUID           string `xml:"guid,omitempty"`
	IsCommissioner bool   `xml:"is_commissioner,omitempty"`
	IsCurrentLogin bool   `xml:"is_current_login,omitempty"`
	Email          string `xml:"email,omitempty"`
	ImageURL       string `xml:"image_url,omitempty"`
	FeloScore      int    `xml:"felo_score,omitempty"`
	FeloTier       string `xml:"felo_tier,omitempty"`
}

// TeamStats contains stats for a team.
type TeamStats struct {
	CoverageType string `xml:"coverage_type,omitempty"`
	Week         int    `xml:"week,omitempty"`
	Stats        *Stats `xml:"stats,omitempty"`
}

// TeamPoints contaons the points scored by a team.
type TeamPoints struct {
	CoverageType string `xml:"coverage_type,omitempty"`
	Week         int    `xml:"week,omitempty"`
	Total        int    `xml:"total,omitempty"`
}

// TeamRemainingGames contains information on the remaining games of a team.
type TeamRemainingGames struct {
	CoverageType string `xml:"coverage_type,omitempty"`
	Week         int    `xml:"week,omitempty"`
	Total        *Total `xml:"total,omitempty"`
}

// Total contains information on remaining games.
type Total struct {
	RemainingGames int `xml:"remaining_games,omitempty"`
	LiveGames      int `xml:"live_games,omitempty"`
	CompletedGames int `xml:"completed_games,omitempty"`
}

// TeamStandings contains information about a Team's ranking in their league.
type TeamStandings struct {
	Rank                    int                      `xml:"rank,omitempty"`
	PlayoffSeed             int                      `xml:"playoff_seed,omitempty"`
	GamesBack               string                   `xml:"games_back,omitempty"`
	OutcomeTotals           *OutcomeTotals           `xml:"outcome_totals,omitempty"`
	DivisionalOutcomeTotals *DivisionalOutcomeTotals `xml:"divisional_outcome_totals,omitempty"`
}

// OutcomeTotals contains information on the outcomes of a Team's matchups.
type OutcomeTotals struct {
	Wins       int     `xml:"wins,omitempty"`
	Losses     int     `xml:"losses,omitempty"`
	Ties       int     `xml:"ties,omitempty"`
	Percentage float32 `xml:"percentage,omitempty"`
}

// DivisionalOutcomeTotals contains information on the outcomes of a Team's matchups in their division.
type DivisionalOutcomeTotals struct {
	Wins   int `xml:"wins,omitempty"`
	Losses int `xml:"losses,omitempty"`
	Ties   int `xml:"ties,omitempty"`
}

// Roster contains information on a Team's roster.
type Roster struct {
	CoverageType string   `xml:"coverage_type,omitempty"`
	Date         string   `xml:"date,omitempty"`
	IsEditable   bool     `xml:"is_editable,omitempty"`
	Players      *Players `xml:"players,omitempty"`
}

// RosterPositions contains a list of roster positions .
type RosterPositions struct {
	RosterPosition []RosterPosition `xml:"roster_position,omitempty"`
}

// RosterPosition contains information about roster position.
type RosterPosition struct {
	Position           string `xml:"position,omitempty"`
	PositionType       string `xml:"position_type,omitempty"`
	Count              int    `xml:"count,omitempty"`
	Abbreviation       string `xml:"abbreviation,omitempty"`
	DisplayName        string `xml:"display_name,omitempty"`
	IsBench            bool   `xml:"is_bench,omitempty"`
	IsStartingPosition bool   `xml:"is_starting_position,omitempty"`
}

// Settings contains information about a league's settings.
type Settings struct {
	DraftType                  string           `xml:"draft_type,omitempty"`
	ScoringType                string           `xml:"scoring_type,omitempty"`
	UsesPlayoff                bool             `xml:"uses_playoff,omitempty"`
	PlayoffStartWeek           int              `xml:"playoff_start_week,omitempty"`
	UsesPlayoffReseeding       bool             `xml:"uses_playoff_reseeding,omitempty"`
	UsesLockEliminatedTeams    bool             `xml:"uses_lock_eliminated_teams,omitempty"`
	UsesFaab                   bool             `xml:"uses_faab,omitempty"`
	TradeEndDate               string           `xml:"trade_end_date,omitempty"`
	TradeRatifyType            string           `xml:"trade_ratify_type,omitempty"`
	TradeRejectTime            int              `xml:"trade_reject_time,omitempty"`
	IsAuctionDraft             bool             `xml:"is_auction_draft,omitempty"`
	PersistentURL              string           `xml:"persistent_url,omitempty"`
	HasPlayoffConsolationGames bool             `xml:"has_playoff_consolation_games,omitempty"`
	NumPlayoffTeams            int              `xml:"num_playoff_teams,omitempty"`
	NumPlayoffConsolationTeams int              `xml:"num_playoff_consolation_teams,omitempty"`
	HasMultiweekChampionship   bool             `xml:"has_multiweek_championship,omitempty"`
	UsesRosterImport           bool             `xml:"uses_roster_import,omitempty"`
	RosterImportDeadline       string           `xml:"roster_import_deadline,omitempty"`
	WaiverType                 string           `xml:"waiver_type,omitempty"`
	WaiverRule                 string           `xml:"waiver_rule,omitempty"`
	DraftTime                  string           `xml:"draft_time,omitempty"`
	DraftPickTime              string           `xml:"draft_pick_time,omitempty"`
	PostDraftPlayers           string           `xml:"post_draft_players,omitempty"`
	MaxTeams                   int              `xml:"max_teams,omitempty"`
	WaiverTime                 string           `xml:"waiver_time,omitempty"`
	PlayerPool                 string           `xml:"player_pool,omitempty"`
	CantCutList                string           `xml:"cant_cut_list,omitempty"`
	DraftTogether              string           `xml:"draft_together,omitempty"`
	CanTradeDraftPicks         bool             `xml:"can_trade_draft_picks,omitempty"`
	SendbirdChannelURL         string           `xml:"sendbird_channel_url,omitempty"`
	MaxWeeklyAdds              int              `xml:"max_weekly_adds,omitempty"`
	RosterPositions            *RosterPositions `xml:"roster_positions,omitempty"`
	StatCategories             *StatCategories  `xml:"stat_categories,omitempty"`
	StatModifiers              *StatModifiers   `xml:"stat_modifiers,omitempty"`
	Divisions                  *Divisions       `xml:"divisions,omitempty"`
}

// Divisions is a list of divisions.
type Divisions struct {
	Division []Division `xml:"division,omitempty"`
}

// Division contains information about a division.
type Division struct {
	DivisionID int    `xml:"division_id,omitempty"`
	Name       string `xml:"name,omitempty"`
}

// PositionsTypes is a list of position types.
type PositionsTypes struct {
	PositionType *PositionType `xml:"position_type,omitempty"`
}

// PositionType contains information on a position type.
type PositionType struct {
	Type        string `xml:"type,omitempty"`
	DisplayName string `xml:"display_name,omitempty"`
}

// GameWeeks is a list of game weeks.
type GameWeeks struct {
	Count    int       `xml:"count,attr,omitempty"`
	GameWeek *GameWeek `xml:"game_week,omitempty"`
}

// GameWeek contains information on a week in a game.
type GameWeek struct {
	Week        int    `xml:"week,omitempty"`
	DisplayName string `xml:"display_name,omitempty"`
	Start       string `xml:"start,omitempty"`
	End         string `xml:"end,omitempty"`
}

// PercentOwned contains information on the percent of teams that own a player
// across Yahoo leagues.
type PercentOwned struct {
	CoverageType string `xml:"coverage_type,omitempty"`
	Week         int    `xml:"week,omitempty"`
	Value        int    `xml:"value,omitempty"`
	Delta        int    `xml:"delta,omitempty"`
}

// DraftAnalysis contains information on how a player was drafted across Yahoo leagues.
type DraftAnalysis struct {
	AveragePick    string `xml:"average_pick,omitempty"`
	AverageRound   string `xml:"average_round,omitempty"`
	AverageCost    string `xml:"average_cost,omitempty"`
	PercentDrafted string `xml:"percent_drafted,omitempty"`
}

// Ownership contains information on the ownership status of player within a league.
type Ownership struct {
	OwnershipType string `xml:"ownership_type,omitempty"`
	OwnerTeamKey  string `xml:"owner_team_key,omitempty"`
	OwnerTeamName string `xml:"owner_team_name,omitempty"`
	WaiverDate    string `xml:"waiver_date,omitempty"`
	DisplayDate   string `xml:"display_date,omitempty"`
	Team          *Team  `xml:"team,omitempty"`
}

// DraftResults is a list of draft results from a league's draft.
type DraftResults struct {
	Count       int          `xml:"count,attr,omitempty"`
	DraftResult *DraftResult `xml:"draft_result,omitempty"`
}

// DraftResult contains information about a draft pick from a league's draft.
type DraftResult struct {
	Pick      string `xml:"pick,omitempty"`
	Round     string `xml:"round,omitempty"`
	TeamKey   string `xml:"team_key,omitempty"`
	PlayerKey string `xml:"player_key,omitempty"`
}

// RecommendedTradePartners is a list of recommended trade partners.
type RecommendedTradePartners struct {
	RecommendedTradePartner []RecommendedTradePartner `xml:"recommended_trade_partner,omitempty"`
}

// RecommendedTradePartner is a trade partner recommended by Yahoo.
type RecommendedTradePartner struct {
	TeamKey string `xml:"team_key,omitempty"`
}

// PositionalRanks is a list of positional ranks.
type PositionalRanks struct {
	PositionalRank *PositionalRank `xml:"positional_rank,omitempty"`
}

// PositionalRank holds the team's comparitive rank information for a postion.
type PositionalRank struct {
	Position        string           `xml:"position,omitempty"`
	Rank            string           `xml:"rank,omitempty"`
	Grade           string           `xml:"grade,omitempty"`
	StartingPlayers *StartingPlayers `xml:"starting_players,omitempty"`
	BenchPlayers    *BenchPlayers    `xml:"bench_players,omitempty"`
}

// StartingPlayers is a list of starting players.
type StartingPlayers struct {
	Player []Player `xml:"player,omitempty"`
}

// BenchPlayers is a list of bench players.
type BenchPlayers struct {
	Player []Player `xml:"player,omitempty"`
}
