#!/bin/bash
ps aux|grep gemini-api|grep -v grep|awk '{print $2}'|xargs kill -9
echo "" > nohup.out