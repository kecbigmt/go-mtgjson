package mtgjson

import (
  "os"
  "io"
  "fmt"
  "bytes"
  "testing"
  "path/filepath"
)

type dummyGetSetAPI struct{}

func (api *dummyGetSetAPI) doRequest(path string) (io.Reader, error) {
  f, err := os.Open(filepath.Join("testdata", "set.json"))
  if err != nil {
    return nil, fmt.Errorf("os.Open: %s", err)
  }
  defer f.Close()

  var buf bytes.Buffer
  if _, err := io.Copy(&buf, f); err != nil {
    return nil, fmt.Errorf("io.Copy: %s", err)
  }

  return &buf, nil
}

func TestGetSet(t *testing.T) {
  c := &Client{&dummyGetSetAPI{}}
  set, err := c.GetSet("ELD")
  if err != nil {
    t.Fatalf("failed to get set: c.GetSet: %s", err)
  }

  if set.Code != "ELD" {
    t.Fatalf("got wrong code")
  }
  if set.Name != "Throne of Eldraine" {
    t.Fatalf("got wrong name")
  }
  if set.Type != "expansion" {
    t.Fatalf("got wrong type")
  }
  if set.Cards[0].Artist != "David Gaillet" {
    t.Fatalf("got wrong card artist")
  }
  if len(set.Cards[0].ForeignData) != 10 {
    t.Fatalf("got wrong foreign data")
  }
}
