# uaparser-bench

Tested on iMac, 3.2 GHz Intel Core i5. Equality is not tested on this benchmark.

    -> % go test -v -bench . -benchmem
    goos: darwin
    goarch: amd64
    pkg: github.com/suzuken/uaparser-bench
    BenchmarkWoothee/ios-4            500000              2636 ns/op             182 B/op          5 allocs/op
    BenchmarkWoothee/android-4        500000              2615 ns/op             176 B/op          3 allocs/op
    BenchmarkWoothee/kddi-4          1000000              1662 ns/op             144 B/op          2 allocs/op
    BenchmarkWoothee/googlebot-4    10000000               158 ns/op             112 B/op          1 allocs/op
    BenchmarkWoothee/vivaldi-4       1000000              2104 ns/op             176 B/op          3 allocs/op
    BenchmarkUAParser/ios-4          1000000              2265 ns/op             224 B/op          2 allocs/op
    BenchmarkUAParser/android-4       300000              3958 ns/op             224 B/op          2 allocs/op
    BenchmarkUAParser/kddi-4         1000000              1621 ns/op             144 B/op          2 allocs/op
    BenchmarkUAParser/googlebot-4     500000              3403 ns/op             160 B/op          2 allocs/op
    BenchmarkUAParser/vivaldi-4      1000000              1153 ns/op             224 B/op          2 allocs/op
    BenchmarkMssola/ios-4             500000              2614 ns/op            1768 B/op         38 allocs/op
    BenchmarkMssola/android-4         500000              2420 ns/op            1768 B/op         38 allocs/op
    BenchmarkMssola/kddi-4           1000000              1306 ns/op             848 B/op         20 allocs/op
    BenchmarkMssola/googlebot-4      1000000              1814 ns/op             576 B/op         14 allocs/op
    BenchmarkMssola/vivaldi-4         500000              2649 ns/op            1848 B/op         40 allocs/op
    BenchmarkXojoc/ios-4              300000              5326 ns/op            1676 B/op         30 allocs/op
    BenchmarkXojoc/android-4          500000              3623 ns/op            1672 B/op         26 allocs/op
    BenchmarkXojoc/kddi-4             500000              2401 ns/op            2664 B/op         33 allocs/op
    BenchmarkXojoc/googlebot-4       2000000               891 ns/op             565 B/op         11 allocs/op
    BenchmarkXojoc/vivaldi-4          500000              3881 ns/op            1672 B/op         26 allocs/op
    PASS
    ok      github.com/suzuken/uaparser-bench       32.790s


## LICENSE

Public Domain
