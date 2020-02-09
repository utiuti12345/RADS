#!/bin/sh -xe

WEEK=`date +"%u"`
TODAY=`date +"%d"`
YESTERDAY_WEEK=`date +"%u" -D %s -d "$(($(date +%s) - 86400 * 1))"`
YESTERDAY=`date +"%d" -D %s -d "$(($(date +%s) - 86400 * 1))"`
if [ $WEEK = 6 -o $WEEK = 7 ]; then
    echo "土日なので実施しない"
    exit 0
elif [ $YESTERDAY = 01 -a $YESTERDAY_WEEK != 6 -a $YESTERDAY_WEEK != 7 ]; then
    echo "実行済なので実施しない"
    exit 0
elif [ $YESTERDAY = 02 -a $YESTERDAY_WEEK != 6 -a $YESTERDAY_WEEK != 7 ]; then
    echo "実行済なので実施しない"
    exit 0
fi

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