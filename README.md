## Sintrop: A Blockchain for Social and Environmental Impact Applications

Blockchain technology has the potential to significantly impact the fight against climate change
and humanity’s social problems, many different applications are possible to be built with smart
contracts. The problem is that decentralized blockchains process very few transactions per second,
so projects with real-world impact need to compete for transactions with memecoins, NFTs and
others speculative low-value tokens and projects. This competition increases gas fees, limiting the number of impact business models able to operate. What is needed is a decentralized blockchain
infrastructure focused on applications that can make the world better. We propose a public peer-to-peer impact blockchain, running with proof-of-work consensus algorithm with the aim of
operating mostly with renewable energy.

## Tutorials

- How to run a Sintrop node [link](./tutorials/nodes/how-to-run-sintrop.md).
- Run Sintrop with docker [link](./tutorials/nodes/run-sintrop-with-docker.md).
- How to run a Sequoia testnet node [link](./tutorials/nodes/how-to-run-testnet-sequoia-node.md).
- Run Sequoia with docker [link](./tutorials/nodes/run-sequoia-with-docker.md).

## CoreGeth: An Ethereum Protocol Provider

Forked from CoreGeth protocol.

### Documentation

- CoreGeth documentation is available [here](https://etclabscore.github.io/core-geth).
  + Getting Started [Installation](https://etclabscore.github.io/core-geth/getting-started/installation) and [CLI](https://etclabscore.github.io/core-geth/getting-started/run-cli)
  + [JSONRPC API](https://etclabscore.github.io/core-geth/JSON-RPC-API)
  + [Developers](https://etclabscore.github.io/core-geth/developers/build-from-source)
  + [Tutorials](https://etclabscore.github.io/core-geth/tutorials/private-network)
- Further [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum) documentation about can be found [here](https://geth.ethereum.org/docs/).
- Documentation about documentation lives [here](./docs/developers/documentation.md).

## Contribution

Thank you for considering to help out with the source code! We welcome contributions
from anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to go-sintrop, please fork, fix, commit and send a pull request
for developers to review and merge into the main code base. If you wish to submit more complex changes though, submit a proposal first to ensure those changes are in line with the general philosophy of the project and/or get some early feedback which can make both your efforts much lighter as well as our review and merge procedures quick and simple.

Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting)
   guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary)
   guidelines.
 * Pull requests need to be based on and opened against the `master` or `release` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "eth, rpc: make trace configs optional"

Please see the [Developers' Guide](https://github.com/ethereum/go-ethereum/wiki/Developers'-Guide)
for more details on configuring your environment, managing project dependencies, and testing procedures.

## License

The go-sintrop library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-sintrop binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
