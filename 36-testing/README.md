# Unit Test
```go test .```

```go test . -v```

# Test Coverage
```go test -coverprofile=coverage.out```

```go tool cover -html=coverage.out```

# Benchmark Test
```go test -bench=.```

```go test -bench=. -run=ˆ#```

```go test -bench=. -run=ˆ# -count=10```

```go test -bench=. -run=ˆ# -benchmem```

# Fuzzing Test
```go test -fuzz=. -fuzztime=5s```
