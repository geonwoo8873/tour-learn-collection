# 1. Amazon Machine Image [`AMI`]

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

## 1.1 AMI 유형 및 특성

#### 시작 권한

시작 권한은 AMI를 사용하여 인스턴스를 시작할 수 있는 사용자를 결정한다. 시작 권한을 부여하면 AMI를 다른 사용자와 공유할 수 있으며, AMI 소유자만 시작 권한을 지정하여 가용성을 결정할 수 있다.

| 시작 권한 [`Launch permission`] | 설명 [`Description`] |
| ------------------------------- | -------------------- |
| Pulic | 소유자는 모든 AWS 계정에 시작 권한을 부여한다. |
| Explicit (`명시적`) | 소유자는 특정 AWS 계정, 조직 또는 OU (`조직 단위`)에 시작 권한을 부여한다. |
| Implicit (`암묵적`) | 소유자는 AMI에 대한 암묵적인 시작 권한을 갖는다. |

## 1.2 루트 볼륨 유형 [`Root Vloume Type`]

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

## 1.3 가상화 유형 [`Virtualization types`]

Amazon Machine Image는 PV (`Paravirtualization` ,`반가상화`) 또는 HVM (`Hardware Virtualization Maichen`, 하드웨어 가상 머신)의 두 가지 유형의 가상화를 사용한다. PV AMI와 HVM AMI의 주요 차이점은 부팅 진행 방식과 성능을 위한 특수 하드웨어 확장 (`리소스`)을 활용 여부에 결정되며, Windows AMI는 HVM AMI으로 분류된다.

**HVM (`Hardware Virtualization Maichen`) **

HVM AMI는 이미지 루트 블록 디바이스의 마스터 부트 레코드를 실행하여 가상화된 H/W 및 부트 세트를 함께 제공된다. 가상화 유형은 운영 체제 미설치 H/W 초기 실행 시점 처럼 가상 머신에서 OS를 수정하지 않고 실행할 수 있기 때문에 EC2 호스트 시스템은 게스트에게 제공되는 기본 하드웨어의 일부 또는 전체를 에뮬레이트를 한다.

모든 최신 인스턴스 유형은 HVM AMI를 지원하며 HVM 게스트는 H/W 확장을 활용하여 호스트 시스템의 기본 하드웨어에 빠르게 액세스할 수 있다. 또한 향상된 네트워킹과 GPU 처리를 사용해야 하지만, 특수 네트워크 및 GPU 디바이스에 명령을 전달하려면 OS가 기본 하드웨어 팻랫폼에 액세스할 수 있어야 하며 HVM 가상화가 이러한 액세스를 제공한다.

**PV (`Paravirtualization`)**

PV AMI는 PV-GRUB 특수 부트 로더를 통해 부팅되며, 로더는 부팅 주기를 시작한 후 사용자 이미지의 menu.1st 파일에 지정된 커널을 체인 로드한다. 반가상화 게스트는 가상화를 명시적으로 지원하지 않는 하드웨어에서 실행할 수 있다.

C1, C2, M1, M2, M3, T1 등과 같은 전 세대 인스턴스 유형은 PV AMI를 지원하지만, 최신 세대 인스턴스 유형은 PV AMI를 지원하지 않는다. 또한 향상된 네트워킹 또는 GPU 처리와 같은 특수 하드웨어 확장을 활용할 수 없다는 단점이 존재한다.
