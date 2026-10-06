# 1. Storage

## 1.1 Disk Size

```bash
df -<options>
```

```bash
Filesystem      1K-blocks     Used  Available Use% Mounted on
none              7994668        0    7994668   0% /usr/lib/modules/6.18.40.1-microsoft-standard-WSL2
none              7994668        4    7994664   1% /mnt/wsl
drivers         498975740 88742988  410232752  18% /usr/lib/wsl/drivers
/dev/sdd       1055762868  1360332 1000699064   1% /
none              7994668       36    7994632   1% /mnt/wslg
none              7994668        0    7994668   0% /usr/lib/wsl/lib
rootfs            7990236     3352    7986884   1% /init
none              7994668      536    7994132   1% /run
none              7994668        0    7994668   0% /run/lock
none              7994668        0    7994668   0% /run/shm
none              7994668      100    7994568   1% /mnt/wslg/versions.txt
none              7994668      100    7994568   1% /mnt/wslg/doc
C:\             498975740 88742988  410232752  18% /mnt/c
none                 1024        0       1024   0% /run/credentials/systemd-journald.service
tmpfs             7994668        0    7994668   0% /tmp
none                 1024        0       1024   0% /run/credentials/systemd-resolved.service
none                 1024        0       1024   0% /run/credentials/getty@tty1.service
tmpfs             1598932       12    1598920   1% /run/user/1000
```

<details>
<summary>Disk size output type</summary>

```bash
df -g
```

</details>