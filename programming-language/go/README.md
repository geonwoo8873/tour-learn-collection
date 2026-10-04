# Go

## 1. Golang installation

### 1.1 Linux

**1. 설치 파일 다운로드 (Install file download)**

```sh
wget -<options> https://go.dev/dl/go<version>.linux-<socket_type>.tar.gz
```

```sh
# Currently latest version of 1.27.1
wget https://go.dev/dl/go1.27.1.linux-amd64.tar.gz
```

> [!TIP]
> **`wget`은 웹에서 특정 데이터 파일을 가져오는(`GET`) Linux 명령어이다.**
> 

<details>
<summary>Linux wget command options</summary>

| Options                | Description                                                                                          | Example                                                 |
| ---------------------- | ---------------------------------------------------------------------------------------------------- | ------------------------------------------------------- |
| -O (`Output-Document`) | 파일을 지정한 다른 이름으로 저장한다. (예시 : cp <file_name> <replace_file_name>)                    | wget -O <replace_file_name> <request_download_file_url> |
| -c (`Continue`)        | 다운로드 중 중단이 되었거나 중단이 된 상태라면 이어서 진행하도록 한다.                               | wget -c <request_download_file_url>                     |
| -b (`Background`)      | 다운로드를 백그라운드 프로세스로 전환하여 진행한다. (파일의 사이즈가 크거나 진행이 느린 경우에 유용) | wget -b <request_download_file_url>                     |
| -r (`Recursive`)       | 웹사이트를 재귀적으로 순회해 해당하는 모든 연결 파일을 받는다.                                       | wget -r <request_download_file_url_path>                |
| -i (`Input-File`)      | 요청하는 파일에 있는 모든 URL 리스트를 다운로드한다.                                                 | wget -i <replace_download_url_file_name>                |
| -v (`Verbose`)         | 다운로드가 내부적으로 수행하는 작업과 상태 정보를 구체적으로 출력한다.                               | wget -v <request_download_file_url>                     |
| --limit-rate           | 다운로드 속도를 제한해 네트워크 대역폭을 조절하여 리소스 자원을 효율화한다.                          |                                                         |

</details>

**2. 다운로드 압축해제 (Downloaded to Extract)**

```sh
sudo <compression_type> -<options> <path> -<options> <go_version_pkg_url>
```

```sh
sudo tar -C /usr/local -xzf go1.27.1.linux-amd64.tar.gz
```

**3. User 프로필 경로 추가 (Add export path to in the User profile)**

```sh
vi .profile
```

```sh
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

**4. Go 버전 확인**

```sh
go version
```

**5. Linux 실행 파일 생성**

```sh
GOOS=linux GOARCH=amd64 go build
```

## 2.

#### 2.1 Golang 프로젝트 구성 파일 확장자 [project configure file extension]

<details>
<summary>Go project conifg command options</summary>

| Name  | Description                 |
| ----- | --------------------------- |
| Doc   | 코드 문서 표시              |
| Fmt   | 소스 코드 파일 포맷팅       |
| Get   | Go 패키지 다운로드          |
| Mod   | 프로그램 종속성으로 작업    |
| Test  | 테스트 실행                 |
| cover | 코드 테스트 커버리지 보고   |
| Vet   | 코드 정적 분석 실행         |
| Pprof | 코드 프로파일링 보고서 생성 |

</details>


---

## Reference

* [Go Programming Language Installation Page](https://go.dev/dl/)