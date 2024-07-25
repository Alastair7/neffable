## NEFFABLE

Neffable is an app made for soulmates. Connect with your special one and send him/her how you feel with a notification.

### REQUIREMENTS

- Android Studio
- Docker.
- VS Code or your favorite code editor.
- Golang >= 1.21.3

### SETUP BACKEND SERVER

1. Clone this repo.
2. Go to the repo root folder >> ./neffable.
3. Run `docker-compose up -d`
4. Try it by going to http:localhost:5001/api/test/ping.

### HOW TO TEST THE SERVER LOCALLY

1. Run the docker compose container.
2. Stop the app container.
3. Go to ./neffable/backend
4. `go run main.go`
