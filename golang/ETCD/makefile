BINARY=ETCD

apply:
	kubectl apply -f etcd.yaml

init:
	kubectl exec -it etcd-0 -- etcdctl put my-key "1764Snippets ETCD"
	kubectl exec -it etcd-0 -- etcdctl get my-key
	kubectl exec -it etcd-1 -- etcdctl get my-key
	kubectl exec -it etcd-2 -- etcdctl get my-key

delete:
	kubectl delete -f etcd.yaml

describe:
	kubectl describe statefulset etcd

status:
	kubectl get pods -l app=etcd

logs:
	kubectl logs -l app=etcd --tail=50

shell:
	kubectl exec -it etcd-0 -- /bin/sh

wait:
	kubectl wait --for=condition=Ready pod -l app=etcd --timeout=300s

build-go:
	go build -o $(BINARY)X $(BINARY).go

run-go:
	sleep 4
	kubectl port-forward svc/etcd 2379:2379 & 
	./$(BINARY)X &

change-config:
	sleep 2
	kubectl exec -it etcd-0 -- etcdctl put config '{"setting1": "Changed 1764Snippets config!", "setting2": 1764}'
	kubectl exec -it etcd-0 -- etcdctl get config
	kubectl exec -it etcd-1 -- etcdctl get config
	kubectl exec -it etcd-2 -- etcdctl get config

etcd-demo: build-go apply run-go change-config

clean-go:
	rm -f $(BINARY).go

clean-binaries:
	rm -f ./*X