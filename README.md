# web-traffic-workstation
A websocket mechanism to display connection related information like start time, end time, number of packets transferred, number of bytes transferred for each connection. This is done via websockets using nodejs as frontend and golang as backend to send data to browser.

![Watch demo](https://github.com/MillionEyes/web-traffic-workstation/blob/main/recording/Screen%20Recording%202022-06-25%20at%207.14.14%20PM.mov)

## Components:
## Frontend
A NodeJS app which opens a socket connection to the backend server.

1. Navigate to frontend directory in terminal
2. Run "npm install"
3. Run "npm start"

## Backend
A Golang server listening for incoming socket connections.

1. Navigate to backend directory in terminal
2. Run "go run sock.go"

## Working
1. Navigate to http://localhost:3000
2. All new incoming transport layer connection information will be present on the dashboard and old connection information will be continuosly updated with new information inplace.


## How can it be better?
1. Use server sent events since this is a one way communication.
2. Check if wireshark already does this in which case this is useless but fun to experiment with.
