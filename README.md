# dump1090 forwarder

## Build

```fish
env GOOS=linux GOARCH=arm GOARM=7 go build
```

## Install

### Executable

Place compiled executable `dump1090-forward` to `/home/pi/dump1090-forward`.

### Systemd

Make sure you are running `dump1090-mutability` in your system.

Create systemd unit file `/etc/systemd/system/dump1090-forward
`:

```systemd
[Unit]
Description=Read dump1090 port 30005 and forward it to another dump1090 port 30004
Requires=dump1090-mutability.service

[Service]
Type=simple
ExecStart=/home/pi/dump1090-forward
Restart=always
User=pi

[Install]
Alias=dump1090-forward
```

Reload and start:

```sh
sudo systemctl daemon-reload
sudo systemctl enable dump1090-forward
sudo systemctl start dump1090-forward
```

Check logs:

```sh
journalctl -u dump1090-forward.service
```
