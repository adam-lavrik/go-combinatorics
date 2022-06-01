# go-combinatorics
Package containing efficient implementation of combinatorial functions in Go.

# Summary
All 

# Types
`T` (alias for `uint64`) - type of combinatorial function parameters and results.

# Functions

## Factorial
Factorial of n; is equal to number of permutations of n objects.
```
n! = n * (n - 1) * (n - 2) * (n - 3) * ... * 1
0! = 1
```

## Subfactorial
Subfactorial of n; is equal to number of derangements of n objects.
```
!n = n! / 0! - n! / 1! + n! / 2! - ... +\- n! / (n - 1)! -\+ n! / n!
!0 = 1
```

## Superfactorial
Superfactorial of n.
```
sf(n) = 1! * 2! * ... * n!
sf(0) = 1
```

## Multifactorial
Multifactorial of n.
```
n!...(m) = n * (n - m) * (n - 2m) * ... (while (n - km) > 0)
n!...(0) = 0
0!...(m) = 1
```

## RisingFactorial
Rising factorial of n by m.
```
n!^m = n * (n + 1) * ... * (n + m - 1)
0!^0 = 0
```

## PCount
Number of permutations of m objects from n without repetitions; is equal to falling factorial of n by m.
```
P(n, m) = n! / (n - m)! = n * (n - 1) * ... * (n - m + 1)
P(n, m) = 0 if n < m
```

## PrCount
Number of permutations of m objects from n with repetitions.
```
Pr(m0, m1, ...) = (m0 + m1 + ...)! / m0! / m1! / ...
Pr() = Pr(0, 0, ...) = 1
```

## CCount
Number of combinations of m objects from n without repetitions.
```
C(n, m) = n! / (n - m)! / m! = n / 1 * (n - 1) / 2 * ... * (n - m + 1) / n
C(n, m) = 0 if n < m
```

## CrCount
Number of combinations of m objects from n with repetitions.
```
Cr(n, m) = C(n + m - 1, m)
Cr(0, 0) = 0
```

## WCount
Number of m-letter words of n-letter alphabet.
```
W(n, m) = n ** m
W(0, 0) = 1
```
