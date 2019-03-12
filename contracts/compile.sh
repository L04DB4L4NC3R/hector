#!/bin/bash

../solc/solcjs --abi Store.sol | awk '/JSON ABI/{x=1;next}x' > Store.abi
../solc/solcjs --bin Store.sol | awk '/Binary:/{x=1;next}x' > Store.bin
rm Store.bin Store.abi
abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go