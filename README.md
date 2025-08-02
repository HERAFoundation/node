# Sporechain Node

## Running the ATProtocol Identity Provider (AIP)

```
docker build -t aip ./aip
```

```
docker run -p 8080:8080 \
  -e EXTERNAL_BASE=http://localhost:8080 \
  -e DPOP_NONCE_SEED=$(openssl rand -hex 32) \
  aip
```
