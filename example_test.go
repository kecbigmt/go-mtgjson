package mtgjson_test

import (
  "os"
  "fmt"

  "github.com/kecbimgt/go-mtgjson"
)

func ExampleGetSetList() {
  c := mtgjson.New()
  setList, err := c.GetSetList()
  if err != nil {
    fmt.Println("error: %s", err)
  }

  // get GRN from SetList
  var setGRN mtgjson.Set
  for _, set := range setList {
    if set.Code == "GRN" {
      setGRN = set
    }
  }
  if set == nil {
    fmt.Println("GRN not found")
  }

  os.Stdout.Write(setGRN.Name)
  // Output:
  // Guilds of Ravnica
}
