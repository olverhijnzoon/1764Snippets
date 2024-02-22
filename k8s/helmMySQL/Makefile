install:
	helm install helm-mysql ./

logs:
	kubectl logs deployment/helm-mysql

prep:
	brew install mysql-client
	brew --prefix mysql-client

# as MySQL root user
test:
	@echo "Running MySQL test..."
	@cd /opt/homebrew/opt/mysql-client/bin && ./mysql -h 127.0.0.1 -P 30010 -p

uninstall:
	helm uninstall helm-mysql