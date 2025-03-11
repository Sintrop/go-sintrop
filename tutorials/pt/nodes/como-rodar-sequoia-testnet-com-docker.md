## Como rodar Sequoia-testnet com Docker

## Pre requisitos
- Docker instalado


## Rodar o projeto

### Buildar go-sintrop
```
docker build -t go-sintrop .
# Build without cache
docker build --no-cache -t go-sintrop .
```

### Executar go-sintrop
```
docker run -p=30303:30303 -p=8545:8545 -it -v /home/user/sequoia_volume:/go-sintrop/sequoia_node  go-sintrop
## Change /home/user/sequoia_volume to your dir
```
## GETH

### Iniciar um node

Change miner.etherbase to your wallet address and run the following command:

```
geth --identity Sequoia --datadir ./sequoia_node \
  --sequoia \
  --syncmode "full" \
  --networkid 1600 \
  --cache=1024 \
  --port 30303 \
  -authrpc.addr localhost --authrpc.port 8551 \
  --http.vhosts=* --http.addr "0.0.0.0" --http.port 8545 --http=true \
  --miner.threads=1 \
  --miner.etherbase=0x0000000000000000000000000000000000000000 \
  console
```

### Start a archive node
```
geth --identity Sequoia --datadir ./sequoia_node \
  --sequoia \
  --syncmode "full" \
  --networkid 1600 \
  --cache=1024 \
  --port 30303 \
  --authrpc.addr localhost --authrpc.port 8551 \
  --http.vhosts=* --http.addr "0.0.0.0" --http.port 8545 --http=true --http.api debug,net,eth,web3,txpool --ws=true --ws.addr 0.0.0.0 --ws.port 8546 --ws.origins "*" \
  --miner.threads=1 \
  --miner.etherbase=0x0000000000000000000000000000000000000000 \
  --gcmode=archive \
  console
```

### Operate network

```
balance = web3.fromWei(eth.getBalance("0x0000000000000000000000000000000000000000), "ether");
eth.blockNumber
web3.eth.getBlock(eth.blockNumber)
```
## MINER

### Start and stop mining
```
  miner.start()
  miner.stop()
```
