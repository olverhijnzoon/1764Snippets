run:
	ansible-playbook -i ansible/inventory.ini ansible/vyosConfig.yml

# init vyos manually for this basic vyos snippet
# prep_vyos by entering device with vyos:vyos then
# 	configure
# 	set interfaces ethernet eth0 address 172.16.176.2/24
#	set service ssh port '22'
#	commit
#	save
#	exit
# ssh vyos@172.16.176.2 -p 22