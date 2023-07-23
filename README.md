# Notify-list
Accomplish of notifications for list of the tasks in **cron** format

Describe your tasks on **$HOME/.config/notify-list/list.json**

Or simply run:
```bash
notify-list -list example-list.json
```

For autostart provides systemd user unit.  
Also possible use *.xinitrc*  or *dex*.  

#### Dependency
Install **notify-send** on linux 

### Installation
```bash
make check
make build
make install
```

### Configuration 
```bash 
mkdir $HOME/.config/notify-list
cp example-list.json $HOME/.config/notify-list/list.json
```
And edit you're tasks

### Remove
```bash
make clean
make uninstall
```
