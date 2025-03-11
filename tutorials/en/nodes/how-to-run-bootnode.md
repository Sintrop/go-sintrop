## How to run a Sintrop bootnode

### Requirements
- Fixed public IP;
- Lastest version of [sintrop/go-sintrop](https://github.com/sintrop/go-sintrop);
- Forwarding port 30303 on the modem to the local machine;

### Run a node

After downloading the latest version of go-sintrop, run the following command at the project folder to start the node.

```
./geth --identity Sintrop --datadir ./sintrop_node --sintrop --syncmode "full" --networkid 250225 --cache=1024 --port 30303 console
```

```
./geth --identity Sintrop --datadir ./sintrop_node --sintrop --syncmode "full" --networkid 250225 --cache=1024 --port 30303 --http.port 8545 --http=true --miner.threads=1 --miner.etherbase=0x0000000000000000000000000000000000000000 console
```

### Add bootnode to sourcecode

To see your enode, type "admin" at the geth console. Copy your enode. 

Clone the go-sintrop repository, create a new branch to update bootnode. Add your enode at the /param/bootnodes_sintrop.go

Publish the branch and open a Pull Request for review.