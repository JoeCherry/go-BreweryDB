package brewerydb

import (
    "errors"
)

type AdjunctService struct {
    client *Client
}

type AdjunctResponse struct {
    CurrentPage int
    NumberOfPages int
    TotalResults int
    Data []Adjunct
    Status string
}

type AdjunctByIdResponse struct {
    Message string
    Data Adjunct
    Status string
}

type Adjunct struct {
    Id int
    Name string
    Category string
    CategoryDisplay string
    CreateDate string
}

func (a *AdjunctService) Get(options string) (*AdjunctResponse, error) {
    adjuncts := new(AdjunctResponse)

    req, err := a.client.NewRequest("GET", "adjuncts", options)
    if err != nil {
        return nil, err
    }

    resp, err := a.client.Do(req, adjuncts)
    if err != nil {
        return nil, err
    }

    return resp.(*AdjunctResponse), nil
}

func (a *AdjunctService) GetById(id string) (*AdjunctByIdResponse, error) {
    adjunct := new(AdjunctByIdResponse)

    if id == "" {
        return adjunct, errors.New("No id provided")
    }

    req, err := a.client.NewRequest("GET", "adjunct/" + id, "")
    if err != nil {
        return nil, err
    }

    resp, err := a.client.Do(req, adjunct)
    if err != nil {
        return nil, err
    }

    return resp.(*AdjunctByIdResponse), nil
}
