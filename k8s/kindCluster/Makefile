context: create
	kubectl cluster-info --context kind-cluster1764

create: install
	kind create cluster --name cluster1764 --config cluster1764.yaml

install:
	brew install kind
	open -a Docker

clean:
	kind delete cluster --name kind-cluster1764