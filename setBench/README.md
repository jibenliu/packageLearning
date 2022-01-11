
go test -v -bench=. -count=3 -benchmem | tee result.txt



go get -u golang.org/x/perf/cmd/benchstat



benchstat result.txt