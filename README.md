# Redir to ALA

Abuse the ALA to easily download old packages.

This webserver behaves like an Arch Linux mirror, but it will actually redirect requests to https://archive.archlinux.org/.

## How to use

Run the server somewhere. In this example I'll just assume you run it locally on 8080.

Add this to `/etc/pacman.d/mirrorlist` (NOT AT THE TOP):
```
Server = http://localhost:8080/$repo/os/$arch
```
This should never be your first mirror. It's slow and it will abuse the underlying server which is meant for occasional archival downloads, not normal package serving.

## Why?

When a package is updated, old versions of the package are promptly removed from mirrors. So if you want to install the old version of the package, normally, you can't.

Why would you want to install the old version? The old version may be the latest in your local copy of the package repo database. To install the new version, you must first do `pacman -Sy`, which you should never do without doing `pacman -Syu`. And maybe you don't want to do a whole-ass system upgrade right now, and possibly need to restart your computer.

I [previously wrote about this dilemma](https://notes.cg505.com/arch-kernel-update-reboot/).

## This is cursed.

Sorry. Please don't use this as the first URL in your mirrorlist.
