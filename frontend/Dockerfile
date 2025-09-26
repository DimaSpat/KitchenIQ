FROM oven/bun:latest AS builder

WORKDIR /app

COPY package.json bun.lock ./

RUN bun install

COPY . .


RUN bun run build

FROM oven/bun:1.2.22-slim

WORKDIR /app

COPY --from=builder /app/dist ./dist
COPY --from=builder /app/server ./server
COPY --from=builder /app/package.json ./

RUN bun install --production

CMD [ "bun", "run", "serve" ]