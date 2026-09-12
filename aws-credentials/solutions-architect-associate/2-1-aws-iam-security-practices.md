# 1. IAM 신원과 자격 증명 비교

AWS 신원 및 접근 관리에서 신원은 IAM 사용자, 역할, 그룹으로 분류되어 이것들의 아이덴티티는 AWS가 생성한 루트 사용자 외에 추가로 사용된다. 통산 일반적인 작업 혹은 관리 작업의 경우에도 루트 사용자를 사용하지 않을 것을 AWS측에서 권장하며, 추가 사용자를 제공해서 필요한 작업을 수행하는 데 필요한 권한을 해당 사용자에게 부여한다.

보안 강화를 위해 루트 액세스를 중앙 집중화하여 AWS Organizations로 관리되는 AWS 계정의 루트 사용자 자격 증명을 중앙에서 관리 및 보호하는 것이 좋으며 [중앙에서 멤버 계정에 대한 루트 액세스 관리](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#id_root-user-access-management)를 통해 장기적인 루트 사용자 자격 증명 복구를 중앙에서 제거하고 방지해 대규모 단위의 루트 액세스를 방지할 수 있다.

> [!NOTE]
> * **IAM Resource**
>   * IAM 서비스는 다음 리소스를 저장하며 이 리소스들은 콘솔에서 추가, 편집, 제거할 수 있다.
> * **IAM Identity**
>   * AWS가 인증에 사용하는 IAM 리소스, 리소스 기반 정책에서 엔터티를 보안 주체로 지정한다.

#### IAM Resource

* IAM 사용자
* IAM 그룹
* IAM 역할
* 권한 정책
* 자격 증명 공급자 객체

#### IAM Identity

* IAM 사용자
* IAM 역할

#### IAM 자격 증명

<img width="50%" height="25%" alt="image" src="https://docs.aws.amazon.com/images/IAM/latest/UserGuide/images/iam-terms-2.png" />

정책에서 권한을 부여받아 작업을 수행하고 리소스에 액세스할 수 있는 IAM 리소스이다.

* ID
  * IAM 사용자
  * IAM 그룹
  * IAM 역할

#### 보안 주체

AWS 리소스에 대한 작업 또는 연산을 요청할 수 있는 AWS 계정 루트 사용자, IAM 사용자 또는 IAM 역할에 해당된다. 사용자, 워크로드, 페더레이션 보안 주체 및 수임된 역할이 포함되어 인증 후 IAM은 보안 주체 유형에 따라 AWS에 요청할 수 있는 영구 또는 임시 보안 인증 정보를 보안 주체에게 부여한다.



---

## 참조

[AWS Account root user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#id_root-user-access-management)