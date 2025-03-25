FROM node:18-alpine AS base

FROM base AS installer
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
FROM base AS builder

WORKDIR /app
COPY --from=installer ./app/node_modules ./node_modules
COPY . .
RUN npm run build


FROM base AS runner
WORKDIR /app

COPY --from=builder /app/.output ./

EXPOSE 3000

CMD ["node", "server/index.mjs"]







