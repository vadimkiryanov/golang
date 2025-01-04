package coincap

import "fmt"

type AssetsResponse struct {
	Assets    []AssetData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}
type AssetResponse struct {
	Asset     AssetData `json:"data"`
	Timestamp int64     `json:"timestamp"`
}
type AssetData struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"MarketCapUsd"`
	VolumeUsd24Hr     string `json:"VolumeUsd24Hr"`
	PriceUsd          string `json:"PriceUsd"`
	ChangePercent24Hr string `json:"ChangePercent24Hr"`
	Vwap24Hr          string `json:"Vwap24Hr"`
}

func (asset AssetData) GetInfo() string {
	var result = fmt.Sprintf("[ID:%s] | [Name:%s] | [Symbol:%s]  \n", asset.ID, asset.Name, asset.Symbol)

	return result
}
