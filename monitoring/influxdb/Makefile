deploy:
	kubectl apply -f influxdb-deployment.yaml
	kubectl apply -f influxdb-service.yaml
	kubectl apply -f nodeexporter-deployment.yaml
	kubectl apply -f nodeexporter-service.yaml

clean:
	kubectl delete -f influxdb-deployment.yaml
	kubectl delete -f influxdb-service.yaml
	kubectl delete -f nodeexporter-deployment.yaml
	kubectl delete -f nodeexporter-service.yaml

# scraper nodeExporter http://node-exporter:9100/metrics