test:
	go test -v
	go build
	go run main.go README.md > out.js
	node --version
	node test.js

clean:
	rm -f out.js
