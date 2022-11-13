package yfquery

const (
	addDropTransaction = `
  <fantasy_content>
    <transaction>
      <type>add/drop</type>
      <players>
        <player>
          <player_key>%s</player_key>
          <transaction_data>
            <type>add</type>
            <destination_team_key>%s</destination_team_key>
          </transaction_data>
        </player>
        <player>
          <player_key>%s</player_key>
          <transaction_data>
            <type>drop</type>
            <source_team_key>%s</source_team_key>
          </transaction_data>
        </player>
      </players>
    </transaction>
  </fantasy_content>`

	addTransaction = `
  <fantasy_content>
    <transaction>
      <type>add</type>
      <player>
        <player_key>%s</player_key>
        <transaction_data>
          <type>add</type>
          <destination_team_key>%s</destination_team_key>
        </transaction_data>
      </player>
    </transaction>
  </fantasy_content>`

	dropTransaction = `
  <fantasy_content>
    <transaction>
      <type>drop</type>
      <player>
        <player_key>%s</player_key>
        <transaction_data>
          <type>drop</type>
          <source_team_key>%s</source_team_key>
        </transaction_data>
      </player>
    </transaction>
  </fantasy_content>`

	updateWaiverPriorFAABTransaction = `
  <fantasy_content>
    <transaction>
      <transaction_key>%s</transaction_key>
      <type>waiver</type>
      <waiver_priority>%d</waiver_priority>
      <faab_bid>%d</faab_bid>
    </transaction>
  </fantasy_content>`

	updateWaiverPriorTransaction = `
  <fantasy_content>
    <transaction>
      <transaction_key>%s</transaction_key>
      <type>waiver</type>
      <waiver_priority>%d</waiver_priority>
    </transaction>
  </fantasy_content>`

	updateWaiverFAABTransaction = `
  <fantasy_content>
    <transaction>
      <transaction_key>%s</transaction_key>
      <type>waiver</type>
      <faab_bid>%d</faab_bid>
    </transaction>
  </fantasy_content>`

	tradeTransaction = `
  <fantasy_content>
    <transaction>
      <transaction_key>%s</transaction_key>
      <type>pending_trade</type>
      <action>%s</action>
      <trade_note>%s</trade_note>
      <voter_team_key>%s</voter_team_key>
    </transaction>
  </fantasy_content>`

	cancelTransaction = `
  <fantasy_content>
    <transaction>
      <transaction_key>%s</transaction_key>
      <type>%s</type>
    </transaction>
  </fantasy_content>`
)
