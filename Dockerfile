FROM node:8

RUN mkdir -p /opt/flight2bq
WORKDIR /opt/flight2bq

ADD package.json /opt/flight2bq/package.json
ADD yarn.lock /opt/flight2bq/yarn.lock
RUN yarn install --production
ADD . /opt/flight2bq

CMD node index.js