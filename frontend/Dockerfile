# --- Stage 1: Build the Vue application ---
FROM node:24-alpine AS builder

WORKDIR /frontend

# Install dependencies based on package lock files
COPY package*.json ./
RUN npm install

# Copy source code and build for production
COPY . .
RUN npm run build

# --- Stage 2: Serve with Nginx ---
FROM nginx:alpine

# Copy the built static assets from the builder stage to Nginx's public folder
COPY --from=builder /frontend/dist /usr/share/nginx/html

# Expose port 80 inside the container
EXPOSE 80

# Start Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
