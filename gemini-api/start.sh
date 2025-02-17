#!/bin/bash
export API_KEY=AIzaSyDj4nWVYvErFPH3DZM5epn0hdXsc3u32-I
ps aux|grep gemini-api|grep -v grep|awk '{print $2}'|xargs kill -9
nohup ./gemini-api &
tail -f nohup.out