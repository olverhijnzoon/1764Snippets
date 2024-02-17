install_argocd:
# prep
	kubectl create namespace argocd || true
	kubectl config set-context --current --namespace=argocd

# apply
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# retrieving the admin password and stuff: https://argo-cd.readthedocs.io/en/stable/getting_started/
	kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 --decode; echo

# localhost:30180
	kubectl apply -f argo-service.yaml

# lets setup a local Gogs with a Postgres
deploy_gogs:
# be aware of POSTGRES_USER and password from gogs-deployment.yaml then doing initial Gogs setup
# essentially use password "gogs" and change application url from 3000 -> 30000
# localhost:30000
	kubectl apply -f gogs-deployment.yaml

# create a user e.g. "test" and get a token to create a repo in local Gogs
create_repo:
	curl -H "Authorization: token 68223c4bcdd4fcbda556dcc12df8d54a1077a029" -H "Content-Type: application/json" -d '{"name": "1764snippets_inception"}' http://localhost:30000/api/v1/user/repos

# init 1764snippets_inception repo with test:test credentials
create_init_commit:
	rm -rf ./1764snippets_inception; \
	git clone http://test:test@localhost:30000/test/1764snippets_inception; \
	cd 1764snippets_inception && \
	touch README.md && \
	git add README.md && \
	git commit -m "first commit" && \
	git push

# lets do changes to the copy_1764nginx-deployment and push it to 1764snippets_inception repo
create_commit:
	cp copy_1764nginx-deployment.yaml 1764snippets_inception/1764nginx-deployment.yaml; \
	cd 1764snippets_inception && \
	git add 1764nginx-deployment.yaml && \
	git commit -m "Update index.html" && \
	git push

# ArgoCD app will be watching for changes in the local Gogs 1764snippets_inception repo
create_1764nginx:
	kubectl apply -f 1764nginx-argoapp.yaml
	#kubectl port-forward svc/nginx 8042:80

clean_argo_service:
	kubectl delete -f argo-service.yaml

clean_gogs:
	kubectl delete -f gogs-deployment.yaml