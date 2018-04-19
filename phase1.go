package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/y0ssar1an/archinstall/sh"
	"golang.org/x/sys/unix"
)

const (
	targetDisk    = "/dev/nvme0n1"
	bootPartition = targetDisk + "p1"
	rootPartition = targetDisk + "p2"
	cryptDevice   = "/dev/mapper/cryptroot"

	mirrorListURL  = "https://www.archlinux.org/mirrorlist/?country=US&protocol=https&ip_version=6"
	mirrorListPath = "/etc/pacman.d/mirrorlist"
	makepkgConf    = "/etc/makepkg.conf"
)

func phase1() error {
	log.Println("checking to see if we can reach the internet...")
	reachable, err := hasInternet()
	if err != nil {
		return errors.Wrapf(err, "can't install: internet isn't working")
	}
	if !reachable {
		return errors.New("can't install: internet isn't working")
	}
	log.Println("internet is reachable")

	log.Println("enabling NTP...")
	err = enableNTP()
	if err != nil {
		return err
	}
	log.Println("NTP enabled")

	errc := make(chan error, 1)
	go func() {
		log.Println("updating mirrorlist...")
		errc <- updateMirrorList()
		log.Println("mirror list updated")
	}()

	log.Println("updating /etc/makepkg.conf so pacman will build faster...")
	err = editMakepkgConf()
	if err != nil {
		return err
	}
	log.Println("makepkg.conf updated on the installer")

	log.Printf("partitioning %s...", targetDisk)
	err = partitionDisk(targetDisk)
	if err != nil {
		return errors.Wrap(err, "can't install: failed to format disk")
	}
	log.Printf("%s partitioned", targetDisk)

	log.Println("encrypting root partition...")
	err = cryptsetup(rootPartition, password)
	if err != nil {
		return err
	}
	log.Printf("encrypted root partition %s created", cryptDevice)

	log.Println("formatting partitions...")
	err = formatPartitions()
	if err != nil {
		return err
	}
	log.Println("boot and root partitions formatted")

	log.Println("mounting partitions...")
	err = mountPartitions()
	if err != nil {
		return err
	}
	log.Println("partitions mounted")

	// Wait the updateMirrorList() to complete before running pacstrap.
	err = <-errc
	if err != nil {
		return err
	}

	log.Println("running pacstrap...")
	err = pacstrap()
	if err != nil {
		return err
	}
	log.Println("pacstrap completed successfully")

	log.Println("generating fstab...")
	err = genfstab()
	if err != nil {
		return err
	}
	log.Println("fstab generated")

	log.Println("configuring UEFI bootloader...")
	err = bootloaderConfig()
	if err != nil {
		return err
	}
	log.Println("UEFI bootloader configured")

	return archChroot()
}

// hasInternet verifies the Internet is reachable by making a HEAD request to
// google.com. HTTP is used because the Go stdlib does not provide a native ping
// implementation.
func hasInternet() (bool, error) {
	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Head("https://www.google.com")
	if err != nil {
		return false, err
	}

	return resp.StatusCode == http.StatusOK, nil
}

// enableNTP enables the NTP service, which sets the system clock.
func enableNTP() error {
	return sh.Command(5*time.Second, "timedatectl", "set-ntp", "true").Run()
}

