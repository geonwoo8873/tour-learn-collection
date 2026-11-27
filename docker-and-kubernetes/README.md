# Docker & Kubernetes

# 1. Collection Information

`docker-and-kubernetes` 디렉터리는 Docker와 Kubernetes의 기초 개념부터 컨테이너 이미지 빌드, 배포 운영까지 학습한 내용을 기록하는 공간입니다. 또한 실습에 필요한 Linux 환경 구성, Terraform 기반 인프라 학습, 장애 원인 분석과 해결 과정도 함께 다룹니다.

> [!WARNING]
> 이 디렉터리의 학습 예제에는 계정 정보, API 키, 인증서, 클라우드 자격 증명 또는 기타 민감한 정보가 포함될 수 있습니다. 민감한 값은 파일에 직접 기록하거나 Git에 커밋하지 말고 환경 변수나 비밀 관리 도구를 사용해 주십시오.

# 2. Learning Contents

## Docker

- Docker 설치와 기본 명령어
- 이미지와 컨테이너의 개념
- Dockerfile 작성과 이미지 빌드
- 컨테이너 네트워크와 볼륨
- 컨테이너 레지스트리 활용
- Docker Compose와 멀티 컨테이너 환경
- 빌드와 실행 과정의 문제 해결

## Kubernetes

- Kubernetes 기본 구성 요소와 아키텍처
- Pod, Deployment, Service와 주요 리소스
- ConfigMap과 Secret을 통한 설정 관리
- 워크로드 스케줄링과 컨트롤러
- 상태 확인과 로그 수집
- 클러스터 접근 및 트러블슈팅
- 실습 환경에 애플리케이션 배포

## Linux

- Linux 기본 명령어와 파일 시스템
- 사용자, 권한과 프로세스 관리
- 패키지 설치와 서비스 관리
- 네트워크 확인과 시스템 점검

## Terraform

- Terraform 기본 개념과 실행 흐름
- Provider와 Resource 작성
- 변수, 출력값과 상태 파일 관리
- Docker 또는 클라우드 인프라 구성 자동화

## Troubleshooting

- Docker 빌드 및 실행 오류
- Kubernetes 리소스 및 Pod 장애
- Linux 환경 문제
- 로그 분석과 원인 추적
- 재현 가능한 해결 절차 정리

# 3. Secret Guide

실습 과정에서 비밀번호, 토큰, 키 또는 자격 증명 파일을 사용해야 할 때는 다음 원칙을 따릅니다.

- 예제의 실제 값은 `.env.example` 또는 템플릿 파일에 변수 이름만 기록합니다.
- 실제 비밀 값은 환경 변수 또는 로컬 비밀 관리 도구로 주입합니다.
- 인증서와 자격 증명 파일은 예외적으로 필요한 경우에만 로컬에 보관하고 Git 추적 대상에서 제외합니다.
- 커밋 전 `.gitignore`, 변경 목록과 저장소 검색을 확인하여 민감한 값이 포함되지 않았는지 점검합니다.
- 민감한 값이 노출된 경우 즉시 폐기하거나 재발급하고 관련 이력을 확인합니다.

# 4. WorkTree

```text
docker-and-kubernetes/
├── README.md
├── docker/
│   ├── 000-installation.md
│   └── 001-introduction.md
├── kubernetes/
│   ├── 000-installration.md
│   ├── 001-introduction.md
│   └── 002-controller.md
├── linux/
│   └── README.md
├── terraform/
│   └── README.md
├── troubleshooting/
│   ├── docker.md
│   ├── linux.md
│   └── img/
└── example/
    ├── mysql/
    └── server/
```