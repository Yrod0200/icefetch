package modules

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func getSliceFromCommandStdout(cmd []string) ([]string, error) {

	command := exec.Command(cmd[0], cmd[1:]...)
	var out bytes.Buffer
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	return lines, nil

}

var PackageManagersListCommands = map[string][]string{
	"apt":     {"dpkg-query", "-W", "-f=${binary:Package}\n"},
	"dnf":     {"rpm", "-qa"},
	"zypper":  {"rpm", "-qa"},
	"pacman":  {"pacman", "-Q"},
	"apk":     {"apk", "info"},
	"portage": {"bash", "-c", "set -o pipefail; ls /var/db/pkg/*/*"},
	"flatpak": {"bash", "-c", "set -o pipefail; flatpak list --app"},
	"snap":    {"bash", "-c", "set -o pipefail; snap list | tail -n +2"},
	"nix":     {"nix-env", "-q"},
}

func removeAptHeaders(slice []string) []string {

	var cleanSlice []string

	for _, line := range slice {

		if strings.Contains(line, "ii") {

			cleanSlice = append(cleanSlice, line)

		}

	}

	return cleanSlice

}

func getPackageAmountByCommandOutput(manager string, lines []string) int {

	return len(lines)

}

func Package() string {

	var PackageCommandOrder = viper.GetStringSlice("packages.order")
	var PackageCommandFormat = viper.GetString("packages.format")

	if len(PackageCommandOrder) == 0 {
		PackageCommandOrder = []string{"apt", "dnf", "zypper", "pacman", "apk", "portage", "flatpak", "snap", "nix"}
	}

	var formattedPackageAmount string

	for _, manager := range PackageCommandOrder {

		command := PackageManagersListCommands[manager]

		if len(command) == 0 {
			continue
		}

		lines, err := getSliceFromCommandStdout(command)

		if err != nil {
			continue
		}

		amount := getPackageAmountByCommandOutput(manager, lines)

		if amount > 0 {
			formattedPackageAmount += (strings.Replace(strings.Replace(PackageCommandFormat, "$manager", manager, -1), "$amount", fmt.Sprint(amount), -1))
		}

	}

	return formattedPackageAmount

}
