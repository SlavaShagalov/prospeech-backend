FROM node:19 AS build
ARG FRONTEND_CODE_PATH="./frontend"
WORKDIR /home/code
#RUN adduser node sudo
COPY $FRONTEND_CODE_PATH .
RUN npm install
RUN npm run build

FROM nginx:1.22.1
COPY --from=build /home/code/build /var/www/build
CMD ["nginx", "-g", "daemon off;"]
