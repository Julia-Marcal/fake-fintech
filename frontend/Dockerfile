# Use Node.js 20.11.1 as the base image for building the Angular app
FROM node:20.11.1

# Install Angular CLI globally (use the version specified in the project)
RUN npm install -g @angular/cli@latest

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Expose port 4200
EXPOSE 4200

# Start Angular development server
CMD ["ng", "serve", "--host", "0.0.0.0", "--port", "80"]