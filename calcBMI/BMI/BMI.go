// 本文件处理BMI计算相关功能
package BMI

func CalcBMI(height, weight float64) (BMI float64) {
	BMI = weight / (height * height)
	return
}
