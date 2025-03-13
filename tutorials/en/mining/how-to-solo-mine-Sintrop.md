## How to solo mine Sintrop

### Requirements
- Sintrop node running;
- Local 8545 port allowed;


### Set miner.etherbase

Before start mining, you need to set your miner.etherbase, the wallet address that will receive the rewards when finding new blocks. 

To do it, you can pass the etherbase flag at the geth command when starting the node, or set it manually with the command:

```
miner.setEtherbase('0xyourwalletaddress')
```

### Start and stop mining with CPU

To start mining with CPU, run the following command at the geth console: 

```
  miner.start()
  miner.stop()
```

### Mining with GPU

To mine with GPU, it is needed to use a third party mining software. The are many options avaiable, and our community is developing an updated version of ethminer, but yet not avaiable. In this tutorial, we will show how to mine using ethminer software.

First, download the latest version of [ethminer](https://github.com/ethereum-mining/ethminer/releases/tag/v0.18.0). 

Extract the software, open a new terminal at the miner folder and run the following command to start mining with your graphics card:

```
./ethminer -G -P http://localhost:8545
```

### Renewable energy

We encourage all nodes and miners to use renewable energy to power up their machines. Try to find solar, wind, hydro and other renewable energy sources. Our goal is to be a (mostly) renewable energy powered network.

### Additional information

Make sure that you have properly installed the drivers to your GPUs. The mining software works better at Ubuntu.

Make sure to properly overclock the cards and set the appropriate power limit to avoid consume energy without improving your hashrate. 
