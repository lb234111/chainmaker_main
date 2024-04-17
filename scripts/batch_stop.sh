hosts=(10.15.9.6 10.15.9.7 10.15.9.8)

echo "============= Batch stop all chain ============="
for i in ${hosts[@]}; do
	workspace="~/chainmaker-master"
	ssh chenxing@$i "cd $workspace && ./stop.sh"
done