#!/usr/bin/env bash

KOKKA_HOME="/apps/kokka"
cd $KOKKA_HOME

echo "$KOKKA_HOME"
echo "Fixing session file permissions..."

# Fix sessionfile permissions
# chmod 775 $KOKKA_HOME/data/*.datc

echo "Session file permissions fixed!"