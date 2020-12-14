FROM node:15.3.0-alpine3.12
EXPOSE 4000
ENV TODO_SERVICE_URL       ''

WORKDIR /app
ADD /package.json /app/
ADD /package-lock.json /app/
RUN npm install

ADD /index.js /app/
ADD /startup.sh /app/
ADD /schema.graphql /app/

CMD [ "/app/startup.sh" ]

