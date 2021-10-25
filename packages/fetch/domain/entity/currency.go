package entity

type Currency struct {
	Base   string `json:"base"`
	Result struct {
		USD float64 `json:"USD"`
	} `json:"result"`
}
