run: prep
	ansible-playbook anotherSimpleAnsible.yml --check --diff

prep:
	@(ansible-galaxy collection list community.kubernetes > /dev/null 2>&1 && echo "community.kubernetes collection already installed") || ansible-galaxy collection install community.kubernetes
