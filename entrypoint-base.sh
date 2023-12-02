#!/bin/sh

# Start the news service
./news_service &

# Start the post_story service
./post_story_service &

# Start KrakenD
krakend run -c /etc/krakend/krakend.json

