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

# we need to somehow determine here if we should source the `.path` or not, probably through exit codes or something
"$SCRIPTPATH/$EXECUTABLE" "$@"

EXIT_CODE=$?

if [ $EXIT_CODE -eq 20 ]; then
	if [ -f "$SCRIPTPATH/.path" ]; then
	. "$SCRIPTPATH/.path"
	fi

	if [ -f "$SCRIPTPATH/.env" ]; then
	. "$SCRIPTPATH/.env"
	fi
elif [ $EXIT_CODE -eq 21 ]; then
	# Special Handlers maybe
else
fi
