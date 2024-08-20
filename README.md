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
sudo -E ./keymulate
```

## Windows Support

Windows support is currently a work in progress. Stay tuned for updates!

## Contributing

We welcome contributions! Please feel free to submit a Pull Request.
