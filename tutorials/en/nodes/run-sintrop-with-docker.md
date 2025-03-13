## Run Sintrop with docker

## Run Sequoia with docker

## Pre requisites
- Docker installed


## Run Project

### Build go_sintrop
```
docker build -t go_sintrop .
```

```
# Build without cache
docker build --no-cache -t go_sintrop .
```

### Run go_sintrop
```
docker run -p=30303:30303 -p=8545:8545 -it -v /home/user/sintrop_node:/go-sintrop/sintrop_node  go_sintrop
```

Change /home/user/sintrop_node to your dir
```
## GETH

### Start a node

Change miner.etherbase to your wallet address and run the following command:

```
geth --identity Sintrop --datadir ./sintrop_node \
  --sintrop \
  --syncmode "full" \
  --networkid 250225 \
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
geth --identity Sintrop --datadir ./sintrop_node \
  --sintrop \
  --syncmode "full" \
  --networkid 250225 \
  --cache=1024 \
  --port 30303 \
  --authrpc.addr localhost --authrpc.port 8551 \
  --http.vhosts=* --http.addr "0.0.0.0" --http.port 8545 --http=true --http.api debug,net,eth,web3,txpool --ws=true --ws.addr 0.0.0.0 --ws.port 8546 --ws.origins "*" \
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
