# Docker Desktop and Engine Installation

## Ubuntu

### 1. Need rquirement other terminall install

```bash
sudo apt install gnome-terminal
```

### 2. 공식 도커 서명키 발급 및 서명 [Add Offical docker gpg key]

```bash
sudo apt update
sudo apt install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
```

### 3. 공식 도커 리포지토리와 리소스 추가 다운로드 [Add Offical docker repository resource]

```bash
sudo tee /etc/apt/sources.list.d/docker.sources <<EOF
Types: deb
URIs: https://download.docker.com/linux/ubuntu
Suites: $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}")
Components: stable
Architectures: $(dpkg --print-architecture)
Signed-By: /etc/apt/keyrings/docker.asc
EOF
```

#### APT Update

```bash
sudo apt update
```

### 4. 도커 활성화 및 자동 시작

```bash
sudo systemctl status docker
```

```bash
sudo systemctl start docker && sudo systemctl enable docker
```

#### 선택. 일반 사용자 도커 엔진 사용 권한 부여

> [!WARNING]
> 해당 권한을 부여할 시 반드시 이용에 적합한 사용자인지 확인을 해야하고, 이는 시스템에 악영향을 끼칠 수 있다.

```bash
sudo chmod 666 /var/run/docker.sock
```

#### 선택. Docker desktop 설치

```bash
 sudo apt-get update
 curl -sl https://desktop.docker.com/linux/main/amd64/docker-desktop-amd64.deb?utm_source=docker&utm_medium=webreferral&utm_campaign=docs-driven-download-linux-amd64
 sudo apt install ./docker-desktop-amd64.deb
```
**사용중인 도커 엔진 또는 데스크탑 목록 확인**

```bash
docker context ls
```
```bash
NAME              DESCRIPTION                               DOCKER ENDPOINT                                    ERROR
default           Current DOCKER_HOST based configuration   unix:///var/run/docker.sock                        
desktop-linux *   Docker Desktop                            unix:///home/geonwoo/.docker/desktop/docker.sock   
```

**도커 엔진 혹은 데스크탑 사용 전환**
```bash
docker context use default<or desktop-linux>
```
