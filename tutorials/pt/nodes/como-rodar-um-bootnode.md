## Como rodar um bootnode

### Requisitos
- IP público fixo;
- Última versão do [sintrop/go-sintrop](https://github.com/sintrop/go-sintrop);
- Encaminhando de porta da porta 30303 do modem para a máquina local;

### Rodar o bootnode

Após baixar a versão mais recente do go-sintrop, execute o seguinte comando na pasta do projeto para iniciar o nó.

```
./geth --identity Sintrop --datadir ./sintrop_node --sintrop --syncmode "full" --networkid 250225 --cache=1024 --port 30303 console
```

```
./geth --identity Sintrop --datadir ./sintrop_node --sintrop --syncmode "full" --networkid 250225 --cache=1024 --port 30303 --http.port 8545 --http=true --miner.threads=1 --miner.etherbase=0x0000000000000000000000000000000000000000 console
```

### Adicionar o bootnode no sourcecode

Para ver seu enode, digite "admin" no console geth. Copie seu enode.

Clone o repositório go-sintrop, crie uma nova ramificação para atualizar o bootnode. Adicione seu enode em /param/bootnodes_sintrop.go

Publique a ramificação e abra um Pull Request para revisão.