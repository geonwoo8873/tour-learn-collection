# OpenSSH Install

> ![NOTE]
> **현재 작성된 문서는 root 기준입니다. Vmware이 아닌 외부 환경에서는 절대 root에서 작업하면 안되며, 일절 권한의 악용을 지양합니다.**

## 1. System package update

```sh
sudo apt update && sudo apt upgarde
```

```sh
sudo apt install openssh-server
```
### 1.1 Vitual OS Connect

```sh
ssh <username>@<ipaddress>
```

### 1.2 SSH Status check

```sh
root@test-local:/# systemctl status ssh
```

```sh
Warning: The unit file, source configuration file or drop-ins of ssh.service changed on disk. Run 'systemctl daemon-reload' to reload units.
● ssh.service - OpenBSD Secure Shell server
     Loaded: loaded (/usr/lib/systemd/system/ssh.service; enabled; preset: enabled)
     Active: active (running) since Thu 2026-09-10 06:46:10 UTC; 5min ago <= SSH connect feature active status
 Invocation: 5cc333da36464b5***************
TriggeredBy: ● ssh.socket
       Docs: man:sshd(8)
             man:sshd_config(5)
    Process: 2829 ExecStartPre=/usr/sbin/sshd -t (code=exited, status=0/SUCCESS)
   Main PID: 2831 (sshd)
      Tasks: 1 (limit: 3735)
     Memory: 7.5M (peak: 9.4M)
        CPU: 92ms
     CGroup: /system.slice/ssh.service
             └─2831 "sshd: /usr/sbin/sshd -D [listener] 0 of 10-100 startups"

Sep 10 06:46:10 test-local systemd[1]: Starting ssh.service - OpenBSD Secure Shell server...
Sep 10 06:46:10 test-local sshd[2831]: Server listening on 0.0.0.0 port 22.
Sep 10 06:46:10 test-local sshd[2831]: Server listening on :: port 22.
Sep 10 06:46:10 test-local systemd[1]: Started ssh.service - OpenBSD Secure Shell server.
Sep 10 06:46:16 test-local sshd-session[2834]: Accepted password for geonwoo from 000.000.000.0 port 00000 ssh2
Sep 10 06:46:16 test-local sshd-session[2834]: pam_unix(sshd:session): session opened for user geonwoo(uid=1000) by geonwoo(uid=0)
Sep 10 06:46:40 test-local sshd-session[2919]: Accepted password for geonwoo from 000.000.000.0 port 00000 ssh2
Sep 10 06:46:40 test-local sshd-session[2919]: pam_unix(sshd:session): session opened for user geonwoo(uid=1000) by geonwoo(uid=0)
```

### 1.3 SSH Service enable

#### Disabled status service SSH

```sh
sudo systemctl start ssh
```

#### Auto start service SSH

```sh
sudo systemctl enable ssh
```

## 2. Firewall

### 2.1 Firewall status check

```sh
root@test-local:~# ufw status

Status: inactive
```

### 2.2 Firewall feature disable

```sh

```