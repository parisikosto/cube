package linux

import (
	"fmt"
	"os"
)

// InstallDockerPrerequisites installs the packages required before adding the Docker repository.
func InstallDockerPrerequisites() error {
	if err := AptUpdate(); err != nil {
		return err
	}

	return runCmdInteractive(
		"$ sudo apt install ca-certificates curl gnupg",
		"sudo", "apt", "install", "ca-certificates", "curl", "gnupg",
	)
}

// SetupDockerGPGKey creates the keyrings directory, downloads and stores the Docker GPG key.
func SetupDockerGPGKey() error {
	if err := runCmd(
		"$ sudo install -m 0755 -d /etc/apt/keyrings",
		"sudo", "install", "-m", "0755", "-d", "/etc/apt/keyrings",
	); err != nil {
		return err
	}

	if err := runShell(
		"$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg",
		"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg",
	); err != nil {
		return err
	}

	return runCmd(
		"$ sudo chmod a+r /etc/apt/keyrings/docker.gpg",
		"sudo", "chmod", "a+r", "/etc/apt/keyrings/docker.gpg",
	)
}

// AddDockerRepository adds the Docker apt repository to the system sources.
func AddDockerRepository() error {
	return runShell(
		`$ echo "deb [arch=... signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu ... stable" | sudo tee /etc/apt/sources.list.d/docker.list`,
		`echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null`,
	)
}

// VerifyDockerRepository checks that docker-ce is available from the Docker repo.
func VerifyDockerRepository() error {
	return runCmd("$ apt-cache policy docker-ce", "apt-cache", "policy", "docker-ce")
}

// InstallDockerCE installs Docker Community Edition.
func InstallDockerCE() error {
	if err := AptUpdate(); err != nil {
		return err
	}

	return runCmdInteractive(
		"$ sudo apt install docker-ce",
		"sudo", "apt", "install", "docker-ce",
	)
}

// VerifyDockerRunning checks that the Docker service is active.
func VerifyDockerRunning() error {
	return runCmd("$ sudo systemctl status docker --no-pager", "sudo", "systemctl", "status", "docker", "--no-pager")
}

// CheckDockerGroup prints the groups of the current user to verify docker group membership.
func CheckDockerGroup() error {
	return runCmd("$ id -nG", "id", "-nG")
}

// AddUserToDockerGroup adds the current user to the docker group.
// Returns the username that was added.
func AddUserToDockerGroup() (string, error) {
	user := os.Getenv("USER")
	if user == "" {
		user = "${USER}"
	}

	err := runCmd(
		fmt.Sprintf("$ sudo usermod -aG docker %s", user),
		"sudo", "usermod", "-aG", "docker", user,
	)
	return user, err
}

// DockerHelloWorld runs the docker hello-world container to verify Docker is working.
func DockerHelloWorld() error {
	return runCmd("$ docker run hello-world", "docker", "run", "hello-world")
}

// PruneDocker removes all unused Docker resources: system, containers, volumes, networks, and images.
func PruneDocker() error {
	if err := runCmd("$ docker system prune", "docker", "system", "prune"); err != nil {
		return err
	}
	if err := runCmd("$ docker container prune", "docker", "container", "prune"); err != nil {
		return err
	}
	if err := runCmd("$ docker volume prune", "docker", "volume", "prune"); err != nil {
		return err
	}
	if err := runCmd("$ docker network prune", "docker", "network", "prune"); err != nil {
		return err
	}
	return runCmd("$ docker image prune -a", "docker", "image", "prune", "-a")
}

// UninstallDocker removes Docker CE and cleans up unused dependencies.
func UninstallDocker() error {
	if err := runCmdInteractive("$ sudo apt remove docker-ce", "sudo", "apt", "remove", "docker-ce"); err != nil {
		return err
	}

	return runCmdInteractive("$ sudo apt autoremove", "sudo", "apt", "autoremove")
}

// InstallDocker runs the full Docker CE installation sequence.
func InstallDocker() error {
	if err := InstallDockerPrerequisites(); err != nil {
		return err
	}

	if err := SetupDockerGPGKey(); err != nil {
		return err
	}

	if err := AddDockerRepository(); err != nil {
		return err
	}

	if err := AptUpdate(); err != nil {
		return err
	}

	if err := VerifyDockerRepository(); err != nil {
		return err
	}

	if err := InstallDockerCE(); err != nil {
		return err
	}

	return VerifyDockerRunning()
}
