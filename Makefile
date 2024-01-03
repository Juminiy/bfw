main.out: cmd/main.go
	go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o cmd/main.out $<
	go tool objdump -s "main.main" main.out > main.S