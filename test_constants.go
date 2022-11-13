package yfquery

const (
	rosterTestResp = `
  <?xml version="1.0" encoding="UTF-8"?>
  <fantasy_content xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng" xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/team/253.l.102614.t.10/roster/players" time="110.02206802368ms" copyright="Data provided by Yahoo! and STATS, LLC">
    <team>
      <team_key>253.l.102614.t.10</team_key>
      <team_id>10</team_id>
      <name>Matt Dzaman</name>
      <url>https://baseball.fantasysports.yahoo.com/b1/102614/10</url>
      <team_logos>
        <team_logo>
          <size>medium</size>
          <url>https://l.yimg.com/a/i/us/sp/fn/mlb/gr/icon_12_2.gif</url>
        </team_logo>
      </team_logos>
      <managers>
        <manager>
          <manager_id>10</manager_id>
          <nickname>Sean Montgomery</nickname>
          <guid>VZVEVUCLSJAHSM73FMJ4BYFIKU</guid>
          <is_current_login>1</is_current_login>
        </manager>
      </managers>
      <roster>
        <coverage_type>date</coverage_type>
        <date>2011-07-22</date>
        <players count="3">
          <player>
            <player_key>253.p.7569</player_key>
          </player>
          <player>
            <player_key>253.p.7054</player_key>
          </player>
          <player>
            <player_key>253.p.7382</player_key>
          </player>
        </players>
      </roster>
    </team>
  </fantasy_content>`

	standingsTestResp = `
  <?xml version="1.0" encoding="UTF-8"?>
  <fantasy_content xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng" xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/league/223.l.431/standings" time="201.46489143372ms" copyright="Data provided by Yahoo! and STATS, LLC">
    <league>
      <league_key>223.l.431</league_key>
      <league_id>431</league_id>
      <name>Y! Friends and Family League</name>
      <url>https://football.fantasysports.yahoo.com/archive/pnfl/2009/431</url>
      <draft_status>postdraft</draft_status>
      <num_teams>4</num_teams>
      <edit_key>17</edit_key>
      <weekly_deadline/>
      <league_update_timestamp>1262595518</league_update_timestamp>
      <scoring_type>head</scoring_type>
      <current_week>16</current_week>
      <start_week>1</start_week>
      <end_week>16</end_week>
      <is_finished>1</is_finished>
      <standings>
        <teams count="4">
          <team>
            <team_key>223.l.431.t.10</team_key>
          </team>
          <team>
            <team_key>223.l.431.t.5</team_key>
          </team>
          <team>
            <team_key>223.l.431.t.8</team_key>
          </team>
          <team>
            <team_key>223.l.431.t.12</team_key>
          </team>
        </teams>
      </standings>
    </league>
  </fantasy_content>`

	searchTestResp = `
  <?xml version="1.0" encoding="UTF-8"?>
  <fantasy_content xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/league/410.l.16883/players;search=Jalen Green" time="25.722980499268ms" copyright="Data provided by Yahoo! and STATS, LLC" refresh_rate="60" xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng">
   <league>
    <league_key>410.l.16883</league_key>
    <league_id>16883</league_id>
    <name>NBA Fantasy 2K22</name>
    <season>2021</season>
    <players count="1">
     <player>
      <player_key>410.p.6513</player_key>
     </player>
    </players>
   </league>
  </fantasy_content>
  `

	gameTestResp = `<?xml version="1.0" encoding="UTF-8"?>
     <fantasy_content xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/game/nba" xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" time="30.575037002563ms" copyright="Data provided by Yahoo! and STATS, LLC" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng">
      <game>
        <game_key>410</game_key>
        <game_id>410</game_id>
        <name>Basketball</name>
        <code>nba</code>
        <type>full</type>
        <url>https://football.fantasysports.yahoo.com/f1</url>
        <season>2021</season>
      </game>
    </fantasy_content>`

	leagueTestResp = `<?xml version="1.0" encoding="UTF-8"?>
  <fantasy_content xml:lang="en-US" yahoo:uri="http://fantasysports.yahooapis.com/fantasy/v2/users;use_login=1/games;game_keys=nba/leagues" time="37.668943405151ms" copyright="Data provided by Yahoo! and STATS, LLC" refresh_rate="60" xmlns:yahoo="http://www.yahooapis.com/v1/base.rng" xmlns="http://fantasysports.yahooapis.com/fantasy/v2/base.rng">
   <users count="1">
    <user>
     <guid>EKFDPDVSJIGZD64VAL6WYCIH2I</guid>
     <games count="1">
      <game>
       <game_key>410</game_key>
       <game_id>410</game_id>
       <name>Basketball</name>
       <code>nba</code>
       <type>full</type>
       <url>https://basketball.fantasysports.yahoo.com/nba</url>
       <season>2021</season>
       <is_registration_over>0</is_registration_over>
       <is_game_over>0</is_game_over>
       <is_offseason>0</is_offseason>
       <leagues count="3">
        <league>
         <league_key>410.l.16883</league_key>
        </league>
        <league>
         <league_key>410.l.61777</league_key>
        </league>
        <league>
         <league_key>410.l.159928</league_key>
        </league>
       </leagues>
      </game>
     </games>
    </user>
   </users>
  </fantasy_content>`

	leagueFullTestResp = `<league>
    <league_key>410.l.16883</league_key>
    <league_id>16883</league_id>
    <name>NBA Fantasy 2K22</name>
    <url>https://basketball.fantasysports.yahoo.com/nba/16883</url>
    <logo_url>https://yahoofantasysports-res.cloudinary.com/image/upload/t_s192sq/fantasy-logos/0743c1232a845a00b408b165a683c19ad0ee273236a45766137c9698234246bf.jpg</logo_url>
    <draft_status>postdraft</draft_status>
    <num_teams>12</num_teams>
    <edit_key>2021-11-09</edit_key>
    <weekly_deadline>intraday</weekly_deadline>
    <league_update_timestamp>1636441834</league_update_timestamp>
    <scoring_type>headone</scoring_type>
    <league_type>private</league_type>
    <renew>402_22765</renew>
    <short_invitation_url>https://basketball.fantasysports.yahoo.com/nba/16883/invitation?key=1f14749a8282d491&amp;ikey=bc6fb9e93fd791bb</short_invitation_url>
    <allow_add_to_dl_extra_pos>1</allow_add_to_dl_extra_pos>
    <is_pro_league>0</is_pro_league>
    <is_cash_league>0</is_cash_league>
    <current_week>4</current_week>
    <start_week>1</start_week>
    <start_date>2021-10-19</start_date>
    <end_week>24</end_week>
    <end_date>2022-04-10</end_date>
    <game_code>nba</game_code>
    <season>2021</season>
   </league>
  `

	teamFullTestResp = `
	<team>
    <team_key>410.l.16883.t.1</team_key>
    <team_id>1</team_id>
    <name>Bring Me A Shot</name>
    <is_owned_by_current_login>1</is_owned_by_current_login>
    <url>https://basketball.fantasysports.yahoo.com/nba/16883/1</url>
    <team_logos>
     <team_logo>
      <size>large</size>
      <url>https://yahoofantasysports-res.cloudinary.com/image/upload/t_s192sq/fantasy-logos/6d1c4b1ab7f7d94e33ea9f4a3306381954a2880b244fe54ca8dfea504c7be242.jpg</url>
     </team_logo>
    </team_logos>
    <waiver_priority>9</waiver_priority>
    <number_of_moves>23</number_of_moves>
    <number_of_trades>0</number_of_trades>
    <roster_adds>
     <coverage_type>week</coverage_type>
     <coverage_value>4</coverage_value>
     <value>3</value>
    </roster_adds>
    <league_scoring_type>headone</league_scoring_type>
    <draft_position>4</draft_position>
    <has_draft_grade>0</has_draft_grade>
    <managers>
     <manager>
      <manager_id>1</manager_id>
      <nickname>Fabio</nickname>
      <is_commissioner>1</is_commissioner>
      <is_current_login>1</is_current_login>
      <email>example@gmail.com</email>
      <image_url>https://s.yimg.com/ag/images/default_user_profile_pic_64sq.jpg</image_url>
      <felo_score>890</felo_score>
      <felo_tier>platinum</felo_tier>
     </manager>
    </managers>
   </team>
	`

	playerFullTestResp = `
	<player>
	 <player_key>410.p.6065</player_key>
	 <player_id>6065</player_id>
	 <name>
		<full>Shake Milton</full>
		<first>Shake</first>
		<last>Milton</last>
		<ascii_first>Shake</ascii_first>
		<ascii_last>Milton</ascii_last>
	 </name>
	 <editorial_player_key>nba.p.6065</editorial_player_key>
	 <editorial_team_key>nba.t.20</editorial_team_key>
	 <editorial_team_full_name>Philadelphia 76ers</editorial_team_full_name>
	 <editorial_team_abbr>PHI</editorial_team_abbr>
	 <uniform_number>18</uniform_number>
	 <display_position>PG,SG</display_position>
	 <headshot>
		<url>https://s.yimg.com/iu/api/res/1.2/PTF3UNtaGJzwH3Ah22R0Ow--~C/YXBwaWQ9eXNwb3J0cztjaD0yMzM2O2NyPTE7Y3c9MTc5MDtkeD04NTc7ZHk9MDtmaT11bGNyb3A7aD02MDtxPTEwMDt3PTQ2/https://s.yimg.com/xe/i/us/sp/v/nba_cutout/players_l/10142021/6065.png</url>
		<size>small</size>
	 </headshot>
	 <image_url>https://s.yimg.com/iu/api/res/1.2/PTF3UNtaGJzwH3Ah22R0Ow--~C/YXBwaWQ9eXNwb3J0cztjaD0yMzM2O2NyPTE7Y3c9MTc5MDtkeD04NTc7ZHk9MDtmaT11bGNyb3A7aD02MDtxPTEwMDt3PTQ2/https://s.yimg.com/xe/i/us/sp/v/nba_cutout/players_l/10142021/6065.png</image_url>
	 <is_undroppable>0</is_undroppable>
	 <position_type>P</position_type>
	 <primary_position>PG</primary_position>
	 <eligible_positions>
		<position>PG</position>
		<position>SG</position>
		<position>G</position>
		<position>Util</position>
	 </eligible_positions>
	 <has_player_notes>1</has_player_notes>
	 <player_notes_last_timestamp>1636515545</player_notes_last_timestamp>
	</player>
	`

	transactionFullTestResp = `
	<transaction>
	 <transaction_key>410.l.16883.tr.227</transaction_key>
	 <transaction_id>227</transaction_id>
	 <type>add/drop</type>
	 <status>successful</status>
	 <timestamp>1636674697</timestamp>
	 <players count="2">
		<player>
		 <player_key>410.p.6450</player_key>
		 <player_id>6450</player_id>
		 <name>
			<full>Paul Reed</full>
		 </name>
		 <transaction_data>
			<type>add</type>
			<source_type>freeagents</source_type>
			<destination_type>team</destination_type>
			<destination_team_key>410.l.16883.t.8</destination_team_key>
			<destination_team_name>Anti-Vax and INJ</destination_team_name>
		 </transaction_data>
		</player>
		<player>
		 <player_key>410.p.4488</player_key>
		 <player_id>4488</player_id>
		 <name>
			<full>George Hill</full>
		 </name>
		 <transaction_data>
			<type>drop</type>
			<source_type>team</source_type>
			<source_team_key>410.l.16883.t.8</source_team_key>
			<source_team_name>Anti-Vax and INJ</source_team_name>
			<destination_type>waivers</destination_type>
		 </transaction_data>
		</player>
	 </players>
	</transaction>
	`

	transactionsTestResp = `
	<transactions>
	  <transaction>
	  	<transaction_key>257.l.193.pt.1</transaction_key>
		</transaction>
		<transaction>
	  	<transaction_key>257.l.193.pt.2</transaction_key>
		</transaction>
		<transaction>
	  	<transaction_key>257.l.193.pt.3</transaction_key>
		</transaction>
	</transactions>
	`

	gamesTestResp = `
	<games>
	  <game>
		  <code>nba</code>
		</game>
		<game>
		  <code>nfl</code>
		</game>
		<game>
		  <code>nhl</code>
		</game>
	</games>
	`
)
