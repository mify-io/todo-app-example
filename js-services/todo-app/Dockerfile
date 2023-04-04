FROM node:lts-alpine as builder

WORKDIR /app
ARG TODO_BACKEND_CLIENT_ENDPOINT
ENV TODO_BACKEND_CLIENT_ENDPOINT=${TODO_BACKEND_CLIENT_ENDPOINT}

COPY . .

WORKDIR /app/todo-app

RUN yarn install \
  --prefer-offline \
  --frozen-lockfile \
  --non-interactive \
  --production=false

RUN yarn build

RUN rm -rf node_modules && \
  NODE_ENV=production yarn install \
  --prefer-offline \
  --pure-lockfile \
  --non-interactive \
  --production=true

RUN npm prune --production

FROM node:lts-alpine

WORKDIR /app

COPY --from=builder /app .

WORKDIR /app/todo-app

ENV HOST 0.0.0.0
ENV PORT 80
EXPOSE 80

CMD [ "yarn", "start" ]
