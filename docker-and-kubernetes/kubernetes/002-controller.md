# 1. Controller 객체

컨트롤러 객체는 다른 리소스를 관리하는 쿠버네티스의 핵심 리소스다. 다르게 보면 Docker Compose에서 docker compose up, down, stop 등 다양한 액션을 취할 수 있기에 쿠버네티스 API와 연동하며 시스템의 현재 상태를 모니터링하다 특정 파드나 컨테이너의 상태가 기준에서 차이가 생기면 필요에 따라 그 차이를 바로 잡는다. 내부에는 여러 가지 컨트룰러 객체가 존재하며 주로 파드를 관리하는 디플로이먼트(`Deployment`)가 해결을 해준다.

> [!NOTE]
> 파드는 단순한 객체로 고립된 한 쌍의 애플리케이션으로 각 파드는 서로 다른 노드에 배정되기 때문에, 특정 노드가 문제가 발생하면 파드는 유실되며, 유실된 파드를 새 파드로 대체하지 않기 때문에 여러 파드를 실행하며, 고가용성을 확보하고자 하려해도 다른 노드에 흩어져서 실행된다는 강제성이나 보장이 없다.

앞서 [001-introduction.md #12-deployment-생성](001-introduction.md/#12-deployment-생성)에서 기재한 명령어를 입력하면 디플로이먼트이 생성되고, 생성한 파드 이름은 무작위 문자열로 부여한다. 반드시 생성 후 `kubetctl get pods`나 `kubectl get deployment`를 통해서 확인 해야한다. 모든 쿠버네티스 리소스는 간단한 키-값 쌍 형태의 레이블을 가지며, 원하는 데이터를 담는 데 사용할 수 있다.

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

레이블을 이용하여 리소스 간 관계를 파악하는 것은 흔히 자주 사용되는 패턴으로 리소스는 생애 주기를 거치며 생성이나 수정, 삭제와 함께 레이블을 부여 받는다. 또한 컨트룰러 객체는 레이블 셀럭터를 통해 자신이 관리하는 리소스인지 판단한다. 컨트룰러가 자신이 관리하는 리소스 목록을 직접 유지하지 않더라도 레이블 셀렉터가 컨트롤러 객체의 정의에 포함되어 있기 때문에 언제라도 쿠버네티스 API를 통해 자신이 관장하는 리소스를 찾아볼 수 있다는 유동적인 설계 방식이다.

> 디플로이먼트는 파드와 직접적으로 관계를 갖지 않아 레이블 셀렉터와 일치하는 파드가 하나라도 존재하기만 하면 되기 때문에 파드의 레이블이 수정되면 디플로이먼트는 해당 파드를 인지하지 못하게 된다.

