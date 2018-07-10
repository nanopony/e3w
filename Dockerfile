FROM golang:1.8 as backend
RUN mkdir -p /go/src/github.com/nanopony/web_etcd
ADD . /go/src/github.com/nanopony/web_etcd
WORKDIR /go/src/github.com/nanopony/web_etcd
RUN CGO_ENABLED=0 go build

FROM node:8 as frontend
RUN mkdir /app
ADD static /app
WORKDIR /app
RUN npm --registry=https://registry.npm.taobao.org \
--cache=$HOME/.npm/.cache/cnpm \
--disturl=https://npm.taobao.org/mirrors/node \
--userconfig=$HOME/.cnpmrc install && npm run publish

FROM alpine:latest
RUN mkdir -p /app/static/dist /app/conf
COPY --from=backend /go/src/github.com/nanopony/web_etcd/web_etcd /app
COPY --from=frontend /app/dist /app/static/dist
COPY conf/config.docker.ini /app/conf/config.default.ini
EXPOSE 8080
WORKDIR /app
CMD ["./web_etcd"]
