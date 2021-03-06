package primenumbers_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	//nolint:gci
	primenumbers "prime-numbers-go"

	"github.com/cucumber/godog"
)

type PrimeCalculation struct {
	primes []uint64
}

func TestFeatures(t *testing.T) {
	t.Parallel()

	suite := godog.TestSuite{
		Name:                 "BDD Tests",
		ScenarioInitializer:  InitializeScenario,
		TestSuiteInitializer: nil,
		//nolint:exhaustivestruct
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

func convertToIntegers(commaSeparatedString string) []uint64 {
	converted := []uint64{}

	for _, intStr := range strings.Split(commaSeparatedString, ",") {
		j, err := strconv.Atoi(strings.TrimSpace(intStr))
		if err != nil {
			panic(err)
		}

		converted = append(converted, uint64(j))
	}

	return converted
}

func (primeCalculation *PrimeCalculation) calculatedPrimes(start, stop int) error {
	primeCalculation.primes = primenumbers.CalculatePrimes(uint64(start), uint64(stop))

	return nil
}

func (primeCalculation *PrimeCalculation) calculatedPrimesShouldBe(expectedPrimes string) error {
	expectedPrimesAsIntegers := convertToIntegers(expectedPrimes)

	for i, prime := range primeCalculation.primes {
		if expectedPrimesAsIntegers[i] != prime {
			//nolint:goerr113
			return fmt.Errorf("expected primes %v and calculated primes %v are not equal",
				expectedPrimesAsIntegers, primeCalculation.primes)
		}
	}

	return nil
}

func (primeCalculation *PrimeCalculation) primeCalculationThrowsException() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	primeCalculation := &PrimeCalculation{[]uint64{}}

	ctx.Step(`^I calculate the prime numbers between (-?[0-9]{0,10}) and (-?[0-9]{0,10})$`,
		primeCalculation.calculatedPrimes)
	ctx.Step(`^the calculated prime numbers should be (\d.*)$`,
		primeCalculation.calculatedPrimesShouldBe)
	ctx.Step(`^the calculated prime numbers should be Exception$`,
		primeCalculation.primeCalculationThrowsException)
}
