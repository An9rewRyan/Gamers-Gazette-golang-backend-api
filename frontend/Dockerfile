FROM node:latest AS node_builder
WORKDIR "/frontend"
ADD frontend .
RUN ls
RUN npm install
EXPOSE 3000
CMD npm start