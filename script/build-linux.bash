#!/usr/bin/env bash

##
## Global config
##
DIRNAME_CLI=/usr/bin/dirname
GO_CLI=go
PWD_CLI=/bin/pwd
UNAME_CLI=/bin/uname
SCRIPT_DIR=$(cd -- "$(${DIRNAME_CLI} "${BASH_SOURCE[0]:-${(%):-%x}}")" &> /dev/null && ${PWD_CLI})
SYSTEM_PLATFORM=$($UNAME_CLI)
PROJECT_DIR=$SCRIPT_DIR/..

##
## Project config
##
####
#### Project-level config
####
PROJECT_GO_VERSION="${MY_PROJECT_GO_VERSION:=1.23.12}"
PROJECT_USE_DEFAULT_GO_CLI="${MY_PROJECT_USE_DEFAULT_GO_CLI:=OFF}"

##
## My variables
##
MY_PROJECT_GO_BIN=""
MY_PROJECT_GO_BIN_DETECTED=false
MY_PROJECT_GO_CLI=$GO_CLI$PROJECT_GO_VERSION
MY_PROJECT_GO_CLI_DETECTED=false
MY_PROJECT_GO_VERSION=$PROJECT_GO_VERSION
if [ "ON" = "$PROJECT_USE_DEFAULT_GO_CLI" ]; then
    MY_PROJECT_GO_CLI=$GO_CLI
    MY_PROJECT_GO_CLI_DETECTED=true
    MY_PROJECT_GO_VERSION=default
fi



## Print build information
echo "[$SYSTEM_PLATFORM] Project Information: GO version: $MY_PROJECT_GO_VERSION"



## Detect GO SDK
echo "[$SYSTEM_PLATFORM] Detecting GO SDK ..."
$GO_CLI version 1>/dev/null 2>&1
MY_CHECK_RESULT=$?
if [ $MY_CHECK_RESULT -ne 0 ] ; then
    echo "[$SYSTEM_PLATFORM] Detecting GO SDK ... NOT FOUND"
    exit 1
fi
echo "[$SYSTEM_PLATFORM] Detecting GO SDK ... FOUND"



## Detect GOBIN
if [ "$MY_PROJECT_GO_BIN_DETECTED" = "false" ] ; then
    echo "[$SYSTEM_PLATFORM] Detecting GOBIN ..."
    MY_GOBIN=$($GO_CLI env GOBIN)
    MY_CHECK_RESULT=$?
    if [ $MY_CHECK_RESULT -ne 0 ] ; then
        echo "[$SYSTEM_PLATFORM] Detecting GOBIN ... FAILED (ExitCode: $MY_CHECK_RESULT)"
        MY_PROJECT_GO_BIN_DETECTED=false
    else
        if [ -z "$MY_GOBIN" ] ; then
            echo "[$SYSTEM_PLATFORM] Detecting GOBIN ... NOT FOUND"
            MY_PROJECT_GO_BIN_DETECTED=false
        else
            echo "[$SYSTEM_PLATFORM] Detecting GOBIN ... FOUND ( $MY_GOBIN )"
            MY_PROJECT_GO_BIN=$MY_GOBIN
            MY_PROJECT_GO_BIN_DETECTED=true
        fi
    fi
fi



## Detect GOPATH
if [ "$MY_PROJECT_GO_BIN_DETECTED" = "false" ] ; then
    echo "[$SYSTEM_PLATFORM] Detecting GOPATH ..."
    MY_GOPATH=$($GO_CLI env GOPATH)
    MY_CHECK_RESULT=$?
    if [ $MY_CHECK_RESULT -ne 0 ] ; then
        echo "[$SYSTEM_PLATFORM] Detecting GOPATH ... FAILED (ExitCode: $MY_CHECK_RESULT)"
        MY_PROJECT_GO_BIN_DETECTED=false
    else
        if [ -z "$MY_GOPATH" ] ; then
            echo "[$SYSTEM_PLATFORM] Detecting GOPATH ... NOT FOUND"
            MY_PROJECT_GO_BIN_DETECTED=false
        else
            echo "[$SYSTEM_PLATFORM] Detecting GOPATH ... FOUND ( $MY_GOPATH )"
            MY_PROJECT_GO_BIN="$MY_GOPATH/bin"
            MY_PROJECT_GO_BIN_DETECTED=true
        fi
    fi
