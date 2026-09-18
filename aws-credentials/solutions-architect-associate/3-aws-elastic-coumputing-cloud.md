# 1. Elastic Coumputing Cloud [`EC2`]

Amazon Elastic Compute Cloud (`Amazon EC2`)는 Amazon Web Service (`AWS`) 클라우드에서 온디맨드 확장 가능 컴퓨팅 용량을 제공하여, 사용자가 하드웨어 비용을 절감함 으로써 애플리케이션을 더욱 빠르게 개발 및 배포에 집중할 수 있게 해준다. 이는 원하는 수의 가상 서버를 구축하고 보안 및 네트워킹을 구성해 스토리지를 관리하여, 용량을 `스케일 업` 하거나 유동성을 이용하여 월간이나 연간 프로세스 또는 웹 사이트 트래픽 급증 등 다양한 이슈에도 컴퓨팅 사용량이 많은 작업을 처리할 수 있게 해준다.

EC2 Instance는 AWS 클라우드의 가상 서버로 Instance를 시작할 때 지정하는 유형에 따라 가용할 수 있는 하드웨어가 결정되어 서로 다른 Compute, Memory, Network, Storage 리소스의 균형을 제공한다.

![](https://docs.aws.amazon.com/images/AWSEC2/latest/UserGuide/images/instance-types.png)

# 2. Amazon EC2 인스턴스 유형

> [!CAUTION]
> **U-9tb1, U-12tb1, U-18tb1, U-24tb1의 Instance 유형은 더 이상 새로 생성할 수 없으며, 워크로드에 저장된 메모리 인스턴스가 필요한 경우 U7i Instance 유형을 대신 사용하는 것을 권장한다.**

`Amazon EC2`는 `Cpu`, `Memory`, `Storage`, `Network` 등 호스트 컴퓨터의 일부 리소스를 특정 인스턴스에 전용을 할당되어, 네트워크 및 디스크 하위 시스템과 같은 기타 리소스를 인스턴스 간에 공유한다. 호스트 컴퓨터의 각 인스턴스가 공유 리소스 중 하나를 최대한 가용률이 높은 경우 리소스는 각 인스턴스에 분배하지만, 리소스 사용률이 저조할 경우 리소스에 여유가 있는 한 특정 인스턴스가 해당 리소스를 더 많이 소비할 수 있게 해준다.

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



---

## 참조

* [Amazon Elastic Computing Clout Instance Pricing](https://aws.amazon.com/ko/ec2/pricing/)
* [Amazon Elastic Computing Cloud Instance Typs](https://aws.amazon.com/ko/ec2/instance-types/)
* [Amazon Elastic Computing Cloud Perfomance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)