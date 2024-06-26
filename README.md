# go-ntpd-server-statistics
Obtains the amount of UNIQUE public IPs that access our self-hosted NTPD server during a interval of seconds.

## HOW
During an interval of time takes the IP public address from every packet arriving to our server (ntp protocol, port 123/UDP). 

Only count **unique IPs** and not amount of packets.

## USE
Requires install LibPcap development libs:

```
apt install libpcap-dev
```

**IMPORTANT:** Modify config.yml to match your interface and public ip to listen. 

Just run the thing:

```
go run main.go
````

![imagen](https://github.com/nireitdev/go-ntpd-server-statistics/assets/85206635/bc78f046-63b2-4452-a4d4-6a2f518a71bd)

