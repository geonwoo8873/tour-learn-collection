# Example Go RESTFul API Server Tutorial

>[!WARNING]
> 현재 Windows 11 부터 `Smart App Controle` 보안 기능이 탑재되어 `goplz.exe`가 서명되지 않아 차단 당하는 현상이 존재 합니다. 이는 예외처리를 사전에 진행하고 트레이스가 적용되는지 점검 부탁드립니다. 관련 사례 https://github.com/geonwoo8873/tour-learn-collection/issues/18

해당 디렉터리는 Go의 Restful API 서버로 구성되어 있습니다. Kubernetes 예제로 사용되며 여러 파드에서 활성화 되는 과정을 확인 위함으로 코드와 명세서가 작성되니 참고 부탁드립니다.

>[!IMPORTANT]
> Go Install은 우선 앞서 도커 컨테이너화 하기전 서버를 시작하여 확인하고 싶은 경우에만 설치 진행 해주시면 됩니다.

#### Docker Image Build

```bash
docker build . -t server:v1.0.1 --no-cache
```

```bash
docker buildx build . -t server:v1.0.1 --no-cache
```

> [!TIP]
> `buildx`는 Docker BuildKit을 기반으로 동작하는 빌드 도구로, 기존 docker build 기능을 확장하여 멀티 플랫폼 이미지 생성, 원격 캐시 활용, 병렬 빌드 및 고급 빌드 기능을 제공한다.

#### Connecting Test

```bash
curl localhost:8080
```