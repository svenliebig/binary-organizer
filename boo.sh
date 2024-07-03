#!/bin/bash

# add this alias to your .bashrc or .bash_profile
# alias boo='source /path/to/boo.sh'

VERSION="0.0.1"
EXECUTABLE="boo-$VERSION"

SCRIPTPATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

# check if ./boo exists
if [ -f "$SCRIPTPATH/$EXECUTABLE" ]; then
else
	# change directory to use the go.mod
	cd $SCRIPTPATH

	# build the binary once
	go build -o "$EXECUTABLE" "$SCRIPTPATH/main.go"

	# change back to the original directory
	cd -
fi

"$SCRIPTPATH/$EXECUTABLE" "$@"

