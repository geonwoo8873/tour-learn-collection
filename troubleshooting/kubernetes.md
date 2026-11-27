# 로컬 Kubernetes 배포 문제 기록

2026-11-22 Docker Desktop Kubernetes에서 Compose 구성을 옮기면서 확인한
문제와 대응을 정리한다. 명령은 별도 표시가 없으면 저장소 루트의 PowerShell에서
실행한다. 이벤트는 과거 기록을 포함하므로 현재 Pod 상태와 로그를 함께 확인한다.

## 확인된 상태

2026-11-22 20:29경 확인 결과:

- MySQL은 초기화 후 연결을 받을 준비가 되었고, Flyway Job은 `agent_kb`에
  V1과 V2 마이그레이션을 적용했다.
- Qdrant 초기화 Job은 완료되었다. AI와 embedding Pod는 준비 상태가 되었다.
- Vite, ingest, Prometheus, Grafana Pod는 실행 중이었다.
- Spring은 MySQL 복구 후 재시작해 `ocrdb`에 13개 마이그레이션을 적용했고,
  Pod가 준비 상태가 되었다. Deployment 롤링 업데이트도 완료되었다.

Job의 `Completed`와 `0/1` 표시는 작업이 끝난 상태이며, 실행 중인 서비스를
의미하지 않는다.

## 빌드와 이미지 태그

### 상대 경로와 Docker build context

`utils` 디렉터리에서 `..\main-vite-web`을 사용하면 형제 저장소가 아닌
`main-infra-manage\main-vite-web`을 찾게 되어 빌드 컨텍스트 오류가 났다.
[이미지 빌드 스크립트](../utils/image-tag-manage.ps1)는 스크립트 위치를
기준으로 형제 저장소를 찾는다.

Vite 이미지 빌드에서는 `nginx.kubernetes.conf`가 웹 프로젝트의
`.dockerignore`에 포함되어 `COPY`에 실패했다. 해당 설정 파일을 제외 목록에서
제거했으며, Docker의 빌드 검사로 경고가 사라진 것을 확인했다.
`EXPOSE`는 포트 정보를 표시할 뿐 Nginx의 실제 `listen` 포트를 바꾸지 않는다.
Nginx 설정, 컨테이너 포트와 Service의 `targetPort`를 맞춰야 한다.

### 새 태그 반영과 ImagePullBackOff

같은 태그의 이미지만 다시 빌드하면 Deployment의 Pod 템플릿이 바뀌지 않아
자동 롤링 업데이트가 시작되지 않는다. 스크립트는 빌드마다 시간 기반 태그를
사용하며, 모든 이미지 빌드가 성공한 후 해당 매니페스트를 갱신한다.
빌드 실패 시 매니페스트는 갱신하지 않는다.

```powershell
.\utils\image-tag-manage.ps1
```

스크립트는 Kubernetes 적용이나 노드 이미지 로드를 수행하지 않는다. 이미지가
스케줄링된 노드에서 사용 가능한지 확인해야 한다. `ImagePullBackOff`는
이미지 이름·태그 불일치, 노드에 이미지가 없음, 레지스트리 접근 실패 등으로
발생할 수 있다. 클러스터 종류에 맞게 이미지를 로드하거나 레지스트리를 사용한다.

이번에는 이전 태그를 참조한 Pod에서 pull 오류가 있었지만, 이후 태그의 Vite,
AI와 embedding 이미지는 정상적으로 시작되었다. 롤링 업데이트 중에는 새 Pod가
준비될 때까지 이전 ReplicaSet이 남을 수 있다.

```powershell
kubectl get pods -o wide
kubectl describe pod <pod-name>
kubectl get events --sort-by=.lastTimestamp
```

기존 `qdrant-init` Job은 Pod 템플릿을 수정할 수 없어 이미지 변경 후 Job만
재생성했다. 완료된 Job은 매니페스트를 다시 적용해도 자동으로 재실행되지 않는다.
Job 삭제와 PVC 삭제는 별개지만, 삭제 전 대상과 영향을 확인한다.

## Kustomize와 공통 리소스

공통 파일은 `kubernetes\base`에 있다. 상위 디렉터리 파일을 직접 참조하기보다
`base`를 Kustomization 디렉터리로 참조한다. 루트 전체 구성은 `base`를 한 번
포함하고, 서비스별 단독 구성은 `kubernetes\standalone`에서 `base`와 해당
서비스를 조합한다. 같은 리소스를 하나의 렌더링 결과에 중복 포함하지 않는다.

```powershell
kubectl kustomize .\kubernetes
kubectl kustomize .\kubernetes\standalone\mysql
```

