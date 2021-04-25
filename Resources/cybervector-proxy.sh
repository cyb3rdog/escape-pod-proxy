#!/bin/bash

# Setup EscapePod config file
sudo sed -i -e '/^\(ENABLE_EXTENSIONS=\).*/{s//\1true/;:a;n;ba;q}' -e '$aENABLE_EXTENSIONS=true' /etc/escape-pod.conf
sudo sed -i -e '/^\(ESCAPEPOD_EXTENDER_TARGET=\).*/{s//\1127.0.0.1:8089/;:a;n;ba;q}' -e '$aESCAPEPOD_EXTENDER_TARGET=127.0.0.1:8089' /etc/escape-pod.conf
sudo sed -i -e '/^\(ESCAPEPOD_EXTENDER_DISABLE_TLS=\).*/{s//\1true/;:a;n;ba;q}' -e '$aESCAPEPOD_EXTENDER_DISABLE_TLS=true' /etc/escape-pod.conf

# Restart Escapepod service
sudo systemctl restart escape_pod

# Setup Cyb3rVector extension proxy attributes
chmod 755 /usr/local/escapepod/bin/cybervector-proxy

# Setup Cyb3rVector extension config file 
sudo mv /usr/local/escapepod/bin/cybervector-proxy.conf /etc/cybervector-proxy.conf

# Setup Cyb3rVector extension proxy service
sudo mv /usr/local/escapepod/bin/cybervector-proxy.service /usr/lib/systemd/system/cybervector-proxy.service
sudo chown root:root /usr/lib/systemd/system/cybervector-proxy.service

# Start Cyb3rVector extension proxy service
sudo systemctl daemon-reload
sudo systemctl enable cybervector-proxy
sudo systemctl start cybervector-proxy
