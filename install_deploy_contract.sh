# bin/bash
# 安装部署合约
set -euo
INFO(){
    local content=$*
    local DATE_NOW=`date "+%Y-%m-%d %H:%M:%S"`
    echo -e "\033[32m[INFO ${DATE_NOW}][chainmaker] ${content} \033[0m"
}
INFO "go mod tidy"
go mod tidy
INFO "编译chainmaker合约"
bash ./scripts/build_contract.sh
INFO "部署合约"
go run deploy/main.go