fi
if [ "$MY_PROJECT_GO_BIN_DETECTED" = "false" ] ; then
    echo "[$SYSTEM_PLATFORM] Detecting GOBIN or GOPATH ... NOT FOUND"
    exit 1
fi



## Add GOBIN or GOPATH/bin into PATH
echo "[$SYSTEM_PLATFORM] Adding GOBIN or GOPATH/bin into PATH ..."
if [[ $PATH == *$MY_PROJECT_GO_BIN* ]]; then
    echo "[$SYSTEM_PLATFORM] Adding GOBIN or GOPATH/bin into PATH ... SKIPPED"
else
    PATH=$MY_PROJECT_GO_BIN:$PATH
    echo "[$SYSTEM_PLATFORM] Adding GOBIN or GOPATH/bin into PATH ... DONE"
fi



## Detect project GO SDK
if [ "$MY_PROJECT_GO_CLI_DETECTED" = "false" ] ; then
    echo "[$SYSTEM_PLATFORM] Detecting project GO SDK ..."
    $MY_PROJECT_GO_CLI version 1>/dev/null 2>&1
    MY_CHECK_RESULT=$?
    if [ $MY_CHECK_RESULT -ne 0 ] ; then
        echo "[$SYSTEM_PLATFORM] Detecting project GO SDK ... NOT FOUND"
        MY_PROJECT_GO_CLI_DETECTED=false
    else
        echo "[$SYSTEM_PLATFORM] Detecting project GO SDK ... FOUND"
        MY_PROJECT_GO_CLI_DETECTED=true
    fi
fi



## Restore project GO SDK
if [ "$MY_PROJECT_GO_CLI_DETECTED" = "false" ] ; then
    echo "[$SYSTEM_PLATFORM] Restoring project GO SDK ..."

    # Create SDK Downloader
    echo "[$SYSTEM_PLATFORM] Restoring project GO SDK ... Creating SDK downloader"
    $GO_CLI install golang.org/dl/go$PROJECT_GO_VERSION@latest
    MY_CHECK_RESULT=$?
    if [ $MY_CHECK_RESULT -ne 0 ] ; then
        echo "[$SYSTEM_PLATFORM] Restoring project GO SDK ... Creating SDK downloader ... FAILED"
        exit 1
    fi
    echo "[$SYSTEM_PLATFORM] Restoring project GO SDK ... Creating SDK downloader ... DONE"

    # Download SDK
    echo "[$SYSTEM_PLATFORM] Restoring project GO SDK ... Downloading SDK"
    $MY_PROJECT_GO_CLI download
    MY_CHECK_RESULT=$?
    if [ $MY_CHECK_RESULT -ne 0 ] ; then
        echo "[$SYSTEM_PLATFORM] Restoring project GO SDK ... Downloading SDK ... FAILED"
        exit 1
    fi
    echo "[$SYSTEM_PLATFORM] Restoring project GO SDK ... Downloading SDK ... DONE"
fi



## Build project
echo "[$SYSTEM_PLATFORM] Building project ..."
cd -- "$PROJECT_DIR" || exit 1

## Build project - Fetch project dependencies
echo "[$SYSTEM_PLATFORM] Building project ... Fetching project dependencies ..."
$MY_PROJECT_GO_CLI mod tidy
MY_CHECK_RESULT=$?
if [ $MY_CHECK_RESULT -ne 0 ] ; then
    echo "[$SYSTEM_PLATFORM] Building project ... Fetching project dependencies ... FAILED (ExitCode: $MY_CHECK_RESULT)"
    exit 1
fi
echo "[$SYSTEM_PLATFORM] Building project ... Fetching project dependencies ... DONE"

## Build project - Run go test
echo "[$SYSTEM_PLATFORM] Building project ... Running go test ..."
$MY_PROJECT_GO_CLI test -short -cover ./...
MY_CHECK_RESULT=$?
if [ $MY_CHECK_RESULT -ne 0 ] ; then
    echo "[$SYSTEM_PLATFORM] Building project ... Running go test ... FAILED (ExitCode: $MY_CHECK_RESULT)"
    exit 1
fi
echo "[$SYSTEM_PLATFORM] Building project ... Running go test ... DONE"

echo "[$SYSTEM_PLATFORM] Building project ... DONE"
