## Como rodar node Sintrop com Docker

## Rodar Sequoia com Docker

## Requisitos
- Docker instalado


## Rodar o node

### Fazer o build go-sintrop
```
docker build -t go-sintrop .
# Build without cache
docker build --no-cache -t go-sintrop .
```

### Executar go-sintrop
```
docker run -p=30303:30303 -p=8545:8545 -it -v /home/user/sintrop_node:/go-sintrop/sintrop_node  go-sintrop

## Trocar /home/user/sintrop_node para o seu diretório
```
## GETH

### Iniciar o node

Trocar miner.etherbase para o seu endereço de carteira e executar o comando abaixo:

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

### Rodar um archive node
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

### Operar a rede

```
balance = web3.fromWei(eth.getBalance("0x0000000000000000000000000000000000000000), "ether");
eth.blockNumber
web3.eth.getBlock(eth.blockNumber)
```
## MINERAÇÂO

### Iniciar e interromper a mineração
```
  miner.start()
  miner.stop()
```
