# 1. Elastic Coumputing Cloud [`EC2`]

Amazon Elastic Compute Cloud (`Amazon EC2`)는 Amazon Web Service (`AWS`) 클라우드에서 온디맨드 확장 가능 컴퓨팅 용량을 제공하여, 사용자가 하드웨어 비용을 절감함 으로써 애플리케이션을 더욱 빠르게 개발 및 배포에 집중할 수 있게 해준다. 이는 원하는 수의 가상 서버를 구축하고 보안 및 네트워킹을 구성해 스토리지를 관리하여, 용량을 `스케일 업` 하거나 유동성을 이용하여 월간이나 연간 프로세스 또는 웹 사이트 트래픽 급증 등 다양한 이슈에도 컴퓨팅 사용량이 많은 작업을 처리할 수 있게 해준다.

EC2 Instance는 AWS 클라우드의 가상 서버로 Instance를 시작할 때 지정하는 유형에 따라 가용할 수 있는 하드웨어가 결정되어 서로 다른 Compute, Memory, Network, Storage 리소스의 균형을 제공한다.

![](https://docs.aws.amazon.com/images/AWSEC2/latest/UserGuide/images/instance-types.png)

# 2. Amazon EC2 인스턴스 유형

> [!CAUTION]
> **U-9tb1, U-12tb1, U-18tb1, U-24tb1의 Instance 유형은 더 이상 새로 생성할 수 없으며, 워크로드에 저장된 메모리 인스턴스가 필요한 경우 U7i Instance 유형을 대신 사용하는 것을 권장한다.**

`Amazon EC2`는 `Cpu`, `Memory`, `Storage`, `Netword` 등 호스트 컴퓨터의 일부 리소스를 특정 인스턴스에 전용을 할당되어, 네트워크 및 디스크 하위 시스템과 같은 기타 리소스를 인스턴스 간에 공유한다. 호스트 컴퓨터의 각 인스턴스가 공유 리소스 중 하나를 최대한 가용률이 높은 경우 리소스는 각 인스턴스에 분배하지만, 리소스 사용률이 저조할 경우 리소스에 여유가 있는 한 특정 인스턴스가 해당 리소스를 더 많이 소비할 수 있게 해준다.

각 인스턴스 유형은 공유 리소스의 최소 성능을 상황에 따라 제공량이 다르다. I/O 성능이 높은 인스턴스 유형에는 더 많은 공유 리소스가 할당되어 성능의 변동성이 그만큼 감소하지만 대부분의 애플리케이션에 대해서는 보통 수준의 성능만으로 충분하기 때문에 일관적인 성능이 필요한 애플리케이션에 대해서는 높은 유형의 인스턴스를 사용하는 것이 좋다.

## 2.1 EC2 Instance Type 명명 규칙

**Instance Serise**
| Serise      | Description                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------- |
| A           | Arm 기반 AWS Graviton Process 구동                                                                    |
| C           | 컴퓨팅 최적화                                                                                         |
| D           | 고밀도 스토리지                                                                                       |
| F           | FPGA                                                                                                  |
| G           | 그래픽 집약적                                                                                         |
| Hpc         | 고성능 컴퓨팅                                                                                         |
| I / Im / Is | 스토리지 최적화 / 최적화된 스토리지 (`CPU, Storage = 1:4`) / 최적화된 스토리지 (`CPU, Storage = 1:6`) |
| Inf         | AWS 추론                                                                                              |
| M           | 범용                                                                                                  |
| Mac         | macOS                                                                                                 |
| P           | GPU 가속                                                                                              |
| R           | 메모리 최적화                                                                                         |
| T           | 버스트 가능한 성능                                                                                    |
| Trn         | AWS Trainium                                                                                          |
| U           | 고용량 메모리                                                                                         |
| VT          | 비디오 트랜스 코딩                                                                                    |
| X           | 메모리 잡약적                                                                                         |
| Z           | 고용량 메모리                                                                                         |

**Instance Options**
| Options          | Description                         |
| ---------------- | ----------------------------------- |
| a                | AMD Process                         |
| b * 00 / gb * 00 | NVIDIA Blackwell Gpu로 가속화       |
| g                | AWS Graviton Process                |
| i                | Intel Process                       |
| m* / m* pro      | Apple Chipset                       |
| b                | 블록 스토리지 최적화                |
| d                | 인스턴스 저장소 볼륨                |
| e                | 추가 인스턴스 스토리지, 메모리, GPU |
| flex             | Flex 인스턴스                       |
| n                | 네트워크 및 EBS 최적화              |
| q                | Qualcomm 추론 액셀러레이터          |
| * tb             | 고용량 메모리 인스턴스의 메모리     |
| z                | 높은 CPU 주파수                     |

## 2.1 EC2 Instance 성능

#### 고정 성능 인스턴스

고정 성능의 인스턴스는 고정 CPU 리소스를 제공하여 언제든지 워크로드에 필요한 동안 전체 CPU 성능을 일관된 상태에서 제공하고 유지할 수 있다. 비디오 인코딩과 같은 애플리케이션, 대용량 웹사이트 또는 HPC 애플리케이션을 위해 일관되게 높은 CPU 성능이 필요하다면, 고정 성능 인스터스 유형이 적합하다.

#### 성능 버스트 가능 인스턴스

버스트 가능 성능 (`T`) 인스턴스는 기본 수준의 CPU 성능외에 기준 이상으로 퍼포먼스를 높일 수 있는 기능을 제공한다. 기준 CPU는 대규모 마이크로 서비스, 웹 서버, 중소 규모의 데이터베이스, 데이터 로깅, 코드 리포지토리, 가상 데스크톱, 개발 및 테스트 환경과 같은 대부분의 범용 워크로드의 요구 사항을 충족하도록 설계되어있다.

#### Flex 인스턴스

`C7i-flex`, `C8i-flex`, `M7i-flex`, `M8i-flex`, `R8i-flex`와 같은 인스턴스는 리소스의 균형을 제공하며 광범위한 범용 애플리케이션을 실행하는 가장 비용 측면에서 효율적인 방법을 제공한다. 이러한 인스턴스에서 안정적인 리소스를 제공중 40%의 기준 CPU 성능을 제공하며, 이는 대부분 범용 워크로드에 대한 컴퓨팅 요구를 사항을 충족하도록 설계했기 때문에 더 많은 성능이 필요한 경우 이러한 인스턴스에서는 기준 CPU 성능을 초과해 약 24시간 동안 95%에서 최대 100%의 성능을 제공할 수 있는 기능을 제공한다.

## 2.2 EC2 Instance 기능

* **인스턴스 (`Instance`)**
  * 가상 서버
* **Amazon Machine Images (`AMIs`)**
  * 서버에 필요한 구성 요소 (`운영 체제와 추가 소프트웨어 포함`)를 패키징하는 인스턴스용 사전 구성 템플릿
* **인스턴스 유형**
  * 인스턴스의 다양한 CPU, Memory, Storage, Network 및 Graphic H/W 구성

# 3. Amazon Machine Image [`AMI`]

Amazon Machine Image (`AMI`)는 Amazon EC2 Intance를 설정하고 부팅하는 데 필요한 S/W를 제공하는 이미지 서비스다. 각 AMI에는 시작하는 인스턴스에 연결할 블록 디바이스를 지정하는 매핑이 포함되어 시작할 때 AMI를 지정해야 하기 때문에 AMI는 선택한 인스턴스 유형과 호환되어야 한다. AWS에서 제공하는 AMI, Public AMI, AMI Marketplace에서 구매한 AMI를 사용할 수 있다.

**AMI는 다음과 같은 경우에만 사용할 수 있다.**
* Region
* OS
* Process Architecture
* Launch permissions [`시작 권한`]
* Root Vloume Type
* Virtualization types

![](https://docs.aws.amazon.com/images/AWSEC2/latest/UserGuide/images/launch-from-ami.png)

동일한 구성의 인스턴스가 여러 개 필요할 때는 한 AMI에서 시작할 수 있다.

## 3.1 AMI 유형 및 특성

#### 시작 권한

시작 권한은 AMI를 사용하여 인스턴스를 시작할 수 있는 사용자를 결정한다. 시작 권한을 부여하면 AMI를 다른 사용자와 공유할 수 있으며, AMI 소유자만 시작 권한을 지정하여 가용성을 결정할 수 있다.

| 시작 권한 [`Launch permission`] | 설명 [`Description`] |
| ------------------------------- | -------------------- |
| Pulic | 소유자는 모든 AWS 계정에 시작 권한을 부여한다. |
| Explicit (`명시적`) | 소유자는 특정 AWS 계정, 조직 또는 OU (`조직 단위`)에 시작 권한을 부여한다. |
| Implicit (`암묵적`) | 소유자는 AMI에 대한 암묵적인 시작 권한을 갖는다. |

## 3.2 루트 볼륨 유형 [`Root Vloume Type`]

**모든 AMI는 Amazon EBS에 의해 지원되는 유형 또는 Amazon S3에 의해 지원되는 유형으로 분류된다.**
* **Amazon EBS 지원 AMI**
  * AMI에서 시작된 인스턴스의 루트 볼륨은 Amazon EBS 스냅샷 (`SnapShot`)에서 생성된 Amazon Elastic Block Store (`EBS`) 볼륨이며 Linux와 Windows AMI 모두 지원된다.
* **Amazon S3 지원 AMI**
  * AMI에서 시작된 인스턴스의 루트 볼륨은 Amazon S3에 저장된 템플릿으로 부터 생성된 인스턴스 저장소 볼륨이며, Linux AMI에서만 지원하지만 Windows AMI의 루트 볼륨에 대한 인스턴스 저장소는 지원하지 않는다.

#### EBS 지원 AMI과 S3 지원 AMI의 정리

**Amazon EBS 지원 AMI**

Amazon EBS 지원 AMI는 EBS 볼륨이 루트 볼륨이며 인스턴스의 부팅 시간이 `S3 지원 AMI` 보다 빠른 `평균 1분 이하로 부팅 완료`되며, 데이터의 지속이 기본적으로 인스턴스가 `종료될 때 루트 볼륨이 동시에 삭제`되며 `다른 EBS 볼륨의 데이터는 종료 후에도 유지`된다. 중지 상태일 때 인스턴스가 중지되고 실행 중이지 않은 경우에도 루트 볼륨은 EBS에 유지된다. 

인스턴스 유형, Kernel내 리소스 (`RAM`, `Disk`, `User`) 데이터는 변경될 수 있으며 비용 부과 유형에서 `인스턴스 사용량`, `EBS 볼륨 사용량` 및 `스냅샷`으로 저장하는 것에 대한 비용이 청구되며, 생성 과정도 단일 명령 혹은 호출을 사용하면된다.

**Amazon S3 지원 AMI**

Amazon S3 지원 AMI는 `인스턴스 저장소 불륨`으로 인스턴스의 부팅 시간이 EBS 불륨 형식보다 상대적으로 느려 `평균 5분 이하로 부팅 완료`되며, 데이터 지속성이 모든 `인스턴스의 수명 주기를 따르기 때문`에 별도로 `중지를 할 수 없으며 실행되거나 종료되어 삭제되는 방식외 지속되지 않는다.` EBS 지원과 유사하게 `인스턴스 사용량`과 `Amazon S3에 AMI를 저장`하는 것에 대한 비용만 청구되고, AMI 도구를 별도 설치 및 사용해야 한다.

## 3.3 가상화 유형 [`Virtualization types`]

Amazon Machine Image는 PV (`Paravirtualization` ,`반가상화`) 또는 HVM (`Hardware Virtualization Maichen`, 하드웨어 가상 머신)의 두 가지 유형의 가상화를 사용한다. PV AMI와 HVM AMI의 주요 차이점은 부팅 진행 방식과 성능을 위한 특수 하드웨어 확장 (`리소스`)을 활용 여부에 결정되며, Windows AMI는 HVM AMI으로 분류된다.

**HVM (`Hardware Virtualization Maichen`) **

HVM AMI는 이미지 루트 블록 디바이스의 마스터 부트 레코드를 실행하여 가상화된 H/W 및 부트 세트를 함께 제공된다. 가상화 유형은 운영 체제 미설치 H/W 초기 실행 시점 처럼 가상 머신에서 OS를 수정하지 않고 실행할 수 있기 때문에 EC2 호스트 시스템은 게스트에게 제공되는 기본 하드웨어의 일부 또는 전체를 에뮬레이트를 한다.

모든 최신 인스턴스 유형은 HVM AMI를 지원하며 HVM 게스트는 H/W 확장을 활용하여 호스트 시스템의 기본 하드웨어에 빠르게 액세스할 수 있다. 또한 향상된 네트워킹과 GPU 처리를 사용해야 하지만, 특수 네트워크 및 GPU 디바이스에 명령을 전달하려면 OS가 기본 하드웨어 팻랫폼에 액세스할 수 있어야 하며 HVM 가상화가 이러한 액세스를 제공한다.

**PV (`Paravirtualization`)**

PV AMI는 PV-GRUB 특수 부트 로더를 통해 부팅되며, 로더는 부팅 주기를 시작한 후 사용자 이미지의 menu.1st 파일에 지정된 커널을 체인 로드한다. 반가상화 게스트는 가상화를 명시적으로 지원하지 않는 하드웨어에서 실행할 수 있다.

C1, C2, M1, M2, M3, T1 등과 같은 전 세대 인스턴스 유형은 PV AMI를 지원하지만, 최신 세대 인스턴스 유형은 PV AMI를 지원하지 않는다. 또한 향상된 네트워킹 또는 GPU 처리와 같은 특수 하드웨어 확장을 활용할 수 없다는 단점이 존재한다.

---

## 참조

* [Amazon Elastic Computing Clout Instance Pricing](https://aws.amazon.com/ko/ec2/pricing/)
* [Amazon Elastic Computing Cloud Instance Typs](https://aws.amazon.com/ko/ec2/instance-types/)
* [Amazon Elastic Computing Cloud Perfomance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)