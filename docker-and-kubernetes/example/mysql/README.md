# MySQL on Kubernetes 예제

`api-config.yaml` 하나에 ConfigMap, Secret, PVC, Service, StatefulSet을 정의한 MySQL 예제.

## 구성

```
ConfigMap (mysql-config)  ─┐
Secret    (mysql-secret)  ─┼─> StatefulSet (mysql) ─> Pod (mysql-0)
PVC       (mypvc)         ─┘          ▲
Service   (mysql-service) ── selector: app=mysql
```

## api-config.yaml

```yaml
##########################################
# MySQL Kubernetes Architecture Layout
# Configuration
#   -> Public [ConfigMap]
#   -> Private [Secret]
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
  MYSQL_PASSWORD: "1234"
  MYSQL_ROOT_PASSWORD: "1234"

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
  name: mysql-service
  labels:
    app: mysql-service
spec:
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
        - name: mypvc
          mountPath: /docker-entrypoint-initdb.d/dummy.sql
        ports:
        - containerPort: 3306
      volumes:
        - name: mypvc
          persistentVolumeClaim:
            claimName: mypvc
```

## 사용 명령

```powershell
# 문법 검증
kubectl apply --dry-run=client -f api-config.yaml

# 적용 / 확인
kubectl apply -f api-config.yaml
kubectl get pods
kubectl logs mysql-0

# 삭제 (PVC 포함, DB 데이터도 삭제됨)
kubectl delete -f api-config.yaml
```

- StatefulSet만 삭제하면 PVC는 남는다. 데이터를 지우려면 PVC를 따로 삭제해야 한다.
- 파드를 줄이려면 `kubectl scale statefulset mysql --replicas=0`을 사용한다.

## 데이터 확인

### 1. 파드 안에서 mysql 클라이언트 실행

```powershell
kubectl exec -it mysql-0 -- mysql -uroot -p
```

비밀번호는 Secret의 `MYSQL_ROOT_PASSWORD` 값이다. 비밀번호를 명령에 직접 쓰지 않으려면 위처럼 프롬프트로 입력한다.

```sql
SHOW DATABASES;
USE testdb;
SHOW TABLES;
SELECT * FROM users;
```

### 2. 한 줄로 쿼리 실행

```powershell
kubectl exec mysql-0 -- mysql -uroot -p1234 -e "SHOW DATABASES;"
```

- 비밀번호가 셸 히스토리에 남으므로 학습용으로만 사용한다.

### 3. dummy.sql 적용

초기화 경로에 마운트하지 않았다면 파드에 직접 넣어 실행한다.

```powershell
kubectl cp dummy.sql mysql-0:/tmp/dummy.sql
kubectl exec mysql-0 -- sh -c "mysql -uroot -p1234 mydatabase < /tmp/dummy.sql"
```

- `dummy.sql`의 `USE testdb;`는 ConfigMap의 `MYSQL_DATABASE: mydatabase`와 다르다. 둘 중 하나를 맞춘다.

### 4. 로컬에서 접속 (포트 포워딩)

```powershell
kubectl port-forward pod/mysql-0 3306:3306
```

다른 터미널에서 로컬 클라이언트로 접속한다.

```powershell
mysql -h 127.0.0.1 -P 3306 -uroot -p
```

### 5. 환경변수와 Secret 값 확인

```powershell
kubectl exec mysql-0 -- env | Select-String MYSQL
kubectl get secret mysql-secret -o jsonpath="{.data.MYSQL_ROOT_PASSWORD}"
```

- Secret 조회 값은 base64이다. 디코딩은 `[Text.Encoding]::UTF8.GetString([Convert]::FromBase64String("<값>"))`로 한다.

### 6. 데이터 유지 확인

```powershell
kubectl delete pod mysql-0
kubectl get pods -w
kubectl exec -it mysql-0 -- mysql -uroot -p -e "SHOW DATABASES;"
```

- 파드가 다시 만들어져도 PVC가 `/var/lib/mysql`에 마운트되어 있으면 데이터가 유지된다.

## 발생한 문제와 원인, 해결 방안

