BINARY=golanghashicorpvault
IMAGE=localhost:31005/avault

build_go:
	go build -o $(BINARY)X $(BINARY).go

run_go:
	./$(BINARY)X

god: build_go run_go

build_image:
	docker build -t ${IMAGE} .

push_image: build_image
	docker push ${IMAGE}

deploy_vault:
	kubectl apply -f vault-deployment.yaml
	kubectl apply -f vault-service.yaml
#kubectl port-forward svc/vault 8200:8200
clean_go:
	rm -f $(BINARY)X

clean_binaries:
	rm -f ./*X

clean:
	kubectl delete -f vault-deployment.yaml
	kubectl delete -f vault-service.yaml
