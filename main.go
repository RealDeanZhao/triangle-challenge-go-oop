package main

import (
	"fmt"
	"os"
	"strconv"
)

type triangle struct {
	a, b, c float64
	t       string
}

type userInput struct {
	a, b, c float64
}

type typeChecker interface {
	check(u userInput) bool
}

type equilateralChecker struct {
	u userInput
}

func (e *equilateralChecker) check() bool {
	if allSidesAreEqual(e.u) {
		return true
	}
	return false
}

type isoscelesChecker struct {
	u userInput
}

func (i *isoscelesChecker) check() bool {
	if !allSidesAreEqual(i.u) {
		if twoSidesAreEqual(i.u) {
			return true
		}
	}
	return false
}

type scaleneChecker struct {
	u userInput
}

func (s *scaleneChecker) check() bool {
	if !allSidesAreEqual(s.u) && !twoSidesAreEqual(s.u) {
		return true
	}
	return false
}

func allSidesAreEqual(u userInput) bool {
	if u.a != u.b || u.a != u.c || u.b != u.c {
		return false
	}
	return true
}

func twoSidesAreEqual(u userInput) bool {
	if u.a == u.b || u.a == u.c || u.b == u.c {
		return true
	}
	return false
}

type triangleValidator interface {
	validate() bool
}

type validator struct {
	u userInput
}

func (v *validator) validate() bool {
	if !allSidesArePositive(v.u) {
		return false
	}
	if !twoSidesAreGreaterThanTheThird(v.u) {
		return false
	}
	return true
}

func allSidesArePositive(u userInput) bool {
	if u.a < 0 || u.b < 0 || u.c < 0 {
		return false
	}
	return true
}

func twoSidesAreGreaterThanTheThird(u userInput) bool {
	if u.a > u.b-u.c && u.b > u.a-u.c && u.b > u.c-u.a {
		return true
	}
	return false
}

func convertUserInput(a string, b string, c string) userInput {
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
	return userInput{
		a: floatA,
		b: floatB,
		c: floatC,
	}
}

func main() {
	input := convertUserInput(os.Args[1], os.Args[2], os.Args[3])
	validator := validator{
		u: input,
	}
	isValid := validator.validate()
	fmt.Println("Is it a valid triangle?:", isValid)
	if isValid {
		equilateralChecker := equilateralChecker{
			u: input,
		}
		isEquilateral := equilateralChecker.check()
		fmt.Println("Is it a Equilateral?:", isEquilateral)

		isoscelesChecker := isoscelesChecker{
			u: input,
		}
		isIsosceles := isoscelesChecker.check()
		fmt.Println("Is it a Isosceles?:", isIsosceles)

		scaleneChecker := scaleneChecker{
			u: input,
		}
		isScalene := scaleneChecker.check()
		fmt.Println("Is it a Scalene?:", isScalene)
	}
}
