const express = require('express');
const umbral = require("@nucypher/umbral-pre");
const swaggerJsdoc = require('swagger-jsdoc');
const swaggerUi = require('swagger-ui-express');

const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());

// Swagger setup
const options = {
  definition: {
    openapi: '3.0.0',
    info: {
      title: 'Rubix SmartEncryptDecrypt API',
      version: '1.0.0',
      description: 'An API to demonstrate Rubix Smart Contract encryption and decryption',
    },
  },
  apis: ['./server.js'],
};
const specs = swaggerJsdoc(options);
app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(specs));

function bytesToHex(bytes) {
  return Array.from(bytes, (byte) => {
    return ('0' + (byte & 0xFF).toString(16)).slice(-2);
  }).join('');
}

function hexToBytes(hex) {
  const bytes = [];
  for (let i = 0; i < hex.length; i += 2) {
    bytes.push(parseInt(hex.substr(i, 2), 16));
  }
  return new Uint8Array(bytes);
}

/**
 * @swagger
 * /generate-secret-key:
 *   get:
 *     summary: Generate a secret key and its corresponding public key
 *     responses:
 *       200:
 *         description: Returns the secret key and public key in hexadecimal format
 */
app.get('/generate-secret-key', (req, res) => {
  let secretKey = umbral.SecretKey.random();
  let secretKeyBytes = secretKey.toBEBytes(); //Secret key to Big Endian bytes as noted from the rust-umbral-pre documentation
  console.log(secretKeyBytes)
  let secretKeyHex = bytesToHex(secretKeyBytes)
  console.log(secretKeyHex)
  let publicKey = secretKey.publicKey().toCompressedBytes();
  let publicKeyHex = bytesToHex(publicKey);
  
  res.json({ secretKey: secretKeyHex, publicKey: publicKeyHex });
});

/**
 * @swagger
 * /encrypt:
 *   post:
 *     summary: Encrypt a plaintext using a given public key
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               public_key:
 *                 type: string
 *               plaintext:
 *                 type: string
 *     responses:
 *       200:
 *         description: Returns the encrypted data and the capsule
 *       400:
 *         description: Missing required parameters
 */
app.post('/encrypt', (req, res) => {
    let { public_key, plaintext } = req.body;
    console.log(req.body)
    let plaintext_bytes = Buffer.from(plaintext, 'utf8');
    let publicKeyBytes = hexToBytes(public_key);
    console.log(`Public key is ${public_key}`);
    let publicKey = umbral.PublicKey.fromCompressedBytes(publicKeyBytes);

  
    if (!publicKey || !plaintext_bytes) {
      return res.status(400).json({ error: 'Missing required parameters' });
    }
  
    const [capsule, ciphertext] = umbral.encrypt(publicKey, plaintext_bytes);
    let capsuleBytes = capsule.toBytes();
    let capsuleHex = bytesToHex(capsuleBytes);
    let cipherTextHex = bytesToHex(ciphertext)
    return res.json({ capsule: capsuleHex, ciphertext: cipherTextHex });
});

/**
 * @swagger
 * /decrypt:
 *   post:
 *     summary: Decrypt a ciphertext using a given secret key and capsule
 *     requestBody:
 *       required: true
 *       content:
 *         application/json:
 *           schema:
 *             type: object
 *             properties:
 *               secretKey:
 *                 type: string
 *               capsule:
 *                 type: string
 *               ciphertext:
 *                 type: string
 *     responses:
 *       200:
 *         description: Returns the decrypted plaintext
 *       400:
 *         description: Missing required parameters
 */
app.post('/decrypt', (req, res) => {
    let dec = new TextDecoder("utf-8");
    const { secretKey, capsule, ciphertext } = req.body;
    let secretKeyBytes = hexToBytes(secretKey)
    let capsuleBytes = hexToBytes(capsule)
    let ciphertextBytes = hexToBytes(ciphertext)
    let capsuleObtained = umbral.Capsule.fromBytes(capsuleBytes)
    console.log("Secret Key recreated Bytes:", secretKeyBytes)
    let secretKeyObtained = umbral.SecretKey.fromBEBytes(secretKeyBytes);
    console.log(secretKeyObtained)
  
    if (!secretKeyObtained || !capsuleObtained || !ciphertext) {
      return res.status(400).json({ error: 'Missing required parameters' });
    }
  
    let plaintextBytes = umbral.decryptOriginal(secretKeyObtained, capsuleObtained, ciphertextBytes);
    let plaintext = dec.decode(plaintextBytes)

    return res.json({ plaintext });
});

app.post('/api/createdid', async (req, res) => {
  try {
    // Extract parameters from the request body, assuming it contains "port" and "didImagepath"
    const { port, didImagepath } = req.body;

    // Call the createDID function
    const response = await rubixUtil.createDID(port, didImagepath);

    // Respond with a success message
    res.json({ response });
  } catch (error) {
    // Handle errors and respond with an error message
    console.error('Error generating DID:', error.message);
    res.status(500).json({ error: 'An error occurred while generating DID' });
  }
});

app.post('/api/generate-smart-contract', async (req, res) => {
  try {
    const { did, wasmPath, schemaPath, rawCodePath, port } = req.body;

    // Call the generateSmartContract function
    const response = await rubixUtil.generateSmartContract(did, wasmPath, schemaPath, rawCodePath, port);

    // Respond with a success message
    res.json({ response });
  } catch (error) {
    console.error('Error:', error.message);
    res.status(500).json({ error: 'An error occurred while generating the smart contract' });
  }
});

app.post('/api/deploy-smart-contract', async (req, res) => {
  try {
    const {
      comment,
      deployerAddress,
      quorumType,
      rbtAmount,
      smartContractToken,
      port: deployPort, // Rename "port" to "deployPort" to avoid conflict
    } = req.body;

    // Call the deploySmartContract function directly
    const response = await rubixUtil.deploySmartContract(
      comment,
      deployerAddress,
      quorumType,
      rbtAmount,
      smartContractToken,
      deployPort
    );

    // Respond with the generated ID
    res.json({ response });
  } catch (error) {
    console.error('Error:', error.message);
    res.status(500).json({ error: 'An error occurred while deploying the smart contract' });
  }
});

app.post('/api/execute-smart-contract', async (req, res) => {
  try {
    const {
      comment,
      executorAddress,
      quorumType,
      smartContractData,
      smartContractToken,
      port: executionPort, // Rename "port" to "executionPort" to avoid conflict
    } = req.body;

    // Call the executeSmartContract function directly
    const response = await rubixUtil.executeSmartContract(
      comment,
      executorAddress,
      quorumType,
      smartContractData,
      smartContractToken,
      executionPort
    );

    // Respond with a success message
    res.json({ response });
  } catch (error) {
    console.error('Error:', error.message);
    res.status(500).json({ error: 'An error occurred while executing the smart contract' });
  }
});

app.post('/api/subscribe-contract', async (req, res) => {
  try {
    const { contractToken, port: subscribePort } = req.body;

    // Call the subscribeSmartContract function directly
    const response = await rubixUtil.subscribeSmartContract(contractToken, subscribePort);

    // Respond with a success message
    res.json({ response });

    
  } catch (error) {
    console.error('Error:', error.message);
    res.status(500).json({ error: 'An error occurred while subscribing to the smart contract' });
  }
});


app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
  console.log(`Swagger UI available at http://localhost:${port}/api-docs`);
});
