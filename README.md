# myip

Get your external IP address from command line.

[![Build status](https://dev.azure.com/nekomimiswitch/General/_apis/build/status/myip)](https://dev.azure.com/nekomimiswitch/General/_build/latest?definitionId=35)

## Usage

Command line argument `-4`/`-6` can be used to set IPv4/IPv6 preference with some methods.

### STUN

* This is the default method
* `stun.l.google.com:19302` is the default server
* Connection over UDP only

```shell
myip --method STUN --server stun.l.google.com:19302
```

### ip.sb HTTPS API

```shell
myip --method ip.sb
```

### OpenDNS DNS Query

```shell
myip --method OpenDNS
```

### OpenDNS HTTPS API

* `-4`/`-6` doesn't work

```shell
myip --method OpenDNS-API
```

### Donation

If this project is helpful to you, please consider buying me a coffee.

[![Buy Me A Coffee](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/Jamesits) or [PayPal](https://paypal.me/Jamesits)