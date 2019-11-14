package lib

import (
	"fmt"
	"strconv"
)

func Round2BitFloat64(origin float64) (target float64, err error)  {
	target, err = strconv.ParseFloat(fmt.Sprintf("%.2f", origin), 64)
	return
}