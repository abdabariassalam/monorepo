package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bariasabda/monorepo/packages/fetch/constants"
	"github.com/bariasabda/monorepo/packages/fetch/domain/entity"
	"github.com/patrickmn/go-cache"
)

type CurrencyConverterResponse struct {
	Uuid         *string    `json:"uuid"`
	Komoditas    *string    `json:"komoditas"`
	AreaProvinsi *string    `json:"area_provinsi"`
	AreaKota     *string    `json:"area_kota"`
	Size         *string    `json:"size"`
	Price        *string    `json:"price"`
	PriceUSD     *string    `json:"price_usd"`
	TglParsed    *time.Time `json:"tgl_parsed"`
	Timestamp    *string    `json:"timestamp"`
}

func (s *service) CurrencyConverter(reqToken string) (*[]CurrencyConverterResponse, error) {
	_, err := s.VerifyToken(reqToken)
	if err != nil {
		return nil, err
	}
	var response []CurrencyConverterResponse
	c := cache.New(5*time.Minute, 10*time.Minute)

	// get currency converter
	curr, found := c.Get("curr")
	if !found {
		curr, err := s.repo.CurrencyConverter(constants.IDRCurrency, constants.USDCurrency)
		if err != nil {
			return nil, err
		}
		c.Set("curr", curr.Result.USD, cache.DefaultExpiration)
	}
	currFloat, _ := curr.(float64)

	// get data resource
	data, err := s.repo.GetResource()
	if err != nil {
		return nil, err
	}

	for _, item := range *data {
		if item.Price != nil {
			priceFloat, err := strconv.ParseFloat(*item.Price, 64)
			if err != nil {
				break
			}
			priceUSD := fmt.Sprintf("%.2f", (priceFloat * currFloat))
			response = append(response, composeResourcetoResponse(item, &priceUSD))
			continue
		}
		response = append(response, composeResourcetoResponse(item, nil))

	}

	return &response, nil
}

func composeResourcetoResponse(item entity.Data, priceUSD *string) CurrencyConverterResponse {
	return CurrencyConverterResponse{
		Uuid:         item.Uuid,
		Komoditas:    item.Komoditas,
		AreaProvinsi: item.AreaProvinsi,
		AreaKota:     item.AreaKota,
		Size:         item.Size,
		Price:        item.Price,
		PriceUSD:     priceUSD,
		TglParsed:    item.TglParsed,
		Timestamp:    item.Timestamp,
	}

}
