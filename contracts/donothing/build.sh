#!/bin/bash

contract_name="donothing"

# echo "please input contract name: "
# read contract_name
go build -ldflags="-s -w" -o $contract_name
7z a $contract_name $contract_name