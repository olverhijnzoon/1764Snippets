run:
	ansible-playbook -i ansible/inventory.ini ansible/vyos_command.yml -vvv

# init vyos manually for this basic vyos snippet
# prep_vyos by entering device with vyos:vyos then
#	show interfaces
#	show interfaces ethernet eth0
# 	configure
# 	set interfaces ethernet eth0 address 172.16.176.13/24
#	set service ssh port '22'
#	commit
#	save
#	exit
# ssh vyos@172.16.176.13 -p 22