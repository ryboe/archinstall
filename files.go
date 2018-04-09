package main

const (
	bbrConf = `
net.core.default_qdisc=fq
net.ipv4.tcp_congestion_control=bbr
`

	kbdrateService = `
[Unit]
Description=Set the keyboard repeat rate for the Linux console

[Service]
Type=oneshot
RemainAfterExit=true
StandardInput=tty
StandardOutput=tty
ExecStart=/usr/bin/kbdrate -s -d 250 -r 30

[Install]
WantedBy=multi-user.target
`

	keymap = `
# us.map
keymaps 0-2,4-6,8-9,12
alt_is_meta
include "qwerty-layout"
include "linux-with-alt-and-altgr"
include "compose.latin1"
include "euro1.map"
strings as usual

keycode   1 = Escape
keycode   2 = one              exclam
keycode   3 = two              at               at               nul              nul
keycode   4 = three            numbersign
	control	keycode   4 = Escape
keycode   5 = four             dollar           dollar           Control_backslash
keycode   6 = five             percent
	control	keycode   6 = Control_bracketright
keycode   7 = six              asciicircum
	control	keycode   7 = Control_asciicircum
keycode   8 = seven            ampersand        braceleft        Control_underscore
keycode   9 = eight            asterisk         bracketleft      Delete
keycode  10 = nine             parenleft        bracketright
keycode  11 = zero             parenright       braceright
keycode  12 = minus            underscore       backslash        Control_underscore Control_underscore
keycode  13 = equal            plus
keycode  14 = Delete
keycode  15 = Tab
	shift	keycode  15 = Meta_Tab
keycode  26 = bracketleft      braceleft
	control	keycode  26 = Escape
keycode  27 = bracketright     braceright       asciitilde       Control_bracketright
keycode  28 = Return
	alt	keycode  28 = Meta_Control_m
keycode  29 = Control
keycode  39 = semicolon        colon
keycode  40 = apostrophe       quotedbl
	control	keycode  40 = Control_g
keycode  41 = grave            asciitilde
	control	keycode  41 = nul
keycode  42 = Shift
keycode  43 = backslash        bar
	control	keycode  43 = Control_backslash
keycode  51 = comma            less
keycode  52 = period           greater
keycode  53 = slash            question
	control keycode  53 = Control_underscore
	control shift keycode  53 = Delete
keycode  54 = Shift
keycode  56 = Alt
keycode  57 = space
	control	keycode  57 = nul
keycode  58 = Escape
keycode  86 = less             greater          bar
keycode  97 = Control
`

	localeConf = `
LANG=en_US.UTF-8
LC_COLLATE=en_US.UTF-8
`

	mkinitcpioConf = `
# vim:set ft=sh
# MODULES
# The following modules are loaded before any boot hooks are
# run.  Advanced users may wish to specify all system modules
# in this array.  For instance:
#     MODULES=(piix ide_disk reiserfs)
MODULES=(i915)

# BINARIES
# This setting includes any additional binaries a given user may
# wish into the CPIO image.  This is run last, so it may be used to
# override the actual binaries included by a given hook
# BINARIES are dependency parsed, so you may safely ignore libraries
BINARIES=()

# FILES
# This setting is similar to BINARIES above, however, files are added
# as-is and are not parsed in any way.  This is useful for config files.
FILES=()

# HOOKS
# This is the most important setting in this file.  The HOOKS control the
# modules and scripts added to the image, and what happens at boot time.
# Order is important, and it is recommended that you do not change the
# order in which HOOKS are added.  Run 'mkinitcpio -H <hook name>' for
# help on a given hook.
# 'base' is _required_ unless you know precisely what you are doing.
# 'udev' is _required_ in order to automatically load modules
# 'filesystems' is _required_ unless you specify your fs modules in MODULES
# Examples:
##   This setup specifies all modules in the MODULES setting above.
##   No raid, lvm2, or encrypted root is needed.
#    HOOKS=(base)
#
##   This setup will autodetect all modules for your system and should
##   work as a sane default
#    HOOKS=(base udev autodetect block filesystems)
#
##   This setup will generate a 'full' image which supports most systems.
##   No autodetection is done.
#    HOOKS=(base udev block filesystems)
#
##   This setup assembles a pata mdadm array with an encrypted root FS.
##   Note: See 'mkinitcpio -H mdadm' for more information on raid devices.
#    HOOKS=(base udev block mdadm encrypt filesystems)
#
##   This setup loads an lvm2 volume group on a usb device.
#    HOOKS=(base udev block lvm2 filesystems)
#
##   NOTE: If you have /usr on a separate partition, you MUST include the
#    usr, fsck and shutdown hooks.
HOOKS=(base udev autodetect modconf keyboard consolefont keymap block encrypt filesystems fsck)

# COMPRESSION
# Use this to compress the initramfs image. By default, gzip compression
# is used. Use 'cat' to create an uncompressed image.
#COMPRESSION="gzip"
#COMPRESSION="bzip2"
#COMPRESSION="lzma"
#COMPRESSION="xz"
#COMPRESSION="lzop"
#COMPRESSION="lz4"

# COMPRESSION_OPTIONS
# Additional options for the compressor
#COMPRESSION_OPTIONS=()
`
	resolvedConf = `
[Resolve]
DNS=2606:4700:4700::1111
FallbackDNS=2606:4700:4700::1001
LLMNR=yes
MulticastDNS=yes
Cache=yes
`

	rngd = `
RNGD_OPTS="--no-tpm=1"
`

	sudoers = `
## sudoers file.
##
## This file MUST be edited with the 'visudo' command as root.
## Failure to use 'visudo' may result in syntax or file permission errors
## that prevent sudo from running.
##
## See the sudoers man page for the details on how to write a sudoers file.
##
##
## Host alias specification
##
## Groups of machines. These may include host names (optionally with wildcards),
## IP addresses, network numbers or netgroups.
# Host_Alias	WEBSERVERS = www1, www2, www3
##
## User alias specification
##
## Groups of users.  These may consist of user names, uids, Unix groups,
## or netgroups.
# User_Alias	ADMINS = millert, dowdy, mikef
##
## Cmnd alias specification
##
## Groups of commands.  Often used to group related commands together.
# Cmnd_Alias	PROCESSES = /usr/bin/nice, /bin/kill, /usr/bin/renice, \
# 			    /usr/bin/pkill, /usr/bin/top
# Cmnd_Alias	REBOOT = /sbin/halt, /sbin/reboot, /sbin/poweroff
##
## Defaults specification
##
## You may wish to keep some of the following environment variables
## when running commands via sudo.
##
## Locale settings
# Defaults env_keep += "LANG LANGUAGE LINGUAS LC_* _XKB_CHARSET"
##
## Run X applications through sudo; HOME is used to find the
## .Xauthority file.  Note that other programs use HOME to find
## configuration files and this may lead to privilege escalation!
# Defaults env_keep += "HOME"
##
## X11 resource path settings
# Defaults env_keep += "XAPPLRESDIR XFILESEARCHPATH XUSERFILESEARCHPATH"
##
## Desktop path settings
# Defaults env_keep += "QTDIR KDEDIR"
##
## Allow sudo-run commands to inherit the callers' ConsoleKit session
# Defaults env_keep += "XDG_SESSION_COOKIE"
##
## Uncomment to enable special input methods.  Care should be taken as
## this may allow users to subvert the command being run via sudo.
# Defaults env_keep += "XMODIFIERS GTK_IM_MODULE QT_IM_MODULE QT_IM_SWITCHER"
##
## Uncomment to use a hard-coded PATH instead of the user's to find commands
# Defaults secure_path="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
##
## Uncomment to send mail if the user does not enter the correct password.
# Defaults mail_badpass
##
## Uncomment to enable logging of a command's output, except for
## sudoreplay and reboot.  Use sudoreplay to play back logged sessions.
# Defaults log_output
# Defaults!/usr/bin/sudoreplay !log_output
# Defaults!/usr/local/bin/sudoreplay !log_output
# Defaults!REBOOT !log_output
##
## Runas alias specification
##
##
## User privilege specification
##
root ALL=(ALL) ALL
## Uncomment to allow members of group wheel to execute any command
# %wheel ALL=(ALL) ALL
## Same thing without a password
%wheel ALL=(ALL) NOPASSWD: ALL
## Uncomment to allow members of group sudo to execute any command
%sudo	ALL=(ALL) ALL
## Uncomment to allow any user to run sudo if they know the password
## of the user they are running the command as (root by default).
# Defaults targetpw  # Ask for the password of the target user
# ALL ALL=(ALL) ALL  # WARNING: only use this together with 'Defaults targetpw'
## Read drop-in files from /etc/sudoers.d
## (the '#' here does not indicate a comment)
#includedir /etc/sudoers.d
`

	vconsoleConf = `
FONT="ter-132n"
KEYMAP="us-caps-esc"
`

	wiredNetwork = `
[Match]
Name=enp0s31f6

[Network]
DHCP=ipv4

[DHCP]
UseDNS=false
UseNTP=false

[IPv6AcceptRA]
UseDNS=false
`

	zshenv = `
ZDOTDIR=${HOME}/.config/zsh/
`
)
