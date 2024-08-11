# Specify the base image with Node.js
FROM node:16-alpine as build-stage

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json (if it exists)
COPY ../resources/package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the project files
COPY ../resources .

# Build the Vue.js application
RUN npm run build

# Use Nginx for serving static files
FROM nginx:alpine as production-stage

# Copy the built application to the Nginx root directory
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
