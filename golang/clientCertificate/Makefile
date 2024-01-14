BINARY=clientCertificate
SERVER_CERT=server.crt
SERVER_KEY=server.key
CLIENT_CERT=client.crt
CLIENT_KEY=client.key

create_certs:
	openssl req -x509 -newkey rsa:4096 -keyout $(SERVER_KEY) -out $(SERVER_CERT) -days 365 -nodes -subj "/CN=localhost"
	openssl req -newkey rsa:4096 -keyout $(CLIENT_KEY) -out client.csr -nodes -subj "/CN=client"
	openssl x509 -req -in client.csr -CA $(SERVER_CERT) -CAkey $(SERVER_KEY) -CAcreateserial -out $(CLIENT_CERT) -days 365
	rm client.csr

prep:
	@if [ ! -f $(SERVER_CERT) ] || [ ! -f $(SERVER_KEY) ]; then \
		echo "Error: $(SERVER_CERT) and $(SERVER_KEY) must be present to run the server"; \
		exit 1; \
	fi

run_go: build_go prep
	./$(BINARY)X

build_go:
	go build -o $(BINARY)X $(BINARY).go

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X

test: prep
	curl --cert $(CLIENT_CERT) --key $(CLIENT_KEY) --cacert $(SERVER_CERT) https://localhost:8071

test_fail:
	curl --cacert $(SERVER_CERT) https://localhost:8071
