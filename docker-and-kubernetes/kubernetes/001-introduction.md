# Kubernetes

Kubernetes는 클러스터 전체 상태를 Control Plane이 관리하는 구조로, 중앙 저장소를 기준으로 동작하는 SVN처럼 하나의 중앙 상태 저장소를 기준으로 동작한다기보다, Control Plane을 통해 클러스터의 원하는 상태를 관리하고 각 노드에 반영한다. 다만 Kubernetes는 단순 버전 관리 시스템이 아닌 분산 시스템의 상태를 지속적으로 조정(Reconciliation)하는 플랫폼이라는 차이점이 존재한다.

Cluster 관리자는 노드(`Node`)라고 하는 각각의 서버를 관리하여 처리 용량을 확장하기 위해 노드를 추가하거나 서비스에서 제외하는 행위를 차례로 업데이트하는 롤링 업데이트(`Rolling update`)를 적용할 수 있다. 흔히 주위에서 볼 수 있는 Microsoft Azure의 AKS나 Amazon Web Service에서 제공하는 EKS 같은 관리자 대신하여 관리하는 서비스를 이행함 으로써 웹 인터페이스나 CLI으로 편리하게 사용할 수 있도록 구성되어 있다.

```ps1
PS C:\Windows\System32> kubectl get nodes
NAME                    STATUS   ROLES           AGE   VERSION
desktop-control-plane   Ready    control-plane   33m   v1.36.1
desktop-worker          Ready    <none>          33m   v1.36.1
```
<작성자는 Docker Desktop을 통해 Kubernetes를 사용 중이다.>

이러한 서비스들의 목적은 애플리케이션을 실행하는 것으로 `YAML`과 같은 확장자를 지닌 파일에 애플리케이션을 기술하고 Kubernetes API에 전달을 하면, 파일에 기술된 애플리케이션 구성을 파악해 클러스터의 현재 상태와 비교하여 그 차이가 발생한 컨테이너를 추가하거나 또는 제거하는 행동을 취한다. 고가용성을 위해 여러 노드에 흩어져 있어 관리하는 가상 네트워크로 서로 통신할 수 있기 때문에 각각의 노드에 집중해야할 시간이 줄일 수 있는 이점을 가져다 준다.

한 컴포넌트를 여러 컨테이너에서 여러 횟수 실행하여 스케일링을 구현하고 가용성을 확보할 수 있는 한 편, 컨테이너는 클러스터를 구성하는 노드에 흩어져 실행되지만 Docker의 Container Network처럼 다른 노드에 있는 Container라도 표준적인 네트워크 방식으로 서로 통신이 가능하기 때문에 애플리케이션 구성을 파일에 기술하기만 한다면, 그 구성에 따라 쿠버네티스가 애플리케이션의 전반적인 실행부터 관리를 해준다.

특정 노드가 문제가 발생하더라도 해당 노드에서 실행되던 컨테이너는 다른 노드에서 실행된 컨테이너로 대체되고, 컨테이너가 문제가 발생하면 쿠버네티스가 해당 컨테이너를 재시작(`Restart`)를 진행한다. 하지만 부하에 대한 부분은 컨테이너를 제거를 하는 방식보단 추가로 실행하여 부담을 완화해주고 도커 이미지와 쿠버네티스 구성 파일만 존재하면 쿠버네티스 클러스터에서 동일하게 동작하는 자기수복형(`Self-Healing`) 애플리케이션을 만들 수 있다.

> [!NOTE]
> 어떻게 보면 Kubernetes는 Health Check를 기반으로 사람이 개입하지 않더라도 서비스가 기대하는 상태(Desired State)를 유지하도록 돕는 자동화된 운영 플랫폼과 유사하다고 생각된다. 실제로 Docker Compose 역시 Health Check와 Restart 정책을 통해 장애가 발생한 컨테이너를 재시작할 수 있지만, Kubernetes는 단순히 컨테이너를 재시작하는 것에 그치지 않고 Pod 또는 Node 장애 상황에서도 다른 노드에 컨테이너를 재배치하여 서비스의 가용성을 유지한다는 차이점이 존재한다.

클러스터에서 분산 데이터베이스를 통해 API 키나 접속 패스워드 같은 공개되어선 안되는 비밀값들도 함께 저장할 수 있기에 이러한 정보들을 컨테이너에 전달을 하는 역할도 제공하고 있다. 설정들이 일치하다면 어드 환경에서든 동일한 환경의 컨테이너 이미지를 사용할 수 있으며, 함께 마운트되는 스토리지로 외부 데이터를 저장할 수 있으며, 유상태 애플리케이션의 고가용성을 확보할 수 있다.

#### Compose의 Volume 예시
```yaml
services:
    <...>
    volumes:
      - ../db/init:/docker-entrypoint-initdb.d:ro
      - ../db/verify:/opt/verify:ro
      - database-data:/var/lib/mysql
<...>
volumes:
  database-data:
    driver: local
```

