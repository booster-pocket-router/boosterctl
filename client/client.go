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
	URL, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL:    URL,
		httpClient: http.DefaultClient,
	}, nil
}

func (c *Client) Get(endpoint string, query *url.Values) (*http.Response, error) {
	URL, err := url.Parse(c.BaseURL.String() + endpoint)
	if err != nil {
		return nil, err
	}
	if query != nil {
		URL.RawQuery = query.Encode()
	}

	return c.httpClient.Get(URL.String())
}

func (c *Client) Post(endpoint string, body io.Reader) (*http.Response, error) {
	return c.do("POST", endpoint)
}

func (c *Client) Del(endpoint string) (*http.Response, error) {
	return c.do("DELETE", endpoint)
}

func (c *Client) do(method string, endpoint string) (*http.Response, error) {
	URL, err := url.Parse(c.BaseURL.String() + endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, URL.String(), nil)
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func PrettyJSON(src io.Reader, dest io.Writer) error {
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
