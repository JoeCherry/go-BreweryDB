package brewerydb

type ChangeService struct {
    client *Client
}

type ChangesResponse struct {
    Status string
    NumberOfPages int
    Data []Change
    CurrentPage int
}
