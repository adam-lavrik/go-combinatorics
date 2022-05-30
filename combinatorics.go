/*
Copyright 2022 Adam Lavrik <lavrik.adam@gmail.com>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific
language governing permissions and limitations under the License.
*/
package combinatorics

type Count = uint64

// Factorial returns factorial
// n! = n * (n - 1) * (n - 2) * (n - 3) * ... * 1
// 0! = 1
func Factorial(n Count) Count {
	return product(2, n)
}

// Subfactorial returns subfactorial
// !n = n! / 0! - n! / 1! + n! / 2! - ... +\- n! / (n - 1)! -\+ 1
func Subfactorial(n Count) Count {
	// s0 = 1 + n * (n - 1) + n * (n - 1) * (n - 2) * (n - 3) + ...
	// s1 = n + n * (n - 1) * (n - 2) + ...
	// !n = s0 - s1 if n is even or s1 - s0 if n is odd.
	var p, s0, s1 Count = 1, 1, 0
	for {
		if n == 0 {
			return s0 - s1
		}
		p *= n
		s1 += p
		n--
		if n == 0 {
			return s1 - s0
		}
		p *= n
		s0 += p
		n--
	}
}

// Superfactorial returns superfactorial
// sf(n) = 1! * 2! * ... * n!
// sf(0) = 1
func Superfactorial(n Count) Count {
    var i, p, pp Count = 1, 1, 1
    for i < n {
        i++
        p *= i
        pp *= p
    }
    return pp
}

// Multifactorial returns multifactorial
// n!...(m) = n * (n - m) * (n - 2m) * ... (while (n - km) > 0)
// 0!...(m) = 1
// n!...(m) = 0 if m = 0
func Multifactorial(n, m Count) Count {
	if m == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	p := n
	for m < n {
		n -= m
		p *= n
	}
	return p
}

// RisingFactorial returns rising factorial
// n!^m = n * (n + 1) * ... * (n + m - 1)
// 0!^0 = 0
func RisingFactorial(n, m Count) Count {
	if n | m == 0 {
		return 0
	}
	var p Count = 1
	m += n
	for m > 1 {
		m--
		p *= m
	}
	return p
}

// PCount returns number of permutations without repetitions
// P(n, m) = n! / (n - m)! = n * (n - 1) * ... * (n - m + 1)
// P(n, m) = 0 if n < m
func PCount(n, m Count) Count {
	if n < m {
		return 0
	}
	return product(n - m + 1, n)
}

// PrCount returns number of permutations with repetitions
// Pr(m0, m1, ...) = (m0 + m1 + ...)! / m0! / m1! / ...
// Pr() = Pr(0, 0, ...) = 1
func PrCount(m ...Count) Count {
	i := len(m) - 1
	if i < 1 {
		return 1
	}
	var p, s Count = 1, m[i] + 1
	var sx Count
	for i > 0 {
		i--
		sx = s + m[i]
		p *= dividedProduct(s, sx - 1)
		s = sx
	}
	return p
}

// CCount returns number of combinations without repetitions
// C(n, m) = n! / (n - m)! / m! = n / 1 * (n - 1) / 2 * ... * (n - m + 1) / n
// C(n, m) = 0 if n < m
func CCount(n, m Count) Count {
	if n < m {
		return 0
	}
	return dividedProduct(n - m + 1, n)
}

// CrCount returns number of combinations with repetitions
// Cr(n, m) = C(n + m - 1, m)
// Cr(0, 0) = 0
func CrCount(n, m Count) Count {
	if n | m == 0 {
		return 0
	}
	return dividedProduct(n, n + m - 1)
}

// WCount returns number of m-letter words of n-letter alphabet 
// W(n, m) = n ** m
// W(0, 0) = 1
func WCount(n, m Count) Count {
	var p Count = 1
	for m != 0 {
		if m & 1 == 1 {
			p *= n
		}
		n *= n
		m >>= 1
	}
	return p
}

// product returns l * (l + 1) * ... * u
func product(l, u Count) Count {
	var p Count = 1
	for l <= u {
		p *= l
		l++
	}
	return p
}

// dividedProduct returns l / 1 * (l + 1) / 2 * ... * u / (u - l + 1)
func dividedProduct(l, u Count) Count {
	var p Count = 1
	var d Count = 0
	for l <= u {
		d++
		p = p * l / d
		l++
	}
	return p
}
