POC = $(wildcard poc/*.go)

main: clean main.go
	ln -s ../../../bin bin
	go build -o bin/$@
clean:
	rm -f bin/*