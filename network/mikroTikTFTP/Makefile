# update inventory with ip of router
run:
	ansible-playbook -i ansible/inventory.ini ansible/TFTPMikroTik.yml

#prep:
#	ssh admin@172.16.176.134
#	pip3 install librouteros --break-system-packages
# put ignite.sh into "Files" in the WebFig GUI 

test:
	@echo "get startup\nquit" | tftp 172.16.176.134

clean:
	rm startup