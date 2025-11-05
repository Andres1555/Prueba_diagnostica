package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func pascalRow(n int) []int {
	row := make([]int, n+1)
	row[0] = 1
	for i := 1; i <= n; i++ {
		row[i] = row[i-1] * (n - i + 1) / i
	}
	return row
}

func evaluatePolynomial(coeffs []int, x int) int {
	result := 0
	for i, c := range coeffs {
		result += c * int(math.Pow(float64(x), float64(i)))
	}
	return result
}

func main() {
	start := time.Now()
	n := 100
	x := 2
	coeffs := pascalRow(n)
	result := evaluatePolynomial(coeffs, x)

	file, _ := os.Create("Resultado1go.txt")
	defer file.Close()

	fmt.Fprintf(file, "Coeficientes: %v\n", coeffs)
	fmt.Fprintf(file, "f(%d) = %d\n", x, result)
	fmt.Fprintf(file, "Tiempo: %v\n", time.Since(start))
}
