MIXTOOL_BINARY:=$(GOPATH)/bin/mixtool
JB_BINARY:=$(GOPATH)/bin/jb


.PHONY: all
all: kube-prometheus/manifests sample-app/.docker-image

sample-app/sample-app: sample-app/main.go
	go build -o sample-app/sample-app sample-app/main.go

sample-app/.docker-image: sample-app/sample-app
	docker build -t quay.io/mxinden/self-service-monitoring-sample-app:v$(shell cat sample-app/VERSION) -f sample-app/Dockerfile/ sample-app
	touch $@

kube-prometheus/manifests: $(MIXTOOL_BINARY) kube-prometheus/vendor kube-prometheus/example.jsonnet
	cd kube-prometheus && mixtool build -m manifests/ example.jsonnet

kube-prometheus/vendor: $(JB_BINARY) kube-prometheus/jsonnetfile.json kube-prometheus/jsonnetfile.lock.json
	cd kube-prometheus && $(JB_BINARY) install

.PHONY: clean
clean:
	rm sample-app/.docker-image
	rm -r kube-prometheus/vendor
	rm -r kube-prometheus/manifests


# Binaries

$(JB_BINARY):
	go get -u github.com/jsonnet-bundler/jsonnet-bundler/cmd/jb

$(MIXTOOL_BINARY):
	go get -u github.com/metalmatze/mixtool/cmd/mixtool
