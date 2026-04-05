package helpers

import (
	"fmt"
	"icefetch/src/modules"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Module struct {
	Format string `mapstructure:"format"`
}

type Config struct {
	Modules map[string]Module `mapstructure:"modules"`
}

type FetchFunc func() string

var Registry = map[string]FetchFunc{
	"hostname": modules.HostAndUser,
	"os":       modules.FullOSName,
	"break":    modules.Break,
	"shell":    modules.Shell,
	"packages": modules.Package,
	"wm-de":    modules.DesktopEnvironment,
	"kernel":   modules.Kernel,
	"cpu":      modules.Cpu,
	"host":     modules.Host,
	"ram":      modules.Ram,
}

func findAsciiFile(name string) (string, error) {

	paths := []string{
		filepath.Join(".", "ascii", "default", name),
		filepath.Join(".", "config", "ascii", "default", name),
		filepath.Join(os.Getenv("HOME"), ".config", "icefetch", "ascii", "default", name),
		filepath.Join("/etc", "icefetch", "ascii", "default", name),
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {

			return path, nil
		}
	}

	return "", fmt.Errorf("Unable to find logo in any standard location")
}

var Cfg Config

func readConfig() {

	viper.SetConfigName("icefetch")

	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("$HOME/.config/icefetch")
	viper.AddConfigPath("/etc/icefetch")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.SetConfigType("toml")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w ", err))
	}

	err = viper.Unmarshal(&Cfg)

	if err != nil {
		panic(err)
	}

}

func GenerateInfoSlice() []string {

	readConfig()

	var slice []string

	order := viper.GetStringSlice("order")

	for _, name := range order {

		fn, fnExists := Registry[name]

		conf, confExists := Cfg.Modules[name]

		if fnExists {

			data := fn()
			var formattedString string

			if !confExists || len(conf.Format) == 0 {
				formattedString = data
			} else {

				formattedString = strings.Replace(conf.Format, "$content", data, -1)

				if !strings.Contains(conf.Format, "$content") && name != "break" {
					formattedString = data
				}
			}

			slice = append(slice, formattedString)
		}
	}

	return slice
}

func GenerateLogoSlice() []string {

	osLogoFile, _ := GetOSReleaseValue("ID")
	osLogoFile += ".ascii"

	configLogo := viper.GetString("logo.path")
	var err error
	if configLogo == "" {
		configLogo, err = findAsciiFile(osLogoFile)

		if err != nil {
			configLogo, err = findAsciiFile("tux.ascii")

			if err != nil {
				panic(err)
			}
		}
	}

	logo := GetLogoSlice(configLogo)

	return logo

}