// partitionDisk deletes the partition table on the given disk and creates a new
// table with a boot and root partition. It returns an error if there's a
// problem creating the new partitions.
func partitionDisk(disk string) error {
	const (
		timeout         = 10 * time.Second
		bootStartSector = "2MiB"
		bootEndSector   = "512MiB"
	)

	// Make a GUID partition table.
	err := sh.Command(timeout, "parted", "--script", disk, "mktable", "gpt").Run()
	if err != nil {
		return err
	}

	// Create the boot partition, aka the EFI System Partition (ESP).
	err = sh.Command(timeout, "parted", "--align", "optimal", disk, "mkpart", "ESP", "fat32", bootStartSector, bootEndSector).Run()
	if err != nil {
		return err
	}

	// Create the root partition.
	err = sh.Command(timeout, "parted", "--align", "optimal", disk, "mkpart", "primary", "ext4", bootEndSector, "100%").Run()
	if err != nil {
		return err
	}

	// Set the ESP flag on the boot partition.
	err = sh.Command(timeout, "parted", disk, "set", "1", "esp", "on").Run()
	if err != nil {
		return err
	}

	// Verify the partitions are correctly aligned.
	err = sh.Command(timeout, "parted", disk, "align-check", "opt", "1").Run()
	if err != nil {
		return err
	}

	return sh.Command(timeout, "parted", disk, "align-check", "opt", "2").Run()
}

// cryptsetup creates a new encrypted partition on the given disk.
func cryptsetup(disk, pwd string) error {
	const timeout = 30 * time.Second
	cmd := sh.Command(timeout, "cryptsetup", "luksFormat", "--type", "luks2", disk)
	cmd.Stdin = strings.NewReader(pwd)

	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = sh.Command(timeout, "cryptsetup", "open", disk, path.Base(cryptDevice))
	cmd.Stdin = strings.NewReader(pwd)
	return cmd.Run()
}

// format partitions formats the boot and root partitions.
func formatPartitions() error {
	// Format the boot partition as FAT32.
	const timeout = 1 * time.Minute
	if err := sh.Command(timeout, "mkfs.vfat", "-F", "32", bootPartition).Run(); err != nil {
		return err
	}

	// Format the root partition as ext4.
	return sh.Command(timeout, "mkfs.ext4", cryptDevice).Run()
}

// mountPartitions mounts the boot and root partitions.
func mountPartitions() error {
	err := unix.Mount(cryptDevice, "/mnt", "ext4", 0, "")
	if err != nil {
		return errors.Wrapf(err, "mount %s /mnt failed", cryptDevice, "/mnt")
	}

	err = os.Mkdir("/mnt/boot", 0644)
	if err != nil {
		return err
	}

	err = unix.Mount(bootPartition, "/mnt/boot", "vfat", 0, "")
	if err != nil {
		return errors.Wrapf(err, "mount %s %s failed", bootPartition, "/mnt/boot")
	}

	return nil
}

// updateMirrorlist fetches the latest US mirrors.
func updateMirrorList() (err error) {
	resp, err := httpGet(mirrorListURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tmpList, err := ioutil.TempFile("/tmp", "mirrorlist")
	if err != nil {
		return err
	}
	defer func() {
		if cerr := tmpList.Close(); err == nil {
			err = cerr
		}
	}()

	err = uncommentMirrors(tmpList, resp.Body)
	if err != nil {
		return err
	}

	return rankMirrors(tmpList.Name())
}

// httpGet sends a GET request to the given URL. It times out after 30s.
func httpGet(url string) (*http.Response, error) {
	client := http.Client{Timeout: 30 * time.Second}

	return client.Get(url)
}

// uncommentMirrors deletes the first character of every line in the given
// file reader. The first character is a comment character.
func uncommentMirrors(dstList io.Writer, srcList io.Reader) error {
	scanner := bufio.NewScanner(srcList)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			fmt.Fprintln(dstList, line[1:])
		}
	}

	return scanner.Err()
}

// rankMirrors sorts the mirrors in the given mirror list by performance.
// Then it copies the sorted mirror list to /etc/pacman.d/mirrorlist.
func rankMirrors(srcListPath string) (err error) {
	f, err := os.Create(mirrorListPath)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); err == nil {
			err = cerr
		}
	}()

	cmd := sh.Command(1*time.Minute, "rankmirrors", "-n", "5", srcListPath)
	cmd.RedirectStdout(f)

	return cmd.Run()
}

