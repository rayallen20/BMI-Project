package proteinProvide

var EatFood map[string]float64 = map[string]float64{
	"floor": 1,
	"upper": 1.2,
}

var DoNotEatFood map[string]float64 = map[string]float64{
	"floor": 1.2,
	"upper": 1.5,
}

func CalcProteinProvide(weight float64, isEatFood bool) (proteinProvideMap map[string]float64) {
	var floor float64
	var upper float64
	if isEatFood {
		floor = weight * EatFood["floor"]
		upper = weight * EatFood["upper"]
	} else {
		floor = weight * DoNotEatFood["floor"]
		upper = weight * DoNotEatFood["upper"]
	}

	proteinProvideMap = map[string]float64{}
	proteinProvideMap["floor"] = floor
	proteinProvideMap["upper"] = upper
	return
}
