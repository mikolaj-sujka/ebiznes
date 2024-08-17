# Stage 1: Build the React app
FROM node:14 AS build
WORKDIR /app

# Copy package.json and package-lock.json to install dependencies
COPY ../Zadanie_5_Frontend/package*.json ./ 
RUN npm install  # Install dependencies

# Copy the rest of the app's source code
COPY ../Zadanie_5_Frontend/frontned-app/ ./  

# Run the build script
RUN npm run build  # Build the React app

# Stage 2: Serve the app with Nginx
FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
