# Transform Encryption

This Rust program demonstrates how to use the Recrypt library to encrypt, transform, and decrypt data using public and private keys.

## Overview
The program performs the following steps:

create_key_pair() - create new Encryption KeyPairs and Signing Keypairs using rng returns
    pub_key: [u8; 33],
    pvt_key: [u8; 32],
    sign_key: [u8; 32],
    ver_key: [u8; 33],

encrypt_data() - encrypt data to delegators pub key and splits into fragments. Retruns ciphertext and fragments

decrypt_data() - decrypt data to be called by delegator by passing cipher text and relevant private key

## Usage
Clone the repository or download the source code.

Compile and run the program:
`cargo run`

Review the output and any assertion errors to understand the behavior of the umbral-rs library.

## Dependencies
This program depends on the following crates:

[umbral-pre](https://crates.io/crates/umbral-pre/0.2.0): A crate that provides cryptographic primitives for data encryption, transformation, and decryption. We are using v0.2.0 for compatitbiltiy reasons.