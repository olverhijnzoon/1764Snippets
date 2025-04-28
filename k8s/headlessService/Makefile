IMAGE_NAME=headlessservice
IMAGE_TAG=latest
IMAGE=localhost:31005/headlessservice

push_image: build_image
	docker push ${IMAGE}

build_image:
	docker build -t ${IMAGE} .

apply:
	@echo "Applying Kubernetes resources..."
	kubectl apply -f ./deployment.yaml
	kubectl apply -f ./service.yaml
	@echo "Deployment and Service applied."

delete:
	@echo "Deleting Kubernetes resources..."
	kubectl delete -f ./service.yaml
	kubectl delete -f ./deployment.yaml
	@echo "Deployment and Service deleted."

status:
	@echo "Checking status..."
	kubectl get deployment headless-service-deployment
	kubectl get service headless-service-service
	kubectl get pods -l component=headless-service-1764
.PHONY: apply delete status

test:
	kubectl run tmp-shell --rm -it --image=busybox -- sh
	# nslookup headless-service-service
	# wget -O - 10.x.x.x:14269