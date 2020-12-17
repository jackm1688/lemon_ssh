#!/bin/bash

#cd /usr/local/etcd-v3.3.25;nohup ./etcd --config-file config.yaml &
DIR=/usr/local/etcd-v3.3.25
cd $DIR
echo "================ start etcd server================="
PID=`ps aux |grep etcd |grep "config.yaml" |grep -v grep|awk -F " " '{print $2}'`
if [ $PID > 0 ];then
   echo " etcd service is already running!"
   exit 0
fi

nohup ./etcd --config-file config.yaml 1>>./start.log 2>&1  &
stat=$?
if [ $stat == 0 ]
then
   pid=`ps aux |grep etcd |grep "config.yaml" |grep -v grep|awk -F " " '{print $2}'`
   echo "etcd service start success, pid:$pid"
   exit 0
else
  echo "etcd server start failed"
fi
