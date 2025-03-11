## Comandos e operações

Princiapis comandos para interagir com a blockchain no console.

### Lista de informações e funções disponíveis

```
eth
```

### Funções admin

```
admin
```

### Adicionar um novo peer

```
admin.addPeer("enode:30303")
```

### Ver saldo (wei)

```
eth.getBalance('0xwallet')
```

### Ver saldo de Sintrops

```
web3.fromWei(eth.getBalance("0xwalletaddress"), 'ether')
```

### Bloco atual

```
eth.blockNumber
```

### Enviar Sintrop

```
eth.sendTransaction({from:eth.coinbase, to:'0xaddress', value:web3.toWei(1,"ether"), gas:21000});
```

