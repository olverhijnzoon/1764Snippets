PROJECT_ID=25839cde-c992-4764-bd5e-6650b5d27246
PROJECT_NAME=basic

start: run
	sleep 5
	curl -X POST http://localhost:3080/v2/projects/$(PROJECT_ID)/nodes/start

run:
	open -a GNS3 $(PROJECT_NAME).gns3

connect:
	telnet localhost 5001

connect-and-trace:
	curl -X POST http://localhost:3080/v2/projects/$(PROJECT_ID)/nodes/start
	sleep 5
	chmod +x trace_command.expect
	./trace_command.expect