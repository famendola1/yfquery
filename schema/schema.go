package schema

// FantasyContent models the top-level structure returned by the Yahoo Fantasy API.
type FantasyContent struct {
	Lang        string  `xml:"lang,attr"`
	URI         string  `xml:"uri,attr"`
	Time        string  `xml:"time,attr"`
	Copyright   string  `xml:"copyright,attr"`
	RefreshRate string  `xml:"refresh_rate,attr"`
	Yahoo       string  `xml:"yahoo,attr"`
	Xmlns       string  `xml:"xmlns,attr"`
	Games       Games   `xml:"games"`
	Game        Game    `xml:"game"`
	Leagues     Leagues `xml:"leagues"`
	League      League  `xml:"league"`
	Teams       Teams   `xml:"teams"`
	Team        Team    `xml:"team"`
	Users       Users   `xml:"users"`
}

// Users is a list of users.
type Users struct {
	Count int  `xml:"count,attr"`
	User  User `xml:"user"`
}

// User models a user in Users.
type User struct {
	GUID    string  `xml:"guid"`
	Games   Games   `xml:"games"`
	Leagues Leagues `xml:"leagues"`
	Teams   Teams   `xml:"teams"`
}

// Player models a Yahoo player.
type Player struct {
	PlayerKey                string              `xml:"player_key"`
	PlayerID                 int                 `xml:"player_id"`
	Name                     Name                `xml:"name"`
	EditorialPlayerKey       string              `xml:"editorial_player_key"`
	EditorialTeamKey         string              `xml:"editorial_team_key"`
	EditorialTeamFullName    string              `xml:"editorial_team_full_name"`
	EditorialTeamAbbr        string              `xml:"editorial_team_abbr"`
	UniformNumber            int                 `xml:"uniform_number"`
	DisplayPosition          string              `xml:"display_position"`
	Headshot                 Headshot            `xml:"headshot"`
	ImageURL                 string              `xml:"image_url"`
	IsUndroppable            bool                `xml:"is_undroppable"`
	PositionType             string              `xml:"position_type"`
	PrimaryPosition          string              `xml:"primary_position"`
	EligiblePositions        EligiblePositions   `xml:"eligible_positions"`
	HasPlayerNotes           bool                `xml:"has_player_notes"`
	PlayerNotesLastTimestamp uint64              `xml:"player_notes_last_timestamp"`
	TransactionData          TransactionData     `xml:"transaction_data"`
	PlayerStats              PlayerStats         `xml:"player_stats"`
	PlayerAdvancedStats      PlayerAdvancedStats `xml:"player_advanced_stats"`
	Status                   string              `xml:"status"`
	StatusFull               string              `xml:"status_full"`
	InjuryNote               string              `xml:"injury_note"`
	HasRecentPlayerNotes     string              `xml:"has_recent_player_notes"`
	SelectedPosition         SelectedPosition    `xml:"selected_position"`
	IsKeeper                 IsKeeper            `xml:"is_keeper"`
	PercentOwned             PercentOwned        `xml:"percent_owned"`
	DraftAnalysis            DraftAnalysis       `xml:"draft_analysis"`
	Ownership                Ownership           `xml:"ownership"`
}

// Name contains information about a Player's name.
type Name struct {
	Full       string `xml:"full"`
	First      string `xml:"first"`
	Last       string `xml:"last"`
	ASCIIFirst string `xml:"ascii_first"`
	ASCIILast  string `xml:"ascii_last"`
}

// Headshot contains information about a Player's Yahoo photo.
type Headshot struct {
	URL  string `xml:"url"`
	Size string `xml:"size"`
}

// EligiblePositions contains all the eligible positions a Player can play.
type EligiblePositions struct {
	Position []string `xml:"position"`
}

// TransactionData contains transaction information for a Player.
type TransactionData struct {
	Type                string `xml:"type"`
	SourceType          string `xml:"source_type"`
	DestinationType     string `xml:"destination_type"`
	DestinationTeamKey  string `xml:"destination_team_key"`
	DestinationTeamName string `xml:"destination_team_name"`
	SourceTeamKey       string `xml:"source_team_key"`
	SourceTeamName      string `xml:"source_team_name"`
}

// PlayerStats contains stat information for a Player.
type PlayerStats struct {
	CoverageType string `xml:"coverage_type"`
	Season       string `xml:"season"`
	Date         string `xml:"date"`
	Stats        Stats  `xml:"stats"`
}

