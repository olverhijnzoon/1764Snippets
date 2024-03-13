run: context
	tilt up

context: create
	kubectl cluster-info --context kind-jmx-test-cluster

create: install
	kind create cluster --name jmx-test-cluster --config cluster.yaml

install:
	brew install kind
	brew install tilt

clean:
	tilt down
	kind delete cluster --name jmx-test-cluster