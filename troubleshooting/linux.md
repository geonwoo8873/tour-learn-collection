### 1. [L-001] Ubuntu note installed

```bash
Installing, this may take a few minutes...
WslRegisterDistribution failed with error: 0x80370114
Error: 0x80370114 The operation could not be started because a required feature is not installed.

Press any key to continue...
```

[L-001](#1-l-001-ubuntu-note-installed)의 문제 발생은 Windows에서 제공하는 `Windows subsystem for linux` 활성화가 안되어 있음에 따라 **ubuntu는 이 기능을 자동으로 활성화 시킬 수 없기 때문에 발생한 오류**로 사용자는 해당 기능을 활성화 시켜주면된다. 다만 주의해야할 점은 일부 디바이스에서는 `Hyper-V`나 `Windows subsystem for linux`를 활성화 시키면 Kernel- `BSOD`가 발생할 수 있으니 주의해야 한다.

### 2. [L-002] bin9-dnstools

<details>

```bash
geonwoo@WORK-DESKTOP:/$ sudo apt install bind9-dnsutils -y
```

```bash
Installing:
  bind9-dnsutils

Installing dependencies:
  bind9-host  bind9-libs  liblmdb0  libmaxminddb0  liburcu8t64  libuv1t64

Suggested packages:
  mmdb-bin

Summary:
  Upgrading: 0, Installing: 7, Removing: 0, Not Upgrading: 0
  Download size: 1525 kB / 1761 kB
  Space needed: 5092 kB / 1025 GB available

Ign:1 http://archive.ubuntu.com/ubuntu resolute-updates/main amd64 bind9-libs amd64 1:9.20.24-1ubuntu0.2
Ign:2 http://archive.ubuntu.com/ubuntu resolute-updates/main amd64 bind9-host amd64 1:9.20.24-1ubuntu0.2
Ign:3 http://archive.ubuntu.com/ubuntu resolute-updates/main amd64 bind9-dnsutils amd64 1:9.20.24-1ubuntu0.2
Err:1 http://security.ubuntu.com/ubuntu resolute-updates/main amd64 bind9-libs amd64 1:9.20.24-1ubuntu0.2
  404  Not Found [IP: 91.189.91.83 80]
  404  Not Found [IP: 91.189.92.23 80]
Err:2 http://security.ubuntu.com/ubuntu resolute-updates/main amd64 bind9-host amd64 1:9.20.24-1ubuntu0.2
  404  Not Found [IP: 91.189.91.83 80]
  404  Not Found [IP: 91.189.92.23 80]
Err:3 http://security.ubuntu.com/ubuntu resolute-updates/main amd64 bind9-dnsutils amd64 1:9.20.24-1ubuntu0.2
  404  Not Found [IP: 91.189.91.83 80]
  404  Not Found [IP: 91.189.92.23 80]
Error: Failed to fetch http://security.ubuntu.com/ubuntu/pool/main/b/bind9/bind9-libs_9.20.24-1ubuntu0.2_amd64.deb  404  Not Found [IP: 91.189.92.23 80]
Error: Failed to fetch http://security.ubuntu.com/ubuntu/pool/main/b/bind9/bind9-host_9.20.24-1ubuntu0.2_amd64.deb  404  Not Found [IP: 91.189.92.23 80]
Error: Failed to fetch http://security.ubuntu.com/ubuntu/pool/main/b/bind9/bind9-dnsutils_9.20.24-1ubuntu0.2_amd64.deb  404  Not Found [IP: 91.189.92.23 80]
Error: Unable to fetch some archives, maybe run apt update or try with --fix-missing?
```

</details>

#### Solution

```bash
geonwoo@WORK-DESKTOP:/$ sudo apt update && sudo apt install bind9-dnsutils -y
```


**Result Output**

<details>

```bash
Hit:1 http://archive.ubuntu.com/ubuntu resolute InRelease
Get:2 http://archive.ubuntu.com/ubuntu resolute-updates InRelease [137 kB]
Get:3 http://security.ubuntu.com/ubuntu resolute-security InRelease [137 kB]
Get:4 http://archive.ubuntu.com/ubuntu resolute-backports InRelease [137 kB]
Get:5 http://security.ubuntu.com/ubuntu resolute-security/main amd64 Packages [561 kB]
Get:6 http://archive.ubuntu.com/ubuntu resolute/universe amd64 Packages [16.0 MB]
Get:7 http://security.ubuntu.com/ubuntu resolute-security/main Translation-en [132 kB]
Get:8 http://security.ubuntu.com/ubuntu resolute-security/main amd64 Components [46.7 kB]
Get:9 http://security.ubuntu.com/ubuntu resolute-security/universe amd64 Packages [198 kB]
Get:10 http://security.ubuntu.com/ubuntu resolute-security/universe Translation-en [62.6 kB]
Get:11 http://security.ubuntu.com/ubuntu resolute-security/universe amd64 Components [55.7 kB]
Get:12 http://security.ubuntu.com/ubuntu resolute-security/universe amd64 c-n-f Metadata [3552 B]
Get:13 http://security.ubuntu.com/ubuntu resolute-security/restricted amd64 Packages [445 kB]
Get:14 http://security.ubuntu.com/ubuntu resolute-security/restricted Translation-en [89.7 kB]
Get:15 http://security.ubuntu.com/ubuntu resolute-security/multiverse amd64 Packages [10.8 kB]
Get:16 http://security.ubuntu.com/ubuntu resolute-security/multiverse Translation-en [2844 B]
Get:17 http://security.ubuntu.com/ubuntu resolute-security/multiverse amd64 Components [212 B]
Get:18 http://security.ubuntu.com/ubuntu resolute-security/multiverse amd64 c-n-f Metadata [120 B]
Get:19 http://archive.ubuntu.com/ubuntu resolute/universe Translation-en [6329 kB]
Get:20 http://archive.ubuntu.com/ubuntu resolute/universe amd64 Components [4556 kB]
Get:21 http://archive.ubuntu.com/ubuntu resolute/universe amd64 c-n-f Metadata [313 kB]
Get:22 http://archive.ubuntu.com/ubuntu resolute/multiverse amd64 Packages [290 kB]
Get:23 http://archive.ubuntu.com/ubuntu resolute/multiverse Translation-en [127 kB]
Get:24 http://archive.ubuntu.com/ubuntu resolute/multiverse amd64 Components [50.0 kB]
Get:25 http://archive.ubuntu.com/ubuntu resolute/multiverse amd64 c-n-f Metadata [8276 B]
Get:26 http://archive.ubuntu.com/ubuntu resolute-updates/main amd64 Packages [728 kB]
Get:27 http://archive.ubuntu.com/ubuntu resolute-updates/main Translation-en [170 kB]
Get:28 http://archive.ubuntu.com/ubuntu resolute-updates/main amd64 Components [98.4 kB]
Get:29 http://archive.ubuntu.com/ubuntu resolute-updates/universe amd64 Packages [320 kB]
Get:30 http://archive.ubuntu.com/ubuntu resolute-updates/universe Translation-en [103 kB]
Get:31 http://archive.ubuntu.com/ubuntu resolute-updates/universe amd64 Components [201 kB]
Get:32 http://archive.ubuntu.com/ubuntu resolute-updates/universe amd64 c-n-f Metadata [4704 B]
Get:33 http://archive.ubuntu.com/ubuntu resolute-updates/restricted amd64 Packages [490 kB]
Get:34 http://archive.ubuntu.com/ubuntu resolute-updates/restricted Translation-en [97.7 kB]
Get:35 http://archive.ubuntu.com/ubuntu resolute-updates/multiverse amd64 Packages [11.9 kB]
Get:36 http://archive.ubuntu.com/ubuntu resolute-updates/multiverse Translation-en [3276 B]
Get:37 http://archive.ubuntu.com/ubuntu resolute-updates/multiverse amd64 Components [216 B]
Get:38 http://archive.ubuntu.com/ubuntu resolute-updates/multiverse amd64 c-n-f Metadata [256 B]
Get:39 http://archive.ubuntu.com/ubuntu resolute-backports/main amd64 Components [212 B]
Get:40 http://archive.ubuntu.com/ubuntu resolute-backports/main amd64 c-n-f Metadata [112 B]
Get:41 http://archive.ubuntu.com/ubuntu resolute-backports/universe amd64 Packages [3136 B]
Get:42 http://archive.ubuntu.com/ubuntu resolute-backports/universe Translation-en [7988 B]
Get:43 http://archive.ubuntu.com/ubuntu resolute-backports/universe amd64 Components [1060 B]
Get:44 http://archive.ubuntu.com/ubuntu resolute-backports/universe amd64 c-n-f Metadata [116 B]
Get:45 http://archive.ubuntu.com/ubuntu resolute-backports/restricted amd64 Components [216 B]
Get:46 http://archive.ubuntu.com/ubuntu resolute-backports/restricted amd64 c-n-f Metadata [120 B]
Get:47 http://archive.ubuntu.com/ubuntu resolute-backports/multiverse amd64 Components [216 B]
Get:48 http://archive.ubuntu.com/ubuntu resolute-backports/multiverse amd64 c-n-f Metadata [120 B]
Fetched 31.9 MB in 15s (2061 kB/s)
98 packages can be upgraded. Run 'apt list --upgradable' to see them.
Installing:
  bind9-dnsutils

Installing dependencies:
  bind9-host  bind9-libs  liblmdb0  libmaxminddb0  liburcu8t64  libuv1t64

Suggested packages:
  mmdb-bin

Summary:
  Upgrading: 0, Installing: 7, Removing: 0, Not Upgrading: 98
  Download size: 1526 kB / 1762 kB
  Space needed: 5092 kB / 1024 GB available

Continue? [Y/n] y
Get:1 http://archive.ubuntu.com/ubuntu resolute-updates/main amd64 bind9-libs amd64 1:9.20.24-1ubuntu0.3 [1313 kB]
Get:2 http://archive.ubuntu.com/ubuntu resolute-updates/main amd64 bind9-host amd64 1:9.20.24-1ubuntu0.3 [49.7 kB]
Get:3 http://archive.ubuntu.com/ubuntu resolute-updates/main amd64 bind9-dnsutils amd64 1:9.20.24-1ubuntu0.3 [164 kB]
Fetched 1526 kB in 2s (685 kB/s)
debconf: unable to initialize frontend: Dialog
debconf: (Dialog frontend requires a screen at least 13 lines tall and 31 columns wide.)
debconf: falling back to frontend: Readline
Selecting previously unselected package liblmdb0:amd64.
(Reading database ... 35936 files and directories currently installed.)
Preparing to unpack .../0-liblmdb0_0.9.31-1build2_amd64.deb ...
Unpacking liblmdb0:amd64 (0.9.31-1build2) ...
Selecting previously unselected package libmaxminddb0:amd64.
Preparing to unpack .../1-libmaxminddb0_1.12.2-1build2_amd64.deb ...
Unpacking libmaxminddb0:amd64 (1.12.2-1build2) ...
Selecting previously unselected package liburcu8t64:amd64.
Preparing to unpack .../2-liburcu8t64_0.15.6-1_amd64.deb ...
Unpacking liburcu8t64:amd64 (0.15.6-1) ...
Selecting previously unselected package libuv1t64:amd64.
Preparing to unpack .../3-libuv1t64_1.51.0-2ubuntu1_amd64.deb ...
Unpacking libuv1t64:amd64 (1.51.0-2ubuntu1) ...
Selecting previously unselected package bind9-libs:amd64.
Preparing to unpack .../4-bind9-libs_1%3a9.20.24-1ubuntu0.3_amd64.deb ...
Unpacking bind9-libs:amd64 (1:9.20.24-1ubuntu0.3) ...
Selecting previously unselected package bind9-host.
Preparing to unpack .../5-bind9-host_1%3a9.20.24-1ubuntu0.3_amd64.deb ...
Unpacking bind9-host (1:9.20.24-1ubuntu0.3) ...
Selecting previously unselected package bind9-dnsutils.
Preparing to unpack .../6-bind9-dnsutils_1%3a9.20.24-1ubuntu0.3_amd64.deb ...
Unpacking bind9-dnsutils (1:9.20.24-1ubuntu0.3) ...
Setting up liblmdb0:amd64 (0.9.31-1build2) ...
Setting up liburcu8t64:amd64 (0.15.6-1) ...
Setting up libmaxminddb0:amd64 (1.12.2-1build2) ...
Setting up libuv1t64:amd64 (1.51.0-2ubuntu1) ...
Setting up bind9-libs:amd64 (1:9.20.24-1ubuntu0.3) ...
Setting up bind9-host (1:9.20.24-1ubuntu0.3) ...
Setting up bind9-dnsutils (1:9.20.24-1ubuntu0.3) ...
Processing triggers for man-db (2.13.1-1build1) ...
Processing triggers for libc-bin (2.43-2ubuntu2.3) ...
```

</details>

[L-002](#2-l-002)의 문제는 Windows 환경의 WSL (`Ubuntu`)에서 Package (`pkg`) 설치 시 404 Not Found가 발생한것으로, 설치 로그 분석 결과, APT가 `bind9-dnsutils resolute-updates/main amd64 bind9-libs amd64 1:9.20.24-1ubuntu0.2` 버전으로 요청되어 해당 버전이 존재하지 않아 apt update를 진행이 완료되더라도 `1:9.20.24-1ubuntu0.3` 버전으로 설치 되지 않았다. 재시도 결과 정상 버전으로 설치는 완료 되었으나, 본 이슈는 네트워크나 DNS 문제보다 갱신되지 않았던 `APT Metadata`와 실제 `Ubuntu Repository`의 **Package Version간 불일치**로 판단된다.