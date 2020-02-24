package mtgjson

import (
  "io"
  "fmt"
  "bytes"
  "net/http"
)

// Client gets json from mtgjson and return parsed struct
type Client struct {
  API api
}

// New return new Client pointer
func New() *Client {
  return &Client{&API{}}
}

type api interface {
  doRequest(string) (io.Reader, error)
}

// API does request to https://www.mtgjson.com
type API struct{}

func (api *API) doRequest(path string) (io.Reader, error) {
  resp, err := http.Get(fmt.Sprintf("https://%s%s", host, path))
  if err != nil {
    return nil, fmt.Errorf("http.Get: %s", err)
  }

  defer resp.Body.Close()
  var buf bytes.Buffer
  if _, err := io.Copy(&buf, resp.Body); err != nil {
    return nil, fmt.Errorf("io.Copy: %s", err)
  }

  return &buf, nil
}
