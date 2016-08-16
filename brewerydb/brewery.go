package brewerydb

type BreweryService struct {
    client *Client
}

type BreweriesResponse struct {
    Status string
    NumberOfPages int
    Data []Brewery
    CurrentPage int
}

type BreweryByIdResponse struct {
    Status string
    Data Brewery
    Message string
}

type Brewery struct {
    Id string
    Descrption string
    Name string
    CreateDate string
    MailingListUrl string
    UpdateDate string
    Images Images
    Established string
    IsOrganic string
    Website string
    Status string
    StatusDisplay string
}

type Images struct {
    Medium string
    Large string
    Icon string
}

// Get listings of all the breweries
func (b *BreweryService) Get(options string) (*BreweriesResponse, error) {
    breweries := new(BreweriesResponse)

    req, err := b.client.NewRequest("GET", "breweries", options)
    if err != nil {
        return nil, err
    }

    resp, err := b.client.Do(req, breweries)
    if err != nil {
        return nil, err
    }

    return resp.(*BreweriesResponse), nil
}

//Get a single brewery by its id
func (b *BreweryService) GetById(id string) (*BreweryByIdResponse, error) {
    brewery := new(BreweryByIdResponse)

    req, err := b.client.NewRequest("GET", "brewery/" + id, "")
    if err != nil {
        return nil, err
    }

    resp, err := b.client.Do(req, brewery)
    if err != nil {
        return nil, err
    }

    return resp.(*BreweryByIdResponse), nil
}
