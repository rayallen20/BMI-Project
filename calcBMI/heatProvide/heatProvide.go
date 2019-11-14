// 本文件用于计算热量供给相关信息
package heatProvide

const CoefficientHeatProvide float64 = 0.7

var RecordHeatProvide map[string]map[string]map[string]float64 = map[string]map[string]map[string]float64{
	"veryLight": map[string]map[string]float64{
		"light": map[string]float64{
			"floor": 25,
			"upper": 30,
		},
		"normal": map[string]float64{
			"floor": 20,
			"upper": 25,
		},
		"overweight": map[string]float64{
			"floor": 15,
			"upper": 20,
		},
	},

	"light": map[string]map[string]float64{
		"light": map[string]float64{
			"floor": 30,
			"upper": 35,
		},
		"normal": map[string]float64{
			"floor": 25,
			"upper": 30,
		},
		"overweight": map[string]float64{
			"floor": 20,
			"upper": 25,
		},
	},

	"middle": map[string]map[string]float64{
		"light": map[string]float64{
			"floor": 35,
			"upper": 40,
		},
		"normal": map[string]float64{
			"floor": 30,
			"upper": 35,
		},
		"overweight": map[string]float64{
			"floor": 25,
			"upper": 30,
		},
	},

	"heavy": map[string]map[string]float64{
		"light": map[string]float64{
			"floor": 40,
			"upper": 45,
		},
		"normal": map[string]float64{
			"floor": 35,
			"upper": 40,
		},
		"overweight": map[string]float64{
			"floor": 30,
			"upper": 35,
		},
	},
}

func CalcHeatProvide(intensity, bodilyForm string, weight float64) (heatProvide map[string]float64) {
	floor := RecordHeatProvide[intensity][bodilyForm]["floor"]
	upper := RecordHeatProvide[intensity][bodilyForm]["upper"]

	heatProvide = map[string]float64{
		"floor": weight * floor * CoefficientHeatProvide,
		"upper": weight * upper * CoefficientHeatProvide,
	}

	return heatProvide
}
