# Base python package
FROM python:3.8-slim-buster

# Working directory
WORKDIR /flask-api

# Copy the dependencies
COPY requirements.txt .

# Install the dependencies
RUN pip3 install -r requirements.txt

# Copy the files
COPY . .

# Open frontend port
EXPOSE 5454

# Executable commands
CMD exec python3 tazmouApp.py