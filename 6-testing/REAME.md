Simple: \
`go test .` 

Verbose: \
`go test -v .` 

Check coverage: \
`go test -coverprofile=coverage.out` 

Check coverage and shows the uncover code part: \
`go tool cover -html=coverage.out` 

Test with fuzzing, testing only fuzzing test: \
`go test -fuzz=. -run=^#` 


Test with fuzzing with fuzztime
`go test -fuzz=. -fuzztime=5s -run=^#`