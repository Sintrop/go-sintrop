## Como rodar um node na Sintrop

### Baixe o software GO-SINTROP

Baixe a versão mais recente do go-sintrop no repositório do projeto [sintrop/go-sintrop](https://github.com/sintrop/go-sintrop).

Clique na aba 'Releases' para baixar a versão mais recente. Escolha o arquivo certo de acordo com seu sistema operacional.

### Rodar Sintrop

Extraia o arquivo e acesse a pasta do projeto em um terminal.

Altere miner.etherbase para o endereço da sua carteira e execute o seguinte comando:

```
./geth --identity Sintrop --datadir ./sintrop_node --sintrop --syncmode "full" --networkid 250225 --cache=1024 --port 25225 --http.vhosts=* --http.addr "0.0.0.0" --http.port 8545 --http=true --miner.threads=1 --miner.etherbase=0x0000000000000000000000000000000000000000 console
```
### Pausar o node

Para pausar e interromper corretamente o node utilize o comando:

```
ctrl + d
```

### MINERAÇÃO

### Iniciar e interromper a mineração

Para iniciar a mineração com CPU, execute o seguinte comando no console:

```
  miner.start()
  miner.stop()
```
