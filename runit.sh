#!/bin/sh

APP_SERVER_PORT=4243
APP_SERVER=$(dig @resolver1.opendns.com ANY myip.opendns.com +short -4)
# APP_SERVER=localhost # Uncomment to test
VUE_APP_PORT=8080
VUE_APP_SERVER_HTTP=http://"$APP_SERVER":"$APP_SERVER_PORT"

cat << EOF > gomoku-env.list
APP_SERVER_PORT=$APP_SERVER_PORT
APP_SERVER=$APP_SERVER
VUE_APP_PORT=$VUE_APP_PORT
EOF

# Build the compose file, and pass this as a build variable to ui
VUE_APP_SERVER_HTTP=$VUE_APP_SERVER_HTTP docker-compose up --build