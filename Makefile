GO=$(shell which go)

dev-compile:
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo -o goleak .
dev-build: dev-compile
	docker build -t dockerq/goleak:dev -f Dockerfile.dev .
dev-push: dev-build
	docker push dockerq/goleak:dev
dev-run:
	docker run -d --name goleak --net host -v /var/run/docker.sock:/var/run/docker.sock -v /proc:/linker/proc dockerq/goleak:dev \
        /goleak -i 2000
