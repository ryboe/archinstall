package main

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/y0ssar1an/archinstall/sh"
)

const (
	backgroundImageURL = "https://i.imgur.com/TbFzJFV.png" // arch_linux_black_3840x2160.png
	hostname           = "birdwell"
	keymapDir          = "/usr/share/kbd/keymaps/i386/qwerty/"
	keymapFilename     = "us-caps-esc.map.gz"
	locale             = "en_US.UTF-8"
	pacmanBuildDir     = "/tmp/makepkg"
	systemdNetworkDir  = "/etc/systemd/network/"
	systemdNetworkName = "50-wired.network"
	username           = "y0ssar1an"
	vimPlugURL         = "https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim"
	yayGitHubPath      = "github.com/Jguer/yay"
	zshPath            = "/usr/bin/zsh"
)

func phase2() error {
	// Tell pacman to build with architecture-native GCC optimizations. Also,
	// tell it to use an in-memory directory for building packages (fast!).
	log.Println("editing makepkg.conf on target disk...")
	err := editMakepkgConf()
	if err != nil {
		return err
	}
	log.Println("makepkg.conf edited on target disk")

	log.Printf("creating pacman build dir %s...", pacmanBuildDir)
	err = createPacmanBuildDir()
	if err != nil {
		return err
	}
	log.Printf("pacman build dir %s created", pacmanBuildDir)

	log.Println("setting locale...")
	err = setLocale()
	if err != nil {
		return err
	}
	log.Println("locale set")

	log.Println("setting hardware clock...")
	err = setHWClock()
	if err != nil {
		return err
	}
	log.Println("hardware clock set")

	log.Println("setting hostname...")
	err = ioutil.WriteFile("/etc/hostname", []byte(hostname+"\n"), 0644)
	if err != nil {
		return err
	}
	log.Printf("hostname %q set", hostname)

	log.Println("setting root password...")
	err = setPassword("root", password)
	if err != nil {
		return err
	}
	log.Println("root password set")

	log.Println("installing essential pkgs...")
	err = installPkgs()
	if err != nil {
		return err
	}
	log.Println("essential pkgs installed")

	log.Println("updating sudoers file...")
	err = ioutil.WriteFile("/etc/sudoers", []byte(sudoers[1:]), 0440)
	if err != nil {
		return err
	}
	log.Println("sudoers file updated")

	log.Printf("creating user %s with default shell %s", username, zshPath)
	err = createUser(username, zshPath)
	if err != nil {
		return err
	}
	log.Printf("user %q created", username)

	log.Printf("setting password for %s...", username)
	err = setPassword(username, password)
	if err != nil {
		return err
	}
	log.Printf("user %q password set", username)

	log.Println("creating systemd network...")
	err = createNetwork()
	if err != nil {
		return err
	}
	log.Println("ethernet network created")

	log.Printf("creating %s keymap...", keymapFilename)
	err = createCapsEscKeymap()
	if err != nil {
		return err
	}
	log.Printf("console keymap %s created", keymapFilename)

	log.Printf("loading console keymap %s...", keymapFilename)
	err = loadKeymap()
	if err != nil {
		return err
	}
	log.Printf("console keymap %s loaded", keymapFilename)

	log.Println("creating vconsole config...")
	err = ioutil.WriteFile("/etc/vconsole.conf", []byte(vconsoleConf[1:]), 0644)
	if err != nil {
		return err
	}
	log.Println("console font and keymap set")

	log.Println("updating initramfs...")
	err = configureInitramfs()
	if err != nil {
		return err
	}
	log.Println("initramfs updated")

	log.Println("updating rngd configuration...")
	err = ioutil.WriteFile("/etc/conf.d/rngd", []byte(rngd[1:]), 0644)
	if err != nil {
		return err
	}
	log.Println("rngd configuration updated")

	log.Println("updating systemd-resolved configuration...")
	err = ioutil.WriteFile("/etc/systemd/resolved.conf", []byte(resolvedConf[1:]), 0644)
	if err != nil {
		return err
	}
	log.Println("systemd-resolved configuration updated")

	log.Println("creating kbdrate.service (which increases the keyboard repeat rate)...")
	err = ioutil.WriteFile("/etc/systemd/system/kbdrate.service", []byte(kbdrateService[1:]), 0644)
	if err != nil {
		return err
	}
	log.Println("console key repeat rate increased")

	log.Println("enabling BBR congestion control...")
	err = enableBBR()
	if err != nil {
		return err
	}
	log.Println("BBR enabled")

	log.Println("enabling kbdrate, rngd, systemd-networkd, systemd-resolved, systemd-timesyncd...")
	err = sh.Command(10*time.Second, "systemctl", "enable", "kbdrate", "rngd", "systemd-networkd", "systemd-resolved", "systemd-timesyncd").Run()
	if err != nil {
		return err
	}
	log.Println("kbdrate, rngd, systemd-networkd, systemd-resolved, systemd-timesyncd enabled")

	log.Println("downloading desktop background image...")
	err = downloadBackgroundImage()
	if err != nil {
		return err
	}
	log.Println("background image downloaded")

	log.Println("configuring fonts...")
	err = configureFonts()
	if err != nil {
		return err
	}
	log.Println("fonts configured")

	log.Println("creating go directories...")
	err = setupGoEnv()
	if err != nil {
		return err
	}
	log.Println("go directories created")

	log.Println("installing rust...")
	err = installRust()
	if err != nil {
		return err
	}
	log.Println("rust installed")

	log.Println("installing yay AUR helper...")
	err = installAURHelper()
	if err != nil {
		return err
	}
	log.Println("AUR helper yay installed")

	log.Println("installing firefox...")
	err = installFirefox()
	if err != nil {
		return err
	}
	log.Println("firefox installed")

	log.Println("installing gitprompt...")
	err = installGitprompt()
	if err != nil {
		return err
	}
	log.Println("gitprompt installed")

	log.Println("downloading configs to ~/.config ...")
	err = createDotConfigDir()
	if err != nil {
		return err
	}
	log.Println("~/.config directory created. config files downloaded")

	log.Println("installing vim-plug...")
	err = installVimPlug()
	if err != nil {
		return err
	}
	log.Println("vim-plug installed")

	log.Printf("setting ownership of /home/%[1]s to %[1]s...\n", username)
	homeDir := path.Join("/home/", username)
	err = recursiveChown(homeDir, username, "users")
	if err != nil {
		return err
	}
	log.Printf("owner of /home/%[1]s set to %[1]s\n", username)

	log.Println("unmuting master volume control...")
	err = unmuteAudio()
	if err != nil {
		return err
	}
	log.Println("master volume control unmuted")

	log.Println("installing GNOME...")
	err = installGNOME()
	if err != nil {
		return err
	}
	log.Println("GNOME installed")

	return nil
}

