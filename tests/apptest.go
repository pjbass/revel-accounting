package tests

import (
	"github.com/revel/revel/testing"
  "accounting/app/models"
  "accounting/app/controllers"
  "net/url"
  "regexp"
  "strings"
  "math/rand"
  "fmt"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestAssetCreation() {
  
  assetName := fmt.Sprintf("Test Asset %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("aorl", "Asset")
  data.Add("balance", fmt.Sprintf("%f", 100.0 * rand.Float64()))
  
  t.PostForm("/assets/add", data)
  t.AssertOk()
  t.AssertContains(assetName)
}

func (t *AppTest) TestLiabilityCreation() {
  
  assetName := fmt.Sprintf("Test Asset %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("aorl", "Liability")
  data.Add("balance", fmt.Sprintf("%f", 100.0 * rand.Float64()))
  
  t.PostForm("/assets/add", data)
  t.AssertOk()
  t.AssertContains(assetName)
}

func (t *AppTest) TestNoName() {
  data := url.Values{}
  
  data.Add("aorl", "Asset")
  data.Add("balance", fmt.Sprintf("%f", rand.Float64()))
  
  t.PostForm("/assets/add", data)
  t.AssertOk()
  t.AssertContains("A name for the asset/liability is required!")
}

func (t *AppTest) TestNoType() {
  assetName := fmt.Sprintf("Test Asset %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("balance", fmt.Sprintf("%f", 100.0 * rand.Float64()))
  
  t.PostForm("/assets/add", data)
  t.AssertOk()
  t.AssertContains("The type (Asset or Liability) is required!")
}

func (t *AppTest) TestNoBalance() {
  assetName := fmt.Sprintf("Test Asset %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("aorl", "Asset")
  
  t.PostForm("/assets/add", data)
  t.AssertOk()
  t.AssertContains("The balance for the asset/liability is required!")
}

func (t *AppTest) TestAlphabetBalance() {
  assetName := fmt.Sprintf("Test Balance %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("aorl", "Asset")
  data.Add("balance", "jibberish")
  
  t.PostForm("/assets/add", data)
  t.AssertOk()
  t.Get("/")
  t.AssertNotContains(assetName)
}

func (t *AppTest) TestBadType() {
  
  assetName := fmt.Sprintf("Test Asset %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("aorl", "bility")
  data.Add("balance", fmt.Sprintf("%f", 100.0 * rand.Float64()))
  
  t.PostForm("/assets/add", data)
  t.AssertOk()
  t.AssertContains("The type must be either an Asset or a Liability")
}

func (t *AppTest) TestTotals() {
  
  var assets = []models.Asset {
    models.Asset{
      Name: "test 1", 
      Type: "Asset", 
      Balance: 450.0,
    },
    models.Asset{
      Name: "test 2", 
      Type: "Liability", 
      Balance: 200.0,
    },
    models.Asset{
      Name: "test 3",
      Type: "Asset",
      Balance: 100.0,
    },
    models.Asset{
      Name: "test 4",
      Type: "Liability",
      Balance: 300.0,
    },
    models.Asset{
      Name: "test 5",
      Type: "Asset",
      Balance: 150.0,
    },
  }
  
  assetTotal, liabilityTotal, netWorth := controllers.CalculateTotals(assets)
  
  t.Assert(int(assetTotal) == 700) 
  t.Assert(int(liabilityTotal) == 500)
  t.Assert(int(netWorth) == 200)
}

func (t *AppTest) TestDelete() {
  
  assetName := fmt.Sprintf("Test Delete %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("aorl", "Liability")
  data.Add("balance", fmt.Sprintf("%f", 100.0 * rand.Float64()))
  
  t.PostForm("/assets/add", data)
  t.AssertOk()
  t.AssertContains(assetName)
  
  delPat := regexp.MustCompile(
      `<th scope="row">` + assetName + `</th>[\s\S]*?<input .*?value="(\d+)"`)
  
  deleteId := delPat.FindStringSubmatch(string(t.ResponseBody))[1]
  
  del := url.Values{}
  del.Add("id", deleteId)
  t.PostForm("/assets/delete", del)
  t.AssertOk()
  t.AssertNotContains(assetName)
  
}

func (t *AppTest) TestDeleteNonExisting() {
  
  t.Get("/")
  body := string(t.ResponseBody)
  deleteId := string(rand.Int())
  
  for strings.Contains(body, deleteId) {
    deleteId = string(rand.Int())
  }
  
  data := url.Values{}
  data.Add("id", deleteId)
  
  t.PostForm("/assets/delete", data)
  t.AssertOk()
  t.AssertContains("Asset not found!")
}

func (t *AppTest) TestDeleteAlphabet() {
  data := url.Values{}
  data.Add("id", "jibberish")
  t.PostForm("/assets/delete", data)
  t.AssertOk()
  t.AssertContains("Asset not found!")
}

func (t *AppTest) TestDeleteNoId() {
  data := url.Values{}
  t.PostForm("/assets/delete", data)
  t.AssertOk()
  t.AssertContains("Asset not found!")
}

func (t *AppTest) After() {
	println("Tear down")
}
