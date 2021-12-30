#!/bin/bash -e
cd `dirname $0`

ECR=xxxxxxxxxxxx.dkr.ecr.ap-northeast-1.amazonaws.com
REPOSITORY=$ENV/api
NOWJST=$(TZ=Asia/Tokyo date '+%Y%m%d_%H%M%S')

aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin $ECR
docker tag $ENV/api:latest $ECR/$REPOSITORY:$NOWJST
docker push $ECR/$REPOSITORY:$NOWJST

cd -
