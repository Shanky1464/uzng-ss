# Step 1: Use an official Node.js runtime as the base image
FROM node:18-alpine AS build

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy package.json and package-lock.json for npm install
COPY package.json package-lock.json ./

# Step 4: Install dependencies
RUN npm install

# Step 5: Copy the entire Vue.js project into the container
COPY . .

# Step 6: Build the Vue.js project for production
RUN npm run build

# Step 7: Use Nginx to serve the built files
FROM nginx:alpine

# Step 8: Copy the build output from the previous stage to the Nginx directory
COPY --from=build /app/dist /usr/share/nginx/html

# Step 9: Expose port 80 for accessing the frontend app
EXPOSE 80

# Step 10: Start Nginx when the container starts
CMD ["nginx", "-g", "daemon off;"]
