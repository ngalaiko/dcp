#!/bin/bash 

TODAY=$(date +'%Y-%m-%d')
TEMPLATE=".template"

if [ ! -d "${TODAY}" ]; then
    mkdir "${TODAY}"
    cp "${TEMPLATE}"/* "${TODAY}/"
fi

git add "${TODAY}"
git commit -m "${TODAY}: started"

vim "${TODAY}/main.go"
