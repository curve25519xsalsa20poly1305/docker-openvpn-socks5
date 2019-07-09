# OpenVPN Docker Tunnel to SOCKS5 Server

Convers OpenVPN connection to SOCKS5 server in Docker. This allows you to have multiple OpenVPN to SOCKS5 proxies in different containers and expose to different host ports.

Supports latest Docker for both Windows, Linux, and MacOS.

### Related Projects

* [openvpn-tunnel](https://hub.docker.com/r/curve25519xsalsa20poly1305/openvpn-tunnel/) ([GitHub](https://github.com/curve25519xsalsa20poly1305/docker-openvpn-tunnel)) - Wraps your program with OpenVPN network tunnel fully contained in Docker. It's the base image of this project.
* [openvpn-socks5](https://hub.docker.com/r/curve25519xsalsa20poly1305/openvpn-socks5/) ([GitHub](https://github.com/curve25519xsalsa20poly1305/docker-openvpn-socks5)) - This project.
* [shadowsocksr-tunnel](https://hub.docker.com/r/curve25519xsalsa20poly1305/shadowsocksr-tunnel/) ([GitHub](https://github.com/curve25519xsalsa20poly1305/docker-shadowsocksr-tunnel)) - Wraps your program with ShadowsocksR network tunnel fully contained in Docker. Also exposes SOCKS5 server to host machine.
* [shadowsocksr-aria2](https://hub.docker.com/r/curve25519xsalsa20poly1305/shadowsocksr-aria2/) ([GitHub](https://github.com/curve25519xsalsa20poly1305/docker-shadowsocksr-aria2)) - Extends `shadowsocksr-tunnel` with `aria2` support.

## What it does?

1. It reads in an OpenVPN configuration file (`.ovpn`) from a mounted file, specified through `OPENVPN_CONFIG` environment variable.
2. It starts the OpenVPN client program to establish the VPN connection.
3. It starts the SOCKS5 server and listen on container-scoped port 1080 on default. SOCKS5 authentication can be enabled with `SOCKS5_USER` and `SOCKS5_PASS` environment variables. `SOCKS5_PORT` can be used to change the default port.
4. It optionally executes the user specified CMD line from `docker run` positional arguments ([see Docker doc](https://docs.docker.com/engine/reference/run/#cmd-default-command-or-options)). The program will use the VPN connection inside the container.
5. If user has provided CMD line, and `DAEMON_MODE` environment variable is not set to `true`, then after running the CMD line, it will shutdown the OpenVPN client and terminate the container.

## How to use?

Since it extends the [openvpn-tunnel](https://hub.docker.com/r/curve25519xsalsa20poly1305/openvpn-tunnel/) ([GitHub](https://github.com/curve25519xsalsa20poly1305/docker-openvpn-tunnel)) image, everything in `openvpn-tunnel`'s document are still applicable on this image. Therefore, here we only provide an example that starts a daemon server that connects to VPN defined in file `./vpn.ovpn` on host machine, and expose the SOCKS5 to host-scoped port at `7777`.

```bash
# Unix
SET NAME="mysocks5"
PORT="7777"
USER="myuser"
PASS="mypass"
docker run --name "${NAME}" -dit --rm --device=/dev/net/tun --cap-add=NET_ADMIN \
    -v "${PWD}":/vpn:ro -e OPENVPN_CONFIG=/vpn/vpn.ovpn \
    -p "${PORT}":1080 \
    -e SOCKS5_USER="${USER}" \
    -e SOCKS5_PASS="${PASS}" \
    curve25519xsalsa20poly1305/openvpn-socks5

# Windows
SET NAME="mysocks5"
SET PORT="7777"
SET USER="myuser"
SET PASS="mypass"
docker run --name "%NAME%" -dit --rm --device=/dev/net/tun --cap-add=NET_ADMIN ^
    -v "%CD%":/vpn:ro -e OPENVPN_CONFIG=/vpn/vpn.ovpn ^
    -p "%PORT%":1080 ^
    -e SOCKS5_USER="%USER%" ^
    -e SOCKS5_PASS="%PASS%" ^
    curve25519xsalsa20poly1305/openvpn-socks5
```

Then on your host machine test it with curl:

```bash
# Unix & Windows
curl ifconfig.co/json -x socks5h://myuser:mypass@127.0.0.1:7777
```

To stop the daemon, run this:

```bash
# Unix
NAME="mysocks5"
docker stop "${NAME}"

# Windows
SET NAME="mysocks5"
docker stop "%NAME%"
```

## Contributing

Please feel free to contribute to this project. But before you do so, just make
sure you understand the following:

1\. Make sure you have access to the official repository of this project where
the maintainer is actively pushing changes. So that all effective changes can go
into the official release pipeline.

2\. Make sure your editor has [EditorConfig](https://editorconfig.org/) plugin
installed and enabled. It's used to unify code formatting style.

3\. Use [Conventional Commits 1.0.0-beta.2](https://conventionalcommits.org/) to
format Git commit messages.

4\. Use [Gitflow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow)
as Git workflow guideline.

5\. Use [Semantic Versioning 2.0.0](https://semver.org/) to tag release
versions.

## License

Copyright © 2019 curve25519xsalsa20poly1305 &lt;<curve25519xsalsa20poly1305@gmail.com>&gt;

This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the COPYING file for more details.
