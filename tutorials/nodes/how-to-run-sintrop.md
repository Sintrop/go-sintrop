## How to run a Sintrop node

### Download GO-SINTROP

Download the latest version o go-sintrop at the project repo [sintrop/go-sintrop](https://github.com/sintrop/go-sintrop).

Click at the 'Releases' tab to download the latest version. Choose the right file accordingly to your operating system.

### Run Sintrop

Extract the file, and access the project path.

Change miner.etherbase to your wallet address and run the following command:

```
./geth --identity Sintrop --datadir ./sintrop_node --sintrop --syncmode "full" --networkid 250225 --cache=1024 --port 30303 --http.vhosts=* --http.addr "0.0.0.0" --http.port 8545 --http=true --miner.threads=1 --miner.etherbase=0x0000000000000000000000000000000000000000 console
```

### MINING


### Start and stop mining

To start mining with CPU, run the following command at the geth console: 

```
  miner.start()
  miner.stop()
```
