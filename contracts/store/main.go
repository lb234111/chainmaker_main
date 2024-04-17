// coding:utf-8
// 存储合约
package main

import (
	"log"
	"strconv"

	pb "chainmaker.org/chainmaker/contract-sdk-go/v2/pb/protogo"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sandbox"
	"chainmaker.org/chainmaker/contract-sdk-go/v2/sdk"
)

type STORE struct{}

func (s *STORE) InitContract() pb.Response {
	return sdk.Success([]byte("success"))
}

// UpgradeContract use to upgrade contract
func (h *STORE) UpgradeContract() pb.Response {
	return sdk.Success([]byte("Upgrade success"))
}

// InvokeContract use to select specific method
func (h *STORE) InvokeContract(method string) pb.Response {
	// according method segment to select contract functions
	switch method {
	case "get":
		return h.get()
	case "store":
		return h.store()
	case "withdraw":
		return h.withdraw()
	case "transfer":
		return h.transfer()
	}
	return sdk.Error("invalid method")
}

func (h *STORE) get() pb.Response {
	args := sdk.Instance.GetArgs()
	key := string(args["addr"])
	value_str, exists, err := sdk.Instance.GetStateFromKeyWithExists(key)
	if err != nil {
		return sdk.Error(err.Error())
	}
	if exists == false {
		return sdk.Error("addr not exists")
	}
	return sdk.Success([]byte(value_str))
}

func (h *STORE) store() pb.Response {
	args := sdk.Instance.GetArgs()
	key := string(args["addr"])
	val := string(args["value"])
	store_value, err := strconv.Atoi(val)
	if err != nil {
		return sdk.Error(err.Error())
	}

	old_value_str, exists, err := sdk.Instance.GetStateFromKeyWithExists(key)
	if err != nil {
		return sdk.Error(err.Error())
	}
	var old_value int
	if exists == false {
		old_value = 0
	} else {
		old_value, err = strconv.Atoi(old_value_str)
		if err != nil {
			return sdk.Error(err.Error())
		}
	}
	new_value := old_value + store_value

	sdk.Instance.PutStateFromKey(key, strconv.Itoa(new_value))
	return sdk.Success([]byte("STORE SUCCESS! " + "store value:" + strconv.Itoa(store_value) + " to addr:" + key))
}

func (h *STORE) withdraw() pb.Response {
	args := sdk.Instance.GetArgs()
	key := string(args["addr"])
	val := string(args["value"])
	withdraw_value, err := strconv.Atoi(val)
	if err != nil {
		return sdk.Error(err.Error())
	}

	old_value_str, exists, err := sdk.Instance.GetStateFromKeyWithExists(key)
	if err != nil {
		return sdk.Error(err.Error())
	}
	var old_value int
	if exists == false {
		return sdk.Error("addr not exists")
	} else {
		old_value, _ = strconv.Atoi(old_value_str)
	}

	new_value := old_value - withdraw_value
	if new_value < 0 {
		return sdk.Error("fail")
	}
	sdk.Instance.PutStateFromKey(key, strconv.Itoa(new_value))
	return sdk.Success([]byte("WITHDRAW SUCCESS! " + "withdraw value:" + strconv.Itoa(withdraw_value) + " from addr:" + key))
}

func (h *STORE) transfer() pb.Response {
	args := sdk.Instance.GetArgs()
	key1 := string(args["addr_from"])
	key2 := string(args["addr_to"])
	val := string(args["value"])

	transfer_value, err := strconv.Atoi(val)
	if err != nil {
		return sdk.Error(err.Error())
	}

	old_value_str1, exists1, err := sdk.Instance.GetStateFromKeyWithExists(key1)
	if err != nil {
		return sdk.Error(err.Error())
	}
	old_value_str2, exists2, err := sdk.Instance.GetStateFromKeyWithExists(key2)
	if err != nil {
		return sdk.Error(err.Error())
	}

	var old_value1, old_value2 int
	if exists1 == false {
		return sdk.Error("addr not exists")
	} else {
		old_value1, _ = strconv.Atoi(old_value_str1)
	}
	if exists2 == false {
		old_value2 = 0
	} else {
		old_value2, _ = strconv.Atoi(old_value_str2)
	}

	new_value1 := old_value1 - transfer_value
	new_value2 := old_value2 + transfer_value
	if new_value1 < 0 {
		return sdk.Error("fail")
	}
	sdk.Instance.PutStateFromKey(key1, strconv.Itoa(new_value1))
	sdk.Instance.PutStateFromKey(key2, strconv.Itoa(new_value2))
	return sdk.Success([]byte("TRANSFER SUCCESS! " + "transfer value:" + val + " from addr:" + key1 + " to addr:" + key2))
}

// main
func main() {
	err := sandbox.Start(new(STORE))
	if err != nil {
		log.Fatal(err)
	}
}
