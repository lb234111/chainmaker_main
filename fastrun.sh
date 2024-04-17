# bin/bash
# 快速启动

set -e
# 下载并编译长安链
make
# 启动链
bash ./restart.sh

# 编译安装合约
bash ./install_deploy_contract.sh 