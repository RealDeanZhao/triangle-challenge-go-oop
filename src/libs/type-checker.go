package libs

// TypeChecker is used to check the type by the UserInput
type TypeChecker interface {
	Check() bool
}

// EquilateralChecker checks if it is a Equilateral
type EquilateralChecker struct {
	UserInput UserInput
}

// Check if it is a Equilateral
func (e *EquilateralChecker) Check() bool {
	if allSidesAreEqual(e.UserInput) {
		return true
	}
	return false
}

// IsoscelesChecker checks if it is an Isosceles
type IsoscelesChecker struct {
	UserInput UserInput
}

// Check if it is an Isosceles
func (i *IsoscelesChecker) Check() bool {
	if !allSidesAreEqual(i.UserInput) {
		if twoSidesAreEqual(i.UserInput) {
			return true
		}
	}
	return false
}

//ScaleneChecker checks if it is a Scalene
type ScaleneChecker struct {
	UserInput UserInput
}

// Check if it is a Scalene
func (s *ScaleneChecker) Check() bool {
	if !allSidesAreEqual(s.UserInput) && !twoSidesAreEqual(s.UserInput) {
		return true
	}
	return false
}

func allSidesAreEqual(u UserInput) bool {
	if u.SideA != u.SideB || u.SideA != u.SideC || u.SideB != u.SideC {
		return false
	}
	return true
}

func twoSidesAreEqual(u UserInput) bool {
	if u.SideA == u.SideB || u.SideA == u.SideC || u.SideB == u.SideC {
		return true
	}
	return false
}
