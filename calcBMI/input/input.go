// 本文件处理输入相关功能
package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var height float64
var weight float64
var intensity string
var intensitySlice []string = []string {"veryLight", "light", "middle", "heavy"}
var isEatStapleFood bool
const ExitFlag = "exit"
const CMToMrRatio = 100

// 处理身高输入
func HeightHandle() (height float64, flag bool) {
	fmt.Printf("please input height(unit:cm):")
	reader := bufio.NewReader(os.Stdin)
	inputHeight, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err = %v\n", err)
	}

	inputHeight = strings.Replace(inputHeight, "\n", "", -1)

	// 输入exit表示退出
	if inputHeight == ExitFlag {
		return 0, true
	}

	height, err = strconv.ParseFloat(inputHeight, 64)
	if err != nil {
		fmt.Printf("convert height to float64 failed, err = %v\n", err)
	}
	return height, false
}

// 处理体重输入
func WeightHandle () (weight float64, flag bool) {
	fmt.Printf("please input weight(unit:kg):")
	reader := bufio.NewReader(os.Stdin)
	inputWeight, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("err = %v\n", err)
	}

	inputWeight = strings.Replace(inputWeight, "\n", "", -1)

	// 输入exit表示退出
	if inputWeight == ExitFlag {
		return 0, true
	}

	weight, err = strconv.ParseFloat(inputWeight, 64)
	if err != nil {
		fmt.Printf("convert weight to float64 failed, err = %v\n", err)
	}
	return weight, false
}

// 处理劳动强度输入
func IntensityHandle() (intensity string, flag bool, err error) {
	fmt.Println("please choose intensity:")
	fmt.Println("1: veryLight")
	fmt.Println("2: light")
	fmt.Println("3: middle")
	fmt.Println("4: heavy")
	reader := bufio.NewReader(os.Stdin)
	inputIntensity, err := reader.ReadString('\n')
	if err != nil {
		return "", false, err
	}

	inputIntensity = strings.Replace(inputIntensity, "\n", "", -1)

	// 输入exit表示退出
	if inputIntensity == ExitFlag {
		return "", true, nil
	}

	intensityToInt, err := strconv.ParseInt(inputIntensity, 10, 64)
	intensity, err = chooseIntensity(intensityToInt)
	if err != nil {
		return "", false, err
	}

	return intensity, false, nil
}

// 选择劳动强度
func chooseIntensity(intensityToInt int64) (intensity string, err error) {
	switch intensityToInt {
	case 1:
		return intensitySlice[0], nil
	case 2:
		return intensitySlice[1], nil
	case 3:
		return intensitySlice[2], nil
	case 4:
		return intensitySlice[3], nil
	default:
		err = fmt.Errorf("input intensity isn't exist in record")
		return "", err
	}
}

// 处理是否断主食
func IsEatStapleFoodHandle() (isEatStapleFood bool, flag bool, err error) {
	fmt.Println("please choose staple food:")
	fmt.Println("1: eat")
	fmt.Println("2: don't eat")

	reader := bufio.NewReader(os.Stdin)
	inputEatStaple, err := reader.ReadString('\n')
	if err != nil {
		return false, false, err
	}

	inputEatStaple = strings.Replace(inputEatStaple, "\n", "", -1)

	// 输入exit表示退出
	if inputEatStaple == ExitFlag {
		return false, true, nil
	}

	eatStapleToInt, err := strconv.ParseInt(inputEatStaple, 10, 64)

	isEatStapleFood, err = chooseEatStaple(eatStapleToInt)
	if err != nil {
		return false, false, err
	}
	return isEatStapleFood, false, nil
}

func chooseEatStaple (eatStapleToInt int64) (bool, error) {
	switch eatStapleToInt {
	case 1:
		return true, nil
	case 2:
		return false, nil
	default:
		err := fmt.Errorf("staple food isn't in record")
		return false, err
	}
}