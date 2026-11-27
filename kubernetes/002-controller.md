# 1. Controller 객체 [Controller object]

컨트롤러는 쿠버네티스 API를 통해 리소스의 현재 상태를 관찰하고, 사용자가 선언한 원하는 상태와 차이가 있으면 이를 지속적으로 조정하는 제어 루프다. 예를 들어 Deployment 컨트롤러는 ReplicaSet을 관리하고, ReplicaSet 컨트롤러는 지정된 수의 Pod를 유지한다. 따라서 Deployment가 Pod를 직접 관리하는 것은 아니다.

> [!NOTE]
> 단독으로 생성한 Pod는 노드 장애 등으로 사라져도 이를 다시 생성할 상위 컨트롤러가 없다. 반면 Deployment가 관리하는 Pod는 ReplicaSet이 원하는 Pod 수를 유지하므로 대체 Pod가 생성될 수 있다. 다만 여러 복제본을 만들었다고 서로 다른 노드에 배치되는 것이 자동으로 보장되지는 않는다. 필요하면 Pod anti-affinity나 topology spread constraints 등을 설정해야 한다.

앞서 [001-introduction.md #12-deployment-생성](001-introduction.md/#12-deployment-생성)의 명령어를 입력하면 Deployment와 그에 속한 ReplicaSet, Pod가 생성된다. Pod 이름에는 무작위 문자열이 포함되므로 `kubectl get pods`나 `kubectl get deployments`로 생성 결과를 확인한다. 레이블은 리소스에 선택적으로 붙이는 키-값 메타데이터이며, 셀렉터를 이용해 리소스를 조회하거나 컨트롤러가 관리 대상을 식별하는 데 사용한다.

## 1.1 JSON 포맷으로 키-값 조회

```bash
kubectl get deploy nginx-app -o jsonpath='{.spec.template.metadata.labels}'
```

```bash
{"app":"nginx-app"}
```

## 1.2 JSON 키-값으로 파드 조회

```bash
kubectl get pods -l app=nginx-app
```

```bash
NAME                         READY   STATUS    RESTARTS   AGE
nginx-app-5cf778f7bb-zhn8g   1/1     Running   0          5m56s
```

## 1.3 사용중인 Deployment YAML 기술된 내용 조회

```bash
kubectl get deploy nginx-app -o yaml     
```

```bash                                  
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"
  creationTimestamp: "2026-11-11T16:17:02Z"
  generation: 1
  labels:
    app: nginx-app
  name: nginx-app
  namespace: default
  resourceVersion: "49133"
  uid: ba8b567d-0387-47e4-bda7-fc5f5bdafc4b
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: nginx-app
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: nginx-app
    spec:
      containers:
      - image: nginx
        imagePullPolicy: Always
        name: nginx
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
status:
  availableReplicas: 1
  conditions:
  - lastTransitionTime: "2026-11-11T16:17:04Z"
    lastUpdateTime: "2026-11-11T16:17:04Z"
    message: Deployment has minimum availability.
    reason: MinimumReplicasAvailable
    status: "True"
    type: Available
  - lastTransitionTime: "2026-11-11T16:17:02Z"
    lastUpdateTime: "2026-11-11T16:17:04Z"
    message: ReplicaSet "nginx-app-5cf778f7bb" has successfully progressed.
    reason: NewReplicaSetAvailable
    status: "True"
    type: Progressing
  observedGeneration: 1
  readyReplicas: 1
  replicas: 1
  terminatingReplicas: 0
  updatedReplicas: 1
```

레이블은 리소스 간의 관계를 찾고 대상을 선택하는 데 자주 사용된다. 이 예제에서는 Deployment가 ReplicaSet을 관리하고, ReplicaSet의 셀렉터가 자신이 관리할 Pod를 식별한다. 따라서 레이블 셀렉터를 사용하면 컨트롤러가 관리하는 리소스를 Kubernetes API를 통해 조회할 수 있다.

> ReplicaSet이 관리하는 Pod의 레이블을 셀렉터와 다르게 바꾸면 해당 Pod는 더 이상 셀렉터에 일치하지 않는다. ReplicaSet은 이를 원하는 복제본 수에 포함하지 않으므로 새 Pod를 만들 수 있다.

## 1.4 디플로이먼트가 생성한 파드의 레이블 수정

```bash
kubectl label pods -l app=nginx-app --overwrite app=nginx
```

```bash
kubectl get pods --show-labels
```

```bash
NAME                         READY   STATUS    RESTARTS      AGE    LABELS
nginx-app-5cf778f7bb-9fwdv   1/1     Running   0             63s    app=nginx
nginx-app-5cf778f7bb-zhn8g   1/1     Running   2 (18h ago)   3d9h   app=nginx-app
```

> [!TIP]
> `get pods`는 기본적으로 NAME, READY, STATUS, RESTARTS, AGE를 보여준다. `-o custom-columns`를 사용하면 레이블 같은 메타데이터도 함께 출력할 수 있다.
> ```bash
> kubectl get pods -l app -o custom-columns=NAME:.metadata.name,LABELS:.metadata.labels
> ```

레이블을 바꾼 뒤 `get pods`를 조회하면 새 Pod가 생성된 것을 볼 수 있다. Pod 이름 때문이 아니라, 기존 Pod의 레이블이 ReplicaSet 셀렉터와 일치하지 않아 원하는 복제본 수에 포함되지 않았기 때문이다. 레이블이 바뀐 기존 Pod는 컨트롤러의 관리 대상에서 벗어날 수 있으므로, 이 방법은 동작을 관찰하는 실습용이며 운영 환경에서 권장하는 변경 방식은 아니다.

```bash
# 원래의 레이블로 수정
kubectl label pods -l app=nginx --overwrite app=nginx-app
```

```bash
NAME                         READY   STATUS    RESTARTS      AGE
nginx-app-5cf778f7bb-zhn8g   1/1     Running   2 (18h ago)   3d9h
```

레이블을 되돌리면 ReplicaSet이 다시 해당 Pod를 관리 대상으로 인식할 수 있다. 관리 대상 Pod 수가 `replicas`보다 많아지면 ReplicaSet은 원하는 수를 맞추기 위해 초과 Pod를 정리한다. 어떤 Pod가 남을지는 컨트롤러의 선택에 달려 있으므로 특정 Pod가 반드시 유지된다고 가정하지 않는다.

# 2. 애플리케이션 매니페스트에 배포 정의하기 [Defining Deployment in an Application Manifest]

애플리케이션 매니페스트는 쿠버네티스의 핵심 기능으로 기본 개념의 접근성이 다소 떨어지며, 배포 문제의 원인을 찾기 위해 수 많은 파일과 기술된 내용을들 찾아야 하기 때문에 정작 설계를 하고 배포하는 사용자조차 이해하기 어려운 영역이다. 하지만 YAML에 기술하는 내용들은 결국 사람의 개입을 최소화하고 일관된 서비스 제공과 자동관리 등 다양한 이점을 제공하기에 오히려 많은 기업과 엔지니어들이 사용한다.

Kubernetes API는 JSON 형식으로 리소스를 주고받으며, `kubectl`은 YAML 매니페스트도 입력으로 받을 수 있다. YAML은 주석을 작성할 수 있고 한 파일에 여러 리소스를 `---`로 구분해 정의할 수 있어 널리 사용된다.

> [!WARNING]
> 아래 예제는 비밀 환경 변수를 노출 시키지 않기 위해 YOUR...와 같은 임시 예제로 대체 했다. 매니페스트를 적용하기 전에 `mysql-secret`을 별도로 준비해야 하며, 필요한 키는 예제 아래에서 설명한다.

```yaml
##########################################
# MySQL Kubernetes Architecture Layout
# Configuration
#   -> Public [ConfigMap]
#   -> Credentials [Secret created separately]
# Volume Mounts
#   -> PersistentVolumeClaim [mypvc]
# Client Service
# Client StatefulSet
##########################################

apiVersion: v1
kind: ConfigMap
metadata:
  name: mysql-config
  labels:
    app: mysql-config

data:
  MYSQL_DATABASE: mydatabase
  MYSQL_USER: myuser
  TZ: Asia/Seoul

---
apiVersion: v1
kind: Secret
metadata:
  name: mysql-secret
  labels:
    app: mysql-secret
type: Opaque

stringData:
  MYSQL_PASSWORD: "YOUR_MYSQL_PASSWORD"
  MYSQL_ROOT_PASSWORD: "YOUR_MYSQL_ROOT_PASSWORD"
  
---
#### PersistentVolumeClaim [PVC]
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mypvc
spec:
  resources:
    requests:
      storage: 250Mi
  volumeMode: Filesystem
  accessModes:
    - ReadWriteOnce
  
---
apiVersion: v1
kind: Service
metadata:
  name: mysql
  labels:
    app: mysql

spec:
  clusterIP: None
  ports:
    - port: 3306
      targetPort: 3306
  selector:
    app: mysql

---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql
  labels:
    app: mysql

spec:
  serviceName: "mysql"
  replicas: 1
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
      - name: mysql
        image: mysql:8-oraclelinux9
        envFrom:
        - configMapRef:
            name: mysql-config
        - secretRef:
            name: mysql-secret
        volumeMounts:
        - name: mypvc
          mountPath: /var/lib/mysql
        ports:
        - containerPort: 3306
      volumes:
        - name: mypvc
          persistentVolumeClaim:
            claimName: mypvc
```

이 예제에서 `mysql-secret`은 민감한 값을 저장소에 넣지 않도록 매니페스트에서 제외했다. 적용 전에 `mysql-secret`을 별도로 생성해야 하며, `MYSQL_PASSWORD`와 `MYSQL_ROOT_PASSWORD` 키가 필요하다. Secret 값은 저장소에 커밋하지 않는다. Kubernetes Secret의 값은 기본적으로 base64 인코딩일 뿐이므로, 접근 권한과 저장 시 암호화 설정도 별도로 고려해야 한다.

새로운 파드를 kubectl을 사용하여 다른 파드와 마찬가지로 다룰 수 잇고 세부 정보를 보거나 네트워크 트래픽을 특정 파드로 전달하는 것도 가능하다. 하지만 매니페스트는 애플리케이션 정의를 공유하기 쉽고, 동일한 배포를 반복할 수 있는 점에서 차이가 존재하지만 결국 명령어를 사용해야 하는 시점은 있기 때문에 앞으로도 명령어와 로그 및 모니터링에 대한 보는 방법을 터득하면 향후 운영 관점에서 상대적인 쉬운 식별이 가능해진다.

매니페스트의 필드는 리소스 종류(`kind`)에 따라 달라진다. 아래는 공통 구조와 Deployment에서 자주 만나는 필드다.

| 필드                                              | 역할                                           | 알아둘 점                                                                     |
| ------------------------------------------------- | ---------------------------------------------- | ----------------------------------------------------------------------------- |
| `apiVersion`                                      | 사용할 Kubernetes API 그룹과 버전              | 리소스 종류에 맞는 API 버전을 사용한다.                                       |
| `kind`                                            | 생성할 리소스 종류                             | 예: `Deployment`, `Service`, `ConfigMap`                                      |
| `metadata.name`                                   | 리소스 이름                                    | 같은 namespace 안에서 고유해야 한다.                                          |
| `metadata.namespace`                              | 리소스가 속할 namespace                        | 생략하면 현재 namespace가 사용된다. 클러스터 범위 리소스에는 적용되지 않는다. |
| `metadata.labels`                                 | 리소스를 선택하고 분류하기 위한 키-값          | Service와 컨트롤러의 셀렉터가 이 레이블을 기준으로 대상을 찾는다.             |
| `metadata.annotations`                            | 도구나 사용자용 추가 메타데이터                | 레이블과 달리 일반적으로 리소스 선택에는 사용하지 않는다.                     |
| `spec`                                            | 리소스의 원하는 상태                           | 세부 구조는 `kind`마다 다르다.                                                |
| `spec.selector`                                   | Deployment가 관리할 Pod를 선택하는 조건        | `template.metadata.labels`와 일치해야 하며, 생성 후 변경할 수 없다.           |
| `spec.template`                                   | 컨트롤러가 생성할 Pod의 템플릿                 | Pod의 `metadata`와 `spec`을 포함한다. 모든 리소스에 있는 필드는 아니다.       |
| `spec.replicas`                                   | 유지할 Pod 개수                                | Deployment 등의 복제 가능한 워크로드에서 사용한다.                            |
| `spec.strategy`                                   | Deployment 업데이트 방식                       | 기본적으로 롤링 업데이트를 사용한다.                                          |
| `spec.template.spec.containers`                   | Pod에서 실행할 컨테이너 목록                   | 각 컨테이너의 이미지, 환경 변수, 포트 등을 지정한다.                          |
| `resources`                                       | 컨테이너의 CPU·메모리 요청량과 제한량          | 스케줄링과 자원 사용 제어에 영향을 준다.                                      |
| `readinessProbe`, `livenessProbe`, `startupProbe` | 준비 완료, 생존, 시작 상태 확인                | 목적이 서로 다르므로 애플리케이션 동작에 맞게 설정한다.                       |
| `volumes`, `volumeMounts`                         | Pod의 스토리지를 정의하고 컨테이너 경로에 연결 | `volumes`는 Pod에, `volumeMounts`는 컨테이너에 지정한다.                      |
| `status`                                          | Kubernetes가 관찰한 현재 상태                  | 보통 직접 작성하지 않으며, 컨트롤러가 갱신한다.                               |

단순히 CLI를 통해 kubectl 명령어를 사용할 때보다 사용자가 인지해야 하는 데이터 정보가 더욱 많아지며, 선언적 스크립트라는 장점이 있기에 쿠버네티스에 해야할 일들을 일일이 지시하는 방식이 아니라 최종 결과가 어떻게 되어야 하는지 알려주고 만드는 과정은 개입하지 않는 방식이다.

## 2.1 Deployment 정의

고수준 리소스를 사용한다면 애플리케이션 매니페스트로 광범위한 프로세스를 구축할 수 있고 YAML 포맷으로 Deployment를 적용할 때 필수적으로 작성해야 하는 부분은 실행할 파드의 정의다. 파드는 자신만의 정의 내용이 별도로 존재하기 때문에 Deployment 정의는 파드에 대한 정의를 포함하게 된다.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-deployment

spec:
  selector:
    matchLabels: example-spec
  template:
    metadata:
      labels:
        app: example-spec
  containers:
    - name: example-container
      image: example-image-name
```

모든 애플리케이션 매니페스트는 kubectl의 `apply` 명령어를 사용하여 동일한 방식으로 배치되고, 복잡한 애플리케이션이라도 하나 또는 여러 개의 YAML 파일에 해당 애플리케이션 기술하고 같은 kubectl 명령으로 배치할 수 있다. 현재 예제에는 없는 리소스 제한과 파드 갯수 등 명시해 어떤 운영을 확인하고자 하는 수준으로 설정값을 지정해 유동적으로 관리할 수 있다.

## 2.2 Pod에서 실행 중인 애플리케이션에 접근

Pod와 Deployment가 애플리케이션의 가용성을 확보하는 역할을 하더라도, 결국 애플리케이션은 컨테이너 속에서 동작하기 때문에 컨테이너 런타임에 따라 직접적인 컨테이너 접근을 허용하지 않을 수 있다. 그렇다해도 kubectl을 통해 파드 안에 있는 컨테이너에 접근할 수 있는 방법이 있으며 쿠버네티스 명령행 도구를 사용하면 컨테이너 안에서 명령을 실행하거나, 애플리케이션 로그를 확인하거나 파일을 복사할 수 있다.

#### 파드의 내부 IP 주소 확인

```bash
kubectl get pod nginx-app-5cf778f7bb-zhn8g -o custom-columns=NAME:metadata.name,POD_IP:status.podIP
```

```bash
NAME                         POD_IP
nginx-app-5cf778f7bb-zhn8g   10.244.1.2
```

#### 파드 내부와 연결을 위한 대화형 셀 실행

```bash
kubectl exec -it [POD_NAME] -- sh
```

> [!TIP]
> docker와 같은 대화형 셀 실행 방법은 exec -it <container-name> -- sh로 동일하지만 구분자를 추가해서 입력해주면 진입이 가능하다.

```bash
error: exec [POD] [COMMAND] is not supported anymore. Use exec [POD] -- [COMMAND] instead
See 'kubectl exec -h' for help and examples
```

> [!NOTE]
> 위와 같은 명령어를 입력 했을 때 발생하는 오류인데 `kubectl exec -it <pod-name> sh` 기준으로 대부분의 명령어에 `--`가 제외되어 있어 구분자가 필요하기 때문에 추가하지 않았다면 한번 추가해서 실행을 해보면 된다.

#### 파드 내부에서 IP 주소 확인

```bash
hostname -i
```

```bash
10.244.1.2
```

#### 파드의 로그 확인

**모든 라인수 만큼 로그 확인**
```bash
kubectl logs [POD_NAME]  
```

```bash 
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Sourcing /docker-entrypoint.d/15-local-resolvers.envsh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2026/11/26 11:48:59 [notice] 1#1: using the "epoll" event method
2026/11/26 11:48:59 [notice] 1#1: nginx/1.31.10
2026/11/26 11:48:59 [notice] 1#1: built by gcc 14.2.0 (Debian 14.2.0-19) 
2026/11/26 11:48:59 [notice] 1#1: OS: Linux 6.18.58.1-microsoft-standard-WSL2
2026/11/26 11:48:59 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1073741816:1073741816
2026/11/26 11:48:59 [notice] 1#1: start worker processes
2026/11/26 11:48:59 [notice] 1#1: start worker process 33
2026/11/26 11:48:59 [notice] 1#1: start worker process 34
2026/11/26 11:48:59 [notice] 1#1: start worker process 35
2026/11/26 11:48:59 [notice] 1#1: start worker process 36
2026/11/26 11:48:59 [notice] 1#1: start worker process 37
2026/11/26 11:48:59 [notice] 1#1: start worker process 38
2026/11/26 11:48:59 [notice] 1#1: start worker process 39
2026/11/26 11:48:59 [notice] 1#1: start worker process 40
2026/11/26 11:48:59 [notice] 1#1: start worker process 41
2026/11/26 11:48:59 [notice] 1#1: start worker process 42
2026/11/26 11:48:59 [notice] 1#1: start worker process 43
2026/11/26 11:48:59 [notice] 1#1: start worker process 44
2026/11/26 11:48:59 [notice] 1#1: start worker process 45
2026/11/26 11:48:59 [notice] 1#1: start worker process 46
2026/11/26 11:48:59 [notice] 1#1: start worker process 47
2026/11/26 11:48:59 [notice] 1#1: start worker process 48
```


**특정 라인수 만큼 로그 확인**
```bash
# kubectl logs [OPTION] [POD_NAME]
kubectl logs --tail=2 [POD_NAME]
```

```bash
2026/11/26 11:48:59 [notice] 1#1: start worker process 47
2026/11/26 11:48:59 [notice] 1#1: start worker process 48
```

`tail`은 로그 파일의 마지막 부분을 확인할 때 사용한다. `kubectl logs --tail=2 [pod_name]` 명령은 가장 최근에 출력된 로그 2줄만 조회한다. 실시간으로 로그를 확인하려면 `-f(--follow)` 옵션을 함께 사용해야 한다. 

> [!NOTE]
> `kubectl get events -w`은 해당 파드의 모든 이벤트 기록을 보는 것이기 때문에 `--tail=2` 같은 특정 로그를 보고 싶은 경우와 혼돈 해서는 안된다.

| option | description |
| -- | -- |
| -f | 파일의 업데이트 내용을 실시간으로 계속 출력 |
| -n [int] | 파일의 마지막줄 부터 지정한 라인수 만큼 내용을 출력 |
| -c [byte] | 파일의 마지막부터 지정한 바이트만큼의 내용을 출력 |
| -p | 파일의 헤더와 상단의 파일 이름만 출력하지 않고 내용만 출력 |
| -v | 출력전 파일의 헤더와 이름 먼저 출력 후 내용 출력 |

> [!TIP]
> 여러 파일의 로그를 확인 하고 싶을 경우 `tail [FILE_NAME1] [FILE_NAME2]`으로 가능하다.

#### [파드의 로그 확인](#파드의-로그-확인)과 기록된 실제 로그와 동일한지 확인

```bash
docker container logs --tail=2 $(docker container ls -q --filter label=io.kubernetes.container.name=[POD_NAME])
```

#### 레이블 셀렉터와 일치한 모든 파드의 로그 열람

```bash
kubectl exec deploy/[DEPLOYMENT_NAME] -- sh -c 'wget -O http://localhost>/dev/null'
```

> [!WARNING]
> `docker container logs`와 `kubectl exec ... -- sh -c` 명령어는 현재 작성자의 터미널에서 확인 할 수 없기에 향후 재시도 진행 후 갱신 될 예정이다. 하지만 VSCode의 확장 프로그램의 kubernetes에서 Workload를 통해 pod 로그를 확인할 수 있는 방법이 존재한다.  
> [`11/27/2026 수정`]: 발생한 문제를 확인한 결과 단순히 nginx를 deploy해서 exec 통해 내부에서 wegt 자체가 존재하지 않기 때문에 명령어가 실행되지 않은 문제였다.

#### Deploy의 운영 체제 정보 출력

```bash
geonwoo@MAIN-DESKTOP:~$ kubectl exec deploy/nginx-app -- cat /etc/os-release
```

```bash
PRETTY_NAME="Debian GNU/Linux 13 (trixie)"
NAME="Debian GNU/Linux"
VERSION_ID="13"
VERSION="13 (trixie)"
VERSION_CODENAME=trixie
DEBIAN_VERSION_FULL=13.8
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"
```