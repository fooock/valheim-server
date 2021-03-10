# `Valheim` server installer

![](https://img.shields.io/badge/Steam-Valheim-purple?style=for-the-badge&logo=steam)
[![Discord](https://img.shields.io/discord/812069608064417842?color=purple&logo=discord&logoColor=white&style=for-the-badge)](https://discord.gg/X9cN2cawgW)

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

First, clone this repository. When done, execute the following command using your own values for the `TARGET_USER` and `TARGET_IP` variables that correspond to **your** server info:

```sh
make TARGET_IP=1.2.3.4 TARGET_USER=root ansible-install
```

>You need to provide a `TARGET_IP`. Without this variable the command will not work.

Now you need to wait the command to finish. Be patient and take a coffee :coffee:.
