#!/bin/bash

# Variables
APP_NAME="commander"
INSTALL_LOCATION="/usr/local/bin"
PLIST_FILE="/Library/LaunchDaemons/com.example.commander.plist"

# Step 1: Compile the Go application
go build -o $APP_NAME main.go

# Step 2: Move the binary to /usr/local/bin
sudo mv $APP_NAME $INSTALL_LOCATION/

# Step 3: Create a LaunchDaemon plist
sudo tee $PLIST_FILE > /dev/null <<EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.example.commander</string>

    <key>ProgramArguments</key>
    <array>
        <string>${INSTALL_LOCATION}/${APP_NAME}</string>
    </array>

    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>
EOF

# Step 4: Set permissions and load the daemon
# sudo chmod 644 $PLIST_FILE
# sudo launchctl load $PLIST_FILE

echo "Installation complete. Commander will run on system boot."
