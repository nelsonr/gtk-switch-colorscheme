package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func getColorScheme() ([]byte, error) {
	scheme, err := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "color-scheme").Output()

	return scheme, err
}

func getIsDarkScheme(scheme string) bool {
	return strings.TrimSpace(scheme) == "'prefer-dark'"
}

func setColorScheme(isDarkScheme bool) error {
	scheme := "prefer-dark"

	if isDarkScheme {
		scheme = "prefer-light"
	}

	err := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "color-scheme", scheme).Run()

	return err
}

func setTheme(isDarkScheme bool) error {
	theme := "Adwaita-dark"

	if isDarkScheme {
		theme = "Adwaita-light"
	}

	err := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", theme).Run()

	return err
}

func main() {
	scheme, err := getColorScheme()
	if err != nil {
		log.Fatalln("Unable to get theme name", err)
	}

	isDarkScheme := getIsDarkScheme(string(scheme))

	err = setColorScheme(isDarkScheme)
	if err != nil {
		log.Fatalln("Error while changing color scheme", err)
	}

	err = setTheme(isDarkScheme)
	if err != nil {
		log.Fatalln("Error while changing theme", err)
	}

	if isDarkScheme {
		fmt.Println("Color scheme changed to: light")
	} else {
		fmt.Println("Color scheme changed to: dark")
	}
}
