package main

import (
	"log"
	"testing"
)

func TestResponseCode(t *testing.T) {
	successFailureErrorRatio := 1
	failureRatioModulo := calculateFailureRationModulo(successFailureErrorRatio)
	responseCode := getResponseCode(1, failureRatioModulo, 200, 500)
	if responseCode != 200 {
		log.Println("responseCode is not 200")
		t.Fail()
	}

	successFailureErrorRatio = 99
	failureRatioModulo = calculateFailureRationModulo(successFailureErrorRatio)
	responseCode = getResponseCode(3, failureRatioModulo, 200, 500)
	if responseCode != 500 {
		log.Println("responseCode is not 500")
		t.Fail()
	}
}

func TestResponseCodeRatio1(t *testing.T) {
	successFailureErrorRatio := 1
	failureRatioModulo := calculateFailureRationModulo(successFailureErrorRatio)
	for i := 1; i <= 101; i++ {
		responseCode := getResponseCode(i, failureRatioModulo, 200, 500)
		if responseCode != 200 {
			log.Println("responseCode is not 200")
			t.Fail()
		}
	}
}
