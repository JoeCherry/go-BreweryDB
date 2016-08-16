package brewerydb

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Client struct {
    client *http.Client
    Key string
    BaseUrl string

    common service

    Adjunct *AdjunctService
    Beers *BeerService
    Brewery *BreweryService
    Category *CategoryService
    Change *ChangeService
}

type service struct {
    client *Client
}

func NewClient(key string) *Client {
    if key == "" {
        return nil
    }

    httpClient := http.DefaultClient
    url := "http://api.brewerydb.com/v2/"

    c := &Client{client: httpClient, Key: key, BaseUrl: url}
    c.common.client = c
    c.Adjunct = (*AdjunctService)(&c.common)
    c.Beers = (*BeerService)(&c.common)
    c.Brewery = (*BreweryService)(&c.common)
    c.Category = (*CategoryService)(&c.common)
    c.ChangeService = (*ChangeService)(&c.common)
    return c
}

func (c *Client) NewRequest(method, endpoint string, options string) (*http.Request, error) {
    url := c.BaseUrl + endpoint + "?"
    if options != "" {
        url += options + "&"
    }
    url += "key=" + c.Key
    fmt.Println(url)
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Add("content-type", "application/x-www-form-urlencoded")
    return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (interface{}, error) {
    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    e := json.NewDecoder(resp.Body).Decode(v)
    if e != nil {
        return nil, err
    }
    return v, nil
}
