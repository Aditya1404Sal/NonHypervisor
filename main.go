package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "child" {
		child()
		return
	}

	configFile := "config.txt"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	config, err := parseConfigFile(configFile)
	if err != nil {
		panic(err)
	}

	if err := runContainer(config); err != nil {
		panic(err)
	}
}

func parseConfigFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			config = append(config, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func runContainer(config []string) error {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, config...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
	}
	return cmd.Run()
}

func child() {
	config := os.Args[2:]
	fmt.Println("Running container with config:", config)

	for _, line := range config {
		parts := strings.Fields(line)
		switch parts[0] {
		case "FROM":
			setupRootFS(parts[1])
		case "ENV":
			setupEnv(parts[1:])
		case "RUN":
			runCommand(parts[1:])
		case "CMD":
			runCommand(parts[1:])
		}
	}

	must(syscall.Unmount("proc", 0))
	must(syscall.Unmount("thing", 0))
}

func setupRootFS(image string) {
	ref, err := name.ParseReference(image)
	must(err)

	img, err := remote.Image(ref)
	must(err)

	// Create a directory for the root filesystem
	rootfs := "/tmp/rootfs"
	must(os.MkdirAll(rootfs, 0755))

	// Save the image layers to the root filesystem
	err = tarball.WriteToFile(rootfs, ref, img)
	must(err)

	must(syscall.Chroot(rootfs))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))
}

func setupEnv(env []string) {
	for _, e := range env {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
		}
	}
}

func runCommand(cmd []string) {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	must(command.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
