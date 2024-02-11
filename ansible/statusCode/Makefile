check: prep
	ansible-playbook statusCode.yml --check --diff

run: prep
	ansible-playbook statusCode.yml

prep:
	@(ansible-galaxy collection list community.kubernetes > /dev/null 2>&1 && echo "community.kubernetes collection already installed") || ansible-galaxy collection install community.kubernetes

delete:
	ansible-playbook statusCodeAbsent.yml