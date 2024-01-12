FROM python:3.9
RUN pip install ansible
RUN pip install hvac
RUN ansible-galaxy collection install community.hashi_vault
COPY lookupAnsible.yml /lookupAnsible.yml
CMD ["ansible-playbook", "/lookupAnsible.yml"]
