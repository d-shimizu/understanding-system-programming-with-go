package main

import (
	"github.com/opencontainers/runc/libcontainer"
	"github.com/opencontainers/runc/libcontainer/configs"
	_ "github.com/opencontainers/runc/libcontainer/nsenter"
	"log"
	"os"
	"runtime"
	"path/filepath"
	"golang.org/x/sys/unix"
)

// main()の前に実行される処理
func init() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		runtime.GOMAXPROCS(1)
		runtime.LockOSThread()
		factory, _ := libcontainer.New("")
		if err := factory.StartInitialization(); err != nil {
			log.Fatal(err)
		}
		panic("--this line should have never been executed, congratulations--")
	}
}

func main() {
	abs, _ := filepath.Abs("./")
	factory, err := libcontainer.New(abs, libcontainer.Cgroupfs, libcontainer.InitArgs(os.Args[0],"init"))
	// _, err := libcontainer.New(abs, libcontainer.Cgroupfs, libcontainer.InitArgs(os.Args[0],"init"))
	if err != nil {
		log.Fatal(err)
		return
	}

	// Container Environment Setting
	capabilities := []string{
		"CAP_CHOWN",
		"CAP_DAC_OVERRIDE",
		"CAP_FSETID",
		"CAP_FOWNER",
		"CAP_MKNOD",
		"CAP_NET_RAW",
		"CAP_SETGID",
		"CAP_SETUID",
		"CAP_SETFCAP",
		"CAP_SETPCAP",
		"CAP_NET_BIND_SERVICE",
		"CAP_SYS_CHROOT",
		"CAP_KILL",
		"CAP_AUDIT_WRITE",
	}

	defaultMountFlags := unix.MS_NOEXEC | unix.MS_NOSUID| unix.MS_NODEV

	// &（アドレス演算子）を使って、configのアドレス(ポインタ)を代入
	// configsパッケージのConig構造体へアクセスしている
	// https://qiita.com/zurazurataicho/items/4a95e0daf0d960cfc2f7
	// https://golang.hateblo.jp/entry/golang-how-to-use-struct
	config := &configs.Config{
		Rootfs: abs+"/rootfs",
		// thread capabilities (man 7 capabilities)
		Capabilities: *configs.Capabilities{
			Bounding: capabilities,
			Effective: capabilities,
			Inheritable: capabilities,
			Permitted: capabilities,
			Ambient: capabilities,
		},
		Namespaces: configs.Namespaces([]configs.Namespace{
			{ Type: configs.NEWNS },
			{ Type: configs.NEWUTS },
			{ Type: configs.NEWPID },
			{ Type: configs.NEWNET },
		}),
		Cgroups: &configs.Cgroup{
			Name: "test-container",
			Parent: "system",
			Resources: &configs.Resources {
				MemorySwappiness: nil,
				AllowAllDevices: nil,
				AllowedDevices: configs.DefaultAllowedDevices,
			},
		},
		MaskPaths: []string{
			"/proc/kcore", "/sys/firmware",
		},
		ReadonlyPaths: []string{
			"/proc/sys", "/proc/sysrq-trigger", "/proc/irq", "/proc/bus",
		},
		Devices: configs.DefaultAutoCreatedDevices,
		Hostname: "testing",
		Mounts: []*configs.Mount{
			{
				Source: "proc",
				Destination: "/proc",
				Device: "proc",
				Flags: defaultMountFlags,
			},
			{
				Source: "tmpfs",
				Destination: "/dev",
				Device: "tmpfs",
				Flags: unix.MS_NOSUID|unix.MS_STRICTATIME,
				Data: "mode=755",
			},
			{
				Source: "devpts",
				Destination: "/dev/pts",
				Device: "devpts",
				Flags: unix.MS_NOSUID | unix.MS_NOEXEC,
				Data: "newinstance, ptmxmode=0666, mode=0622, gid=5",
			},
			{
				Source: "mqueue",
				Destination: "/dev/mqueue",
				Device: "mqueue",
				Flags: defaultMountFlags,
			},
			{
				Source: "sysfs",
				Destination: "/sys",
				Device: "sysfs",
				Flags: defaultMountFlags | unix.MS_RDONLY,
			},
		},
		Networks: []*configs.Network{
			{
				Type: "loopback",
				Address: "127.0.0.1/0",
				Gateway: "localhost",
			},
		},
		Rlimits: []configs.Rlimit{
			{
				Type: unix.RLIMIT_NOFILE,
				Hard: uint64(1025),
				Soft: uint64(1025),
			},
		},
	}

	// Container Creation Process
	container, err := factory.Create("container-id", config)
	if err != nil {
		log.Fatal(err)
		return
	}

	process := &libcontainer.Process{
		Args: []string{"/bin/sh"},
		Env: []string{"PATH=/bin"},
		User: "root",
		Stdin: os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err = container.Run(process)
	if err != nil {
		container.Destroy()
		log.Fatal(err)
		return
	}

	_, err = process.Wait()
	if err != nil {
		log.Fatal(err)
	}

	container.Destroy()
}
