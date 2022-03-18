#!/bin/sh
mkdir -p .ssh

echo "Generating private key ... "
openssl genpkey -algorithm RSA -out .ssh/private.pem -pkeyopt rsa_keygen_bits:2048
echo "Private key generated\n"
chmod +r .ssh/private.pem
echo "Deriving public key"
openssl rsa -in .ssh/private.pem -pubout > .ssh/public.pem
echo "Public key derived\n"