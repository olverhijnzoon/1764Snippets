deploy: deploy_ignition

deploy_ignition:
	docker build -t ignition-tags:20240115 .
	docker tag ignition-tags:20240115 localhost/ignition-tags:20240115
	kubectl apply -f ignition-deployment.yaml
	kubectl apply -f ignition-service.yaml

clean_ignition:
	kubectl delete -f ignition-deployment.yaml
	kubectl delete -f ignition-service.yaml

logs:
	kubectl logs deployment/ignition-deployment

dump:
	kubectl exec -it $$(kubectl get pods -l app=ignition -o jsonpath="{.items[0].metadata.name}") -- /bin/sh -c "echo \".dump\" | sqlite3 /usr/local/bin/ignition/data/db/config.idb" > local_dump.sql

grep:
	sed '/INSERT INTO IMAGES VALUES/d' local_dump.sql > local_dump_clean.sql