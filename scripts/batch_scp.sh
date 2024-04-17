hosts=(10.15.9.5)

echo "============= scp crypto-config ============="
for i in ${hosts[@]}; do
	rm -rf ../config_files/chain_server5/*
	scp -r chenxing@$i:~/chainmaker-master/config_files/chain1/crypto-config ~/chainmaker-master/config_files/chain_server5/crypto-config
done

