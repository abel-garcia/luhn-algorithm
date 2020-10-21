# Luhn algorithm
This is an example of Luhn's algorithm in golang.
## Description
The Luhn algorithm or Luhn formula verifies a number against its included check digit, which is usually appended to a partial account number to generate the full account number. This number must pass the following test:

1. From the rightmost digit (excluding the check digit) and moving left, double the value of every second digit. The check digit is neither doubled nor included in this calculation; the first digit doubled is the digit located immediately left of the check digit. If the result of this doubling operation is greater than 9 (e.g., 8 × 2 = 16), then add the digits of the result (e.g., 16: 1 + 6 = 7, 18: 1 + 8 = 9) or, alternatively, the same final result can be found by subtracting 9 from that result (e.g., 16: 16 − 9 = 7, 18: 18 − 9 = 9).

2. Take the sum of all the digits (including the check digit).

3. If the total modulo 10 is equal to 0 (if the total ends in zero) then the number is valid according to the Luhn formula; otherwise, it is not valid.

[more..](https://en.wikipedia.org/wiki/Luhn_algorithm) 


## Running algorithm

Install golang and run the next command.

```bash
go run luhn.go
```