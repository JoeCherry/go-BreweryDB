package brewerydb

import (
    "errors"
)

type CategoryService struct {
    client *Client
}

type CategoriesResponse struct {
    CurrentPage int
    NumberOfPages int
    TotalResults int
    Data []Category
    Status string
}

type CategoryByIdByResponse struct {
    Message string
    Data Category
    Status string
}

type Category struct {
    Id int
    Name string
    CreateDate string
}

func (c *CategoryService) Get(options string) (*CategoriesResponse, error) {
    categories := new(CategoriesResponse)

    req, err := c.client.NewRequest("GET", "categories", options)
    if err != nil {
        return nil, err
    }

    resp, err := c.client.Do(req, categories)
    if err != nil {
        return nil, err
    }

    return resp.(*CategoriesResponse), nil
}

func (c *CategoryService) GetById(id string) (*CategoryByIdByResponse, error) {
    if id == "" {
        return nil, errors.New("No id provided")
    }
    category := new(CategoryByIdByResponse)

    req, err := c.client.NewRequest("GET", "category" + id, "")
    if err != nil {
        return nil, err
    }

    resp, err := c.client.Do(req, category)
    if err != nil {
        return nil, err
    }

    return resp.(*CategoryByIdByResponse), nil
}
