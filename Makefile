all: clean deps
	env GOOS=linux GOARCH=amd64 \
		go build -o build/nbs-wikipedia-importer importer.go
	env GOOS=linux GOARCH=amd64 \
		go build -o build/nbs-wikipedia-reporter reporter.go

deps:
	go get gopkg.in/mgo.v2

release: all
	rm build/.gitkeep
	tar -C build -cpzvf nbs-wikipedia.tar.gz .
	touch build/.gitkeep


clean:
	rm -f build/*
	rm -f nbs-wikipedia.tar.gz
