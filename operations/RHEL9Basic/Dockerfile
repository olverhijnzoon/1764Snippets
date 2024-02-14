# Red Hat Universal Base Images
# FROM registry.access.redhat.com/ubi9/ubi-micro:latest
FROM registry.access.redhat.com/ubi9/ubi-minimal:latest
# FROM registry.access.redhat.com/ubi9/ubi:latest
# FROM registry.access.redhat.com/ubi9/ubi-init:latest
# RUN dnf update -y && dnf install -y vim
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

