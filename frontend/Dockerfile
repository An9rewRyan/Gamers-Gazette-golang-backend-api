FROM node:latest AS node_builder
WORKDIR "/frontend"
RUN ls
ADD frontend .
RUN ls
RUN npm install
EXPOSE 3000
CMD npm start