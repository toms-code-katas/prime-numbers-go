package primenumbers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	// "github.com/cucumber/godog"
)

func TestCalculatePrimes(t *testing.T) {
	expected_primes := []uint64{2, 3, 5, 7}
	primes := CalculatePrimes(1, 10)
	for i, prime := range primes {
		if expected_primes[i] != prime {
			t.Fatal(fmt.Sprintf("Expected primes %v and calculated primes %v are not equal", expected_primes, primes))
		}
	}
}

func TestCalculatePrimesUsingTable(t *testing.T) {

	type test struct {
		start  uint64
		stop   uint64
		primes []uint64
	}

	tests := []test{
		{start: 1, stop: 10, primes: []uint64{2, 3, 5, 7}},
		{start: 1, stop: 25, primes: []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23}},
	}

	for _, tc := range tests {
		calculatedPrimes := CalculatePrimes(tc.start, tc.stop)
		if !reflect.DeepEqual(tc.primes, calculatedPrimes) {
			t.Fatalf("expected: %v, got: %v", tc.primes, calculatedPrimes)
		}
	}
}

type PrimeCalculation struct {
	primes []uint64
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func convertToInts(commaSepeartedString string) []uint64 {
	var converted = []uint64{}
	for _, intStr := range strings.Split(commaSepeartedString, ",") {
		j, err := strconv.Atoi(strings.TrimSpace(intStr))
		if err != nil {
			panic(err)
		}
		converted = append(converted, uint64(j))
	}
	return converted
}

func (primeCalculation *PrimeCalculation) calculatedPrimes(start, stop int) error {
	primeCalculation.primes = CalculatePrimes(uint64(start), uint64(stop))
	return nil
}

func (primeCalculation *PrimeCalculation) calculatedPrimesShouldBe(expectedPrimes string) error {
	var expectedPrimesAsInts = convertToInts(expectedPrimes)
	primes := CalculatePrimes(1, 10)
	for i, prime := range primeCalculation.primes {
		if expectedPrimesAsInts[i] != prime {
			return fmt.Errorf("Expected primes %v and calculated primes %v are not equal", expectedPrimesAsInts, primes)
		}
	}
	return nil
}

func (primeCalculation *PrimeCalculation) primeCalculationThrowsException() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	primeCalculation := &PrimeCalculation{}

	ctx.Step(`^I calculate the prime numbers between (-?[0-9]{0,10}) and (-?[0-9]{0,10})$`, primeCalculation.calculatedPrimes)
	ctx.Step(`^the calculated prime numbers should be (\d.*)$`, primeCalculation.calculatedPrimesShouldBe)
	ctx.Step(`^the calculated prime numbers should be Exception$`, primeCalculation.primeCalculationThrowsException)
}
