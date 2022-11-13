package yfquery

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/famendola1/yfquery/schema"
)

const (
	endpoint = "https://fantasysports.yahooapis.com/fantasy/v2"
)

type query struct {
	base         string
	resource     string
	isCollection bool
	keys         []string
	outs         []string
	params       []string
	payload      string
}

// ToString generates the endpoint string for the query
func (q *query) ToString() string {
	var uri string

	if q.base != "" {
		uri = q.base + "/"
	}

	if q.isCollection {
		uri += q.resource + "s"
		if len(q.keys) != 0 {
			uri += fmt.Sprintf(";%s_keys=%s", q.resource, strings.Join(q.keys, ","))
		}
	} else {
		uri += q.resource
		if len(q.keys) != 0 {
			uri += ("/" + q.keys[0])
		}
	}

	if len(q.outs) == 1 {
		uri += fmt.Sprintf("/%s", q.outs[0])
	}
	if len(q.outs) > 1 {
		uri += fmt.Sprintf(";out=%s", strings.Join(q.outs, ","))
	}

	if len(q.params) != 0 {
		uri += (";" + strings.Join(q.params, ";"))
	}

	if uri[0] != '/' {
		uri = "/" + uri
	}

	return uri
}

// Get sends a "GET" request to the Yahoo Fantasy endpoint that the query would
// generate using the provided client.
func (q *query) Get(client *http.Client) (*schema.FantasyContent, error) {
	return sendRequest(client, "GET", q.ToString(), "")
}

// Post sends a "POST" request to the Yahoo Fantasy endpoint that the query would
// generate using the provided client along with the payload.
func (q *query) Post(client *http.Client) (*schema.FantasyContent, error) {
	return sendRequest(client, "POST", q.ToString(), q.payload)
}

// Put sends a "PUT" request to the Yahoo Fantasy endpoint that the query would
// generate using the provided client along with the payload.
func (q *query) Put(client *http.Client) (*schema.FantasyContent, error) {
	return sendRequest(client, "PUT", q.ToString(), q.payload)
}

// Delete sends a "DELETE" request to the Yahoo Fantasy endpoint that the query would
// generate using the provided client along with the payload.
func (q *query) Delete(client *http.Client) (*schema.FantasyContent, error) {
	return sendRequest(client, "DELETE", q.ToString(), q.payload)
}

func (q *query) Reset() {
	q.keys = []string{}
	q.outs = []string{}
	q.params = []string{}
	q.payload = ""
}

func sendRequest(client *http.Client, method, uri, payload string) (*schema.FantasyContent, error) {
	var fc schema.FantasyContent
	var err error
	var resp *http.Response

	switch method {
	case "GET":
		resp, err = client.Get((endpoint + uri))
		break
	case "POST":
		resp, err = client.Post((endpoint + uri), "application/xml", strings.NewReader(payload))
		break
	case "PUT":
		req, err2 := http.NewRequest("PUT", (endpoint + uri), strings.NewReader(payload))
		if err2 != nil {
			return nil, err2
		}
		req.Header.Add("content-type", "application/xml")
		resp, err = client.Do(req)
		break
	case "DELETE":
		req, err2 := http.NewRequest("DELETE", (endpoint + uri), strings.NewReader(payload))
		if err2 != nil {
			return nil, err2
		}
		req.Header.Add("content-type", "application/xml")
		resp, err = client.Do(req)
		break
	}

	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, handleError(resp)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if err = parse(buf.String(), "//fantasy_content", &fc); err != nil {
		return nil, err
	}

	return &fc, nil
}

// handleError returns an error containing the error message in the response.
func handleError(resp *http.Response) error {
	doc, err := xmlquery.Parse(resp.Body)
	if err != nil {
		return err
	}

	node, err := xmlquery.Query(doc, "//description")
	if err != nil {
		return err
	}

	return fmt.Errorf("%s: %s", resp.Status, node.InnerText())
}
