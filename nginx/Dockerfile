FROM bitnami/minideb:latest

RUN apt-get update && apt-get -y install --no-install-recommends gnupg ca-certificates apt-transport-https curl libnginx-mod-http-lua gettext-base && rm -rf /var/lib/apt/lists/*

RUN curl -kfsSL https://openresty.org/package/pubkey.gpg | apt-key add -
RUN echo "deb http://openresty.org/package/debian buster openresty" \
    | tee /etc/apt/sources.list.d/openresty.list

RUN apt-get update && apt-get -y install openresty && rm -rf /var/lib/apt/lists/*
ENV PATH="/usr/local/openresty/nginx/sbin:${PATH}"

COPY ./conf/nginx.conf.template /
COPY ./lua/* /usr/local/openresty/lualib

CMD ["/bin/sh" , "-c" , "envsubst < /nginx.conf.template > /usr/local/openresty/nginx/conf/nginx.conf && exec nginx -g 'daemon off;'"]
