FROM node:lts-alpine

WORKDIR /app

#RUN npm install -g http-server
RUN npm install -g vite

COPY frontend/dashboard/package*.json ./

RUN npm install 

COPY . .

CMD ["npm", "run", "dev"]
