// 本文件处理体型计算相关功能
package BodilyForm

import "fmt"

var RecordBodilyForm map[string]map[string]float64 = map[string]map[string]float64 {
	"light": map[string]float64{
		"floor": 0,
		"upper": 18.4,
	},
	"normal": map[string]float64{
		"floor": 18.5,
		"upper": 23.9,
	},
	"overweight": map[string]float64{
		"floor": 24.0,
	},
}

func CalcBodilyFormByBMI(BMI float64) (bodilyForm string, err error) {
	for bodilyForm, scope := range RecordBodilyForm {
		// 除超重的map外 其他map均有上下限
		if scope["upper"] != 0 {
			if BMI >= scope["floor"] && BMI <= scope["upper"] {
				return bodilyForm, nil
			}
		} else {
			// 对于超重map 只比较BMI与floor的大小即可
			if BMI >= scope["floor"] {
				return bodilyForm, nil
			}
		}
	}
	err = fmt.Errorf("error:BMI = %f is't in record scope.\n", BMI)
	return "", err
}