### 1. metadata.name과 labels를 맞춰야 하는가
- 원인: 두 값은 서로 무관하다. `name`은 고유 식별자, `labels`는 셀렉터용 key/value이다.
- 해결: 맞출 필요는 없다. 반드시 일치해야 하는 것은 `selector.matchLabels`와 `template.metadata.labels`, 그리고 Service의 `selector`와 파드 레이블이다.

### 2. StatefulSet `apiVersion` 오류
- 원인: StatefulSet을 `apiVersion: v1`로 작성했다. StatefulSet은 `apps` 그룹이다.
- 해결: `apiVersion: apps/v1`로 수정한다. 에러 메시지에 `StatefulSet in version "v1"`이 보이면 이 문제이다.

### 3. ConfigMap 내용이 파일명이었다
- 원인: `data: name: ".env.config"`는 키 `name`에 문자열을 넣은 것이다. 파일을 읽어오지 않는다. `envFrom`은 ConfigMap의 키를 그대로 환경변수로 만든다.
- 해결: 키를 환경변수 이름으로 직접 작성한다(`MYSQL_DATABASE: ...`). 파일에서 만들려면 `kubectl create configmap mysql-config --from-env-file=.env.config`를 사용한다.

### 4. Secret `illegal base64 data at input byte 0`
- 원인: Secret의 `data` 값은 base64여야 하는데 평문(또는 파일명)을 넣었다. `"1234"`는 base64로 해석되어 에러는 안 나지만 디코딩된 엉뚱한 값이 비밀번호가 된다.
- 해결: 평문은 `stringData`에 작성한다. 비밀번호는 ConfigMap이 아니라 Secret에 둔다.

### 5. PVC `unknown field "spec.mountPath"`
- 원인: `mountPath`는 PVC 필드가 아니다. 마운트 위치는 파드(컨테이너)의 `volumeMounts`에서 지정한다.
- 해결: PVC에서 제거한다. PVC는 용량과 접근 모드만 정의한다.

### 6. `configMapRef`/`secretRef`의 이름이 잘못됨
- 원인: 파일명(`.env.config`, `.env.secrect`)을 넣었다. 참조는 리소스의 `metadata.name`으로 한다. `.`으로 시작하는 이름은 RFC 1123 규칙에도 어긋난다.
- 해결: `mysql-config`, `mysql-secret`으로 수정한다.

### 7. `volumeMounts[].name: Not found: "mypvc"`
- 원인: `volumeMounts`에 대응하는 `volumes`가 없었다. 연결은 `volumeMounts.name` -> `volumes.name` -> `persistentVolumeClaim.claimName` -> PVC 이름 순이다.
- 해결: 파드 스펙에 `volumes`를 추가하고 이름을 맞춘다.

### 8. `volumes` 위치 오류 (PVC의 `spec.volumes`, `containers[0].volumes`)
- 원인: `volumes`를 PVC의 `spec`과 컨테이너 항목 안에 넣었다. `volumes`는 파드 필드이다.
- 해결: `spec.template.spec` 아래, `containers`와 같은 들여쓰기에 둔다.

### 9. `serviceName`과 Service 이름 불일치, 마운트 경로 혼동
- 원인: StatefulSet의 `serviceName: "mysql"`과 Service 이름 `mysql-service`가 다르다. 또한 같은 PVC를 `/var/lib/mysql`과 `/docker-entrypoint-initdb.d/dummy.sql` 두 곳에 마운트했다.
- 해결: `serviceName`을 실제 Service 이름과 맞추고, 파드별 DNS가 필요하면 헤드리스(`clusterIP: None`) Service를 쓴다. 데이터는 `/var/lib/mysql`에 PVC를 마운트하고, 초기화 SQL은 ConfigMap 볼륨으로 `/docker-entrypoint-initdb.d/`에 마운트한다.

### 10. 일반 팁
- `kubectl apply`는 일부 리소스가 실패해도 나머지는 생성된다. 에러가 난 리소스만 고치고 다시 적용하면 된다.
- 터미널 출력이 이전 파일 기준일 수 있으니 저장 후 다시 적용해서 확인한다.
- 문제가 생기면 `kubectl describe`, `kubectl logs`, `kubectl get endpoints` 순으로 확인한다.
