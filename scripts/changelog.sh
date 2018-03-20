#!/bin/bash

REPOSITORY=http://github.com/oskarszura/smarthome
NEWER_TAG=$(git tag | tail -r | sed -n 1p)
OLDER_TAG=$(git tag | tail -r | sed -n 2p)
HEADER="# Changelog from $NEWER_TAG"
DIR=./docs/changelogs

if [ ! -d $DIR ]
    then
    mkdir -p $DIR
fi

LOGFILE=$DIR/CHANGELOG_$NEWER_TAG.md

echo $HEADER > $LOGFILE
echo "### Commits" >> $LOGFILE

if [ -z "$OLDER_TAG" ]
then
    git log $NEWER_TAG --pretty=format:'* [[`%h`]('$REPOSITORY'/commit/%H)] - %s (%an)' >> $LOGFILE
else
    git log $OLDER_TAG..$NEWER_TAG --pretty=format:'* [[`%h`]('$REPOSITORY'/commit/%H)] - %s (%an)' >> $LOGFILE
fi
