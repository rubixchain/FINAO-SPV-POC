# Re-encryption server

To run the server : ``` node server.js ```

There are 3 api endpoints defined in this server:

1. /generate-secret-key 
```
curl http://localhost:3000/generate-secret-key

```
The response here will be the secretKey and publicKey which needs to be stored. The response will be in json format.

2. /encrypt
```
curl -X POST -H "Content-Type: application/json" -d '{
  "public_key": "your_stored_public_key_response_from_previous_api_call",
  "plaintext": "TexT_To_be_encrypted"
}' http://localhost:3000/encrypt

```
The response will be a capsule and ciphertext in json format.

3. /decrypt
```
curl -X POST -H "Content-Type: application/json" -d '{
  "secretKey": "your_stored_secret_key_generated_in_the_first_api",
  "capsule": "Encapsulated symmetric key used to encrypt the plaintext. Generated while encrypting",
  "ciphertext": "The encrypted text"
}' http://localhost:3000/decrypt

```

The response will be the decrypted text.