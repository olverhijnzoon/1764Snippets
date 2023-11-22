#!/bin/bash

# Define the folder name variable
snippetname="anotherSimpleAnsible"

# Create the base folder, roles, and defaults directory
mkdir -p "$snippetname/roles/create_namespace/tasks"
mkdir -p "$snippetname/defaults"

# Create the defaults file with kubeconfig path
cat <<EOF > "$snippetname/defaults/main.yml"
---
kubeconfig: "~/.kube/config"
EOF

# Create the role task
cat <<EOF > "$snippetname/roles/create_namespace/tasks/main.yml"
---
- name: Create a namespace
  community.kubernetes.k8s:
    kubeconfig: "{{ kubeconfig }}"
    name: ansible
    api_version: v1
    kind: Namespace
    state: present

- name: Get information about the namespace
  community.kubernetes.k8s_info:
    kubeconfig: "{{ kubeconfig }}"
    kind: Namespace
    name: ansible
  register: ns_info

- name: Print the namespace information
  debug:
    msg: "{{ ns_info.resources }}"
EOF

# Create the Ansible playbook in the base folder
cat <<EOF > "$snippetname/$snippetname.yml"
---
- name: Kubernetes Playbook
  hosts: localhost
  gather_facts: no

  vars_files:
    - defaults/main.yml

  roles:
    - create_namespace
EOF

# Create the Makefile in the base folder
cat <<EOF > "$snippetname/Makefile"
run: prep
	ansible-playbook $snippetname.yml --check --diff

prep:
	@(ansible-galaxy collection list community.kubernetes > /dev/null 2>&1 && echo "community.kubernetes collection already installed") || ansible-galaxy collection install community.kubernetes
EOF

echo "Folder structure, playbook, roles, and Makefile created successfully."

# Navigate to the folder and execute the playbook
cd "$snippetname"
make run
