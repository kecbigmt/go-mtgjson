package mtgjson

import (
  "io"
  "strings"
  "reflect"
  "testing"
)

type dummyGetSetListAPI struct{}

func (api *dummyGetSetListAPI) doRequest(path string) (io.Reader, error) {
  dummy := `[{"baseSetSize": 0, "code": "P15A", "isOnlineOnly": false, "isPaperOnly": false, "meta": {"date": "2020-02-18", "pricesDate": "2020-02-18", "version": "4.6.2+20200218"}, "name": "15th Anniversary Cards", "releaseDate": "2008-04-01", "totalSetSize": 2, "type": "promo"}, {"baseSetSize": 0, "code": "HTR", "isOnlineOnly": false, "isPaperOnly": false, "meta": {"date": "2020-02-18", "pricesDate": "2020-02-18", "version": "4.6.2+20200218"}, "name": "2016 Heroes of the Realm", "releaseDate": "2017-09-20", "totalSetSize": 3, "type": "memorabilia"}]`
  return strings.NewReader(dummy), nil
}

func TestGetSetList(t *testing.T) {
  c := &Client{&dummyGetSetListAPI{}}
  sets, err := c.GetSetList()
  if err != nil {
    t.Fatalf("failed to get set list")
  }

  var expectedSets []Set
  expectedSets = []Set{
    Set{
      Code: "P15A",
      Name: "15th Anniversary Cards",
      Type: "promo",
      ReleaseDate: "2008-04-01",
      TotalSetSize: 2,
      BaseSetSize: 0,
      IsOnlineOnly: false,
      IsPaperOnly: false,
      Meta: SetMeta{
        Date: "2020-02-18",
        PricesDate: "2020-02-18",
        Version: "4.6.2+20200218",
      },
    },
    Set{
      Code: "HTR",
      Name: "2016 Heroes of the Realm",
      Type: "memorabilia",
      ReleaseDate: "2017-09-20",
      TotalSetSize: 3,
      BaseSetSize: 0,
      IsOnlineOnly: false,
      IsPaperOnly: false,
      Meta: SetMeta{
        Date: "2020-02-18",
        PricesDate: "2020-02-18",
        Version: "4.6.2+20200218",
      },
    },
  }

  if !reflect.DeepEqual(sets, expectedSets) {
    t.Fatalf("got wrong sets")
  }
}
