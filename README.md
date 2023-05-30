<p align="center">
  <img src="./black-logo.svg" width="300">
</p>

<div align="center">

[![version](https://img.shields.io/github/tag/elysium-station/black.svg)](https://github.com/elysium-station/black/releases/latest)
[![CircleCI](https://circleci.com/gh/Elysium-Station/black/tree/master.svg?style=shield)](https://circleci.com/gh/Elysium-Station/black/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/elysium-station/black)](https://goreportcard.com/report/github.com/elysium-station/black)
[![API Reference](https://godoc.org/github.com/Elysium-Station/black?status.svg)](https://godoc.org/github.com/Elysium-Station/black)
[![GitHub](https://img.shields.io/github/license/elysium-station/black.svg)](https://github.com/Elysium-Station/black/blob/master/LICENSE.md)
[![Twitter Follow](https://img.shields.io/twitter/follow/BLACK_CHAIN.svg?label=Follow&style=social)](https://twitter.com/BLACK_CHAIN)
[![Discord Chat](https://img.shields.io/discord/704389840614981673.svg)](https://discord.com/invite/kQzh3Uv)

</div>

<div align="center">

### [Telegram](https://t.me/blacklabs) | [Medium](https://medium.com/elysium-station) | [Discord](https://discord.gg/JJYnuCx)

</div>

Reference implementation of Black, a blockchain for cross-chain DeFi. Built using the [cosmos-sdk](https://github.com/cosmos/cosmos-sdk).

## Mainnet

The current recommended version of the software for mainnet is [v0.23.0](https://github.com/Elysium-Station/black/releases/tag/v0.23.0). The master branch of this repository often contains considerable development work since the last mainnet release and is __not__ runnable on mainnet.

### Installation and Setup
For detailed instructions see [the Black docs](https://docs.black.io/docs/participate/validator-node).

```bash
git checkout v0.23.0
make install
```

End-to-end tests of Black use a tool for generating networks with different configurations: [kvtool](https://github.com/Elysium-Station/kvtool).
This is included as a git submodule at [`tests/e2e/kvtool`](tests/e2e/kvtool/).
When first cloning the repository, if you intend to run the e2e integration tests, you must also
clone the submodules:
```bash
git clone --recurse-submodules https://github.com/Elysium-Station/black.git
```

Or, if you have already cloned the repo: `git submodule update --init`

## Testnet

For further information on joining the testnet, head over to the [testnet repo](https://github.com/Elysium-Station/black-testnets).

## Docs

Black protocol and client documentation can be found in the [Black docs](https://docs.black.io).

If you have technical questions or concerns, ask a developer or community member in the [Black discord](https://discord.com/invite/kQzh3Uv).

## Security

If you find a security issue, please report it to security [at] black.io. Depending on the verification and severity, a bug bounty may be available.

## License

Copyright © Black Labs, Inc. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).
