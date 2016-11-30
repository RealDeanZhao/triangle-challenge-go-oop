package main

import (
	"fmt"
	"os"
	"strconv"

	libs "./libs"
)

func convertUserInput(a string, b string, c string) libs.UserInput {
	floatA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		panic("The side A cannot be converted to float64")
	}
	floatB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		panic("The side B cannot be converted to float64")
	}
	floatC, err := strconv.ParseFloat(c, 64)
	if err != nil {
		panic("The side C cannot be converted to float64")
	}
	return libs.UserInput{
		SideA: floatA,
		SideB: floatB,
		SideC: floatC,
	}
}

func main() {
	input := convertUserInput(os.Args[1], os.Args[2], os.Args[3])
	validator := libs.Validator{
		UserInput: input,
	}
	isValid := validator.Validate()
	fmt.Println("Is it a valid triangle?:", isValid)
	if isValid {
		equilateralChecker := libs.EquilateralChecker{
			UserInput: input,
		}
		isEquilateral := equilateralChecker.Check()
		fmt.Println("Is it a Equilateral?:", isEquilateral)

		isoscelesChecker := libs.IsoscelesChecker{
			UserInput: input,
		}
		isIsosceles := isoscelesChecker.Check()
		fmt.Println("Is it a Isosceles?:", isIsosceles)

		scaleneChecker := libs.ScaleneChecker{
			UserInput: input,
		}
		isScalene := scaleneChecker.Check()
		fmt.Println("Is it a Scalene?:", isScalene)
	}
}
