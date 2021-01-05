FROM node:12.19.0 AS build-env

ADD . /app
WORKDIR /app
RUN npm install --production

FROM gcr.io/distroless/nodejs

COPY --from=build-env /app /app
COPY --from=build-env /app/public/ /public/
WORKDIR /app
EXPOSE 3000
CMD ["Servidor/ServerExpress.js"]