// PlayerAdvancedStats contains stat information for a Player.
type PlayerAdvancedStats struct {
	CoverageType string `xml:"coverage_type"`
	Season       string `xml:"season"`
	Stats        Stats  `xml:"stats"`
}

// SelectedPosition contains information on a Player's selected position.
type SelectedPosition struct {
	CoverageType string `xml:"coverage_type"`
	Date         string `xml:"date"`
	Position     string `xml:"position"`
	IsFlex       bool   `xml:"is_flex"`
}

// IsKeeper contains keeper information for a Player.
type IsKeeper struct {
	Status string `xml:"status"`
	Cost   string `xml:"cost"`
	Kept   string `xml:"kept"`
}

// Players is a list of players.
type Players struct {
	Count  int      `xml:"count,attr"`
	Player []Player `xml:"player"`
}

// Transactions is a list of transactions.
type Transactions struct {
	Transaction Transaction `xml:"transactions"`
}

// Transaction represents a Yahoo fantasy transaction.
type Transaction struct {
	TransactionKey    string  `xml:"transaction_key"`
	TransactionID     int     `xml:"transaction_id"`
	Type              string  `xml:"type"`
	Status            string  `xml:"status"`
	Timestamp         uint64  `xml:"timestamp"`
	WaiverPlayerKey   string  `xml:"waiver_player_key"`
	WaiverTeamKey     string  `xml:"waiver_team_key"`
	WaiverDate        string  `xml:"waiver_date"`
	WaiverPriority    int     `xml:"waiver_priority"`
	TraderTeamKey     string  `xml:"trader_team_key"`
	TradeeTeamKey     string  `xml:"tradee_team_key"`
	TradeProposedTime uint64  `xml:"trade_proposed_time"`
	TradeNode         string  `xml:"trade_note"`
	VoterTeamKey      string  `xml:"voter_team_key"`
	Players           Players `xml:"players"`
}

// Matchups is a list of matchups.
type Matchups struct {
	Matchup []Matchup `xml:"matchup"`
}

// Matchup represents a Yahoo matchup.
type Matchup struct {
	Week          int         `xml:"week"`
	WeekStart     string      `xml:"week_start"`
	WeekEnd       string      `xml:"week_end"`
	Status        string      `xml:"status"`
	IsPlayoffs    bool        `xml:"is_playoffs"`
	IsConsolation bool        `xml:"is_consolation"`
	IsTied        bool        `xml:"is_tied"`
	WinnerTeamKey string      `xml:"winner_team_key"`
	StatWinners   StatWinners `xml:"stat_winners"`
	Teams         Teams       `xml:"teams"`
}

// StatWinners is a list of stat winners.
type StatWinners struct {
	StatWinner []StatWinner `xml:"stat_winner"`
}

// StatWinner represents the winner of a stat category.
type StatWinner struct {
	StatID        int    `xml:"stat_id"`
	WinnerTeamKey string `xml:"winner_team_key"`
	IsTied        bool   `xml:"is_tied"`
}

// StatCategories contains the enabled stat categories for a league.
type StatCategories struct {
	Stats Stats `xml:"stats"`
}

// StatModifiers contains the modifiers for stats.
type StatModifiers struct {
	Stats Stats `xml:"stats"`
}

// StatPositionTypes is a list of stat position types.
type StatPositionTypes struct {
	StatPositionType []StatPositionType `xml:"stat_position_type"`
}

// StatPositionType contains information about a stat position type.
type StatPositionType struct {
	PositionType      string `xml:"position_type"`
	IsOnlyDisplayStat string `xml:"is_only_display_stat"`
}

// Stats is a list of stats.
type Stats struct {
	Stat []Stat `xml:"stat"`
}

// Stat represents a stat category in Yahoo.
type Stat struct {
	StatID            int               `xml:"stat_id"`
	Value             string            `xml:"value"`
	Enabled           bool              `xml:"enabled"`
	Name              string            `xml:"name"`
	DisplayName       string            `xml:"display_name"`
	Group             string            `xml:"group"`
	Abbr              string            `xml:"abbr"`
	SortOrder         string            `xml:"sort_order"`
	PositionType      string            `xml:"position_type"`
	PositionTypes     PositionsTypes    `xml:"position_types"`
	IsOnlyDisplayStat bool              `xml:"is_only_display_stat"`
	StatPositionTypes StatPositionTypes `xml:"stat_position_types"`
	Groups            Groups            `xml:"groups"`
}

