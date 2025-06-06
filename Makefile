# Copyright 2017 The Kubernetes Authors.
# #
# # Licensed under the Apache License, Version 2.0 (the "License");
# # you may not use this file except in compliance with the License.
# # You may obtain a copy of the License at
# #
# #     http://www.apache.org/licenses/LICENSE-2.0
# #
# # Unless required by applicable law or agreed to in writing, software
# # distributed under the License is distributed on an "AS IS" BASIS,
# # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# # See the License for the specific language governing permissions and
# # limitations under the License.

.PHONY: test

export CONTAINER_ENGINE ?= docker

# VERSION is based on a date stamp plus the last commit
VERSION?=v$(shell date +%Y%m%d)-$(shell git describe --tags)
BRANCH?=$(shell git branch --show-current)
SHA1?=$(shell git rev-parse HEAD)
BUILD=$(shell date +%FT%T%z)
LDFLAG_LOCATION=sigs.k8s.io/descheduler/pkg/version
ARCHS = amd64 arm arm64

LDFLAGS=-ldflags "-X ${LDFLAG_LOCATION}.version=${VERSION} -X ${LDFLAG_LOCATION}.buildDate=${BUILD} -X ${LDFLAG_LOCATION}.gitbranch=${BRANCH} -X ${LDFLAG_LOCATION}.gitsha1=${SHA1}"

GOLANGCI_VERSION := v1.64.8
HAS_GOLANGCI := $(shell ls _output/bin/golangci-lint 2> /dev/null)

GOFUMPT_VERSION := v0.7.0
HAS_GOFUMPT := $(shell command -v gofumpt 2> /dev/null)

GO_VERSION := $(shell (command -v jq > /dev/null && (go mod edit -json | jq -r .Go)) || (sed -En 's/^go (.*)$$/\1/p' go.mod))

# REGISTRY is the container registry to push
# into. The default is to push to the staging
# registry, not production.
REGISTRY?=gcr.io/k8s-staging-descheduler

# IMAGE is the image name of descheduler
IMAGE:=descheduler:$(VERSION)

# IMAGE_GCLOUD is the image name of descheduler in the remote registry
IMAGE_GCLOUD:=$(REGISTRY)/descheduler:$(VERSION)

# CURRENT_DIR is the current dir where the Makefile exists
CURRENT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))


# TODO: upload binaries to GCS bucket
#
# In the future binaries can be uploaded to
# GCS bucket gs://k8s-staging-descheduler.

HAS_HELM := $(shell which helm 2> /dev/null)

all: build

build:
	CGO_ENABLED=0 go build ${LDFLAGS} -o _output/bin/descheduler sigs.k8s.io/descheduler/cmd/descheduler

build.amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o _output/bin/descheduler sigs.k8s.io/descheduler/cmd/descheduler

build.arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build ${LDFLAGS} -o _output/bin/descheduler sigs.k8s.io/descheduler/cmd/descheduler

build.arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build ${LDFLAGS} -o _output/bin/descheduler sigs.k8s.io/descheduler/cmd/descheduler

dev-image: build
	$(CONTAINER_ENGINE) build -f Dockerfile.dev -t $(IMAGE) .

image:
	$(CONTAINER_ENGINE) build --build-arg VERSION="$(VERSION)" --build-arg ARCH="amd64" -t $(IMAGE) .

image.amd64:
	$(CONTAINER_ENGINE) build --build-arg VERSION="$(VERSION)" --build-arg ARCH="amd64" -t $(IMAGE)-amd64 .

image.arm:
	$(CONTAINER_ENGINE) build --build-arg VERSION="$(VERSION)" --build-arg ARCH="arm" -t $(IMAGE)-arm .

image.arm64:
	$(CONTAINER_ENGINE) build --build-arg VERSION="$(VERSION)" --build-arg ARCH="arm64" -t $(IMAGE)-arm64 .

push: image
	gcloud auth configure-docker
	$(CONTAINER_ENGINE) tag $(IMAGE) $(IMAGE_GCLOUD)
	$(CONTAINER_ENGINE) push $(IMAGE_GCLOUD)

push-all: image.amd64 image.arm image.arm64
	gcloud auth configure-docker
	for arch in $(ARCHS); do \
		$(CONTAINER_ENGINE) tag $(IMAGE)-$${arch} $(IMAGE_GCLOUD)-$${arch} ;\
		$(CONTAINER_ENGINE) push $(IMAGE_GCLOUD)-$${arch} ;\
	done
	DOCKER_CLI_EXPERIMENTAL=enabled $(CONTAINER_ENGINE) manifest create $(IMAGE_GCLOUD) $(addprefix --amend $(IMAGE_GCLOUD)-, $(ARCHS))
	for arch in $(ARCHS); do \
		DOCKER_CLI_EXPERIMENTAL=enabled $(CONTAINER_ENGINE) manifest annotate --arch $${arch} $(IMAGE_GCLOUD) $(IMAGE_GCLOUD)-$${arch} ;\
	done
	DOCKER_CLI_EXPERIMENTAL=enabled $(CONTAINER_ENGINE) manifest push $(IMAGE_GCLOUD) ;\

clean:
	rm -rf _output
	rm -rf _tmp

verify: verify-govet verify-spelling verify-gofmt verify-vendor lint lint-chart verify-gen

verify-govet:
	./hack/verify-govet.sh

verify-spelling:
	./hack/verify-spelling.sh

verify-gofmt:
	./hack/verify-gofmt.sh

verify-vendor:
	./hack/verify-vendor.sh

verify-docs:
	./hack/verify-docs.sh

test-unit:
	./test/run-unit-tests.sh

test-e2e:
	./test/run-e2e-tests.sh

gen:
	./hack/update-generated-conversions.sh
	./hack/update-generated-deep-copies.sh
	./hack/update-generated-defaulters.sh
	./hack/update-docs.sh

gen-docker:
	$(CONTAINER_ENGINE) run --entrypoint make -it -v $(CURRENT_DIR):/go/src/sigs.k8s.io/descheduler -w /go/src/sigs.k8s.io/descheduler golang:$(GO_VERSION) gen

verify-gen:
	./hack/verify-conversions.sh
	./hack/verify-deep-copies.sh
	./hack/verify-defaulters.sh
	./hack/verify-docs.sh

lint:
ifndef HAS_GOLANGCI
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./_output/bin ${GOLANGCI_VERSION}
endif
	./_output/bin/golangci-lint run -v

fmt:
ifndef HAS_GOFUMPT
	go install mvdan.cc/gofumpt@${GOFUMPT_VERSION}
endif
	gofumpt -w -extra .

# helm

ensure-helm-install:
ifndef HAS_HELM
	curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 && chmod 700 ./get_helm.sh && ./get_helm.sh
endif

lint-chart: ensure-helm-install
	helm lint ./charts/descheduler

build-helm:
	helm package ./charts/descheduler --dependency-update --destination ./bin/chart

test-helm: ensure-helm-install
	./test/run-helm-tests.sh

kind-multi-node:
	kind create cluster --name kind --config ./hack/kind_config.yaml --wait 2m

ct-helm:
	./hack/verify-chart.sh
