#!/bin/bash 

TODAY=$(date +'%Y-%m-%d')

git add "${TODAY}"
git commit -m "${TODAY}: completed"