// Groups is a list of groups.
type Groups struct {
	Group []Group `xml:"group"`
}

// Group contains information on a group.
type Group struct {
	GroupName        string `xml:"group_name"`
	GroupDisplayName string `xml:"group_display_name"`
	GroupAbbr        string `xml:"group_abbr"`
}

// Game represents a Yahoo game
type Game struct {
	GameKey            string          `xml:"game_key"`
	GameID             int             `xml:"game_id"`
	Name               string          `xml:"name"`
	Code               string          `xml:"code"`
	Type               string          `xml:"type"`
	URL                string          `xml:"url"`
	Season             string          `xml:"season"`
	IsRegistrationOver bool            `xml:"is_registration_over"`
	IsGameOver         bool            `xml:"is_game_over"`
	IsOffseason        bool            `xml:"is_offseason"`
	RosterPositions    RosterPositions `xml:"roster_positions"`
	PositionsTypes     PositionsTypes  `xml:"position_types"`
	GameWeeks          GameWeeks       `xml:"game_weeks"`
	StatCategories     StatCategories  `xml:"stat_categories"`
}

// Games is a list Games.
type Games struct {
	Count int    `xml:"count,attr"`
	Game  []Game `xml:"game"`
}

// League represents a Yahoo league.
type League struct {
	LeagueKey             string       `xml:"league_key"`
	LeagueID              int          `xml:"league_id"`
	Name                  string       `xml:"name"`
	URL                   string       `xml:"url"`
	LogoURL               string       `xml:"logo_url"`
	DraftStatus           string       `xml:"draft_status"`
	NumTeams              int          `xml:"num_teams"`
	EditKey               string       `xml:"edit_key"`
	WeeklyDeadline        string       `xml:"weekly_deadline"`
	LeagueUpdateTimestamp string       `xml:"league_update_timestamp"`
	ScoringType           string       `xml:"scoring_type"`
	LeagueType            string       `xml:"league_type"`
	Renew                 string       `xml:"renew"`
	ShortInvitationURL    string       `xml:"short_invitation_url"`
	AllowAddToDlExtraPos  string       `xml:"allow_add_to_dl_extra_pos"`
	IsProLeague           bool         `xml:"is_pro_league"`
	IsCashLeague          bool         `xml:"is_cash_league"`
	CurrentWeek           int          `xml:"current_week"`
	StartWeek             int          `xml:"start_week"`
	StartDate             string       `xml:"start_date"`
	EndWeek               int          `xml:"end_week"`
	EndDate               string       `xml:"end_date"`
	GameCode              string       `xml:"game_code"`
	Season                string       `xml:"season"`
	IsFinished            bool         `xml:"is_finished"`
	Standings             Standings    `xml:"standings"`
	Teams                 Teams        `xml:"teams"`
	Players               Players      `xml:"players"`
	Scoreboard            Scoreboard   `xml:"scoreboard"`
	Settings              Settings     `xml:"settings"`
	DraftResults          DraftResults `xml:"draft_results"`
	Transactions          Transactions `xml:"transactions"`
}

// Leagues is a list Leagues.
type Leagues struct {
	Count  int      `xml:"count,attr"`
	League []League `xml:"league"`
}

// Scoreboard contains all the matchups for a league.
type Scoreboard struct {
	Count    int      `xml:"count,attr"`
	Matchups Matchups `xml:"matchups"`
}

// Standings contains the standings for a league.
type Standings struct {
	Teams Teams `xml:"teams"`
}

// Teams is a list of teams.
type Teams struct {
	Count int    `xml:"count,attr"`
	Team  []Team `xml:"team"`
}

// Team represents a Yahoo team.
type Team struct {
	TeamKey               string             `xml:"team_key"`
	TeamID                int                `xml:"team_id"`
	Name                  string             `xml:"name"`
	IsOwnedByCurrentLogin bool               `xml:"is_owned_by_current_login"`
	URL                   string             `xml:"url"`
	TeamLogos             TeamLogos          `xml:"team_logos"`
	WaiverPriority        int                `xml:"waiver_priority"`
	NumberOfMoves         int                `xml:"number_of_moves"`
	NumberOfTrades        int                `xml:"number_of_trades"`
	RosterAdds            RosterAdds         `xml:"roster_adds"`
	LeagueScoringType     string             `xml:"league_scoring_type"`
	DraftPosition         int                `xml:"draft_position"`
	HasDraftGrade         bool               `xml:"has_draft_grade"`
	Managers              Managers           `xml:"managers"`
	TeamStats             TeamStats          `xml:"team_stats"`
	TeamPoints            TeamPoints         `xml:"team_points"`
	TeamRemainingGames    TeamRemainingGames `xml:"team_remaining_games"`
	ClinchedPlayoffs      bool               `xml:"clinched_playoffs"`
	TeamStandings         TeamStandings      `xml:"team_standings"`
	Roster                Roster             `xml:"roster"`
	Matchups              Matchups           `xml:"matchups"`
}

