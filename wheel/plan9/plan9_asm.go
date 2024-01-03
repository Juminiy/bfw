package plan9

// go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o main.out main.go
// go tool objdump -s "main.main" main.out > main.S
// go tool objdump -s "main.addInt" main.out > addInt.S
