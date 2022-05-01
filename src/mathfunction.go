package src

import (
	"fmt"
	"log"
	"math"
)

// executeFunction executes the function based on the passed name and the value of x.
func executeFunction(functionName string, x float64) {
	switch functionName {

	case "base_2_exponential":
		fmt.Printf("Base-2 Exponential of x: %v\n", math.Exp2(x))

	case "base_e_exponential":
		fmt.Printf("Base-e Exponential of x minus 1: %v\n", math.Expm1(x))

	case "square_root":
		fmt.Printf("Square Root of x: %v\n", math.Sqrt(x))

	case "absolute":
		fmt.Printf("Absolute Value of x: %v\n", math.Abs(x))

	case "logarithm":
		fmt.Printf("Natural Logarithm of x: %v\n", math.Log(x))

	case "decimal_logarithm":
		fmt.Printf("Decimal Logarithm of x: %v\n", math.Log10(x))

	case "binary_logarithm":
		fmt.Printf("Binary Logarithm of x: %v\n", math.Log2(x))

	case "sine":
		fmt.Printf("Sine of the radian argument x: %v\n", math.Sin(x))

	case "cosine":
		fmt.Printf("Cosine of the radian argument x: %v\n", math.Cos(x))

	case "tangent":
		fmt.Printf("Tangent of the radian argument x: %v\n", math.Tan(x))

	default:
		log.Fatal("Error: Function not available!")

	}
}
