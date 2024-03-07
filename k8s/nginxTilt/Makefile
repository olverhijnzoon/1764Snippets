run: context
	tilt up

context: create
	kubectl cluster-info --context kind-nginx-test-cluster

create: install
	kind create cluster --name nginx-test-cluster --config cluster.yaml

install:
	brew install kind
	brew install tilt

clean:
	tilt down
	kind delete cluster --name nginx-test-cluster