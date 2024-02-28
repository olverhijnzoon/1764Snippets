run:
	ansible-playbook -i ansible/inventory.ini ansible/vyosFirewall.yml

# init vyos manually for this basic vyos snippet
# prep_vyos by entering device with vyos:vyos then
# 	configure
# 	set interfaces ethernet eth0 address 172.16.176.3/24
#	set service ssh port '22'
#	commit
#	save
#	exit
# ssh vyos@172.16.176.3 -p 22
# run show configuration commands | grep firewall
# set firewall name 'INBOUND'
# set firewall name 'LOCAL_NET_FILTER'
# set firewall name 'OUTBOUND'
# trouble
# delete interfaces ethernet eth0 firewall