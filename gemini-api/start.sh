#!/bin/bash
export API_KEY=your_api_key
ps aux|grep gemini-api|grep -v grep|awk '{print $2}'|xargs kill -9
nohup ./gemini-api &
tail -f nohup.out