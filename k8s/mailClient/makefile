BINARY=mailer
CC=g++
CFLAGS=-Wall -lcurl
SOURCE=main.cpp
IMAGE=localhost:31005/mailer

deploy: push_image
	kubectl apply -f mailer.yaml
	kubectl apply -f mailhog.yaml

push_image: build_image
	docker push ${IMAGE}

build_image: ${BINARY}X
	docker build -t ${IMAGE} .

run: ${BINARY}X
	./${BINARY}X

${BINARY}X: ${SOURCE}
	${CC} ${CFLAGS} -o $@ $<

clean:
	rm -f ./${BINARY}X
	kubectl delete -f mailer.yaml
	kubectl delete -f mailhog.yaml

# prep_registry: docker run -d -p 5001:5000 --restart=always --name registry registry:2