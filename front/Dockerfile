# Use a Node.js runtime as the base image
FROM node:latest

# Create and set the working directory in the container
WORKDIR /stg

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files to the working directory
COPY . .

# Build the Svelte application
RUN npm run build

# Expose the port that the application will run on (change if necessary)
EXPOSE 50003

# Start the application
ENV PORT 50003
CMD ["node", "build"]

#CMD ["tail", "-f", "/dev/null"]