// Create a directory in /tmp for building pacman packages. /tmp uses tmpfs, an
// in-memory filesystem. Having an in-memory working directory for building
// packages greatly speeds up the process.
func createPacmanBuildDir() error {
	err := os.Mkdir(pacmanBuildDir, 0644)
	if err != nil {
		return err
	}

	wheelGID, err := groupID("wheel")
	if err != nil {
		return err
	}

	// Change the group ownership to wheel, so the y0ssar1an user (member of
	// of wheel) can use this directory to build packages with pacman.
	const rootUID = 0
	return os.Chown(pacmanBuildDir, rootUID, wheelGID)
}

func userID(username string) (int, error) {
	usr, err := user.Lookup("root")
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(usr.Uid)
}

func groupID(groupname string) (int, error) {
	grp, err := user.LookupGroup("wheel")
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(grp.Gid)
}

var (
	localeRx = regexp.MustCompile(`(?m)^\s*#\s*` + locale)
)

func setLocale() error {
	const localeGenPath = "/etc/locale.gen"
	fi, err := os.Stat(localeGenPath)
	if err != nil {
		return err
	}

	bs, err := ioutil.ReadFile(localeGenPath)
	if err != nil {
		return err
	}

	bs = localeRx.ReplaceAllLiteral(bs, []byte(locale))

	err = ioutil.WriteFile(localeGenPath, bs, fi.Mode())
	if err != nil {
		return err
	}

	err = sh.Command(10*time.Second, "locale-gen").Run()
	if err != nil {
		return err
	}

	return ioutil.WriteFile("/etc/locale.conf", []byte(localeConf[1:]), 0644)
}

func setHWClock() error {
	return sh.Command(5*time.Second, "hwclock", "--systohc", "--utc").Run()
}

func setPassword(username, pwd string) error {
	if username == "" {
		return errors.New("no username given")
	}

	cmd := sh.Command(10*time.Second, "chpasswd", username)
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%s:%s", username, pwd))

	return cmd.Run()
}

