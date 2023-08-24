package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func help() {
	fmt.Println("Usage: aws-profile [options]")
	fmt.Println("Options:")
	fmt.Println("  -l list all profiles")
	fmt.Println("  -p <profile> print profile")
	fmt.Println("  -d <profile> delete profile")
	fmt.Println("  -a <profile> add profile")
	fmt.Println("  -e <profile> edit profile")
	fmt.Println("  -h help")
}

func readCredentialsIni() *ini.File {
	// Get the home directory path
	homeDir, _ := os.UserHomeDir()

	path := homeDir + "/.aws/credentials"
	// Check if .aws/credentials exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("File does not exist")
	}

	// Load the file
	cfg, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	return cfg
}

func listProfiles() {
	cfg := readCredentialsIni()

	// Get the section names
	sections := cfg.SectionStrings()

	// Print the section names
	for _, section := range sections {
		fmt.Println(section)
	}
}

func printProfile(profile string) {
	cfg := readCredentialsIni()

	// Get the section
	section, err := cfg.GetSection(profile)
	if err != nil {
		fmt.Printf("Fail to get section: %v", err)
		os.Exit(1)
	}

	// Print the section
	fmt.Println(section)
}

func deleteProfile(profile string) {
	cfg := readCredentialsIni()

	// Delete the section
	cfg.DeleteSection(profile)

	fmt.Println("Deleted profile: " + profile)
	fmt.Printf("cfg: %v", cfg)

	// Save the file
	// err := cfg.SaveTo(cfg.Path())
	// if err != nil {
	// 	fmt.Printf("Fail to save file: %v", err)
	// 	os.Exit(1)
	// }
}

func addProfile(profile string) {
	cfg := readCredentialsIni()

	// Add the section
	cfg.NewSection(profile)

	fmt.Println("Added profile: " + profile)
	fmt.Printf("cfg: %v", cfg)

	// Save the file
	// err := cfg.SaveTo()
	// if err != nil {
	// 	fmt.Printf("Fail to save file: %v", err)
	// 	os.Exit(1)
	// }
}

func editProfile(profile string) {

}

func main() {
	// parse switches
	// -l list all profiles
	// -p <profile> print profile
	// -d <profile> delete profile
	// -a <profile> add profile
	// -e <profile> edit profile
	// -h help
	if len(os.Args) == 1 {
		help()
		os.Exit(1)
	}

	for i, arg := range os.Args {
		switch arg {
		case "-l":
			listProfiles()
		case "-p":
			if len(os.Args) < i+2 {
				help()
				os.Exit(1)
			}
			printProfile(os.Args[i+1])
		case "-d":
			if len(os.Args) < i+2 {
				help()
				os.Exit(1)
			}
			deleteProfile(os.Args[i+1])
		case "-a":
			if len(os.Args) < i+2 {
				help()
				os.Exit(1)
			}
			addProfile(os.Args[i+1])
		case "-e":
			if len(os.Args) < i+2 {
				help()
				os.Exit(1)
			}
			editProfile(os.Args[i+1])
		case "-h":
			help()
			os.Exit(0)
		}
	}
}
