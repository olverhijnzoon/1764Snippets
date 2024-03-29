# Modbus
- connection of devices with Programmable Logic Controller (PLC)
- 1 central monitor/controller "master" and up to 247 field devices ("slaves")
- open protocol created mid-1970s by Modicon (now part of Schneider) for their controllers
- supported by other vendors including Siemens, Allen Bradley, Mitsubishi, and Omron, facilitating its use across a wide range of industrial equipment and PLCs from different manufacturers
- enables communication between industrial equipment from different manufacturers
- Cyclic Redundancy Check (CRC) to ensure the integrity of the data being transmitted between devices
- slave will only transmit information when asked from master to transmit
- slaves have unique address called unit ID
- a SCADA/HMI can also be master with PLCs connected to it being the slaves
- Modbus Remote Terminal Unit (RTU) utilizes RS232 for point-to-point communication under 15 meters
- RS422 and RS485 allow multi-device connection on the same line beyond 15 meters
- Modbus Transmission Control Protocol/Internet Protocol (TCP/IP) uses the server-client architecture with more than one device acting as server and server/clients are connected via regular switch
- regular Ethernet cable can be used for it
- each device in Modbus TCP/IP network will have an IP address (no unit ID then)