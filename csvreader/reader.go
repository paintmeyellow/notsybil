package csvreader

import (
	"encoding/csv"
	"errors"
	"os"

	"notsybil/asset"
)

var ErrInvalidFormat = errors.New("invalid csv format")

func Parse(filename string) ([]*asset.Asset, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var assets []*asset.Asset

	for _, record := range records {
		if len(record) != 4 {
			return nil, ErrInvalidFormat
		}
		amt := record[0]
		ccy := record[1]
		chain := record[2]
		addr := record[3]

		a, err := asset.ToAsset(amt, ccy, chain, addr)
		if err != nil {
			return nil, err
		}
		assets = append(assets, a)
	}

	return assets, nil
}
