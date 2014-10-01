all: clean deps
	env GOOS=linux GOARCH=amd64 \
		go build -o build/nbs-wikipedia-importer importer.go
	env GOOS=linux GOARCH=amd64 \
		go build -o build/nbs-wikipedia-reporter reporter.go

deps:
	go get gopkg.in/mgo.v2

clean:
	rm -f build/*
