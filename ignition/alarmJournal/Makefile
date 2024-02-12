deploy: deploy_ignition

deploy_ignition:
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

SQL=$$(cat ./alarmJournal.sql)

run_sql:
	echo $(SQL)
	kubectl exec -it $$(kubectl get pods -l app=ignition -o jsonpath="{.items[0].metadata.name}") -- /bin/sh -c "echo \"$(SQL)\" | sqlite3 /usr/local/bin/ignition/data/db/config.idb"

dump_local_journal:
	kubectl exec -it $$(kubectl get pods -l app=ignition -o jsonpath="{.items[0].metadata.name}") -- /bin/sh -c "echo \".dump\" | sqlite3 /usr/local/bin/ignition/data/local/alarm-journal/alarmjournal_AlarmJournalInternal1764.idb" > local_journal_dump.sql