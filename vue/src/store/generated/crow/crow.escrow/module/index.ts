// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgItemIncorrect } from "./types/escrow/tx";
import { MsgCancelEscrow } from "./types/escrow/tx";
import { MsgItemDamaged } from "./types/escrow/tx";
import { MsgItemReceived } from "./types/escrow/tx";
import { MsgItemNotReceived } from "./types/escrow/tx";
import { MsgItemShipped } from "./types/escrow/tx";


const types = [
  ["/crow.escrow.MsgItemIncorrect", MsgItemIncorrect],
  ["/crow.escrow.MsgCancelEscrow", MsgCancelEscrow],
  ["/crow.escrow.MsgItemDamaged", MsgItemDamaged],
  ["/crow.escrow.MsgItemReceived", MsgItemReceived],
  ["/crow.escrow.MsgItemNotReceived", MsgItemNotReceived],
  ["/crow.escrow.MsgItemShipped", MsgItemShipped],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgItemIncorrect: (data: MsgItemIncorrect): EncodeObject => ({ typeUrl: "/crow.escrow.MsgItemIncorrect", value: MsgItemIncorrect.fromPartial( data ) }),
    msgCancelEscrow: (data: MsgCancelEscrow): EncodeObject => ({ typeUrl: "/crow.escrow.MsgCancelEscrow", value: MsgCancelEscrow.fromPartial( data ) }),
    msgItemDamaged: (data: MsgItemDamaged): EncodeObject => ({ typeUrl: "/crow.escrow.MsgItemDamaged", value: MsgItemDamaged.fromPartial( data ) }),
    msgItemReceived: (data: MsgItemReceived): EncodeObject => ({ typeUrl: "/crow.escrow.MsgItemReceived", value: MsgItemReceived.fromPartial( data ) }),
    msgItemNotReceived: (data: MsgItemNotReceived): EncodeObject => ({ typeUrl: "/crow.escrow.MsgItemNotReceived", value: MsgItemNotReceived.fromPartial( data ) }),
    msgItemShipped: (data: MsgItemShipped): EncodeObject => ({ typeUrl: "/crow.escrow.MsgItemShipped", value: MsgItemShipped.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
