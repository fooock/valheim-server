# `Valheim` server installer

![](https://img.shields.io/badge/Steam-Valheim-purple?style=for-the-badge&logo=steam)

> :construction: Project in early development

This project contains all required components to deploy a ready to use `Valheim` game server. I made my best effort to had a secure configuration, but I can make mistakes. Review the configuration by yourself.

If you use this software for commercial purposes, please consider making a donation or a horde of bloodthirsty vikings led by Odin will fall on you.

### Requirements

You need to have installed in your system in order to deploy a fully functional game server:

* [`Ansible`](https://www.ansible.com/)
* `make`

>If you are on Windows you can use [`wsl`](https://docs.microsoft.com/windows/wsl/install-win10) in order to install the game server. 

Note that at this time, the game server only supports [`Ubuntu`](https://ubuntu.com/) as base operating system. Choose a Ubuntu compatible cloud provider.

## Installation

>See [using packer](#using-packer) if you want to create server snapshots in your favourite cloud provider.

First, clone this repository. When done, execute the following command using your own values for the `TARGET_USER` and `TARGET_IP` variables that correspond to **your** server info:

```bash
make TARGET_IP=1.2.3.4 TARGET_USER=root ansible-install
```

>You need to provide a `TARGET_IP`. Without this variable the command will not work.

Now you need to wait the command to finish. Be patient and take a coffee :coffee:.

## Configuration

Now that you have your game server installed, you need to change some settings like the gameserver password, world name etc. All this can be configured by editing the `/home/steam/valheim/values.env` file, assuming the default configuration. Default values are:

| Variable      	| Default value       	| Description                              	|
|---------------	|---------------------	|------------------------------------------	|
| `GS_NAME`     	| `My Valheim gameserver` 	| The name of your game server             	|
| `GS_PORT`     	| `2456`              	| Game server port                         	|
| `GS_WORLD`    	| `Dedicated`         	| The world name without spaces             |
| `GS_PASSWORD` 	| `secret`            	| Game server password. Min. length is 5   	|
| `GS_PUBLIC`   	| `1`                 	| For public servers is `1`, otherwise `0` 	|

When done, the only thing that you need to do for the new values to take effect is restart the systemd service of our game server using:

```bash
systemctl restart valheim
```

## Using `Packer`

If you want to use `Packer` to create game server snapshots continue reading. Note that the list of currently supported cloud providers is not complete. If you want to add a new cloud provider, create a [feature request](https://github.com/fooock/valheim-server/issues).

#### [`Hetzner`](https://www.hetzner.com/cloud)

>ID: `hcloud`

To use Hetzner, first you need to generate a new API token with *Read and Write* permissions. Save this token in a local environment variable called `HCLOUD_TOKEN` and then execute this command:

```bash
make CLOUD=hcloud packer-run
```

Now, wait until finish.

#### [`Vagrant`](https://www.vagrantup.com/)

>ID: `vagrant`

You will need to have Vagrant and VirtualBox to use this installation method. Just execute this command and wait until finish:

```bash
make CLOUD=vagrant packer-run
```

The resulting box will be created in `output-vagrant/package.box`

## License

Check [`LICENSE`](LICENSE) for more information.
