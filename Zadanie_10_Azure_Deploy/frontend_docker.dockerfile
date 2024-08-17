# Stage 1: Build the React app
FROM node:14 AS build
WORKDIR /app
COPY ../Zadanie_5_Frontend/package*.json . 
RUN npm install  # Install dependencies
COPY ../Zadanie_5_Frontend .  
RUN npm run build  # Build the React app

# Stage 2: Serve the app with Nginx
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
