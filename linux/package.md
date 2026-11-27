# Package [pkg]

```bash
sudo apt-get update && sudo apt-get install -y
```


## 사용중인 패키지 의존성 검색 [Used package dependency search]

```bash
which <package_name>
dpkg -<options> $(which package_name)
```

<details>

```bash
geonwoo@WORK-DESKTOP:~$ which nslookup && dpkg -S $(which nslookup)
```

```bash
/usr/bin/nslookup
bind9-dnsutils: /usr/bin/nslookup
```

</details>