package yfquery

import "testing"

type testQuery interface {
	ToString() string
}

type testQueryPair struct {
	query testQuery
	want  string
}

func testQueries(t *testing.T, queries []testQueryPair) {
	for _, q := range queries {
		got := q.query.ToString()
		if got != q.want {
			t.Errorf("got %q, want %q", got, q.want)
		}
	}
}
