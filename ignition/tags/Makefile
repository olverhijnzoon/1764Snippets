deploy: deploy_ignition

deploy_ignition:
	docker build -t ignition-tags:20240115 .
	docker tag ignition-tags:20240115 localhost/ignition-tags:20240115
	kubectl apply -f ignition-deployment.yaml
	kubectl apply -f ignition-service.yaml

clean_ignition:
	kubectl delete -f ignition-deployment.yaml
	kubectl delete -f ignition-service.yaml

TABLE=TAGCONFIG                                              
#TABLE=TAGGROUPS                                              
#TABLE=TAGHISTORIANPROVIDERSETTINGS                           
#TABLE=TAGHISTORYPROVIDEREP                                   
#TABLE=TAGHISTORYSPLITTERSETTINGS                             
#TABLE=TAGPROVIDERSETTINGS                                    
#TABLE=TAGREPORTSEARCH                                        
#TABLE=TAGSTORAGEPROPS   
gather_tables:
	kubectl exec -it $$(kubectl get pods -l app=ignition -o jsonpath="{.items[0].metadata.name}") -- /bin/sh -c "echo \".tables\" | sqlite3 /usr/local/bin/ignition/data/db/config.idb"

run_sql:
	kubectl exec -it $$(kubectl get pods -l app=ignition -o jsonpath="{.items[0].metadata.name}") -- /bin/sh -c "echo \"SELECT * FROM $(TABLE);\" | sqlite3 /usr/local/bin/ignition/data/db/config.idb"

FILE=$$(cat ./addTags.sql)

run_sql_script:
	kubectl exec -it $$(kubectl get pods -l app=ignition -o jsonpath="{.items[0].metadata.name}") -- /bin/sh -c "echo \"$(FILE)\" | sqlite3 /usr/local/bin/ignition/data/db/config.idb"

