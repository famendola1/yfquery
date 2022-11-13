package yfquery

import (
	"fmt"
	"strings"
)

// TransactionQuery can be used to query the /transactions or /transaction Yahoo
// Fantasy endpoints.
type TransactionQuery struct {
	query
}

// Transaction returns a TransactionQuery for the /transaction endpoint.
func Transaction() *TransactionQuery {
	return &TransactionQuery{query{resource: "transaction"}}
}

// Transactions returns a TransactionQuery for the /transactions endpoint.
func Transactions() *TransactionQuery {
	return &TransactionQuery{query{resource: "transaction", isCollection: true}}
}

// Keys adds the "transaction_keys" parameter with the given keys to the query.
func (t *TransactionQuery) Keys(keys []string) *TransactionQuery {
	t.keys = append(t.keys, keys...)
	return t
}

// Key sets the "transaction_keys" parameter to the the given key. When querying
// the /transaction endpoint the key will be appended to the query path
// (i.e. /transaction/<key>).
func (t *TransactionQuery) Key(key string) *TransactionQuery {
	t.keys = []string{key}
	return t
}

// Types sets the "transaction_types" parameter with the given types to the
// query. Valid types are add,drop,commish,trade.
func (t *TransactionQuery) Types(types []string) *TransactionQuery {
	t.params = append(t.params, fmt.Sprintf("types=%s", strings.Join(types, ",")))
	return t
}

// Count adds the "count" parameter with the given count to the query. count is
// expected by Yahoo to be a positive integer.
func (t *TransactionQuery) Count(count int) *TransactionQuery {
	t.params = append(t.params, fmt.Sprintf("count=%d", count))
	return t
}

// TeamKey adds the "team_key" parameter with the provided key to the query.
func (t *TransactionQuery) TeamKey(key string) *TransactionQuery {
	t.params = append(t.params, fmt.Sprintf("team_key=%s", key))
	return t
}

// Add sets the payload of the query to an "add" transaction. Use the Post
// method to send the payload. This payload works with the /transactions endpoint.
func (t *TransactionQuery) Add(teamKey, playerKey string) *TransactionQuery {
	t.payload = fmt.Sprintf(addTransaction, playerKey, teamKey)
	return t
}

// Drop sets the payload of the query to a "drop" transaction. Use the Post
// method to send the payload. This payload works with the /transactions endpoint.
func (t *TransactionQuery) Drop(teamKey, playerKey string) *TransactionQuery {
	t.payload = fmt.Sprintf(dropTransaction, playerKey, teamKey)
	return t
}

// AddDrop sets the payload of the query to an "add/drop" transaction. Use the
// Post method to send the payload. This payload works with the /transactions endpoint.
func (t *TransactionQuery) AddDrop(teamKey, addPlayerKey, dropPlayerKey string) *TransactionQuery {
	t.payload = fmt.Sprintf(addDropTransaction, addPlayerKey, teamKey, dropPlayerKey, teamKey)
	return t
}

// UpdateWaiverClaim set the payload of the query to perform a waiver claim update on
// priority and FAAB. Use the Put method to send this payload. This payload works
// with the /transaction endpoint.
func (t *TransactionQuery) UpdateWaiverClaim(transactionKey string, priority, faab int) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(updateWaiverPriorFAABTransaction, transactionKey, priority, faab)
	return t
}

// UpdateWaiverClaimPriority set the payload of the query to perform a waiver claim update on
// priority. Use the Put method to send this payload. This payload works with
// the /transaction endpoint.
func (t *TransactionQuery) UpdateWaiverClaimPriority(transactionKey string, priority int) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(updateWaiverPriorTransaction, transactionKey, priority)
	return t
}

// UpdateWaiverClaimFAAB set the payload of the query to perform a waiver claim update on
// FAAB. Use the Put method to send this payload. This payload works with the
// /transaction endpoint.
func (t *TransactionQuery) UpdateWaiverClaimFAAB(transactionKey string, faab int) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(updateWaiverFAABTransaction, transactionKey, faab)
	return t
}

// AcceptTrade set the payload of the query to perform an accept trade action.
// Use the Put method to send this payload. This payload works with the
// /transaction endpoint.
func (t *TransactionQuery) AcceptTrade(transactionKey, note string) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(tradeTransaction, transactionKey, "accept", note, "")
	return t
}

// RejectTrade set the payload of the query to perform a reject trade action.
// Use the Put method to send this payload. This payload works with the
// /transaction endpoint.
func (t *TransactionQuery) RejectTrade(transactionKey, note string) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(tradeTransaction, transactionKey, "reject", note, "")
	return t
}

// AllowTrade set the payload of the query to perform an allow trade action.
// Use the Put method to send this payload. This payload works with the
// /transaction endpoint.
func (t *TransactionQuery) AllowTrade(transactionKey string) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(tradeTransaction, transactionKey, "allow", "", "")
	return t
}

// DisallowTrade set the payload of the query to perform a disallow trade action.
// Use the Put method to send this payload. This payload works with the
// /transaction endpoint.
func (t *TransactionQuery) DisallowTrade(transactionKey string) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(tradeTransaction, transactionKey, "disallow", "", "")
	return t
}

// VoteAgainstTrade set the payload of the query to perform a vote against trade action.
// Use the Put method to send this payload. This payload works with the
// /transaction endpoint.
func (t *TransactionQuery) VoteAgainstTrade(transactionKey, voterTeamKey string) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(tradeTransaction, transactionKey, "vote_against", "", voterTeamKey)
	return t
}

// CancelTrade set the payload of the query to perform a cancel trade action.
// Use the Delete method to send this payload. This payload works with the
// /transaction endpoint.
func (t *TransactionQuery) CancelTrade(transactionKey string) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(cancelTransaction, transactionKey, "pending_trade")
	return t
}

// CancelWaiverClaim set the payload of the query to perform a cancel waiver claim action.
// Use the Delete method to send this payload. This payload works with the
// /transaction endpoint.
func (t *TransactionQuery) CancelWaiverClaim(transactionKey string) *TransactionQuery {
	t.keys = []string{transactionKey}
	t.payload = fmt.Sprintf(cancelTransaction, transactionKey, "waiver")
	return t
}
