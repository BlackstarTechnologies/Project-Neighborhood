package types

/*
TODO:
	1. Define types:
		- Padlock
		- Address Log
		- Connection Pipe
		- Server
	a. Padlock holds:
		- seed
		- asymetric key:
			- symmetric key
			- public key
			- private key
		- queue:
			- in channel
			- deccrypt() #in
			- out channel
			- encrypt(file) #out
	b. Address Log:
		Address type {domain name, ip, port, public key}
		- Address map[uuid]:Address
	c. Connection pipe:
		- domain name, ip, port, public key
		- Padlock
		-	Server handler
			- request handler
			- response handler
			# a. handle one request at a time
			# or b. send/recieve without closing object

	2. Server structure
	- each peer is a server-client pair
		- 1 server
			- handles incoming (listen and respond)
		- N clients
			- handles outgoing (wait and request)
		- encryption engine

*/
