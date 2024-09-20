# Variadic vs Slice

Test benchmark for variadic vs slice argument in function.

```sh
go test -bench=. -benchmem -memprofile mem.prof -cpuprofile cpu.prof -benchtime=100x
```

```sh
go test -bench ^BenchmarkVariadic$ -benchmem -memprofile variadic_mem.prof -cpuprofile variadic_cpu.prof -benchtime=5s > variadic.bench
go test -bench ^BenchmarkSlice$ -benchmem -memprofile mem.prof -cpuprofile cpu.prof -benchtime=5s > slice.bench
```

```sh
go tool pprof -http=:8080 -no_browser variadic.prof
```
