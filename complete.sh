#!/bin/bash 

TODAY=$(date +'%Y-%m-%d')

git add "problems/${TODAY}"
git commit -m "${TODAY}: completed"
