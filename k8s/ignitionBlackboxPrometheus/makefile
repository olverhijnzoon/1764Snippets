deploy: deploy_ignition deploy_blackbox deploy_prometheus

deploy_ignition:
	kubectl apply -f ignition-deployment.yaml
	kubectl apply -f ignition-service.yaml

deploy_blackbox:
	kubectl apply -f blackbox-config.yaml
	kubectl apply -f blackbox-deployment.yaml
	kubectl apply -f blackbox-service.yaml

deploy_prometheus:
	kubectl create configmap prometheus-config --from-file=prometheus-config.yaml
	kubectl apply -f prometheus-deployment.yaml
	kubectl apply -f prometheus-service.yaml

clean: clean_prometheus clean_blackbox clean_ignition

clean_ignition:
	kubectl delete -f ignition-deployment.yaml
	kubectl delete -f ignition-service.yaml

clean_blackbox:
	kubectl delete -f blackbox-config.yaml	
	kubectl delete -f blackbox-deployment.yaml
	kubectl delete -f blackbox-service.yaml

clean_prometheus:
	kubectl delete configmap prometheus-config
	kubectl delete -f prometheus-deployment.yaml
	kubectl delete -f prometheus-service.yaml

