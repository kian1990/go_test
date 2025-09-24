#!/bin/bash
ps aux|grep kian_my_bot|grep -v grep|awk '{print $2}'|xargs kill -9
source env.sh
nohup ./kian_my_bot &
tail -f nohup.out