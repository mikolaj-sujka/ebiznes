@echo off
set PLAY_PORT=9000

echo Starting Play application on port %PLAY_PORT%...
start cmd /k sbt "run %PLAY_PORT%"

echo Waiting for Play application to start...
timeout /t 10

echo Starting ngrok...
start cmd /k ngrok http %PLAY_PORT%

echo Script is running. Press any key to exit.
pause
