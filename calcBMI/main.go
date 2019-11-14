package main

import (
	"calcBMI/BMI"
	"calcBMI/BodilyForm"
	"calcBMI/heatProvide"
	"calcBMI/input"
	"calcBMI/lib"
	"calcBMI/proteinProvide"
	"fmt"
)

func main() {
	for {
		// 身高
		height, isBreak := input.HeightHandle()
		if isBreak {
			fmt.Println("bye")
			break
		}
		height = height / input.CMToMrRatio

		// 体重
		weight, isBreak := input.WeightHandle()
		if isBreak {
			fmt.Println("bye")
			break
		}

		// 计算BMI
		var bmi float64 = BMI.CalcBMI(height, weight)
		bmi, err := lib.Round2BitFloat64(bmi)
		if err != nil {
			fmt.Printf("round BMI error: %v\n", err)
		}
		fmt.Printf("BMI = %v\n", bmi)

		// 计算体型
		bodilyForm, err := BodilyForm.CalcBodilyFormByBMI(bmi)
		if err != nil {
			fmt.Printf("calc bodilyForm failed, error: %v\n", err)
			break
		}
		fmt.Printf("bodilyForm = %s\n", bodilyForm)

		// 选择劳动强度
		intensity, isBreak, err := input.IntensityHandle()
		if isBreak {
			fmt.Println("bye")
			break
		}
		if err != nil {
			fmt.Printf("err = %v\n", err)
			break
		}

		// 计算热量供给
		heatProvideMap := heatProvide.CalcHeatProvide(intensity, bodilyForm, weight)

		heatProvideMap["floor"], err = lib.Round2BitFloat64(heatProvideMap["floor"])
		if err != nil {
			fmt.Printf("round heat provide floor error: %v\n", err)
		}

		heatProvideMap["upper"], err = lib.Round2BitFloat64(heatProvideMap["upper"])
		if err != nil {
			fmt.Printf("round heat provide upper error: %v\n", err)
		}

		fmt.Printf("heat provide floor = %v\n", heatProvideMap["floor"])
		fmt.Printf("heat provide upper = %v\n", heatProvideMap["upper"])

		// 选择是否断主食
		isEatStapleFood, isBreak, err := input.IsEatStapleFoodHandle()
		if isBreak {
			fmt.Println("bye")
			break
		}
		if err != nil {
			fmt.Printf("err = %v\n", err)
			break
		}

		// 计算蛋白质供给范围
		proteinProvideMap := proteinProvide.CalcProteinProvide(weight, isEatStapleFood)

		proteinProvideMap["floor"], err = lib.Round2BitFloat64(proteinProvideMap["floor"])
		if err != nil {
			fmt.Printf("round protein provide floor error: %v\n", err)
		}

		proteinProvideMap["upper"], err = lib.Round2BitFloat64(proteinProvideMap["upper"])
		if err != nil {
			fmt.Printf("round protein provide upper error: %v\n", err)
		}

		fmt.Printf("protein provide floor = %v\n", proteinProvideMap["floor"])
		fmt.Printf("protein provide upper = %v\n", proteinProvideMap["upper"])
	}
}
