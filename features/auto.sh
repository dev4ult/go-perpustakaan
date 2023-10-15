#!/bin/bash

if [[ $# < 1 ]]; then
    echo "argument must not be less than 1"
    exit
fi

FEATURE="${1,,}"
mkdir $FEATURE
cd $FEATURE

printf "package $FEATURE\n\nimport \"gorm.io/gorm\"\n\ntype ${FEATURE^} struct {\n\tgorm.Model\n\n\tID int\n}" > entities.go

printf "package $FEATURE\n\nimport (\n\t\"github.com/labstack/echo/v4\"\n\n\t)"  > interfaces.go

mkdir dtos handler mocks repository usecase


