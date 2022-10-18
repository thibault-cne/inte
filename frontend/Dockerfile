#Base image
FROM node:lts-alpine

#Install serve package
RUN npm install -g serve

# Set the working directory
WORKDIR /vue-app

# Copy the package.json and package-lock.json
COPY package*.json ./

# install project dependencies
RUN npm install

# Copy the project files
COPY . .

# Build the project
RUN npm run build

# Expose frontend ports
EXPOSE 3000

# Executables
CMD [ "serve", "-s", "dist" ]