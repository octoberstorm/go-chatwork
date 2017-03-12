package gochatwork

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"golang.org/x/net/context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type AppengineClient struct {
	ApiKey  string
	BaseUrl string
	Http
	Ctx context.Context
}

func NewAppengineClient(apiKey string, ctx context.Context) *AppengineClient {
	return &AppengineClient{ApiKey: apiKey, BaseUrl: BaseUrl, Ctx: ctx}
}

func (c *AppengineClient) Get(endpoint string, params map[string]string) []byte {
	return c.execute("GET", endpoint, params)
}

func (c *AppengineClient) Post(endpoint string, params map[string]string) []byte {
	return c.execute("POST", endpoint, params)
}

func (c *AppengineClient) Put(endpoint string, params map[string]string) []byte {
	return c.execute("PUT", endpoint, params)
}

func (c *AppengineClient) Delete(endpoint string, params map[string]string) []byte {
	return c.execute("DELETE", endpoint, params)
}

func (c *AppengineClient) buildUrl(baseUrl, endpoint string, params map[string]string) string {
	query := make([]string, len(params))
	for k := range params {
		query = append(query, k+"="+params[k])
	}
	return baseUrl + endpoint + "?" + strings.Join(query, "&")
}

func (c *AppengineClient) buildBody(params map[string]string) url.Values {
	body := url.Values{}
	for k := range params {
		body.Add(k, params[k])
	}
	return body
}

func (c *AppengineClient) parseBody(resp *http.Response) []byte {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []byte(``)
	}
	return body
}

func (c *AppengineClient) execute(method, endpoint string, params map[string]string) []byte {
	var (
		req        *http.Request
		requestErr error
	)

	client := urlfetch.Client(c.Ctx)
	req, requestErr = http.NewRequest(method, c.buildUrl(c.BaseUrl, endpoint, params), nil)

	if method != "GET" {
		req, requestErr = http.NewRequest(method, c.BaseUrl+endpoint, bytes.NewBufferString(c.buildBody(params).Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, requestErr = http.NewRequest(method, c.buildUrl(c.BaseUrl, endpoint, params), nil)
	}
	if requestErr != nil {
		panic(requestErr)
	}

	req.Header.Add("X-ChatWorkToken", c.ApiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return []byte(``)
	}

	return c.parseBody(resp)
}
