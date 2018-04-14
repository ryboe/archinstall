# archinstall

This is an Arch Linux installer I made for my desktop computer. It's design borrows heavily from [variadico/xpslinux](https://github.com/variadico/xpslinux). It's made for a machine with these specs:
* Coffee Lake 8700K CPU
* Integrated GPU
* ASUS STRIX Z370-G motherboard
* Intel 900P Optane SSD

I don't recommend using it unless you're me. It might make your computer unbootable.

### Instructions
1. boot into the Arch USB installer
2. run these commands
```sh
# We need more than 512MB of space to install git and go.
mount -o remount,size=2G /run/archiso/cowspace

# Install git and go, requirements for the `go get` command.
pacman -Sy git go

# Fetch and build archinstall on the USB installer
go get github.com/y0ssar1an/archinstall
```
