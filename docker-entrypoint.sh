#!/usr/bin/env bash

set -e

PARAMETER_STORE_PREFIX=${PARAMETER_STORE_PREFIX:-}
if [ -n "$PARAMETER_STORE_PREFIX" ]; then
    export DB_USER=$(aws ssm get-parameters --name /${PARAMETER_STORE_PREFIX}/database/sample/master/user --query "Parameters[0].Value" --region ap-northeast-1 --output text)
    export DB_PASSWORD=$(aws ssm get-parameters --name /${PARAMETER_STORE_PREFIX}/database/sample/master/password --with-decryption --query "Parameters[0].Value" --region ap-northeast-1 --output text)
    export DB_HOST=$(aws ssm get-parameters --name /${PARAMETER_STORE_PREFIX}/database/sample/master/host --query "Parameters[0].Value" --region ap-northeast-1 --output text)
    export DB_NAME=$(aws ssm get-parameters --name /${PARAMETER_STORE_PREFIX}/database/sample/master/name --query "Parameters[0].Value" --region ap-northeast-1 --output text)
    export DB_PORT=$(aws ssm get-parameters --name /${PARAMETER_STORE_PREFIX}/database/sample/master/port --query "Parameters[0].Value" --region ap-northeast-1 --output text)
fi

exec "$@"
