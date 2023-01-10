# E2E Encryption

Encrypts data between peers

Simple task descripition:

```
- server and 2 clients
- server maintains an address log
- clients lookup/update addresses
- clients intiate communication by:
    1. generating public/private key pairs
    2. exchanging public keys
    3. encrypting further data packets between peers
- clients m.o:
    1. publishing/subscribing to peer address logs
    2. opening bi-directional buffer pairs (acting as a node)
    3. publishing/fetching pre-encrypted data to/from the IO buffers
```

### Ideas

1. [Proper asymmetric file encryption in Go](https://stackoverflow.com/questions/40353070/proper-asymmetric-file-encryption-in-go)

```
If you need public key asymmetric encryption for data larger than the key size you need to use hybrid encryption. Essentially this is how HTTPS works.

Hybrid encryption is where the data is encrypted with symmetric key encryption such as AES and that key is encrypted with asymmetric key encryption such as RSA or EC (Elliptic Curve) Cryptography.
```

2. [Hybrid cryptosystem](https://en.wikipedia.org/wiki/Hybrid_cryptosystem)

```
In cryptography, a hybrid cryptosystem is one which combines the convenience of a public-key cryptosystem with the efficiency of a symmetric-key cryptosystem. Public-key cryptosystems are convenient in that they do not require the sender and receiver to share a common secret in order to communicate securely. However, they often rely on complicated mathematical computations and are thus generally much more inefficient than comparable symmetric-key cryptosystems. In many applications, the high cost of encrypting long messages in a public-key cryptosystem can be prohibitive. This is addressed by hybrid systems by using a combination of both.

A hybrid cryptosystem can be constructed using any two separate cryptosystems:

    a key encapsulation mechanism, which is a public-key cryptosystem, and
    a data encapsulation scheme, which is a symmetric-key cryptosystem.

The hybrid cryptosystem is itself a public-key system, whose public and private keys are the same as in the key encapsulation scheme.

Note that for very long messages the bulk of the work in encryption/decryption is done by the more efficient symmetric-key scheme, while the inefficient public-key scheme is used only to encrypt/decrypt a short key value.

All practical implementations of public key cryptography today employ the use of a hybrid system. Examples include the TLS protocol  and the SSH protocol, that use a public-key mechanism for key exchange (such as Diffie-Hellman) and a symmetric-key mechanism for data encapsulation (such as AES). The OpenPGP file format and the PKCS #7 file format are other examples.

Hybrid Public Key Encryption (HPKE, published as RFC 9180) is a modern standard for generic hybrid encryption. HPKE is used within multiple IETF protocols, including MLS and TLS Encrypted Hello.

Envelope encryption is an example of a usage of hybrid cryptosystems in cloud computing. In a cloud context, hybrid cryptosystems also enable centralized key management.
```
