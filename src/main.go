package main

import (
	"fmt"
	"os"
	"strconv"

	libs "./libs"
)

func main() {
	input := convertUserInput(os.Args[1], os.Args[2], os.Args[3])
	validator := libs.TriangleValidator{
		UserInput: input,
	}
	isValid := validate(&validator)
	fmt.Println("Is it a valid triangle?:", isValid)
	if isValid {
		equilateralChecker := libs.EquilateralChecker{
			UserInput: input,
		}
		isEquilateral := checkType(&equilateralChecker)
		fmt.Println("Is it a Equilateral?:", isEquilateral)

		isoscelesChecker := libs.IsoscelesChecker{
			UserInput: input,
		}
		isIsosceles := checkType(&isoscelesChecker)
		fmt.Println("Is it a Isosceles?:", isIsosceles)

		scaleneChecker := libs.ScaleneChecker{
			UserInput: input,
		}
		isScalene := checkType(&scaleneChecker)
		fmt.Println("Is it a Scalene?:", isScalene)
	}
}

func convertUserInput(a string, b string, c string) libs.UserInput {
	floatA := convertOneSide(a)
	floatB := convertOneSide(b)
	floatC := convertOneSide(c)

	return libs.UserInput{
		SideA: floatA,
		SideB: floatB,
		SideC: floatC,
	}
}

func convertOneSide(side string) float64 {
	floatSide, err := strconv.ParseFloat(side, 64)
	if err != nil {
		panic("The side A cannot be converted to float64")
	}
	return floatSide
}

func checkType(checker libs.ITypeChecker) bool {
	return checker.Check()
}

func validate(validator libs.ITriangleValidator) bool {
	return validator.Validate()
}
