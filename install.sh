#!/bin/bash

APP_NAME=goautoextractor
INSTALL_DIR=/usr/local/bin
SYSTEMD_DIR=/etc/systemd/system
USER_NAME=$(whoami)

echo "Building $APP_NAME..."
go build -o "$APP_NAME"

echo "Installing to $INSTALL_DIR..."
sudo mv "$APP_NAME" "$INSTALL_DIR"

echo "Setting up config..."
mkdir -p "$HOME/.config/$APP_NAME"
cp ./config/default_config.json "$HOME/.config/$APP_NAME/config.json"

echo "Installing systemd service..."
sudo cp ./systemd/${APP_NAME}.service "$SYSTEMD_DIR/"
sudo sed -i "s/youruser/$USER_NAME/" "$SYSTEMD_DIR/${APP_NAME}.service"

echo "Reloading and enabling service..."
sudo systemctl daemon-reexec
sudo systemctl enable --now "$APP_NAME"

echo "Done! Service is running."
