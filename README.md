# Install instructions
```sh
docker build -t lxldapg:1.0 .
```
## Linux
```sh
docker-compose up
```
## Windows (WLS)
```sh
export DISPLAY=192.168.0.102:0.0 # config with your settings of your Xserver for Windows
docker-compose -f docker-compose_win.yml up
```
