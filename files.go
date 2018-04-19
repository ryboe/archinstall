package main

const (
	bbrConf = `
net.core.default_qdisc=fq
net.ipv4.tcp_congestion_control=bbr
`
	fontsAliasesConf = `
<?xml version="1.0"?>
<!DOCTYPE fontconfig SYSTEM "fonts.dtd">
<fontconfig>
<!--
  Mark common families with their generics so we'll get
  something reasonable
-->

<!--
  Serif faces
 -->
	<alias>
		<family>Nazli</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Lotoos</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Mitra</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Ferdosi</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Badr</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Zar</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Titr</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Jadid</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Kochi Mincho</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>AR PL SungtiL GB</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>AR PL Mingti2L Big5</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>ＭＳ 明朝</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>NanumMyeongjo</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>UnBatang</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Baekmuk Batang</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>MgOpen Canonica</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Sazanami Mincho</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>AR PL ZenKai Uni</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>ZYSong18030</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>FreeSerif</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>SimSun</family>
		<default><family>serif</family></default>
	</alias>
<!--
  Sans-serif faces
 -->
	<alias>
		<family>Arshia</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Elham</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Farnaz</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Nasim</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Sina</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Roya</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Koodak</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Terafik</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Kochi Gothic</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>AR PL KaitiM GB</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>AR PL KaitiM Big5</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>ＭＳ ゴシック</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>NanumGothic</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>UnDotum</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Baekmuk Dotum</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>MgOpen Modata</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Sazanami Gothic</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>AR PL ShanHeiSun Uni</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>ZYSong18030</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>FreeSans</family>
		<default><family>sans-serif</family></default>
	</alias>
<!--
  Monospace faces
 -->
	<alias>
		<family>NSimSun</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>ZYSong18030</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>NanumGothicCoding</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>FreeMono</family>
		<default><family>monospace</family></default>
	</alias>

<!--
  Fantasy faces
 -->
	<alias>
		<family>Homa</family>
		<default><family>fantasy</family></default>
	</alias>
	<alias>
		<family>Kamran</family>
		<default><family>fantasy</family></default>
	</alias>
	<alias>
		<family>Fantezi</family>
		<default><family>fantasy</family></default>
	</alias>
	<alias>
		<family>Tabassom</family>
		<default><family>fantasy</family></default>
	</alias>

<!--
  Cursive faces
 -->
	<alias>
		<family>IranNastaliq</family>
		<default><family>cursive</family></default>
	</alias>
	<alias>
		<family>Nafees Nastaleeq</family>
		<default><family>cursive</family></default>
	</alias>

</fontconfig>
<?xml version="1.0"?>
<!DOCTYPE fontconfig SYSTEM "fonts.dtd">
<fontconfig>
<!-- Keep in sync with 60-generic.conf -->

<!-- Emoji -->

	<alias binding="same">
		<family>Emoji Two</family>
		<default><family>emoji</family></default>
	</alias>
	<alias binding="same">
		<family>Emoji One</family>
		<default><family>emoji</family></default>
	</alias>
	<alias binding="same">
		<family>Noto Color Emoji</family> <!-- Google -->
		<default><family>emoji</family></default>
	</alias>
	<alias binding="same">
		<family>Apple Color Emoji</family> <!-- Apple -->
		<default><family>emoji</family></default>
	</alias>
	<alias binding="same">
		<family>Segoe UI Emoji</family> <!-- Microsoft -->
		<default><family>emoji</family></default>
	</alias>
	<alias binding="same">
		<family>Twitter Color Emoji</family> <!-- Twitter -->
		<default><family>emoji</family></default>
	</alias>
	<alias binding="same">
		<family>EmojiOne Mozilla</family> <!-- Mozilla -->
		<default><family>emoji</family></default>
	</alias>
	<!-- B&W -->
	<alias binding="same">
		<family>Noto Emoji</family> <!-- Google -->
		<default><family>emoji</family></default>
	</alias>
	<alias binding="same">
		<family>Android Emoji</family> <!-- Google -->
		<default><family>emoji</family></default>
	</alias>

	<!-- Add language for emoji, to match other emoji fonts. -->
	<match>
		<test name="family">
			<string>emoji</string>
		</test>
		<edit name="lang" mode="prepend">
			<string>und-zsye</string>
		</edit>
	</match>

	<match>
		<test name="lang">
			<string>und-zsye</string>
		</test>
		<test qual="all" name="family" compare="not_eq">
			<string>emoji</string>
		</test>

		<!-- Add generic family. -->
		<edit name="family" mode="append" binding="strong">
			<string>emoji</string>
		</edit>
	</match>


<!-- Math -->

	<!-- https://en.wikipedia.org/wiki/Category:Mathematical_OpenType_typefaces -->
	<alias binding="same">
		<family>XITS Math</family> <!-- Khaled Hosny -->
		<default><family>math</family></default>
	</alias>
	<alias binding="same">
		<family>STIX Two Math</family> <!-- AMS -->
		<default><family>math</family></default>
	</alias>
	<alias binding="same">
		<family>Cambria Math</family> <!-- Microsoft -->
		<default><family>math</family></default>
	</alias>
	<alias binding="same">
		<family>Latin Modern Math</family> <!-- TeX -->
		<default><family>math</family></default>
	</alias>
	<alias binding="same">
		<family>Minion Math</family> <!-- Adobe -->
		<default><family>math</family></default>
	</alias>
	<alias binding="same">
		<family>Lucida Math</family> <!-- Adobe -->
		<default><family>math</family></default>
	</alias>
	<alias binding="same">
		<family>Asana Math</family>
		<default><family>math</family></default>
	</alias>

	<!-- Add language for math, to match other math fonts. -->
	<match>
		<test name="family">
			<string>math</string>
		</test>
		<edit name="lang" mode="prepend">
			<string>und-zmth</string>
		</edit>
	</match>

	<match>
		<test name="lang">
			<string>und-zmth</string>
		</test>
		<test qual="all" name="family" compare="not_eq">
			<string>math</string>
		</test>

		<!-- Add generic family -->
		<edit name="family" mode="append" binding="strong">
			<string>math</string>
		</edit>
	</match>


</fontconfig>
<?xml version="1.0"?>
<!DOCTYPE fontconfig SYSTEM "fonts.dtd">
<fontconfig>
<!--
  Mark common families with their generics so we'll get
  something reasonable
-->

<!--
  Serif faces
 -->
	<alias>
		<family>Bitstream Vera Serif</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Cambria</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Constantia</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>DejaVu Serif</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Elephant</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Garamond</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Georgia</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Liberation Serif</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Luxi Serif</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>MS Serif</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Nimbus Roman No9 L</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Nimbus Roman</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Palatino Linotype</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Thorndale AMT</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Thorndale</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Times New Roman</family>
		<default><family>serif</family></default>
	</alias>
	<alias>
		<family>Times</family>
		<default><family>serif</family></default>
	</alias>
<!--
  Sans-serif faces
 -->
	<alias>
		<family>Albany AMT</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Albany</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Arial Unicode MS</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Arial</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Bitstream Vera Sans</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Britannic</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Calibri</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Candara</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Century Gothic</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Corbel</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>DejaVu Sans</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Helvetica</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Haettenschweiler</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Liberation Sans</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>MS Sans Serif</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Nimbus Sans L</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Nimbus Sans</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Luxi Sans</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Tahoma</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Trebuchet MS</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Twentieth Century</family>
		<default><family>sans-serif</family></default>
	</alias>
	<alias>
		<family>Verdana</family>
		<default><family>sans-serif</family></default>
	</alias>
<!--
  Monospace faces
 -->
	<alias>
		<family>Andale Mono</family>
		<default><family>monospace</family></default>
	</alias>
 	<alias>
		<family>Bitstream Vera Sans Mono</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Consolas</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Courier New</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Courier</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Cumberland AMT</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Cumberland</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>DejaVu Sans Mono</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Fixedsys</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Inconsolata</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Liberation Mono</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Luxi Mono</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Nimbus Mono L</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Nimbus Mono</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Nimbus Mono PS</family>
		<default><family>monospace</family></default>
	</alias>
	<alias>
		<family>Terminal</family>
		<default><family>monospace</family></default>
	</alias>
<!--
  Fantasy faces
 -->
	<alias>
		<family>Bauhaus Std</family>
		<default><family>fantasy</family></default>
	</alias>
	<alias>
		<family>Cooper Std</family>
		<default><family>fantasy</family></default>
	</alias>
	<alias>
		<family>Copperplate Gothic Std</family>
		<default><family>fantasy</family></default>
	</alias>
	<alias>
		<family>Impact</family>
		<default><family>fantasy</family></default>
	</alias>
<!--
  Cursive faces
  -->
	<alias>
		<family>Comic Sans MS</family>
		<default><family>cursive</family></default>
	</alias>
	<alias>
		<family>ITC Zapf Chancery Std</family>
		<default><family>cursive</family></default>
	</alias>
	<alias>
		<family>Zapfino</family>
		<default><family>cursive</family></default>
	</alias>

  <!-- If the font still has no generic name, add sans-serif -->
  <match>
    <test qual="all" name="family" compare="not_eq">
      <string>sans-serif</string>
    </test>
    <test qual="all" name="family" compare="not_eq">
      <string>serif</string>
    </test>
    <test qual="all" name="family" compare="not_eq">
      <string>monospace</string>
    </test>
    <edit name="family" mode="append_last">
      <string>sans-serif</string>
    </edit>
  </match>
</fontconfig>
`

	fontsLocalConf = `
<?xml version="1.0"?>
<!DOCTYPE fontconfig SYSTEM "fonts.dtd">
<!-- /etc/fonts/fonts.conf file to configure system font access -->
<fontconfig>
  <its:rules xmlns:its="http://www.w3.org/2005/11/its" version="1.0">
    <its:translateRule translate="no" selector="/fontconfig/*[not(self::description)]"/>
  </its:rules>

  <!-- FULLY HINTED, ANTIALIASED FONTS -->
  <match target="font">
    <edit name="antialias" mode="assign"><bool>true</bool></edit>
    <edit name="autohinter" mode="assign"><bool>false</bool></edit>
    <edit name="dpi" mode="assign"><double>163</double></edit>
    <edit name="hintstyle" mode="assign"><const>hintfull</const></edit>
    <edit name="lcdfilter" mode="assign"><const>lcddefault</const></edit>
    <edit name="rgba" mode="assign"><const>rgb</const></edit>
  </match>

  <!-- DISABLE BITMAP FONTS EXCEPT FOR TERMINUS -->
  <selectfont>
    <rejectfont>
	  <pattern>
	    <patelt name="scalable"><bool>false</bool></patelt>
	  </pattern>
	</rejectfont>
	<acceptfont>
	  <pattern>
	    <patelt name="family"><string>Terminus</string></patelt>
	  </pattern>
	</acceptfont>
  </selectfont>

  <


</fontconfig>
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

MODULES=(i915)
BINARIES=()
FILES=()
HOOKS=(base udev autodetect modconf keyboard consolefont keymap block encrypt filesystems fsck)
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
