Feature: Calculate prime numbers

  @calculate-primes
  Scenario Outline: Calculate primes
    When I calculate the prime numbers between <start> and <end>
    Then the calculated prime numbers should be <calculated-primes>

    Examples: Valid inputs
      | start | end  | calculated-primes                                                                              |
      | 0     | 10   | 2,3,5,7                                                                                        |
      | 0     | 100  | 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97 |

    Examples: Invalid inputs
      | start | end | calculated-primes |
      | 10    | 0   | Exception         |
      | -5    | 0   | Exception         |
