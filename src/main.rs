#[macro_use]
extern crate ini;

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
    let ini = ini!(credentials_path);

    // List all the profiles
    let profiles = ini.iter().map(|(k, _)| k).collect::<Vec<_>>();
    println!("Available profiles: {:?}", profiles);

}

