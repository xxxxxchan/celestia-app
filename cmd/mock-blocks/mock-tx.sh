#!/usr/bin/env bash

for (( ; ; ))
do
    celestia-appd tx bank send mock --keyring-backend test --chain-id devnet-2 celes1jc62dulde0rrzh409x3tw0fp2wgnuny8ym8tmj -y
    celestia-appd tx payment payForMessage 0102030405060708 6f79e82b17370c1136426f79e82b17370c1136426f79e82b17370c1136426f79e82b17370c1136426f79e82b17370c1136426f79e82b17370c1136426f79e82b17370c113642 --from zkFART --keyring-backend test --chain-id devnet-2 -y
    sleep 5
done