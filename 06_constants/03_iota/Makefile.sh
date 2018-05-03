#!/bin/bash

# This is a script that encapsulated the "make clean" funcationality so
# that our build process can be cross platform. 

# Map the command to a variable
CMD=$1

# Directory Structure Variables
PROJECT_NAME="p2p-api-1"
PROJECT_VERSION=$(git describe --always --dirty)
PROJECT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
PROJECT_BUILD_TIME=$(date -u)
PROJECT_SHA_1=$(git rev-list -1 HEAD)

DIST_DIR="./dist"
DIST_CONFIG_DIR="$DIST_DIR/config"
DIST_PUBLIC_DIR="$DIST_DIR/public"
LOCAL_PUBLIC_DIR="public"
LOCAL_CONFIG_DIR="config"
TEMPLATE_NAME="bitbucket.centene.com/pdsrtr/rtr-advancement-api-template"
TEMPLATE_DIR="../$TEMPLATE_NAME"
TEMPLATE_CONFIG_DIR="$TEMPLATE_DIR/config"
TEMPLATE_PUBLIC_DIR="$TEMPLATE_DIR/public"

# Linker Flags
LINKER_FLAGS="-X bitbucket.centene.com/pdsrtr/rtr-advancement-api-template/env.Version=$PROJECT_VERSION"
LINKER_FLAGS=$LINKER_FLAGS" -X bitbucket.centene.com/pdsrtr/rtr-advancement-api-template/env.Branch=$PROJECT_BRANCH"
LINKER_FLAGS=$LINKER_FLAGS" -X bitbucket.centene.com/pdsrtr/rtr-advancement-api-template/env.BuildTime=$PROJECT_BUILD_TIME"
LINKER_FLAGS=$LINKER_FLAGS" -X bitbucket.centene.com/pdsrtr/rtr-advancement-api-template/env.SHA1=$PROJECT_SHA_1"

# Shared Functions
makeDeps() {
    makeVerifyTemplate
	go get -v ./...
}

makeDist() {
	if [ ! -d $DIST_DIR ]; then
		mkdir $DIST_DIR
	fi
	
    if [ ! -d $DIST_PUBLIC_DIR ]; then
		mkdir $DIST_PUBLIC_DIR
	fi

	cp -r "$LOCAL_PUBLIC_DIR/" $DIST_PUBLIC_DIR
	cp -r "$TEMPLATE_PUBLIC_DIR/" $DIST_PUBLIC_DIR

	if [ ! -d $DIST_CONFIG_DIR ];then
		mkdir $DIST_CONFIG_DIR
	fi

	cp -r "$TEMPLATE_CONFIG_DIR/" $DIST_CONFIG_DIR
}

makeVerifyTemplate() {
	if [ ! -d $TEMPLATE_DIR ]; then
	    echo $TEMPLATE_DIR
	    exit 1
	fi
	pushd $TEMPLATE_DIR > /dev/null 2>&1
	TEMPLETE_VERSION=$(git rev-parse --abbrev-ref HEAD)
	echo "WARNING: Using $TEMPLATE_NAME branch:[$TEMPLETE_VERSION]"
	echo "Verify this is the correct branch and it is up-to-date."
	popd > /dev/null 2>&1
}

# Paremeters

if [ "$CMD" == "clean" ]; then 
    go clean -x
    rm -rf dist/
    rm -rf vendor/
elif [ "$CMD" == "dist" ]; then 
    makeDist
elif [ "$CMD" == "deps" ]; then 
    makeDeps
elif [ "$CMD" == "test" ]; then 
    makeDeps
	go test -v ./... 
elif [ "$CMD" == "server" ]; then 
    makeDeps
    makeDist
	go build -v -o ${DIST_DIR}/${BINARY} -ldflags=$LINKER_FLAGS
elif [ "$CMD" == "start" ]; then 
	$DIST_DIR/$BINARY
elif [ "$CMD" == "swagger" ]; then 
    echo "Make $CMD"
elif [ "$CMD" == "verify-template" ]; then 
    makeVerifyTemplate
else 
    echo "Command \"$CMD\" not supported"
    echo "Possible Commands are as follows:"
    echo "clean"
    echo "deps"
    echo "dist"
    echo "server"
    echo "start"
    echo "swagger"
    echo "test"
    echo "verify-template"
fi