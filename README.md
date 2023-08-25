# AWS Profile Switcher

App that aids in switching active profile for your awscli.

```
Usage: aws-profile [options]
Options:
  -l list all profiles
  -s <profile> set profile
  -p <profile> print profile
  -d <profile> delete profile
  -a <profile> <accsess key id> <secret access key> add profile
  -h help
```

## Installation

 1. Download release from the [Relases](https://github.com/thorerik/aws-profile-switcher/releases) page.
 2. Add the contents of [trap.zsh](./zsh/trap.zsh) to your `~/.zshrc` or `~/.zprofile`
 3. Start switching profiles