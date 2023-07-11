FROM node:18
WORKDIR /frontend
COPY ./package*.json .
RUN npm install
COPY . .
EXPOSE 8888
CMD [ "npm", "install" ]
CMD [ "npm", "run", "serve" ]
