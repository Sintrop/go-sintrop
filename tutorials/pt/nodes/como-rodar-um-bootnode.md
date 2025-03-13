## Como rodar um bootnode

### Requisitos
- IP público fixo;
- Última versão do [sintrop/go-sintrop](https://github.com/sintrop/go-sintrop);
- Encaminhando de porta da porta 25225 do modem para a máquina local;

### Rodar o bootnode

Após baixar a versão mais recente do go-sintrop, execute o seguinte comando na pasta do projeto para iniciar o nó.

```
./geth --identity Sintrop --datadir ./sintrop_node --sintrop --syncmode "full" --networkid 250225 --cache=1024 --port 25225 --http.port 8545 --http=true console
```

### Adicionar o bootnode no sourcecode

Para ver seu enode, digite "admin" no console geth. Copie seu enode.

Clone o repositório go-sintrop, crie uma nova ramificação para atualizar o bootnode. Adicione seu enode em /param/bootnodes_sintrop.go

Publique a ramificação e abra um Pull Request para revisão.