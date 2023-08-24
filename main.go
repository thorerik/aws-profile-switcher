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
	fmt.Println("  -a <profile> <accsess key id> <secret access key> add profile")
	fmt.Println("  -e <profile> edit profile")
	fmt.Println("  -h help")
}

func getPath() string {
	// Get the home directory path
	homeDir, _ := os.UserHomeDir()

	return homeDir + "/.aws/credentials"
}

func readCredentialsIni() *ini.File {
	// Get the path to the credentials file
	path := getPath()

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

	// Print each key and value
	for _, key := range section.Keys() {
		fmt.Println(key.Name() + "=" + key.Value())
	}
}

func deleteProfile(profile string) {
	cfg := readCredentialsIni()

	// Delete the section
	cfg.DeleteSection(profile)

	fmt.Println("Deleted profile: " + profile)
	fmt.Printf("cfg: %v", cfg)

	// Save the file
	err := cfg.SaveTo(getPath())
	if err != nil {
		fmt.Printf("Fail to save file: %v", err)
		os.Exit(1)
	}
}

func addProfile(profile string, awsAccessKeyID string, awsSecretAccessKey string) {
	cfg := readCredentialsIni()

	// Add the section
	cfg.NewSection(profile)

	// add accesskeyid and secret
	cfg.Section(profile).NewKey("aws_access_key_id", awsAccessKeyID)
	cfg.Section(profile).NewKey("aws_secret_access_key", awsSecretAccessKey)

	fmt.Println("Added profile: " + profile)
	fmt.Printf("cfg: %v", cfg)

	// Save the file
	err := cfg.SaveTo(getPath())
	if err != nil {
		fmt.Printf("Fail to save file: %v", err)
		os.Exit(1)
	}
}

func editProfile(profile string) {

}

func main() {
	// parse switches
	// -l list all profiles
	// -p <profile> print profile
	// -d <profile> delete profile
	// -a <profile> <accsess key id> <secret access key> add profile
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
			if len(os.Args) < i+4 {
				help()
				os.Exit(1)
			}
			addProfile(os.Args[i+1], os.Args[i+2], os.Args[i+3])
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
