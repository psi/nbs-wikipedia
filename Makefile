all:
	go build -o build/nbs-wikipedia-importer importer.go
	go build -o build/nbs-wikipedia-reporter reporter.go
