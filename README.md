# Keymulate

> Where typing meets ASMR.

Keymulate is an innovative keyboard sound simulator that transforms your silent typing experience into a rich, auditory adventure. Whether you're nostalgic for the clicky sounds of mechanical keyboards or just want to add some fun to your typing, Keymulate has got you covered.

## Features

- Custom sound effects for each key press and release
- Support for multiple keyboard sound profiles (e.g., MX Brown, Blue switches)
- Low-latency audio playback for a responsive typing experience
- Cross-platform support (Linux fully supported, Windows support coming soon)

## Building and Running on Linux

### Prerequisites

- Go 1.16 or later
- ALSA development libraries

On Ubuntu or Debian-based systems, you can install ALSA development libraries with:

```bash
sudo apt-get install libasound2-dev
```

### Building

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/keymulate.git
   cd keymulate
   ```

2. Build the project:
   ```bash
   go build -o keymulate
   ```

### Running

To run Keymulate, you'll need root privileges to access the input devices:

```bash
sudo ./keymulate
```

## Setting up as a Systemd Service

To have Keymulate start automatically on boot, you can set it up as a systemd service:

1. Create a systemd service file:
   ```bash
   sudo nano /etc/systemd/system/keymulate.service
   ```

2. Add the following content to the file:
   ```
   [Unit]
   Description=Keymulate Keyboard Sound Simulator
   After=network.target

   [Service]
   ExecStart=/path/to/keymulate
   Restart=always
   User=root

   [Install]
   WantedBy=multi-user.target
   ```

   Replace `/path/to/keymulate` with the actual path to your compiled binary.

3. Save the file and exit the editor.

4. Reload systemd to recognize the new service:
   ```bash
   sudo systemctl daemon-reload
   ```

5. Enable the service to start on boot:
   ```bash
   sudo systemctl enable keymulate.service
   ```

6. Start the service:
   ```bash
   sudo systemctl start keymulate.service
   ```

You can check the status of the service at any time with:
```bash
sudo systemctl status keymulate.service
```

## Windows Support

Windows support is currently a work in progress. Stay tuned for updates!

## Contributing

We welcome contributions! Please feel free to submit a Pull Request.
