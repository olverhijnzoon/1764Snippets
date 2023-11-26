deploy: deploy_ignition

deploy_ignition:
	docker build -t ignition-edge-restore:20231126 .
	docker tag ignition-edge-restore:20231126 localhost/ignition-edge-restore:20231126
	kubectl apply -f ignition-deployment.yaml
	kubectl apply -f ignition-service.yaml

clean:
	kubectl delete -f ignition-deployment.yaml
	kubectl delete -f ignition-service.yaml
