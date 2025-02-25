## How to run a Sintrop bootnode

### Requisitos
- IP fixo público;
- Última versão do [sintrop/go-sintrop](https://github.com/sintrop/go-sintrop);
- Redirecionamento da porta 30303 no modem para a máquina;

### Comandos
- Run bootnode:
```
./geth --identity Sintrop --datadir ./sintrop_node \
  --sintrop \
  --syncmode "full" \
  --networkid 250225 \
  --cache=1024 \
  --port 30303 \
  console
```