func installPkgs() error {
	args := []string{
		"-Sy",
		"--noconfirm",
		"alsa-utils",
		"freetype2-cleartype",
		"fzf",
		"git",
		"go",
		"htop",
		"intel-ucode",
		"libva-intel-driver",
		"libva-utils",
		"mesa",
		"neovim",
		"noto-fonts",
		"noto-fonts-cjk",
		"noto-fonts-emoji",
		"noto-fonts-extra",
		"otf-font-awesome",
		"pulseaudio",
		"pulseaudio-alsa",
		"rng-tools",
		"rustup",
		"terminus-font",
		"tree",
		"ufw",
		"vulkan-intel",
		"zsh",
	}

	cmd := sh.Command(15*time.Minute, "pacman", args...)
	cmd.Stdin = strings.NewReader("1\n")

	return cmd.Run()
}

func createUser(username, defaultShell string) error {
	return sh.Command(10*time.Second, "useradd", "-m", "-g", "users", "-G", "wheel", "-s", defaultShell, username).Run()
}

func createNetwork() error {
	path := path.Join(systemdNetworkDir, systemdNetworkName)

	return ioutil.WriteFile(path, []byte(wiredNetwork[1:]), 0644)
}

func createCapsEscKeymap() (err error) {
	dst := path.Join(keymapDir, keymapFilename)

	outf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := outf.Close(); err == nil {
			err = cerr
		}
	}()

	zipw := gzip.NewWriter(outf)
	defer func() {
		if cerr := zipw.Close(); err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(zipw, strings.NewReader(keymap[1:]))
	return err
}

func loadKeymap() error {
	keymapName := strings.TrimSuffix(keymapFilename, ".map.gz")

	return sh.Command(5*time.Second, "loadkeys", keymapName).Run()
}

func configureInitramfs() error {
	if err := ioutil.WriteFile("/etc/mkinitcpio.conf", []byte(mkinitcpioConf[1:]), 0644); err != nil {
		return err
	}

	return sh.Command(1*time.Minute, "mkinitcpio", "-p", "linux").Run()
}

func configureFonts() error {
	const (
		fontsDir      = "/etc/fonts/conf.d/"
		fontsAvailDir = "/etc/fonts/conf.avail/"
	)

	badLinks := []string{
		"10-hinting-slight.conf",
	}

	for _, lnk := range badLinks {
		badSymlink := path.Join(fontsDir, lnk)
		err := os.Remove(badSymlink)
		if err != nil {
			return err
		}
	}

	links := []string{
		"10-hinting-full.conf",
		"10-sub-pixel-rgb.conf",
		"66-noto-mono.conf",
		"66-noto-sans.conf",
		"66-noto-serif.conf",
	}

	for _, lnk := range links {
		availLnk := path.Join(fontsAvailDir, lnk)
		confLnk := path.Join(fontsDir, lnk)
		err = os.Symlink(availLnk, confLnk)
		if err != nil {
			return err
		}
	}

	return nil
}

func setupGoEnv() error {
	gopath := path.Join("/home/", username, "/go")

	var err error
	for _, dir := range []string{"/bin", "/pkg", "/src"} {
		err = os.MkdirAll(path.Join(gopath, dir), 0644)
		if err != nil {
			return err
		}
	}

	return recursiveChown(gopath, username, "users")
}

func installRust() (err error) {
	const toolchainInstallTimeout = 10 * time.Minute
	err = sh.Command(toolchainInstallTimeout, "rustup", "toolchain", "install", "stable").Run()
	if err != nil {
		return err
	}

	err = sh.Command(toolchainInstallTimeout, "rustup", "toolchain", "install", "nightly").Run()
	if err != nil {
		return err
	}

	const switchToolchainTimeout = 10 * time.Second
	err = sh.Command(switchToolchainTimeout, "rustup", "default", "nightly").Run()
	if err != nil {
		return err
	}
	defer func() {
		if cerr := sh.Command(switchToolchainTimeout, "rustup", "default", "stable").Run(); err == nil {
			err = cerr
		}
	}()

	return installRustUtils()
}

func installRustUtils() (err error) {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer func() {
		if cderr := os.Chdir(wd); err == nil {
			err = cderr
		}
	}()

	return installRipgrep()
}

