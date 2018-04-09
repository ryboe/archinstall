package main

import (
	"flag"
	"log"
	"os"
)

var (
	chroot   = flag.Bool("chroot", false, "pass this flag if you want to run inside arch-chroot")
	password string
)

func main() {
	flag.Parse()

	// Make sure $ARCH_PASSWORD is set before we go any further. We'll need it
	// to set y0ssar1an's password later.
	password = os.Getenv("ARCH_PASSWORD")
	if password == "" {
		log.Fatal("$ARCH_PASSWORD not set")
	}

	// Installation happens in two phases. The first phase takes place on the
	// Arch USB installer. The second phase takes place in the chroot jail on
	// the mounted partition. During the second phase, this binary will be
	// copied into the chroot and run with the -chroot flag set.
	var err error
	if *chroot {
		err = phase2()
	} else {
		err = phase1()
	}
	if err != nil {
		log.Fatal(err)
	}

	if !(*chroot) {
		log.Println("installation complete!")
		log.Println("don't forget to umount -R /mnt before restarting")
	}
}
