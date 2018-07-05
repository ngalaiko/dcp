Problem
-------

This problem was asked by Stripe.

Given an array of integers, find the first missing positive integer in linear time and constant space.
In other words, find the lowest positive integer that does not exist in the array.
The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.


bench
-----
```
goos: darwin
goarch: amd64
pkg: github.com/ngalayko/dcp/problems/2018-07-05
Benchmark/0-4   1000000000               2.18 ns/op            0 B/op          0 allocs/op
Benchmark/100-4                 20000000                65.0 ns/op             0 B/op          0 allocs/op
Benchmark/200-4                 10000000               122 ns/op               0 B/op          0 allocs/op
Benchmark/300-4                 10000000               178 ns/op               0 B/op          0 allocs/op
Benchmark/400-4                 10000000               238 ns/op               0 B/op          0 allocs/op
Benchmark/500-4                  5000000               295 ns/op               0 B/op          0 allocs/op
Benchmark/600-4                  5000000               354 ns/op               0 B/op          0 allocs/op
Benchmark/700-4                  3000000               410 ns/op               0 B/op          0 allocs/op
Benchmark/800-4                  3000000               466 ns/op               0 B/op          0 allocs/op
Benchmark/900-4                  3000000               528 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/ngalayko/dcp/problems/2018-07-05     19.293s
```
