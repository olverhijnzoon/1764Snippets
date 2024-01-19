deploy_artifactory:
	kubectl apply -f artifactory-deployment.yaml
	kubectl apply -f artifactory-service.yaml

clean_artifactory:
	kubectl delete -f artifactory-deployment.yaml
	kubectl delete -f artifactory-service.yaml

logs:
	kubectl logs deployment/artifactory

# admin : password