run:
	tilt up

prep:
	kind create cluster --name test-cluster --config cluster.yaml
	kubectl cluster-info --context kind-test-cluster

clean:
	tilt down
	kind delete cluster --name test-cluster 