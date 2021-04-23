package tests

import (
	"github.com/revel/revel/testing"
  "net/url"
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
  data.Add("balance", fmt.Sprintf("%f", rand.Float64()))
  
  t.PostForm("/App/Add", data)
  t.AssertOk()
  t.AssertContains(assetName)
}

func (t *AppTest) TestLiabilityCreation() {
  
  assetName := fmt.Sprintf("Test Asset %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("aorl", "Liability")
  data.Add("balance", fmt.Sprintf("%f", rand.Float64()))
  
  t.PostForm("/App/Add", data)
  t.AssertOk()
  t.AssertContains(assetName)
}

func (t *AppTest) TestNoName() {
  data := url.Values{}
  
  data.Add("aorl", "Asset")
  data.Add("balance", fmt.Sprintf("%f", rand.Float64()))
  
  t.PostForm("/App/Add", data)
  t.AssertOk()
  t.AssertContains("A name for the asset/liability is required!")
}

func (t *AppTest) TestNoType() {
  assetName := fmt.Sprintf("Test Asset %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("balance", fmt.Sprintf("%f", rand.Float64()))
  
  t.PostForm("/App/Add", data)
  t.AssertOk()
  t.AssertContains("The type ('Asset' or 'Liability') is required!")
}

func (t *AppTest) TestNoBalance() {
  assetName := fmt.Sprintf("Test Asset %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("aorl", "Asset")
  
  t.PostForm("/App/Add", data)
  t.AssertOk()
  t.AssertContains("The balance for the asset/liability is required!")
}

func (t *AppTest) TestBadType() {
  
  assetName := fmt.Sprintf("Test Asset %d", rand.Int())
  data := url.Values{}
  
  data.Add("name", assetName)
  data.Add("aorl", "bility")
  data.Add("balance", fmt.Sprintf("%f", rand.Float64()))
  
  t.PostForm("/App/Add", data)
  t.AssertOk()
  t.AssertContains("The type must be either an 'Asset' or a 'Liability'")
}

func (t *AppTest) After() {
	println("Tear down")
}
