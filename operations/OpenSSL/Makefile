SCRIPT = generate_client_cert.sh
CERTS_DIR=certs

run:
	chmod +x $(SCRIPT)
	./$(SCRIPT) $(CERTS_DIR)

wildrun:
	chmod +x $(SCRIPT)
	./$(SCRIPT) $(CERTS_DIR)

clean:
	rm -rf ./$(CERTS_DIR)

speed:
	openssl speed