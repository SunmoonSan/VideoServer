#! /bin/bash

# Build web UI
cd ~/GolandProjects/src/VideoServer/web
go install
cp ~/GolandProjects/bin/web /Users/san/GolandProjects/bin/video_server_web_ui
cp -R ~/GolandProjects/src/VideoServer/templates /Users/san/GolandProjects/bin/video_server_web_ui