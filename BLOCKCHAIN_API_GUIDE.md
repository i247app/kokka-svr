# Blockchain JSON-RPC Proxy API Guide

This document describes the blockchain JSON-RPC proxy API that has been implemented to allow clients to interact with the blockchain through the Go server instead of calling the JSON-RPC server directly.

## Architecture Flow

```
Client (Browser/App)
    ↓ HTTP Request
Go Server (localhost:8080)
    ↓ JSON-RPC Call
Blockchain RPC (https://x24.i247.com)
    ↓ Response
Go Server
    ↓ HTTP Response
Client
```

## Configuration

The blockchain RPC URL is configured in `.env`:

```env
BLOCKCHAIN_RPC_URL=https://x24.i247.com
```

## Available Endpoints

### 1. Get Block Number
**GET** `/api/blockchain/block-number`

Returns the current block number.

**Response:**
```json
{
  "block_number": "0x1234567"
}
```

**Example:**
```bash
curl http://localhost:8080/api/blockchain/block-number
```

---

### 2. Get Balance
**POST** `/api/blockchain/balance`

Get the balance of an Ethereum address.

**Request:**
```json
{
  "address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "block": "latest"
}
```

**Response:**
```json
{
  "address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "balance": "0x1234567890abcdef",
  "block": "latest"
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/blockchain/balance \
  -H "Content-Type: application/json" \
  -d '{
    "address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "block": "latest"
  }'
```

---

### 3. Get Block
**POST** `/api/blockchain/block`

Get block details by block number.

**Request:**
```json
{
  "block_number": "latest",
  "full_tx": false
}
```

**Response:**
```json
{
  "block": {
    "number": "0x1234567",
    "hash": "0xabc...",
    "transactions": [...],
    ...
  }
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/blockchain/block \
  -H "Content-Type: application/json" \
  -d '{
    "block_number": "latest",
    "full_tx": false
  }'
```

---

### 4. Get Transaction
**POST** `/api/blockchain/transaction`

Get transaction details by transaction hash.

**Request:**
```json
{
  "tx_hash": "0x1234567890abcdef..."
}
```

**Response:**
```json
{
  "transaction": {
    "hash": "0x1234567890abcdef...",
    "from": "0xabc...",
    "to": "0xdef...",
    ...
  }
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/blockchain/transaction \
  -H "Content-Type: application/json" \
  -d '{
    "tx_hash": "0x1234567890abcdef..."
  }'
```

---

### 5. Call Contract
**POST** `/api/blockchain/call`

Execute a read-only smart contract call.

**Request:**
```json
{
  "to": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
  "data": "0x70a08231000000000000000000000000...",
  "block": "latest"
}
```

**Response:**
```json
{
  "result": "0x0000000000000000000000000000000000000000000000000de0b6b3a7640000"
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/blockchain/call \
  -H "Content-Type: application/json" \
  -d '{
    "to": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
    "data": "0x70a08231000000000000000000000000...",
    "block": "latest"
  }'
```

---

### 6. Estimate Gas
**POST** `/api/blockchain/estimate-gas`

Estimate gas required for a transaction.

**Request:**
```json
{
  "from": "0xabc...",
  "to": "0xdef...",
  "data": "0x..."
}
```

**Response:**
```json
{
  "gas_estimate": "0x5208"
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/blockchain/estimate-gas \
  -H "Content-Type: application/json" \
  -d '{
    "from": "0xabc...",
    "to": "0xdef...",
    "data": "0x"
  }'
```

---

### 7. Send Raw Transaction
**POST** `/api/blockchain/send-transaction`

Broadcast a signed transaction to the network.

**Request:**
```json
{
  "signed_tx": "0xf86c808504a817c800825208..."
}
```

**Response:**
```json
{
  "tx_hash": "0x1234567890abcdef..."
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/blockchain/send-transaction \
  -H "Content-Type: application/json" \
  -d '{
    "signed_tx": "0xf86c808504a817c800825208..."
  }'
```

---

### 8. Get Gas Price
**GET** `/api/blockchain/gas-price`

Get the current gas price.

**Response:**
```json
{
  "gas_price": "0x3b9aca00"
}
```

**Example:**
```bash
curl http://localhost:8080/api/blockchain/gas-price
```

---

### 9. Get Chain ID
**GET** `/api/blockchain/chain-id`

Get the chain ID.

**Response:**
```json
{
  "chain_id": "0x1"
}
```

**Example:**
```bash
curl http://localhost:8080/api/blockchain/chain-id
```

---

### 10. Generic RPC Call
**POST** `/api/blockchain/rpc`

Call any JSON-RPC method directly. This is useful for methods not explicitly supported by the other endpoints.

**Request:**
```json
{
  "method": "eth_getBlockByNumber",
  "params": ["latest", false]
}
```

**Response:**
```json
{
  "result": {
    "number": "0x1234567",
    ...
  }
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/blockchain/rpc \
  -H "Content-Type: application/json" \
  -d '{
    "method": "eth_blockNumber",
    "params": []
  }'
```

---

## Running the Server

```bash
# Build the server
make build

# Run the server
make run

# Or run directly
go run ./cmd/server
```

The server will start on `http://localhost:8080` by default.

---

## Error Handling

All endpoints return errors in the following format:

**Error Response (Status 400):**
```json
{
  "error": "invalid ethereum address format"
}
```

**Error Response (Status 500):**
```json
{
  "error": "failed to get balance: ..."
}
```

---

## Implementation Details

### Architecture Components

1. **Blockchain Client** (`internal/driven-adapter/external/blockchain/`)
   - JSON-RPC client that communicates with https://x24.i247.com
   - Uses the shared `http_client` package
   - Implements JSON-RPC 2.0 protocol

2. **Blockchain Service** (`internal/applications/services/blockchain_service.go`)
   - Business logic layer
   - Transforms between DTOs and blockchain client calls

3. **DTOs** (`internal/applications/dtos/blockchain_dto.go`)
   - Request/Response structures for the API

4. **Validators** (`internal/applications/validators/blockchain_validator.go`)
   - Input validation for all requests
   - Validates Ethereum addresses, hashes, hex data, etc.

5. **Controller** (`internal/handlers/http/controller/blockchain_controller.go`)
   - HTTP request handlers
   - Parses requests, validates, calls service, returns responses

### Features

- **Retry Logic**: Automatic retry for failed requests (configured for 500-level errors)
- **Request Logging**: All HTTP requests are logged via interceptors
- **Timing Metrics**: Request timing is automatically tracked
- **Input Validation**: All inputs are validated before being sent to the blockchain
- **Generic RPC Support**: Fallback endpoint for any JSON-RPC method

---

## Testing

You can test the implementation by starting the server and making requests to the endpoints above. Make sure the blockchain RPC URL (https://x24.i247.com) is accessible from your server.

### Quick Test

```bash
# Start the server
make run

# In another terminal, test the block number endpoint
curl http://localhost:8080/api/blockchain/block-number

# Test with a balance request (replace with a valid address)
curl -X POST http://localhost:8080/api/blockchain/balance \
  -H "Content-Type: application/json" \
  -d '{"address": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", "block": "latest"}'
```
