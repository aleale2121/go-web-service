SHELL := /bin/bash

run:
	go run main.go
# ==============================================================================
# Building containers

# Example: $(shell git rev-parse --short HEAD)
VERSION := 1.0

all: service

service:
	docker build \
		-f zarf/docker/dockerfile \
		-t service:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Running from within k8s/kind

KIND_CLUSTER := sales-api-starter-cluster

kind-up:
	kind create cluster \
		--image kindest/node:v1.25.3@sha256:f52781bc0d7a19fb6c405c2af83abfeb311f130707a0e219175677e366cc45d1 \
		--name $(KIND_CLUSTER) 

dev-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

# ------------------------------------------------------------------------------


