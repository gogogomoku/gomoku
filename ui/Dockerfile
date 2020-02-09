FROM node:lts-alpine

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

# install project dependencies
# TODO: Install without devDeps
# (currently fails in naive environment without them)
RUN npm install

ENV NODE_ENV=production

ARG VUE_APP_SERVER_HTTP=http://localhost:4243
ENV VUE_APP_SERVER_HTTP=${VUE_APP_SERVER_HTTP}

# build app for production with minification
# TODO: Actually make it kind of mini...
RUN npm run build

EXPOSE 8080
CMD ["npm", "run", "dist"]