// TeamLogos contains a list of team logos.
type TeamLogos struct {
	TeamLogo []TeamLogo `xml:"team_logo"`
}

// TeamLogo contains information on the logo for a team.
type TeamLogo struct {
	Size string `xml:"size"`
	URL  string `xml:"url"`
}

// RosterAdds contains roster add informatin for a team.
type RosterAdds struct {
	CoverageType  string `xml:"coverage_type"`
	CoverageValue int    `xml:"coverage_value"`
	Value         int    `xml:"value"`
}

// Managers is a list of managers.
type Managers struct {
	Count   int       `xml:"count,attr"`
	Manager []Manager `xml:"manager"`
}

// Manager contains information about the manager of a team.
type Manager struct {
	ManagerID      int    `xml:"manager_id"`
	Nickname       string `xml:"nickname"`
	GUID           string `xml:"guid"`
	IsCommissioner bool   `xml:"is_commissioner"`
	IsCurrentLogin bool   `xml:"is_current_login"`
	Email          string `xml:"email"`
	ImageURL       string `xml:"image_url"`
	FeloScore      int    `xml:"felo_score"`
	FeloTier       string `xml:"felo_tier"`
}

// TeamStats contains stats for a team.
type TeamStats struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Stats        Stats  `xml:"stats"`
}

// TeamPoints contaons the points scored by a team.
type TeamPoints struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Total        int    `xml:"total"`
}

// TeamRemainingGames contains information on the remaining games of a team.
type TeamRemainingGames struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Total        Total  `xml:"total"`
}

// Total contains information on remaining games.
type Total struct {
	RemainingGames int `xml:"remaining_games"`
	LiveGames      int `xml:"live_games"`
	CompletedGames int `xml:"completed_games"`
}

// TeamStandings contains information about a Team's ranking in their league.
type TeamStandings struct {
	Rank                    int                     `xml:"rank"`
	OutcomeTotals           OutcomeTotals           `xml:"outcome_totals"`
	DivisionalOutcomeTotals DivisionalOutcomeTotals `xml:"divisional_outcome_totals"`
}

// OutcomeTotals contains information on the outcomes of a Team's matchups.
type OutcomeTotals struct {
	Wins       int     `xml:"wins"`
	Losses     int     `xml:"losses"`
	Ties       int     `xml:"ties"`
	Percentage float32 `xml:"percentage"`
}

// DivisionalOutcomeTotals contains information on the outcomes of a Team's matchups in their division.
type DivisionalOutcomeTotals struct {
	Wins   int `xml:"wins"`
	Losses int `xml:"losses"`
	Ties   int `xml:"ties"`
}

// Roster contains information on a Team's roster.
type Roster struct {
	CoverageType string  `xml:"coverage_type"`
	Date         string  `xml:"date"`
	IsEditable   bool    `xml:"is_editable"`
	Players      Players `xml:"players"`
}

// RosterPositions contains a list of roster positions .
type RosterPositions struct {
	RosterPosition []RosterPosition `xml:"roster_position"`
}

// RosterPosition contains information about roster position.
type RosterPosition struct {
	Position           string `xml:"position"`
	PositionType       string `xml:"position_type"`
	Count              int    `xml:"count"`
	Abbreviation       string `xml:"abbreviation"`
	DisplayName        string `xml:"display_name"`
	IsBench            bool   `xml:"is_bench"`
	IsStartingPosition bool   `xml:"is_starting_position"`
}

