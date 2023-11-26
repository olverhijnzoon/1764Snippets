#!/bin/bash

# Define the snippet name
snippetname="justAnotherIgnitionSnippet"

# Create the directory
mkdir -p "$snippetname"

# Create the ignition-deployment.yaml file
cat <<EOF >"$snippetname/ignition-deployment.yaml"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ignition-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: ignition
  template:
    metadata:
      labels:
        app: ignition
    spec:
      containers:
      - name: ignition-test
        image: inductiveautomation/ignition:8.1.30
        args:
          - "-n"
          - "docker-test"
          - "-a"
          - "localhost"
          - "-h"
          - "9088"
          - "-s"
          - "9043"
        ports:
        - containerPort: 9088
EOF

# Create the ignition-service.yaml file
cat <<EOF >"$snippetname/ignition-service.yaml"
apiVersion: v1
kind: Service
metadata:
  name: ignition-service
  annotations:
    prometheus.io/scrape: "true"
    prometheus.io/path: "/"
spec:
  type: NodePort
  selector:
    app: ignition
  ports:
  - protocol: TCP
    port: 8088
    targetPort: 8088
    nodePort: 30088
EOF

# Create the Makefile
cat <<EOF >"$snippetname/Makefile"
deploy: deploy_ignition

deploy_ignition:
	kubectl apply -f ignition-deployment.yaml
	kubectl apply -f ignition-service.yaml

clean_ignition:
	kubectl delete -f ignition-deployment.yaml
	kubectl delete -f ignition-service.yaml
EOF

echo "Files created in directory: $snippetname"