var (
	// Matches the CFLAGS setting in /etc/makepkg.conf.
	cflagsRx = regexp.MustCompile(`(?m)^\s*(C(?:XX)?FLAGS)="-march=x86-64\s-mtune=generic`)

	// Matches the BUILDDIR setting in /etc/makepkg.conf.
	buildDirRx = regexp.MustCompile(`(?m)^\s*(?:#\s*)?BUILDDIR=\S*`)
)

func editMakepkgConf() error {
	// We'll need the file mode (permission bits) when we write the new
	// makepkg.conf later.
	fi, err := os.Stat(makepkgConf)
	if err != nil {
		return err
	}

	// Read makepkg.conf into a buffer.
	txt, err := ioutil.ReadFile(makepkgConf)
	if err != nil {
		return err
	}

	// Edit the buffer.
	buildDir := fmt.Sprintf("BUILDDIR=%s", pacmanBuildDir)
	txt = cflagsRx.ReplaceAll(txt, []byte(`$1="-march=native`))
	txt = buildDirRx.ReplaceAllLiteral(txt, []byte(buildDir))

	// Overwrite the existing makepkg.conf with the buffer.
	return ioutil.WriteFile(makepkgConf, txt, fi.Mode())
}

// pacstrap runs the pacstrap command, which installs Arch Linux to the mounted
// partition.
func pacstrap() error {
	return sh.Command(5*time.Minute, "pacstrap", "/mnt", "base", "base-devel").Run()
}

// genfstab generates the filesystem table (fstab) for the mounted partition.
func genfstab() (err error) {
	fstab, err := os.OpenFile("/mnt/etc/fstab", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := fstab.Close(); err == nil {
			err = cerr
		}
	}()

	cmd := sh.Command(10*time.Second, "genfstab", "-U", "/mnt")
	cmd.RedirectStdout(fstab)

	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "failed to generate fstab")
	}

	return nil
}

// bootloaderConfig creates a new entry in the UEFI bootloader.
func bootloaderConfig() error {
	// Print existing bootloader entries.
	err := sh.Command(5*time.Second, "efibootmgr", "-v").Run()
	if err != nil {
		return err
	}

	// Delete 0th bootloader entry. Ignore error if the 0th entry is missing.
	sh.Command(5*time.Second, "efibootmgr", "-b", "0", "-B").Run()

	var blkidBuf strings.Builder
	cmd := sh.Command(5*time.Second, "blkid", "-s", "UUID", "-o", "value", rootPartition)
	cmd.Stdout = &blkidBuf
	err = cmd.Run()
	if err != nil {
		return err
	}

	cryptDeviceUUID := strings.TrimSpace(blkidBuf.String())
	unicode := fmt.Sprintf(`cryptdevice=UUID=%s:cryptroot root=/dev/mapper/cryptroot rootfstype=ext4 rw initrd=/intel-ucode.img initrd=/initramfs-linux.img i915.enable_guc_loading=1 i915.enable_guc_submission=1`, cryptDeviceUUID)

	err = sh.Command(5*time.Second, "efibootmgr", "--create", "--disk", "/dev/nvme0n1", "--part", "1", "--label", "Arch Linux", "--loader", "/vmlinuz-linux", "--unicode", unicode).Run()
	if err != nil {
		return err
	}

	// Print updated bootloader entries.
	return sh.Command(5*time.Second, "efibootmgr", "-v").Run()
}

// archChroot chroots into the mounted partition. This is the last step of phase
// 1. Phase 2 configuration takes place within the chroot jail.
func archChroot() error {
	thisBin := path.Base(os.Args[0])
	binPath := path.Join("/mnt", thisBin)
	srcPath := path.Join("/root/go/bin", path.Base(os.Args[0]))

	err := sh.Copy(binPath, srcPath)
	if err != nil {
		return err
	}

	err = sh.Command(30*time.Minute, "arch-chroot", "/mnt", "/"+thisBin, "-chroot").Run()
	if err != nil {
		return err
	}

	return os.Remove(binPath)
}
