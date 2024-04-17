# bin/bash
# 编译合约
set -euo

INFO(){
    local content=$*
    local DATE_NOW=`date "+%Y-%m-%d %H:%M:%S"`
    echo -e "\033[32m[INFO ${DATE_NOW}][chainmaker] ${content} \033[0m"
}

CURR_PATH="$(pwd)"
INFO "编译donothin合约"
cd $CURR_PATH/contracts/donothing && bash ./build.sh
INFO "编译store合约"
cd $CURR_PATH/contracts/store && bash ./build.sh
