package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Rate модель данных
type Rate struct {
	RateType  Sport
	RateValue float64
}

// UnmarshalJSON - анмаршалинг для нетипичного json
func (r *Rate) UnmarshalJSON(b []byte) error {
	var x map[string]map[string]string
	err := json.Unmarshal(b, &x)
	if err != nil {
		return fmt.Errorf("parsing bla bla error | [%v]", err.Error())
	}

	if data, ok := x["lines"]; ok {
		for rType, rValue := range data {
			r.RateType = NewSport(strings.ToLower(rType))
			val, err := strconv.ParseFloat(rValue, 64)
			if err != nil {
				return fmt.Errorf("conversing bla bla error | [%v]", err.Error())
			}
			r.RateValue = val
		}
	} else {
		return fmt.Errorf("wrong response | [%v]", err.Error())
	}

	return nil
}
