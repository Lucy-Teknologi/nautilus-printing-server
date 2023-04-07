#!/bin/bash

# Copy the .service file to /etc/systemd/system directory
sudo cp nautilus-print-server.service /etc/systemd/system/

# Copy the executable to /usr/bin/ directory
sudo cp nautilus-print-server /usr/bin/

# Set execute permission for the executable
sudo chmod +x /usr/bin/nautilus-print-server

# Reload systemd configuration
sudo systemctl daemon-reload

# Enable the service
sudo systemctl enable nautilus-print-server.service

# Start the service
sudo systemctl start nautilus-print-server.service