// Settings contains information about a league's settings.
type Settings struct {
	DraftType                  string          `xml:"draft_type"`
	ScoringType                string          `xml:"scoring_type"`
	UsesPlayoff                bool            `xml:"uses_playoff"`
	PlayoffStartWeek           int             `xml:"playoff_start_week"`
	UsesPlayoffReseeding       bool            `xml:"uses_playoff_reseeding"`
	UsesLockEliminatedTeams    bool            `xml:"uses_lock_eliminated_teams"`
	UsesFaab                   bool            `xml:"uses_faab"`
	TradeEndDate               string          `xml:"trade_end_date"`
	TradeRatifyType            string          `xml:"trade_ratify_type"`
	TradeRejectTime            int             `xml:"trade_reject_time"`
	IsAuctionDraft             bool            `xml:"is_auction_draft"`
	PersistentURL              string          `xml:"persistent_url"`
	HasPlayoffConsolationGames bool            `xml:"has_playoff_consolation_games"`
	NumPlayoffTeams            int             `xml:"num_playoff_teams"`
	NumPlayoffConsolationTeams int             `xml:"num_playoff_consolation_teams"`
	HasMultiweekChampionship   bool            `xml:"has_multiweek_championship"`
	UsesRosterImport           bool            `xml:"uses_roster_import"`
	RosterImportDeadline       string          `xml:"roster_import_deadline"`
	WaiverType                 string          `xml:"waiver_type"`
	WaiverRule                 string          `xml:"waiver_rule"`
	DraftTime                  string          `xml:"draft_time"`
	DraftPickTime              string          `xml:"draft_pick_time"`
	PostDraftPlayers           string          `xml:"post_draft_players"`
	MaxTeams                   int             `xml:"max_teams"`
	WaiverTime                 string          `xml:"waiver_time"`
	PlayerPool                 string          `xml:"player_pool"`
	CantCutList                string          `xml:"cant_cut_list"`
	DraftTogether              string          `xml:"draft_together"`
	CanTradeDraftPicks         bool            `xml:"can_trade_draft_picks"`
	SendbirdChannelURL         string          `xml:"sendbird_channel_url"`
	MaxWeeklyAdds              int             `xml:"max_weekly_adds"`
	RosterPositions            RosterPositions `xml:"roster_positions"`
	StatCategories             StatCategories  `xml:"stat_categories"`
	StatModifiers              StatModifiers   `xml:"stat_modifiers"`
	Divisions                  Divisions       `xml:"divisions"`
}

// Divisions is a list of divisions.
type Divisions struct {
	Division []Division `xml:"division"`
}

// Division contains information about a division.
type Division struct {
	DivisionID int    `xml:"division_id"`
	Name       string `xml:"name"`
}

// PositionsTypes is a list of position types.
type PositionsTypes struct {
	PositionType PositionType `xml:"position_type"`
}

// PositionType contains information on a position type.
type PositionType struct {
	Type        string `xml:"type"`
	DisplayName string `xml:"display_name"`
}

// GameWeeks is a list of game weeks.
type GameWeeks struct {
	Count    int      `xml:"count,attr"`
	GameWeek GameWeek `xml:"game_week"`
}

// GameWeek contains information on a week in a game.
type GameWeek struct {
	Week        int    `xml:"week"`
	DisplayName string `xml:"display_name"`
	Start       string `xml:"start"`
	End         string `xml:"end"`
}

// PercentOwned contains information on the percent of teams that own a player
// across Yahoo leagues.
type PercentOwned struct {
	CoverageType string `xml:"coverage_type"`
	Week         int    `xml:"week"`
	Value        int    `xml:"value"`
	Delta        int    `xml:"delta"`
}

// DraftAnalysis contains information on how a player was drafted across Yahoo leagues.
type DraftAnalysis struct {
	AveragePick    string `xml:"average_pick"`
	AverageRound   string `xml:"average_round"`
	AverageCost    string `xml:"average_cost"`
	PercentDrafted string `xml:"percent_drafted"`
}

// Ownership contains information on the ownership status of player within a league.
type Ownership struct {
	OwnershipType string `xml:"ownership_type"`
	OwnerTeamKey  string `xml:"owner_team_key"`
	OwnerTeamName string `xml:"owner_team_name"`
	Team          Team   `xml:"team"`
}

// DraftResults is a list of draft results from a league's draft.
type DraftResults struct {
	Count       int         `xml:"count,attr"`
	DraftResult DraftResult `xml:"draft_result"`
}

// DraftResult contains information about a draft pick from a league's draft.
type DraftResult struct {
	Pick      string `xml:"pick"`
	Round     string `xml:"round"`
	TeamKey   string `xml:"team_key"`
	PlayerKey string `xml:"player_key"`
}
