all: clean
	env GOOS=linux GOARCH=amd64 \
		go build -o build/nbs-wikipedia-importer importer.go
	env GOOS=linux GOARCH=amd64 \
	go build -o build/nbs-wikipedia-reporter reporter.go

clean:
	rm -f build/*
