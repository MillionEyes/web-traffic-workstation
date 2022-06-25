# web-traffic-workstation
A proxy service which shows detailed insights about the user web traffic

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
