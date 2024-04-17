CHAINMAKER_GO_PATH=/home/chenxing/chainmaker-go

cd ${CHAINMAKER_GO_PATH}/config/config_tpl/chainconfig

sed -i "s/block_tx_capacity: 100/block_tx_capacity: 5000/g" ./bc*