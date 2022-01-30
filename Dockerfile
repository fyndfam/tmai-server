###############
# stage 1
###############
FROM node:16-alpine3.14 AS BUILD_IMG

WORKDIR /usr/src/app

COPY package.json yarn.lock ./

RUN yarn global add @nestjs/cli

RUN yarn install --prod && yarn cache clean --all

COPY . .
RUN yarn build

###############
# stage 2
###############
FROM node:16-alpine3.14

RUN mkdir /home/node/app && chown -R node:node /home/node/app

WORKDIR /home/node/app

COPY --chown=node:node --from=BUILD_IMG /usr/src/app/dist ./dist
COPY --chown=node:node --from=BUILD_IMG /usr/src/app/node_modules ./node_modules

USER node

EXPOSE 3000

CMD ["node", "./dist/src/main.js"]
