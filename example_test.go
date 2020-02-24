package mtgjson_test

import (
  "fmt"

  "github.com/kecbigmt/go-mtgjson"
)

func ExampleClient_GetSetList() {
  c := mtgjson.New()
  setList, err := c.GetSetList()
  if err != nil {
    fmt.Printf("error: %s", err)
  }

  // get GRN from SetList
  var setGRN mtgjson.Set
  for _, set := range setList {
    if set.Code == "GRN" {
      setGRN = set
    }
  }

  fmt.Print(setGRN.Name)
  // Output:
  // Guilds of Ravnica
}
