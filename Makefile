
docker-build:
	docker build --no-cache -t gcr.io/mchirico/knights-tour:test -f Dockerfile .

push:
	docker push gcr.io/mchirico/knights-tour:test

build:
	go build -v .

run:
	docker run --rm -it -p 3000:3000  gcr.io/mchirico/knights-tour:test
