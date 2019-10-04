FROM  postgres:12-alpine
LABEL maintainer="FG Web (fg-web@samfundet.no)"

RUN apk update && apk upgrade && \
      apk add --no-cache bash git openssh

RUN mkdir -p /postgres/sql
COPY postgres/sql/*.pgsql /postgres/sql/

