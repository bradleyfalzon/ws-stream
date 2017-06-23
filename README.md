# Introduction

Quick hack to demonstrate streaming some binary data from client to server over websockets.

Not test very well at all.

Install:

```
go get github.com/bradleyfalzon/ws-stream
```

Server listen for websocket connection and send to stdout, where ffmpeg receives and writes to a file. This could also
be ffplay to play the audio locally.

```
ws-stream -listen 127.0.0.1:3000 | ffmpeg -i pipe:0 out.spx
```

Client starts ffmpeg and generated 10 seconds of low volume audio sending to websocket URL.

```
ffmpeg -t 10 -re -f lavfi -i "sine=frequency=1000" -f spx - | ws-stream -url ws://localhost:3000/ws
```

Client and server doesn't automatically stop, so you'll need to kill it yourself.
