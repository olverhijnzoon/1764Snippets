exec: deploy_rhel9
	sleep 1
	kubectl exec -it deployment/rhel9-deployment -c rhel9 -- /bin/bash

deploy_rhel9:
	docker build -t rhel9-basic:20240214 .
	docker tag rhel9-basic:20240214 localhost/rhel9-basic:20240214
	kubectl apply -f rhel9-deployment.yaml

clean:
	kubectl delete -f rhel9-deployment.yaml

logs:
	kubectl get pods
	kubectl logs deployment/rhel9-deployment