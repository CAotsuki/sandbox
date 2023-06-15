MYDIR=$(cd $(dirname $0) && pwd)
cd ${MYDIR}

# START: install for mac
#curl -o cloud-sql-proxy \
#  https://storage.googleapis.com/cloud-sql-connectors/cloud-sql-proxy/v2.1.2/cloud-sql-proxy.darwin.amd64
#
#chmod +x cloud-sql-proxy
#mkdir -p bin
#mv cloud-sql-proxy bin/.
# END: install for mac

PROJECT="ca-otsuki-study-gce"
REGION="asia-northeast2"
INSTANCE_NAME="todo"
./bin/cloud-sql-proxy --address 0.0.0.0 --port 5432 ${PROJECT}:${REGION}:${INSTANCE_NAME}
