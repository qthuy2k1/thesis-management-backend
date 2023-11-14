FROM node:alpine
WORKDIR /usr/thesis-course-registeration-system-backend
COPY package.json .
RUN npm install\
        && npm install typescript -g
COPY ./dist .
EXPOSE 5000
CMD ["node", "./app.js"]