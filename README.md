# Digital Signature Illustration

## What
This is an illustration of digital signature. This experiment includes classic alice and bob.
They want to communicate over a network, but we know that it is always possible that someone eavesdroping and change the message. To solve this problem we've used asymmetric keys. So bob would sign the message using his private key and alice (or anyone) could use public key to verify the signature.
Cryptographically we decrypt the message using private key to sign a message and encrypt the message using public key for verication.

![Digital Signature Illustration](assets/digital_signature_illustration.png)

## Why
Broadly ditial signature have 3 purposes
1. Authentication: It ensures the authenticity of the sender.
2. Non-repudiation: It ensures that sender have sent the message and they can not deny later on
3. Integrity: It ensures that the message have not been altered.

## References:
* https://www.youtube.com/watch?v=TmA2QWSLSPg&ab_channel=SunnyClassroom
* https://www.thedigitalcatonline.com/blog/2018/04/25/rsa-keys/
