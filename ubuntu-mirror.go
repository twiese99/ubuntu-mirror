package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	BaseDir string
	MirrorType string
	CountryCode string
	RsyncSource string
	Interval int64
)

func init() {
	flag.StringVar(&BaseDir, "basedir", "/data", "directory to sync")
	flag.StringVar(&MirrorType,"type", "releases", "type of mirror (releases, archive)")
	flag.StringVar(&CountryCode,"country", "", "country code for using a country specific mirror")
	flag.StringVar(&RsyncSource,"source", "", "custom rsync server to use. type of server needs to match the mirror-type")
	flag.Int64Var(&Interval,"interval", 360, "interval to run the sync in minutes")
	flag.Parse()
}

func main() {
	exists, isDir, _ := exists(BaseDir)
	if !exists {
		log.Println(fmt.Errorf("%v does not exist yet, trying to create it", BaseDir))
		err := os.Mkdir(BaseDir, os.FileMode(0644))
		if err != nil {
			log.Panic(fmt.Errorf("creation of %v failed", BaseDir))
		}
	} else {
		if !isDir {
			log.Panic(fmt.Errorf("%v is not an directory", BaseDir))
		}
	}

	rsyncSource := getRsyncSource(MirrorType)
	rsyncScript := getRsyncScript(MirrorType)

	log.Printf("Application setup finished with MirrorType: %v", MirrorType)
	log.Printf("Using BaseDir: %v", BaseDir)
	log.Printf("Using RsyncSource: %v", rsyncSource)
	log.Printf("Using Script: %v", rsyncScript)

	for {
		execCmd(exec.Command("/bin/bash", rsyncScript, rsyncSource, BaseDir))
		log.Printf("Next execution in %v minutes", Interval)
		time.Sleep(time.Duration(Interval) * time.Minute)
	}
}

func execCmd(cmd *exec.Cmd) {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw
	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Panicf("error %s", err)
	}
	log.Println(stdBuffer.String())
}

// exists returns whether the given path exists (first bool) and if the path is an directory (second bool)
func exists(path string) (bool, bool, error) {
	info, err := os.Stat(path)
	if err == nil { return true, info.IsDir(), nil }
	if os.IsNotExist(err) { return false, false, nil }
	return false, false, err
}

func notEmpty(val string, def string) string {
	if len(val) > 0 {
		return val
	} else {
		return def
	}
}

func getRsyncScript(mirrorType string) string {
	switch mirrorType {
	case "releases":
		return "./sync-releases.sh"
	case "archive":
		return "./sync-archive.sh"
	default:
		log.Panic(fmt.Errorf("unknown mirror type: %v", mirrorType))
		return ""
	}
}

func getRsyncSource(mirrorType string) string {
	switch mirrorType {
	case "releases":
		return notEmpty(RsyncSource, getReleasesRsyncSource())
	case "archive":
		return notEmpty(RsyncSource, getArchiveRsyncSource())
	default:
		log.Panic(fmt.Errorf("unknown mirror type: %v", mirrorType))
		return ""
	}
}

func getReleasesRsyncSource() string {
	if len(CountryCode) > 0 {
		return fmt.Sprintf("rsync://%v.rsync.releases.ubuntu.com/releases", CountryCode)
	} else {
		return "rsync://rsync.releases.ubuntu.com/ubuntu-releases"
	}
}

func getArchiveRsyncSource() string {
	if len(CountryCode) > 0 {
		return fmt.Sprintf("rsync://%v.rsync.archive.ubuntu.com/ubuntu", CountryCode)
	} else {
		return "rsync://archive.ubuntu.com/ubuntu"
	}
}
