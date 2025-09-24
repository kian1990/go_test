#!/bin/bash
ps aux|grep monster_k_bot|grep -v grep|awk '{print $2}'|xargs kill -9
source env.sh
nohup ./monster_k_bot &
tail -f nohup.out