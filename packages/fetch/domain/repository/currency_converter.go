package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bariasabda/monorepo/packages/fetch/domain/entity"
)

func (r *repository) CurrencyConverter(from string, to string) (*entity.Currency, error) {
	var curr *entity.Currency
	resp, err := http.Get(fmt.Sprintf("%s/fetch-one?from=%s&to=%s&api_key=%s",r.cfg.CurrencyConverter.BaseUrl,from,to,r.cfg.CurrencyConverter.ApiKey))
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &curr)
	return curr, nil
}
