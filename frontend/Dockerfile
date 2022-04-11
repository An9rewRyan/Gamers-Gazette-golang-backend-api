FROM node:latest AS node_builder
WORKDIR /frontend
RUN ls
ADD . .
RUN ls
RUN npm install
RUN npm run build
RUN npm install -g serve
CMD serve -s build