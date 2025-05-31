package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const (
	lightColorScheme = "prefer-light"
	darkColorScheme = "prefer-dark"
	lightTheme = "Adwaita"
	darkTheme = "Adwaita-dark"
)

func getColorScheme() ([]byte, error) {
	scheme, err := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "color-scheme").Output()

	return scheme, err
}

func setColorScheme(isDark bool) error {
	scheme := darkColorScheme

	if isDark {
		scheme = lightColorScheme
	}

	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "color-scheme", scheme)
	err := cmd.Run()

	return err
}

func setTheme(isDark bool) error {
	theme := darkTheme

	if isDark {
		theme = lightTheme
	}

	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", theme)
	err := cmd.Run()

	return err
}

func main() {
	scheme, err := getColorScheme()
	if err != nil {
		log.Fatalln("Unable to get current color scheme:", err)
	}

	isDark := strings.Contains(string(scheme), darkColorScheme)

	err = setColorScheme(isDark)
	if err != nil {
		log.Fatalln("Error while changing color scheme:", err)
	}

	err = setTheme(isDark)
	if err != nil {
		log.Fatalln("Error while changing theme:", err)
	}

	if isDark {
		fmt.Println("Color scheme changed to: light")
	} else {
		fmt.Println("Color scheme changed to: dark")
	}
}
