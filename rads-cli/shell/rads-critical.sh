#!/bin/sh -xe

# 閏年は考慮していないし、31日に実行するとだめになる
CURRENT_MONTH=`date +"%m"`
DAY=30
if [ $CURRENT_MONTH = 03 ]; then
    DAY=28
fi

YEAR=`date +"%Y" -D %s -d "$(($(date +%s) - 86400 * DAY))"`
MONTH=`date +"%m" -D %s -d "$(($(date +%s) - 86400 * DAY))"`

HOST=$1
ROSTERNAME=$2

./rads-cli -y $YEAR -m $MONTH -f $ROSTERNAME -H $HOST

wget --post-data='{ "Title":"勤務表","Message":"勤務表をドライブにコピーしました" }' --header=Content-Type:application/json $HOST/postSlack