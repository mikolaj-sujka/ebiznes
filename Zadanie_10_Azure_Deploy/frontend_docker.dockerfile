# Stage 1: Build the React app
FROM node:14 AS build
WORKDIR /app
COPY ../Zadanie_5_Frontend .  
RUN npm install && npm run build

# Stage 2: Serve the app with Nginx
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
