package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bariasabda/monorepo/packages/fetch/domain/entity"
)

func (r *repository) GetResource() (*[]entity.Data, error) {
	var data *[]entity.Data
	resp, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &data)
	return data, nil
}
