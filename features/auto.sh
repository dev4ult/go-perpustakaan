#!/bin/bash

if [[ $# != 1 ]]; then
    echo "argument can not be less or more than 1"
    exit 400
fi

FEATURE="${1,}"
mkdir $FEATURE
cd $FEATURE

mkdir dtos handler mocks repository usecase

FILES=("interfaces.go" "entities.go" "dtos/request.go" "dtos/response.go" "handler/controller.go" "usecase/service.go" "repository/model.go")

for FILE in "${FILES[@]}"; do
    sed -e "s/placeholder/${FEATURE}/g" -e"s/Placeholder/${FEATURE^}/g" ../_blueprint/$FILE > $FILE
done

sed -e "s/placeholder/${FEATURE}/g" -e"s/Placeholder/${FEATURE^}/g" ../_blueprint/routes.go > "../../routes/${FEATURE}s.go"

