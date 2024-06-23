## NEFFABLE

Neffable is an app made for soulmates. Connect with your special one and send him/her how you feel with a notification.

### REQUIREMENTS

- Android Studio
- Docker.
- VS Code or your favorite code editor.

### SETUP BACKEND SERVER

1. Clone this repo.
2. Go to the repo folder.
3. Build docker image `docker build -f Dockerfile.backend --tag neffable_server .`
4. Run the image `docker run --rm -d --name neffable_api -p 5001:8080 neffable_server`
5. Try it by going to http:localhost:5001/api/test/ping.

If you want to keep the container alive just remove --rm attributes and the container will exist even when it is stopped.
