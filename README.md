# wp

## Server-side setup

For the most basic TCP tests you can use something like Nginx.
```
# apt-get install nginx
# iptables -A INPUT -p tcp --dport 22 -j ACCEPT
# iptables -A INPUT -p tcp --dport 80 -j ACCEPT
# iptables -t nat -A PREROUTING -p tcp -m multiport ! --dports 22,80 -j REDIRECT --to-port 80
```

For UDP, use NTP.
```
# apt-get install ntp
# iptables -A INPUT -p tcp --dport 22 -j ACCEPT
# iptables -A INPUT -p udp --dport 123 -j ACCEPT
# iptables -t nat -A PREROUTING -p udp -m multiport ! --dports 22,123 -j REDIRECT --to-port 123
```

## Checks
I started working on this a long time ago and never finished it. I have no idea what is implemented at this point. I'll update as I finish it up.
- TCP/UDP ports
- TCP Application Firewall Detection
  - Send application specific traffic over exposed TCP services
- ICMP ports
- TLS
  - SNI Spoofing
- HTTP
  - Host header spoofing
  - Cloud services
    - EC2
    - S3
    - GKE
    - Azure
- DLP
  - Nefarious data samples
  - Data size limits

