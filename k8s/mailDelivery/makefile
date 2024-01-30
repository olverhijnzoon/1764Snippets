IMAGE=localhost:31005/postfix

deploy: push_image
	kubectl apply -f postfix-deployment.yaml

push_image: build_image
	docker push ${IMAGE}

build_image:
	docker build -t ${IMAGE} .

logs:
	kubectl logs deployments/postfix-deployment

clean:
	kubectl delete -f postfix-deployment.yaml

### test within postfix container
# root@postfix-deployment-f6c4c6f94-g7289:/# echo "Subject: Test Mail" | sendmail -v root@localhost
# root@postfix-deployment-f6c4c6f94-g7289:/# cat /var/mail/root
# root@postfix-deployment-f6c4c6f94-g7289:/# mutt
