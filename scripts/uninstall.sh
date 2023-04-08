#!/bin/bash

# Stop and disable the service
sudo systemctl stop nautilus-print-server.service
sudo systemctl disable nautilus-print-server.service

# Remove the .service file from /etc/systemd/system directory
sudo rm /etc/systemd/system/nautilus-print-server.service

# Remove the executable from /usr/bin/ directory
sudo rm /usr/bin/nautilus-print-server

# Reload systemd configuration
sudo systemctl daemon-reload
