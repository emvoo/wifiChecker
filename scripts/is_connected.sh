#!/bin/sh

wget --spider http://google.com

if [ $? -eq 0 ]; then
    echo "online"
fi
