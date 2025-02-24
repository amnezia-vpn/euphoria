# Go Implementation of Euphoria

AmneziaWG is a contemporary version of the WireGuard protocol. It's a fork of WireGuard-Go and offers protection against detection by Deep Packet Inspection (DPI) systems. At the same time, it retains the simplified architecture and high performance of the original.

The precursor, WireGuard, is known for its efficiency but had issues with detection due to its distinctive packet signatures.
AmneziaWG addresses this problem by employing advanced obfuscation methods, allowing its traffic to blend seamlessly with regular internet traffic.
As a result, AmneziaWG maintains high performance while adding an extra layer of stealth, making it a superb choice for those seeking a fast and discreet VPN connection.

## Usage

Simply run:

```
$ amneziawg-go wg1
```

This will create an interface and fork into the background. To remove the interface, use the usual `ip link del wg0`, or if your system does not support removing interfaces directly, you may instead remove the control socket via `rm -f /var/run/amneziawg-go/wg0.sock`, which will result in euphoria shutting down.

To run euphoria without forking to the background, pass `-f` or `--foreground`:

```
$ amneziawg-go -f wg0
```
When an interface is running, you may use [`euphoria-tools`](https://github.com/amnezia-vpn/euphoria-tools) to configure it, as well as the usual `ip(8)` and `ifconfig(8)` commands.

To run with more logging you may set the environment variable `LOG_LEVEL=debug`.

## Platforms

### Linux

This will run on Linux; you should run amnezia-wg instead of using default linux kernel module.

### Windows

This runs on Windows, you should use it from [euphoria-windows](https://github.com/amnezia-vpn/euphoria-windows), which uses this as a module.


## Building

This requires an installation of the latest version of [Go](https://go.dev/).

```
$ git clone https://github.com/amnezia-vpn/euphoria
$ cd euphoria
$ make
```
