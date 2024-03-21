#!/bin/bash

# 1. Setup and Configuration
# ---------------------------

# Get arguments (certs dir, and certificate type)
CERTS_DIR=${1:-client_certs}
CERT_TYPE=${2:-standard} # 'standard' or 'wildcard'

# Root CA Configuration 
CA_SUBJECT="/C=US/ST=California/O=My Company/CN=My Company Root CA"
CA_KEY="ca.key.pem"
CA_CERT="ca.crt.pem"

# Client Configuration (modify as needed)
CLIENT_SUBJECT="/C=US/ST=New York/O=Client Org/CN=client1"  
CLIENT_KEY="client1.key.pem"
CLIENT_CERT="client1.crt.pem"
CLIENT_CSR="client1.csr.pem"

# 2. Create Directories
# ---------------------
mkdir -p "$CERTS_DIR"

# 3. Generate the Root CA
# -----------------------
echo "Generating Root CA key..."
openssl genrsa -out "$CERTS_DIR/$CA_KEY" 2048  

echo "Generating Root CA certificate..."
openssl req -x509 -new -nodes -key "$CERTS_DIR/$CA_KEY" \
    -sha256 -days 1024 -out "$CERTS_DIR/$CA_CERT" \
    -subj "$CA_SUBJECT"

# 4. Generate Client Certificate
# ------------------------------
echo "Generating client key..."
openssl genrsa -out "$CERTS_DIR/$CLIENT_KEY" 2048

echo "Generating client certificate signing request (CSR)..."
if [ "$CERT_TYPE" == "wildcard" ]; then
    # Create a subjectAltName extension file
    echo "DNS.1 = *.example.com" > "$CERTS_DIR/v3.ext" 

    openssl req -new -key "$CERTS_DIR/$CLIENT_KEY" \
        -out "$CERTS_DIR/$CLIENT_CSR" \
        -subj "$CLIENT_SUBJECT" \
        -config <(cat "$CERTS_DIR/v3.ext")
else 
    openssl req -new -key "$CERTS_DIR/$CLIENT_KEY" \
        -out "$CERTS_DIR/$CLIENT_CSR" \
        -subj "$CLIENT_SUBJECT"
fi

echo "Signing client certificate with the CA..."
openssl x509 -req -in "$CERTS_DIR/$CLIENT_CSR" \
    -CA "$CERTS_DIR/$CA_CERT" -CAkey "$CERTS_DIR/$CA_KEY" \
    -CAcreateserial -out "$CERTS_DIR/$CLIENT_CERT" \
    -days 500 -sha256 -extfile "$CERTS_DIR/v3.ext"

# 5. Verification
# ---------------
echo "Verifying client certificate..."
openssl verify -CAfile "$CERTS_DIR/$CA_CERT" "$CERTS_DIR/$CLIENT_CERT"

echo "Detailed client certificate info:"
openssl x509 -in "$CERTS_DIR/$CLIENT_CERT" -text -noout 