func installRipgrep() (err error) {
	err = cloneRustProject("BurntSushi/ripgrep")
	if err != nil {
		return err
	}

	dir := path.Join("/home", username, "rust/github.com/BurntSushi/ripgrep")
	err = os.Chdir(dir)
	if err != nil {
		return err
	}

	cmd := sh.Command(5*time.Minute, "cargo", "build", "--release", "--features", "simd-accel avx-accel")
	cmd.Env = []string{"RUSTFLAGS=-C target-cpu=native"}

	err = cmd.Run()
	if err != nil {
		return err
	}

	dst := path.Join("/home", username, ".cargo/bin/rg")
	src := path.Join(dir, "target/release/rg")

	return sh.Copy(dst, src)
}

func cloneRustProject(name string) error {
	p := path.Join("/home", username, "rust/github.com/", name)
	if err := os.MkdirAll(p, 0755); err != nil {
		return err
	}

	githubURL := fmt.Sprintf("https://github.com/%s", name)
	return sh.Command(2*time.Minute, "git", "clone", githubURL, path.Dir(p)).Run()
}

// enableBBR enables BBR TCP congestion control. BBR prevents bufferbloat and
// speeds up TCP.
func enableBBR() error {
	if err := ioutil.WriteFile("/etc/modules-load.d/modules.conf", []byte("tcp_bbr\n"), 0644); err != nil {
		return err
	}

	return ioutil.WriteFile("/etc/sysctl.d/bbr.conf", []byte(bbrConf[1:]), 0644)
}

func createDotConfigDir() error {
	homeDir := path.Join("/home/", username)
	dotConfigDir := path.Join(homeDir, ".config")

	err := sh.Command(1*time.Minute, "git", "clone", "https://github.com/y0ssar1an/dotconfig", dotConfigDir).Run()
	if err != nil {
		return err
	}

	dotGitDir := path.Join(dotConfigDir, ".git")
	err = os.RemoveAll(dotGitDir)
	if err != nil {
		return err
	}

	// The ~/.zshenv file points zsh at ~/.config/zsh, where the .zshrc file is
	// located.
	zshEnvPath := path.Join(homeDir, ".zshenv")
	return ioutil.WriteFile(zshEnvPath, []byte(zshenv[1:]), 0644)
}

func recursiveChown(path, username, groupname string) error {
	uid, err := userID(username)
	if err != nil {
		return err
	}

	gid, err := groupID(groupname)
	if err != nil {
		return err
	}

	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		return os.Chown(name, uid, gid)
	})
}

func downloadBackgroundImage() (err error) {
	const dst = "/usr/share/backgrounds/arch_linux_black_3840x2160.png"

	return downloadFile(1*time.Minute, dst, backgroundImageURL)
}

func installAURHelper() error {
	return sh.Command(1*time.Minute, "go", "get", "-u", yayGitHubPath).Run()
}

func installGitprompt() error {
	return sh.Command(1*time.Minute, "go", "get", "-u", "github.com/y0ssar1an/gitprompt").Run()
}

func installVimPlug() (err error) {
	home := path.Join("/home/", username)
	vimPlugDst := path.Join(home, ".local/share/nvim/site/autoload/plug.vim")

	return downloadFile(1*time.Minute, vimPlugDst, vimPlugURL)
}

func downloadFile(timeout time.Duration, dst, url string) (err error) {
	client := http.Client{Timeout: timeout}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = os.MkdirAll(path.Dir(dst), 0750)
	if err != nil {
		return err
	}

	outf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0744)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := outf.Close(); err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(outf, resp.Body)
	return err
}

func installFirefox() error {
	return sh.Command(10*time.Minute, "yay", "-Sy", "--noconfirm", "firefox").Run()
}

func unmuteAudio() error {
	err := sh.Command(10*time.Second, "pactl", "set-sink-mute", "0", "0").Run()
	if err != nil {
		return err
	}

	err = sh.Command(10*time.Second, "pactl", "set-sink-volume", "0", "50%").Run()
	if err != nil {
		return err
	}

	err = sh.Command(10*time.Second, "pactl", "set-source-mute", "1", "1").Run()
	if err != nil {
		return err
	}

	return sh.Command(10*time.Second, "pactl", "set-source-volume", "1", "50%").Run()
}

func installGNOME() error {
	return sh.Command(30*time.Minute, "yay", "-Sy", "--noconfirm", "gnome").Run()
}
