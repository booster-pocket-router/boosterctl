// Copyright © 2019 KIM KeepInMind GmbH/srl
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

func New(addr string) (*Client, error) {
	URL, err := url.Parse("http://" + addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:    URL,
		httpClient: http.DefaultClient,
	}, nil
}

type PolicyReq struct {
	SourceID string `json:"source_id"`
	Target   string `json:"target"`
	Issuer   string `json:"issuer"`
	Reason   string `json:"reason"`
}

func (c *Client) AddPolicy(name string, p interface{}) (string, io.Reader, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(&p); err != nil {
		return "", nil, fmt.Errorf("client: unable to encode policy: %v", err)
	}

	resp, err := c.post("/policies/"+name+".json", &buf)
	if err != nil {
		return "", nil, fmt.Errorf("client: unable to make request: %v", err)
	}
	defer resp.Body.Close()

	return c.handleResponse(resp)
}

func (c *Client) DelPolicy(id string) (string, io.Reader, error) {
	resp, err := c.del("/policies/" + id + ".json")
	if err != nil {
		return "", nil, fmt.Errorf("client: unable to make request: %v", err)
	}
	defer resp.Body.Close()

	return c.handleResponse(resp)
}

func (c *Client) ListSources() (string, io.Reader, error) {
	resp, err := c.get("/sources.json", nil)
	if err != nil {
		return "", nil, fmt.Errorf("client: unable to make request: %v", err)
	}
	defer resp.Body.Close()

	return c.handleResponse(resp)
}

func (c *Client) ListPolicies() (string, io.Reader, error) {
	resp, err := c.get("/policies.json", nil)
	if err != nil {
		return "", nil, fmt.Errorf("client: unable to make request: %v", err)
	}
	defer resp.Body.Close()

	return c.handleResponse(resp)
}

func (c *Client) QueryMetrics(query *url.Values) (string, io.Reader, error) {
	resp, err := c.get("/policies.json", query)
	if err != nil {
		return "", nil, fmt.Errorf("client: unable to make request: %v", err)
	}
	defer resp.Body.Close()

	return c.handleResponse(resp)
}

func (c *Client) handleResponse(resp *http.Response) (string, io.Reader, error) {
	var buf bytes.Buffer
	if err := PrettyJSON(&buf, resp.Body); err != nil {
		return resp.Status, nil, fmt.Errorf("client: unable to format json response properly: %v", err)
	}
	return resp.Status, &buf, nil
}

func (c *Client) get(endpoint string, query *url.Values) (*http.Response, error) {
	URL, err := url.Parse(c.BaseURL.String() + endpoint)
	if err != nil {
		return nil, err
	}
	if query != nil {
		URL.RawQuery = query.Encode()
	}

	return c.httpClient.Get(URL.String())
}

func (c *Client) post(endpoint string, body io.Reader) (*http.Response, error) {
	return c.do("POST", endpoint, body)
}

func (c *Client) del(endpoint string) (*http.Response, error) {
	return c.do("DELETE", endpoint, nil)
}

func (c *Client) do(method string, endpoint string, body io.Reader) (*http.Response, error) {
	URL, err := url.Parse(c.BaseURL.String() + endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, URL.String(), body)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func PrettyJSON(dest io.Writer, src io.Reader) error {
	var buf bytes.Buffer
	n, err := io.Copy(&buf, src)
	if err != nil {
		return err
	}
	if n <= 0 {
		return fmt.Errorf("src does not contain data")
	}

	var out bytes.Buffer
	if err := json.Indent(&out, buf.Bytes(), "", "\t"); err != nil {
		return err
	}

	_, err = out.WriteTo(dest)
	return err
}