렌더링 결과에는 Secret이 포함될 수 있으므로 공유하지 않는다. `apply`는 같은
클러스터와 네임스페이스의 동일한 공통 리소스를 다시 적용한다. 반대로
`kubectl delete -k`는 해당 구성에 포함된 공통 Secret과 PVC도 삭제 대상으로
삼으므로 단독 서비스 테스트를 마친 뒤에도 주의한다.

반복된 `NotFound`는 삭제 시점에 해당 리소스가 없다는 뜻이다. 당시 일부
ConfigMap과 Secret은 `annotations`가 `metadata` 밖에 있어 생성 실패 가능성이
확인되었다. 삭제 오류만으로 이전에 삭제됐는지, 생성되지 않았는지는 확정할 수
없다. `annotations`는 `metadata.annotations` 아래에 작성한다.

```powershell
kubectl config current-context
kubectl get configmap,secret,pvc -A
```

## PVC와 로컬 스토리지

`mysql-data-volume`의 `ReadWriteMany` 요청은 당시 `standard` StorageClass의
`rancher.io/local-path` 프로비저너가 지원하지 않아 프로비저닝에 실패했다.
단일 노드에서 사용하는 구성은 `ReadWriteOnce`로 변경했다. 이는 한 **노드**에서
읽기·쓰기가 가능하다는 뜻이며 반드시 Pod 하나만 마운트한다는 뜻은 아니다.
여러 노드에서 동시 공유해야 하면 스토리지의 `ReadWriteMany` 지원을 확인한다.

`WaitForFirstConsumer`는 Pod가 PVC를 사용하기 전에는 정상적으로 나타날 수 있다.
이미 생성된 PVC의 접근 모드는 일반적으로 변경할 수 없어 재생성이 필요하다.
재생성 전 데이터와 사용 중인 워크로드를 확인한다.

이번에 확인한 로컬 StorageClass의 회수 정책은 `Delete`였다. PVC 삭제 시 실제
데이터도 삭제될 수 있고, 새 PVC 생성만으로 기존 데이터가 복구되지는 않는다.
PVC는 백업을 대체하지 않는다.

```powershell
kubectl get storageclass
kubectl get pvc -A
kubectl get pv
```

## MySQL 기동과 마이그레이션

MySQL StatefulSet의 `args`에서 `--default-time-zone`과
`--character-set-server` 끝에 들어간 불필요한 따옴표로 인해 MySQL이
`utf8mb4"` 문자셋을 찾지 못하고 종료했다. 따옴표를 제거해 StatefulSet 설정을
수정했다.

```yaml
args:
  - mysqld
  - --log-bin-trust-function-creators=1
  - --default-time-zone=+09:00
  - --character-set-server=utf8mb4
  - --collation-server=utf8mb4_unicode_ci
```

수정된 StatefulSet 템플릿이 적용된 뒤에도 문제가 있는 이전 Pod가 남아 있었다.
이번에는 해당 Pod만 일반 삭제해 새 템플릿으로 재생성했고, MySQL이 정상 기동했다.
Pod 강제 삭제나 PVC 삭제는 사용하지 않았다.

```powershell
kubectl get statefulset mysql
kubectl get pod mysql-0
kubectl logs mysql-0
kubectl logs job/mysql-migration-job
```

MySQL 서버의 인자와 Flyway의 `migrate` 명령은 다르다. Flyway Job은
`mysql-service:3306`으로 접속해 SQL을 적용하며, 같은 네임스페이스에서
Service 이름을 사용할 수 있다. 초기화 스크립트는 비어 있는 데이터 디렉터리의
최초 기동에만 자동 실행된다.

Flyway 마이그레이션은 스키마 생성 작업이다. 임베딩 생성 및 Qdrant 적재는 별도
서비스와 ingest 과정이 필요하며, 빈 PVC에 Compose 호스트의 입력 데이터가
자동으로 복사되지는 않는다.

## 준비 상태와 후속 확인

AI와 embedding의 `Running`은 컨테이너가 실행 중이라는 뜻이지 애플리케이션이
준비되었다는 뜻은 아니다. 모델 다운로드와 초기화 중 readiness probe가 실패할
수 있으므로 로그와 `READY`를 확인한다.

Spring은 MySQL이 기동하지 않았을 때 Flyway DB 연결 실패로 종료했지만, MySQL
복구 후 DB 연결과 마이그레이션에 성공했다. 다시 실패하면 최신 로그를 확인한다.
이전 로그가 현재 재시작의 원인이라고 단정하지 않는다.

```powershell
kubectl get pods
kubectl get jobs
kubectl logs deployment/embedding --tail=50
kubectl get pods -l app=spring
kubectl logs <spring-pod-name> --previous --tail=100
```

실제 비밀번호, 토큰, Secret 전체 내용과 로그의 민감 정보는 문서나 외부 공유
자료에 포함하지 않는다.