.PHONY: configmap build run clean

BUILD_DIR ?=build

namespace:
	kubectl create ns smartedge-system

build:
	docker build -t my-calico .

run:
	kind load docker-image my-calico
	kubectl -n smartedge-system apply -f deployment.yaml

clean:
	kubectl -n smartedge-system delete deployment rad-calico-demo
	kubectl delete ns smartedge-system
