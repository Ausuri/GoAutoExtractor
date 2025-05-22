#!/bin/bash
WATCH_DIR="~/watchfolder"

inotifywait -m -r -e create --format '%w%f' "$WATCH_DIR" | while read NEWFILE
do
    echo "New file detected: $NEWFILE"
    /path/to/your/program "$NEWFILE"
done