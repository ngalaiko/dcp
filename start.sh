#!/bin/bash 

TODAY=$(date +'%Y-%m-%d')
DIR="problems/${TODAY}"
TEMPLATE=".template"

if [ ! -d "${DIR}" ]; then
    mkdir -p "${DIR}"
    cp "${TEMPLATE}"/* "${DIR}/"
fi

git add "${DIR}"
git commit -m "${TODAY}: started"

vim "${DIR}/main.go"
