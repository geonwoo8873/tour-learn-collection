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

아래 예제는 데이터베이스 비밀번호를 저장소에 넣지 않도록 Secret 정의를 포함하지 않는다. 매니페스트를 적용하기 전에 `mysql-secret`을 별도로 준비해야 하며, 필요한 키는 예제 아래에서 설명한다.

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

* apiVersion: 매니페스트 스크립트의 쿠버네티스 API 버전
* kind: 정의하고자 하는 리소스의 타입을 명시
* metadata: 리소스 이름 등 메타데이터를 정의한다. `name`은 필수이며 `labels`는 선택 사항이다.
* spec: 리소스 종류에 따라 달라지는 원하는 상태를 정의한다. 워크로드 리소스에서는 실행할 컨테이너와 이미지 등을 지정한다.

이 예제에서 `mysql-secret`은 민감한 값을 저장소에 넣지 않도록 매니페스트에서 제외했다. 적용 전에 `mysql-secret`을 별도로 생성해야 하며, `MYSQL_PASSWORD`와 `MYSQL_ROOT_PASSWORD` 키가 필요하다. Secret 값은 저장소에 커밋하지 않는다. Kubernetes Secret의 값은 기본적으로 base64 인코딩일 뿐이므로, 접근 권한과 저장 시 암호화 설정도 별도로 고려해야 한다.

단순히 CLI를 통해 kubectl 명령어를 사용할 때보다 사용자가 인지해야 하는 데이터 정보가 더욱 많아지며, 선언적 스크립트라는 장점이 있기에 쿠버네티스에 해야할 일들을 일일이 지시하는 방식이 아니라 최종 결과가 어떻게 되어야 하는지 알려주고 만드는 과정은 개입하지 않는 방식이다.

새로운 파드를 kubectl을 사용하여 다른 파드와 마찬가지로 다룰 수 잇고 세부 정보를 보거나 네트워크 트래픽을 특정 파드로 전달하는 것도 가능하다. 하지만 매니페스트는 애플리케이션 정의를 공유하기 쉽고, 동일한 배포를 반복할 수 있는 점에서 차이가 존재하지만 결국 명령어를 사용해야 하는 시점은 있기 때문에 앞으로도 명령어와 로그 및 모니터링에 대한 보는 방법을 터득하면 향후 운영 관점에서 상대적인 쉬운 식별이 가능해진다.

## 2.1 Deployment 정의

고수준 리소스를 사용한다면 애플리케이션 매니페스트로 광범위한 프로세스를 구축할 수 있고 YAML 포맷으로 Deployment를 적용할 때 필수적으로 작성해야 하는 부분은 실행할 파드의 정의다. 파드는 자신만의 정의 내용이 별도로 존재하기 때문에 Deployment 정의는 파드에 대한 정의를 포함하게 된다.