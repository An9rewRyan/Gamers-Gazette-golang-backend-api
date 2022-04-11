FROM node:latest AS node_builder
RUN ls
WORKDIR /frontend
RUN ls
ADD . .
RUN ls
RUN npm install
# RUN npm run build
CMD npm start