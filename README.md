# 0xBlade

### Terminal based melee game built in Go with no external dependencies

## Building from source
```bash
git clone https://github.com/a3l6/0xBlade.git
cd 0xBlade
make compile VERSION=v1.0 # enter your version code here
```
The output will be in `bin/VERSION/`.

## Useful Links

ANSI Escape Codes : [here](https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797)

## Weird Issues I've run into

1. For some reason in gameManager.drawScreen() it was throwing a nil pointer dereference. But changing it from a `for i=0;i<len;i++` loop to a `for _, idx := range()` and dereferencing the val pointer worked.
2. Need to make udev rule to access inputs. Run these commands to make udev. `sudo nano /etc/udev/rules.d/99-input.rules` then paste this `KERNEL=="event*", SUBSYSTEM=="input", GROUP="input", MODE="0660"` but replace `GROUP="input"` with `GROUP=<USERNAME>`. Then run `sudo udevadm control --reload-rules
&& sudo udevadm trigger`

## Game State
