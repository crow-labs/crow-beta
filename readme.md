# crow
HackMD doc: https://hackmd.io/@oZRKXxjmQTCueM5T_JZcMQ/rJFEbiio9 (most up to date information)
**crow** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

# CROW 
A marketplace for physical goods using an escrow (hence CROW) based model for payments
## Idea
### Short term
Amazon-like marketplace where listed items and transactions are stored and manipulated on a blockchain

### Mid term
eBay-like marketplace where listed items can be sold in an auction

### Long term
Point of sale transactions for brick and mortar stores.
### Description
- We have producers and users whitelists as well as a marketplace
- A producer is a whitelisted address that can list goods on the marketplace
- A user is a whitelisted address that can buy goods on the marketplace
- The marketplace stores listed items
- Governance settles disputes between users and producers
    - Item is not as described
    - Item did not arrive
    - Item is damaged 
- Staked addresses can vote for governance and are rewarded for participating in settling disputes

## Module Txs
### Marketplace
![marketplaceCommands](https://user-images.githubusercontent.com/99290400/185529026-2eb98223-37c3-499e-947e-2cf2f3ff1813.png)

### Escrow
![escrowCommans](https://user-images.githubusercontent.com/99290400/185529025-b6821af1-db66-4fba-9cb6-a323d0fa0915.png)

### Whitelist
![whitelistCommands](https://user-images.githubusercontent.com/99290400/185529027-84896849-2e89-41f0-9488-c8aae811f42f.png)

## Get started

```
ignite chain serve -r
```

`serve` command installs dependencies, builds, initializes, and starts the blockchain in development.

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
