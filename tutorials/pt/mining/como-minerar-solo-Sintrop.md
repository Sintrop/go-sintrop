## Como minerar Sintrop solo

### Requisitos
- Sintrop node rodando;
- Porta local 8545 liberada;


### Setar miner.etherbase

Antes de começar a minerar, você precisa definir seu miner.etherbase, o endereço da carteira que receberá as recompensas ao encontrar novos blocos.

Para fazer isso, você pode passar o sinalizador etherbase no comando geth ao iniciar o nó, ou defini-lo manualmente com o comando:

```
miner.setEtherbase('0xsuacarteira')
```

### Iniciar e parar mineração com CPU

Para iniciar a mineração com CPU, execute o seguinte comando no console geth:

```
  miner.start()
  miner.stop()
```

### Minerando com GPU / Placa de vídeo

Para minerar com GPU, é necessário usar um software de mineração de terceiros. Há muitas opções disponíveis, e nossa comunidade está desenvolvendo uma versão atualizada do ethminer, mas ainda não está disponível. Neste tutorial, mostraremos como minerar usando o software ethminer.

Primeiro, baixe a versão mais recente do [ethminer](https://github.com/ethereum-mining/ethminer/releases/tag/v0.18.0). 

Extraia o software, abra um novo terminal na pasta do minerador e execute o seguinte comando para começar a minerar sua placa de vídeo:

```
./ethminer -G -P http://localhost:8545
```

### Energia renovável

Incentivamos todos os nós e mineradores a usar energia renovável para alimentar suas máquinas. Tente encontrar energia solar, eólica, hidrelétrica e outras fontes de energia renováveis. Nosso objetivo é ser uma rede alimentada por energia renovável.

### Informações adicionais

Certifique-se de ter instalado corretamente os drivers para suas GPUs. O software de mineração funciona melhor no Ubuntu.

Certifique-se de fazer overclock corretamente nas placas e definir o limite de energia apropriado para evitar consumir energia sem melhorar sua taxa de hash.
