## Network operation commands

Common commands to interact with the blockchain at the node console.

### List of information and functions

```
eth
```

### Admin functions

```
admin
```

### Add a peer

```
admin.addPeer("enode:30303")
```

### Get account balance (wei)

```
eth.getBalance('0xwallet')
```

### Get account balance (eth)

```
web3.fromWei(eth.getBalance("0xwalletaddress"), 'ether')
```

### Current block number

```
eth.blockNumber
```

### Send Sintrop

```
eth.sendTransaction({from:eth.coinbase, to:'0xaddress', value:web3.toWei(1,"ether"), gas:21000});
```

