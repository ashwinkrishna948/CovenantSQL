#!/bin/bash -x
test_case=$1
set -e

TEST_WD=$(cd $(dirname $0)/; pwd)
PROJECT_DIR=$(cd ${TEST_WD}/../../; pwd)
cd ${TEST_WD}

OLD_BIN_DIR=${CACHE_DIR}/${PREV_VERSION}
NEW_BIN_DIR=${PROJECT_DIR}/bin

LOGS_DIR=${JOB_CACHE}/old_${test_case}_log
mkdir -p ${LOGS_DIR}

case $test_case in
    "client")
        CLIENTBIN=${OLD_BIN_DIR}/cql
        BPBIN=${NEW_BIN_DIR}/cqld
        MINERBIN=${NEW_BIN_DIR}/cql-minerd
        ;;
    "bp")
        CLIENTBIN=${NEW_BIN_DIR}/cql
        BPBIN=${OLD_BIN_DIR}/cqld
        MINERBIN=${NEW_BIN_DIR}/cql-minerd
        ;;
    "miner")
        CLIENTBIN=${NEW_BIN_DIR}/cql
        BPBIN=${NEW_BIN_DIR}/cqld
        MINERBIN=${OLD_BIN_DIR}/cql-minerd
        ;;
    *)
        exit 1
        ;;
esac


# start bp
nohup ${BPBIN} -config ${PROJECT_DIR}/test/integration/node_0/config.yaml 2>${LOGS_DIR}/bp0.log &
nohup ${BPBIN} -config ${PROJECT_DIR}/test/integration/node_1/config.yaml 2>${LOGS_DIR}/bp1.log &
nohup ${BPBIN} -config ${PROJECT_DIR}/test/integration/node_2/config.yaml 2>${LOGS_DIR}/bp2.log &

# wait bp start
sleep 20

# start miner
nohup ${MINERBIN} -config ${PROJECT_DIR}/test/integration/node_miner_0/config.yaml 2>${LOGS_DIR}/miner0.log &
nohup ${MINERBIN} -config ${PROJECT_DIR}/test/integration/node_miner_1/config.yaml 2>${LOGS_DIR}/miner1.log &
nohup ${MINERBIN} -config ${PROJECT_DIR}/test/integration/node_miner_2/config.yaml 2>${LOGS_DIR}/miner2.log &

# wait miner start
sleep 20

${CLIENTBIN} -config ${PROJECT_DIR}/test/integration/node_c/config.yaml -get-balance

${CLIENTBIN} -config ${PROJECT_DIR}/test/integration/node_c/config.yaml -create 2 -wait-tx-confirm | tee dsn.txt

#get dsn
dsn=$(cat dsn.txt)
if [ -z "$dsn" ]; then
    exit 1
fi

${CLIENTBIN} -config ${PROJECT_DIR}/test/integration/node_c/config.yaml -dsn ${dsn} \
    -command 'create table test_for_new_account(column1 int);'

${CLIENTBIN} -config ${PROJECT_DIR}/test/integration/node_c/config.yaml -dsn ${dsn} \
    -command 'show tables;' | tee result.log

grep "1 row" result.log

