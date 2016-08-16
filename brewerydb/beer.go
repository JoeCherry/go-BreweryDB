package brewerydb

import (
    "errors"
)

type BeerService struct {
    client *Client
}
//Response struct for http://api.brewerydb.com/v2/beers/?key=<user key>&params=[params]
type BeersResponse struct {
    CurrentPage int
    NumberOfPages int
    TotalResults int
    Data []Beer
    Status string
}
//Response struct for http://api.brewerydb.com/v2/beer/:beerid/key?=<user key>
type BeerByIdResponse struct {
    Message string
    Data Beer
    Status string
}

type Beer struct {
    Id string
    Name string
    NameDisplay string
    Descrption string
    Abv string
    Ibu string
    GlasswareId int
    AvailableId int
    StyleId int
    SsOrganic string
    Labels Label
    Status string
    StatusDisplay string
    CreateDate string
    UpdateDate string
    Glass Glass
    Available Available
    Style Style
}

type Label struct {
    Icon string
    Medium string
    Large string
}

type Available struct {
    Id int
    Name string
    Descrption string
}

// Makes a get request to http://api.brewerydb.com/v2/beers/?key=<user key>&params=[params]
// Returns a list of all the beers in the database
func (b *BeerService) Get(options string) (*BeersResponse, error) {
    beers := new(BeersResponse)

    req, err := b.client.NewRequest("GET", "beers", options)
    if err != nil {
        return nil, err
    }

    resp, err := b.client.Do(req, beers)
    if err != nil {
        return nil, err
    }

    return resp.(*BeersResponse), nil
}

//makes a get request to http://api.brewerydb.com/v2/beer/:beerid/key?=<user key>
//returns a single beer based on its id
func (b *BeerService) GetById(id string) (*BeerByIdResponse, error) {
    beer := new(BeerByIdResponse)

    if id == "" {
        return nil, errors.New("No id provided")
    }

    req, err := b.client.NewRequest("GET", "beer/" + id, "")
    if err != nil {
        return nil, err
    }

    resp, err := b.client.Do(req, beer)
    if err != nil {
        return nil, err
    }

    return resp.(*BeerByIdResponse), nil
}
