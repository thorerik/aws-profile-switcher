use pretty_ini::{ini, ini_file};

extern crate dirs;

fn main() {
    let home = match dirs::home_dir() {
        Some(path) => path,
        None => panic!("Impossible to get your home dir!"),
    };

    let credentials_path = home.join(".aws/credentials");
    let credentials_path = credentials_path.to_str().unwrap();

    let config_path = home.join(".aws/config");
    let config_path = config_path.to_str().unwrap();

    // Check that the .aws/credentials file exists
    let mut file = ini_file::IniFile::default();
    file.set_path(credentials_path);

    // List all the profiles
    let profiles = file.sections();
    println!("Available profiles: {:?}", profiles);

    // Parse arguments
    let args: Vec<String> = std::env::args().collect();
    let profile = match args.len() {
        1 => {
            println!("No profile specified, using default");
            "default"
        }
        2 => &args[1],
        _ => panic!("Too many arguments!"),
    };

    // Check that the profile exists
    if !profiles.contains(&&profile.to_string()) {
        panic!("Profile {} not found!", profile);
    }

    // Update the config file
    let mut config = ini!(config_path);

    config

    println!("{:?}", config);
}

