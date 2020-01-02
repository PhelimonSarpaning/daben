package daben

import (
	"strings"
	"testing"
)

//test pickYear empty input
func TestPickYearEmpty(t *testing.T) {
	//test for empty argument
	emptyResult := PickYear("")
	expectedResult := "error : String cannot be empty"
	if emptyResult[0] != expectedResult {
		t.Errorf("PickYear(\"\") failed, expected %v , got %v", "error : String cannot be empty", emptyResult[0])
	} else {
		t.Logf("PickYear(\"\") success, expected %v , got %v", "error : String cannot be empty", emptyResult[0])
	}
}

//test PickYear valid input
func TestPickYearValid(t *testing.T) {
	//test for valid argument
	validResult := PickYear("Kasner was born  on July 17, 1954 .After the 1989 fall of the Berlin Wall.Following the 2005 national elections.")
	var expectedResult = []string{"1954", "1989", "2005"}

	arrayResult := arrayCheck(validResult, expectedResult)
	if arrayResult == false {
		t.Errorf("PickYear(\"Kasner was born  on July 17, 1954 .After the 1989 fall of the Berlin Wall.Following the 2005 national elections.\") failed, expected %v , got %v", "String cannot be empty", validResult)
	} else {
		t.Logf("PickYear(\"Kasner was born  on July 17, 1954 .After the 1989 fall of the Berlin Wall.Following the 2005 national elections.\") success, expected %v , got %v", "String cannot be empty", validResult)
	}
}

// loop through the array to compare all the items in the array
func arrayCheck(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	for i, match := range a {
		//trimming spaces because it prints out with soace infront
		if strings.TrimSpace(match) != b[i] {
			return false
		}
	}
	return true
}

//test for PickMonth function
func TestPickMonthValid(t *testing.T) {
	//test for valid argument
	validResult := PickMonth("Kasner was born  on July 17, 1954 .After the 1989 fall of the Berlin Wall.Following the 2005 national elections.")
	var expectedResult = []string{"July"}

	arrayResult := arrayCheck(validResult, expectedResult)
	if arrayResult == false {
		t.Errorf("PickYear(\"Kasner was born  on July 17, 1954 .After the 1989 fall of the Berlin Wall.Following the 2005 national elections.\") failed, expected %v , got %v", "String cannot be empty", validResult)
	} else {
		t.Logf("PickYear(\"Kasner was born  on July 17, 1954 .After the 1989 fall of the Berlin Wall.Following the 2005 national elections.\") success, expected %v , got %v", "String cannot be empty", validResult)
	}
}

func TestPickDateValid(t *testing.T) {
	//test for valid argument
	validResult := PickDate("Kasner was born  on July 17, 1954 .After the 1989 fall of the Berlin Wall.Following the 2005 national elections.")
	var expectedResult = []string{"July 17, 1954", "1989", "2005"}

	arrayResult := arrayCheck(validResult, expectedResult)
	if arrayResult == false {
		t.Errorf("PickYear(\"Kasner was born  on July 17, 1954 .After the 1989 fall of the Berlin Wall.Following the 2005 national elections.\") failed, expected %v , got %v", "String cannot be empty", validResult)
	} else {
		t.Logf("PickYear(\"Kasner was born  on July 17, 1954 .After the 1989 fall of the Berlin Wall.Following the 2005 national elections.\") success, expected %v , got %v", "String cannot be empty", validResult)
	}
}
