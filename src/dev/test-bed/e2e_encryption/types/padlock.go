package types

/*
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
*/

type AsymetricKey struct {
	symetrickey string
	publickey   string
	privatekey  string
}

type Queue struct {
	in_channel  chan string
	out_channel chan string
}

type Padlock struct {
	seed         string
	asymetrickey AsymetricKey
	queue        Queue
}
