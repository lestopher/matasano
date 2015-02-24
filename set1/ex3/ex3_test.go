package ex3

import "testing"

func TestLengthCharsOnly(t *testing.T) {
	const testString1 = "hello world"
	const expectedTestString1Length = 10
	const testString2 = ",.,.,.,.,0998h"
	const expectedTestString2Length = 1
	const testString3 = "The Rain In Spain Falls Mainly On The PLAINS!!"
	const expectedTestString3Length = 36

	testString1Length := LengthCharsOnly(testString1)
	testString2Length := LengthCharsOnly(testString2)
	testString3Length := LengthCharsOnly(testString3)

	if testString1Length != expectedTestString1Length {
		t.Errorf("Expected testString1 to be %d, got %d\n", testString1Length, expectedTestString1Length)
	}

	if testString2Length != expectedTestString2Length {
		t.Errorf("Expected testString2 to be %d, got %d\n", testString2Length, expectedTestString2Length)
	}

	if testString3Length != expectedTestString3Length {
		t.Errorf("Expected testString3 to be %d, got %d\n", testString3Length, expectedTestString3Length)
	}
}

func TestChiSquareSum(t *testing.T) {
	const testString = "Defend the east wall of the castle"
	const expectedEnglishSum = 18.528310082299488
	const expectedUniformSum = 55.571428571428605

	englishDist, uniformDist := ChiSquareSum(testString)

	if englishDist != expectedEnglishSum {
		t.Errorf("Expected %f, but got %f\n", expectedEnglishSum, englishDist)
	}

	if uniformDist != expectedUniformSum {
		t.Errorf("Expected %f, but got %f\n", expectedUniformSum, uniformDist)
	}
}
