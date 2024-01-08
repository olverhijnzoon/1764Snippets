deploy_artifactory:
	kubectl apply -f registry-deployment.yaml
	kubectl apply -f registry-service.yaml

clean_artifactory:
	kubectl delete -f registry-deployment.yaml
	kubectl delete -f registry-service.yaml