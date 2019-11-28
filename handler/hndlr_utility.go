package hndlr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/MyRetail/common"
)

type Resp struct {
	Result          string `json:"Result"`
	RespCode        string `json:"RespCode"`
	RespDescription string `json:"RespDescription"`
}

type GetProductOutput struct {
	Resp
	Product common.Product `json:"product"`
}

type UpdateProdInput struct {
	Product common.Product `json:"product"`
}

type UpdateProdOutput struct {
	Resp
}

func GetProduct(id string) (*common.Product, error) {
	if id == "" {
		return nil, errors.New("EMPTY_PRODUCT_ID")
	}
	return getProductFromRemote(id)
}

func getProductFromRemote(id string) (*common.Product, error) {
	prodGetUrl := common.PRODUCT_BASE_URL + id + common.PRODUCT_PARAMS
	resp, e := http.Get(prodGetUrl)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	p := result["product"].(map[string]interface{})
	item := p["item"].(map[string]interface{})
	prod_desc := item["product_description"].(map[string]interface{})
	title := prod_desc["title"].(string)
	price := getPriceOfProductRemote(id)
	v, _ := strconv.ParseInt(id, 10, 64)
	prod := new(common.Product)
	prod.ID = uint64(v)
	prod.Name = title
	prod.Price = *price
	fmt.Println("current product", prod)
	return prod, nil
}

func getPriceOfProductRemote(id string) *common.Cost {
	c := new(common.Cost)
	c.Value = 25
	c.Currency = "USD"
	return c
}
