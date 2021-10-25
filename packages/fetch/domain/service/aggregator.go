package service

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/bariasabda/monorepo/packages/fetch/constants"
	"github.com/montanaflynn/stats"
)

type AggregatorResponse struct {
	AreaProvinsi string      `json:"area_provinsi"`
	Aggregate    []Aggregate `json:"aggregate"`
}

type Aggregate struct {
	Tanggal string `json:"tanggal"`
	Min     string `json:"min"`
	Max     string `json:"max"`
	Median  string `json:"median"`
	Avg     string `json:"avg"`
}

// menampung data yang di butuhkan untuk aggregate
type NewAggData struct {
	TglParsed time.Time `json:"tgl_parsed"`
	Price     string    `json:"price"`
}

func (s *service) Aggregator(reqToken string) (*[]AggregatorResponse, error) {
	user, err := s.VerifyToken(reqToken)
	if err != nil {
		return nil, err
	} else if user.Role != constants.RoleAdmin {
		return nil, constants.ErrRoleNotAdmin
	}
	var response []AggregatorResponse
	areaProvinsi := make(map[string][]NewAggData)

	// get data resource
	data, err := s.repo.GetResource()
	if err != nil {
		return nil, err
	}

	// menyiapkan data
	for _, item := range *data {
		// check if areaprovinsi kosong dan price not a float, remove data
		if item.Price != nil {
			_, err := strconv.ParseFloat(*item.Price, 64)
			if item.AreaProvinsi != nil && err == nil && item.TglParsed != nil && *item.TglParsed != (time.Time{}) {
				areaProvinsi[*item.AreaProvinsi] = append(areaProvinsi[*item.AreaProvinsi], NewAggData{
					TglParsed: *item.TglParsed,
					Price:     *item.Price,
				})
			}
		}
	}

	// memproses data menjadi data weekly
	for i, item := range areaProvinsi {
		agg := processAggregatre(item)

		response = append(response, AggregatorResponse{
			AreaProvinsi: i,
			Aggregate:    agg,
		})
	}

	return &response, nil
}

func processAggregatre(value []NewAggData) []Aggregate {
	// menyiap kan weekly aggregator
	arrayAggregate := []Aggregate{}
	minDate := value[0].TglParsed
	maxDate := value[0].TglParsed
	for _, v := range value {
		if v.TglParsed.String() < minDate.String() {
			minDate = v.TglParsed
		}
		if v.TglParsed.String() > maxDate.String() {
			maxDate = v.TglParsed
		}
	}

	// menyiap kan data aggregate weekly
	diffDays := maxDate.Sub(minDate).Hours() / 24
	tempMinDate := minDate
	tempWeeklyPrice := make(map[string][]float64)
	for _, v := range value {
		if int64(math.Ceil(diffDays/7)) > 1 {
			for i := 0; i <= int(math.Ceil(diffDays/7)); i++ {
				if v.TglParsed.String() >= tempMinDate.String() && v.TglParsed.String() <= tempMinDate.AddDate(0, 0, 7).String() {
					price, _ := strconv.ParseFloat(v.Price, 64)
					tempWeeklyPrice[tempMinDate.Format("2006/01/02")+"-"+tempMinDate.AddDate(0, 0, 7).Format("2006/01/02")] =
						append(tempWeeklyPrice[tempMinDate.Format("2006/01/02")+"-"+tempMinDate.AddDate(0, 0, 7).Format("2006/01/02")],
							price)
					break
				}
				tempMinDate = tempMinDate.AddDate(0, 0, 7)
			}
		} else {
			price, _ := strconv.ParseFloat(v.Price, 64)
			tempWeeklyPrice[tempMinDate.Format("2006/01/02")+"-"+tempMinDate.AddDate(0, 0, 7).Format("2006/01/02")] =
				append(tempWeeklyPrice[tempMinDate.Format("2006/01/02")+"-"+tempMinDate.AddDate(0, 0, 7).Format("2006/01/02")],
					price)

		}
	}

	// memproses aggregate weekly untuk mendapatkan min,max,med,avg
	for i, v := range tempWeeklyPrice {
		min, _ := stats.Min(v)
		max, _ := stats.Max(v)
		median, _ := stats.Median(v)
		avg, _ := stats.Mean(v)
		arrayAggregate = append(arrayAggregate, Aggregate{
			Tanggal: i,
			Min:     fmt.Sprintf("%.f", min),
			Max:     fmt.Sprintf("%.f", max),
			Median:  fmt.Sprintf("%.f", median),
			Avg:     fmt.Sprintf("%.f", avg),
		})
	}

	return arrayAggregate
}
