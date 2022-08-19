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
#### Create Listing
![txCreateListing](https://user-images.githubusercontent.com/99290400/185530980-41f81088-fed9-40c7-95b4-db8901086501.png)
#### Create Order
![txCreateOrder](https://user-images.githubusercontent.com/99290400/185530982-ebaaf5eb-1ef6-4ca3-8d8c-f5615ae288ba.png)


### Escrow
![escrowCommans](https://user-images.githubusercontent.com/99290400/185529025-b6821af1-db66-4fba-9cb6-a323d0fa0915.png)
#### Cancel Escrow
![txCancelEscrow](https://user-images.githubusercontent.com/99290400/185531120-8ede8809-48ef-4c24-87fc-a9a9f0826709.png)
#### Item Damaged
![txItemDamaged](https://user-images.githubusercontent.com/99290400/185531122-42046aa8-3be4-4be3-82a3-26ebf249f2a2.png)
#### Item Incorrect
![txItemIncorrect](https://user-images.githubusercontent.com/99290400/185531124-de687509-276b-4765-ae26-a5087f1562cb.png)
#### Item Not Received
![txItemNotReceived](https://user-images.githubusercontent.com/99290400/185531125-9018de1b-dde6-4861-82f4-7240b1fc9b31.png)
#### Item Received
![txItemReceived](https://user-images.githubusercontent.com/99290400/185531126-52a7ba45-4702-49da-9363-f3837adc0bdb.png)
#### Item Shipped
![txItemShipped](https://user-images.githubusercontent.com/99290400/185531127-a5d5ce62-9f89-4650-bb80-1f7106c72456.png)


### Whitelist
![whitelistCommands](https://user-images.githubusercontent.com/99290400/185529027-84896849-2e89-41f0-9488-c8aae811f42f.png)
#### Create Producer
![txCreateProducer](https://user-images.githubusercontent.com/99290400/185531145-666e8401-095d-41c5-a46d-fd38bebfc258.png)
#### Create User
![txCreateUser](https://user-images.githubusercontent.com/99290400/185531149-49d21080-359a-43c6-a6c3-8389018242b4.png)


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
