IMAGE=localhost:31005/lookupvault

run: push_image
	kubectl apply -f lookupAnsible-deployment.yaml
	sleep 8
	kubectl logs deployment/ansible-deployment
	
delete:
	kubectl delete -f lookupAnsible-deployment.yaml	

build_image:
	docker build -t ${IMAGE} .

push_image: build_image
	docker push ${IMAGE}

deploy_vault:
	kubectl apply -f vault-deployment.yaml
	kubectl apply -f vault-service.yaml
	sleep 8
	kubectl logs deployment/vault

clean_vault:
	kubectl delete -f vault-deployment.yaml
	kubectl delete -f vault-service.yaml