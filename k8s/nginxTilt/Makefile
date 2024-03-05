run:
	tilt up

context:
	kubectl cluster-info --context kind-nginx-test-cluster

create-kind-cluster: install
	kind create cluster --name nginx-test-cluster

install:
	brew install kind
	brew install tilt

clean:
	tilt down
	kind delete cluster --name nginx-test-cluster