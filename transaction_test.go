package yfquery

import (
	"fmt"
	"testing"
)

func TestTransactionsQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Transactions(),
				"/transactions",
			},
			{
				Transactions().Keys([]string{"223.l.431.tr.26", "257.l.193.w.c.2_6390"}),
				"/transactions;transaction_keys=223.l.431.tr.26,257.l.193.w.c.2_6390",
			},
			{
				Transactions().Key("223.l.431.tr.26"),
				"/transactions;transaction_keys=223.l.431.tr.26",
			},
			{
				Transactions().Key("223.l.431.tr.26").Key("257.l.193.w.c.2_6390"),
				"/transactions;transaction_keys=257.l.193.w.c.2_6390",
			},
			{
				Transactions().Types([]string{"add", "drop"}),
				"/transactions;types=add,drop",
			},
			{
				Transactions().Count(25),
				"/transactions;count=25",
			},
			{
				Transactions().TeamKey("nba.l.12345.t.1"),
				"/transactions;team_key=nba.l.12345.t.1",
			},
		})
}

func TestTransactionQuery(t *testing.T) {
	testQueries(t,
		[]testQueryPair{
			{
				Transaction(),
				"/transaction",
			},
			{
				Transaction().Key("223.l.431.tr.26"),
				"/transaction/223.l.431.tr.26",
			},
		})
}

func TestTransactionsQueryPayload(t *testing.T) {
	q := Transactions().Add("nba.l.12345", "nba.p.6030")
	got := q.payload
	want := fmt.Sprintf(addTransaction, "nba.p.6030", "nba.l.12345")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transactions().Drop("nba.l.12345", "nba.p.6030")
	got = q.payload
	want = fmt.Sprintf(dropTransaction, "nba.p.6030", "nba.l.12345")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transactions().AddDrop("nba.l.12345", "nba.p.6030", "nba.p.6031")
	got = q.payload
	want = fmt.Sprintf(addDropTransaction, "nba.p.6030", "nba.l.12345", "nba.p.6031", "nba.l.12345")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}
}

func TestTransactionQueryPayload(t *testing.T) {
	q := Transaction().UpdateWaiverClaim("248.l.55438.w.c.2_6093", 1, 20)
	got := q.payload
	want := fmt.Sprintf(updateWaiverPriorFAABTransaction, "248.l.55438.w.c.2_6093", 1, 20)

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transaction().UpdateWaiverClaimPriority("248.l.55438.w.c.2_6093", 1)
	got = q.payload
	want = fmt.Sprintf(updateWaiverPriorTransaction, "248.l.55438.w.c.2_6093", 1)

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transaction().UpdateWaiverClaimFAAB("248.l.55438.w.c.2_6093", 20)
	got = q.payload
	want = fmt.Sprintf(updateWaiverFAABTransaction, "248.l.55438.w.c.2_6093", 20)

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transaction().AcceptTrade("248.l.55438.pt.11", "deal!")
	got = q.payload
	want = fmt.Sprintf(tradeTransaction, "248.l.55438.pt.11", "accept", "deal!", "")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transaction().RejectTrade("248.l.55438.pt.11", "bad trade bro")
	got = q.payload
	want = fmt.Sprintf(tradeTransaction, "248.l.55438.pt.11", "reject", "bad trade bro", "")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transaction().AllowTrade("248.l.55438.pt.11")
	got = q.payload
	want = fmt.Sprintf(tradeTransaction, "248.l.55438.pt.11", "allow", "", "")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transaction().DisallowTrade("248.l.55438.pt.11")
	got = q.payload
	want = fmt.Sprintf(tradeTransaction, "248.l.55438.pt.11", "disallow", "", "")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transaction().VoteAgainstTrade("248.l.55438.pt.11", "248.l.55438.t.1")
	got = q.payload
	want = fmt.Sprintf(tradeTransaction, "248.l.55438.pt.11", "vote_against", "", "248.l.55438.t.1")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transaction().CancelTrade("248.l.55438.pt.11")
	got = q.payload
	want = fmt.Sprintf(cancelTransaction, "248.l.55438.pt.11", "pending_trade")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}

	q = Transaction().CancelWaiverClaim("248.l.55438.w.c.2_6093")
	got = q.payload
	want = fmt.Sprintf(cancelTransaction, "248.l.55438.w.c.2_6093", "waiver")

	if got != want {
		t.Errorf("got:\n%q\n\nwant:\n %q", got, want)
	}
}
