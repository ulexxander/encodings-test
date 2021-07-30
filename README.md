### Just played a bit with jwt and gob

Did not expect gob to be that slow, hm.

This jwt library is pretty cool, API is great.

- `encoding/gob`
- `github.com/gbrlsnchs/jwt/v3`

```
go test -benchmem -cpuprofile cpu.out -bench .

goos: linux
goarch: amd64
pkg: gitlab.com/ulexxander/encodingz
cpu: AMD Ryzen 7 2700 Eight-Core Processor
BenchmarkJWTGbrlsnchs-16          271519              4291 ns/op            1002 B/op         21 allocs/op
BenchmarkGob-16                    53192             22420 ns/op            7641 B/op        189 allocs/op
PASS
ok      gitlab.com/ulexxander/encodingz 2.817s
```
