FROM inductiveautomation/ignition:8.1.30
USER root
RUN mkdir /tmp/backup
RUN chown ignition:ignition /tmp/backup
RUN chmod 755 /tmp/backup
USER ignition
COPY ignition-edge-test.gwbk /tmp/backup