> [!TIP]
> 애플리케이션 기술한 YAML 파일은 애플리케이션 매니페스트(Application Manifest)라고 하며, 파일 내용이 애플리케이션을 구성하는 모든 컴포넌트의 목록이기 때문이기도 한다. 이를 구성하는 컴포넌트를 리소스(Resource)라고 한다. 하지만 리소스에는 여러 종류가 존재하기 때문에 향후 구체적인 내용을 작성할 예정이다.

# 1. 쿠버네티스에서 컨테이너 실행과 관리

## 1.1 Nginx 이미지 및 단일 Pod 생성 

```bash
kubectl run nginx --image nginx
```

```bash
pod/nginx created
```

현재 예제에서는 Docker Hub에 공개된 공식 nginx 이미지를 사용하므로 단순히 nginx라는 이름만 지정해도 컨테이너를 실행할 수 있다. 하지만 실제 운영 환경에서는 특정 버전의 이미지를 사용하기 위해 v1.0.0과 같은 버전 태그를 사용하거나, CI/CD 파이프라인에서 생성된 Git Commit SHA 값 일부를 이미지 태그로 사용하여 배포 이력을 추적하기도 한다. 따라서 쿠버네티스에서 애플리케이션을 배포할 때는 이미지 이름뿐만 아니라 태그와 버전이 의도한 버전인지 함께 확인하는 것이 중요하다.

## 1.2 Deployment 생성

```bash
kubectl create deployment nginx-app --image=nginx
```

```bash
deployment.apps/nginx-app created
```

## 1.3 가동중인 Pod 조회

```bash
kubectl get pods
```

```bash
NAME                         READY   STATUS    RESTARTS   AGE
nginx                        1/1     Running   0          52s
nginx-app-5cf778f7bb-62fkj   1/1     Running   0          34s
```

> [!TIP]
> Pod phase는 Pending, Running, Succeeded, Failed, Unknown이며, kubectl get pods의 STATUS에는 CrashLoopBackOff 같은 상태 이유도 표시될 수 있다.
> 또한 RESTARTS는 환경 변수값이 잘못 전달됐거나, 코드 문제로 재시작 과정을 거치는데 주의해야 하는 것은 compose로 직접 하는 변경 사항을 적용하는 것이 아닌 쿠버네티스 API에서 변경점을 확인해서 롤링 업데이트 방식으로 변경될 수 있다.

## 1.4 특정 Pod 상세 정보 조회

```bash
kubectl describe pod nginx
```

```bash
Name:             nginx
Namespace:        default
Priority:         0
Service Account:  default
Node:             desktop-worker/172.18.0.3
Start Time:       Wed, 11 Nov 2026 23:15:20 +0900
Labels:           run=nginx
Annotations:      <none>
Status:           Running
IP:               10.244.1.4
IPs:
  IP:  10.244.1.4
Containers:
  nginx:
    Container ID:   containerd://fc88405cd9da5b7912fff091e8e3a5cafdb302753ca119e636e4c8032a2f798a
    Image:          nginx
    Image ID:       docker.io/library/nginx@sha256:165ddde8751e75a1c907d24ef10d9f6deb08b57e0fe56caf67512486aff76d82
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Wed, 11 Nov 2026 23:15:23 +0900
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-ngjmg (ro)
Conditions:
  Type                        Status
  PodReadyToStartContainers   True 
  Initialized                 True 
  Ready                       True 
  ContainersReady             True 
  PodScheduled                True 
Volumes:
  kube-api-access-ngjmg:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    Optional:                false
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  31s   default-scheduler  Successfully assigned default/nginx to desktop-worker
  Normal  Pulling    30s   kubelet            spec.containers{nginx}: Pulling image "nginx"
  Normal  Pulled     28s   kubelet            spec.containers{nginx}: Successfully pulled image "nginx" in 1.805s (1.805s including waiting). Image size: 63933865 bytes.
  Normal  Created    28s   kubelet            spec.containers{nginx}: Container created
  Normal  Started    28s   kubelet            spec.containers{nginx}: Container started
```

쿠버네티스는 직접 컨테이너를 실행하지 않으며, 컨테이너를 생성할 책임을 해당 노드에 설치된 컨테이너 런타임에 맡기는 형태로 특정 컨테이너 런타임은 도커가 주 매개체로 될 순 있지만, 그외 다른 형태로 이루어져 있을 수 있다. 결국 파드가 컨테이너를 추상화 시킨 것인 이유도 있다고 볼 수 있다. 파드는 쿠버네티스가 관리하는 리소스고, 컨테이너는 외부에서 관리되어 `kubectl` 명령어를 통해 다양한 정보를 받아 확인할 수 있다.

## 1.5 Pod에 있는 Container 접속

```bash
kubectl port-forward pod/nginx 8080:80
```

```bash
Forwarding from 127.0.0.1:8080 -> 80
Forwarding from [::1]:8080 -> 80
Handling connection for 8080
```

#### 다른 터미널에서 응답 확인

```bash
curl.exe http://localhost:8080
```

> Port-Forward는 도커 컴포즈나 혹은 컨테이너 생성시에 -p를 더해 해당 포트를 외부로 노출 시키는 방법과 동일