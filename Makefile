main: cmd/main.go
	go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o cmd/main.out $<
	go tool objdump -s "main.main" main.out > main.S

main0: cmd/main.go
	go tool compile -S -c 5 -cpuprofile cpu.profile.txt $<

main_asm: cmd/main.go
	go tool