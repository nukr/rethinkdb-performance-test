VERSION_FILE=VERSION
VERSION=`cat $(VERSION_FILE)`

build:
	GOOS=linux go build
docker:
	docker build -t asia.gcr.io/instant-matter-785/rethinkdb_perf_tool:$(VERSION) .
