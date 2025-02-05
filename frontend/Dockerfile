# frontend/Dockerfile
FROM node:22.13.1 AS frontend-builder

# Update npm globally
RUN npm install -g npm@11.1.0

RUN npm install -g @angular/cli

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json to leverage Docker layer caching
COPY package.json package-lock.json ./

# Install dependencies (force installation if there are conflicts)
RUN npm install --legacy-peer-deps

# Copy the rest of the application code
COPY . .

# Expose the port (optional, since Docker Compose will handle this)
EXPOSE 4200

CMD ["ng", "serve", "--host", "0.0.0.0"]