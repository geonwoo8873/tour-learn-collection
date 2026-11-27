# Linux Kubernetes Installation

# Ubuntu Install

## 1. kubectl 바이너리 설치 [Kubectl Binary Install]

```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```

<details>

### Option. Other Version Install

```bash
curl -LO https://dl.k8s.io/release/v1.37.0/bin/linux/amd64/kubectl
```

</details>

## 2. (Option) kubectl 바이너리 검증 [Kubectl Binary Verification]

```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
```

### 2.1 Kubectl 바이너리를 체크섬 파일과 대조 및 검증

```bash
echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check
```

<details>

#### Successful

```bash
kubectl: OK
```

#### Fail

```bash
kubectl: FAILED
sha256sum: WARNING: 1 computed checksum did NOT match
```

</details>

## 3. Kubectl 설치 [Kubectl Install]

```bash
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```
<details>

### Option. Root 권한 없는 사용자 설치

```bash
chmod +x kubectl
mkdir -p $PATH/.local/bin
mv ./kubectl $PATH/.local/bin/kubectl
```

</details>

## 3. 기본 패키지 관리 도구 사용하여 설치

### 3.1 apt 패키지 색인을 업데이트하고 쿠버네티스 apt 리포지터리를 사용하는 데 필요한 패키지들을 설치

```bash
sudo apt-get update && \
sudo apt-get install -y apt-transport-https ca-certificates curl gnupg
```

### 3.2 쿠버네티스 패키지 리포지터리용 공개 서명 키를 다운로드

```bash
# sudo mkdir -p -m 755 /etc/apt/keyrings
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.37/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
sudo chmod 644 /etc/apt/keyrings/kubernetes-apt-keyring.gpg
```

### 3.3 쿠버네티스 apt 리포지터리를 추가 및 그외 마이너 버전으로 변경

```bash
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.37/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo chmod 644 /etc/apt/sources.list.d/kubernetes.list
```

### 3.4 apt 패키지 색인을 업데이트하고 kubectl을 설치

```bash
sudo apt-get update && \
sudo apt-get install -y kubectl
```

## 4. Kubectl 구성 확인

### 4.1 Cluster 상태 조회

```bash
kubectl cluster-info
```

#### Option. 클러스터 접근 불가 상태로 구성 검증

```bash
kubectl cluster-info dump
```

## 5. Kuberc 구성

### 5.1 Kuberc convert plugin 설치

```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl-convert"
```

