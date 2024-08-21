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

3. Set file capabilities:
    ```bash
    sudo setcap cap_dac_read_search+ep main
    ```

### Running

```bash
./keymulate
```

### Set up as systemd service (thanks to [@ShashwatAgrawal20](https://github.com/ShashwatAgrawal20/))
1. Go to user's systemd directory:
    ```bash
    cd ~/.config/systemd/user/
    ```

2. Create keymulate.service:
    ```bash
    touch keymulate.service
    ```

3. Describe the service:
    ```
    [Unit]
    Description=Keymulate Service
    Requires=pipewire.service pipewire-pulse.service wireplumber.service
    After=pipewire.service pipewire-pulse.service wireplumber.service

    [Service]
    ExecStart=/path/to/keymulate/binary
    WorkingDirectory=/path/to/keymulate/directory
    Restart=on-failure

    [Install]
    WantedBy=default.target
    ```  

## Windows Support

Windows support is currently a work in progress. Stay tuned for updates!

## Contributing

We welcome contributions! Please feel free to submit a Pull Request.
