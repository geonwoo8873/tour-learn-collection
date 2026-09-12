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

<img width="25%" height="25%" alt="image" src="https://docs.aws.amazon.com/images/IAM/latest/UserGuide/images/iam-terms-2.png" />

정책에서 권한을 부여받아 작업을 수행하고 리소스에 액세스할 수 있는 IAM 리소스이다.

* ID
  * IAM 사용자
  * IAM 그룹
  * IAM 역할

#### 보안 주체

AWS 리소스에 대한 작업 또는 연산을 요청할 수 있는 AWS 계정 루트 사용자, IAM 사용자 또는 IAM 역할에 해당된다. 사용자, 워크로드, 페더레이션 보안 주체 및 수임된 역할이 포함되어 인증 후 IAM은 보안 주체 유형에 따라 AWS에 요청할 수 있는 영구 또는 임시 보안 인증 정보를 보안 주체에게 부여한다.

* 워크로드는 애플리케이션, 프로세스, 운영 도구 및 기타 구성 요소 같이 비즈니스 가치를 창출하는 리소스 및 코드 모음
* 페더레이션 위탁자는 Active Directory, Okta, Microsoft Entra ID 등의 다른 ID 제공업체가 ID와 자격 증명을 관리하는 사용자
* IAM 역할은 계정에서 생성할 수 있고 해당 자격 증명이 수행할 수 있는 작업과 수행할 수 없는 작업을 결정하는 특정 권한을 지닌 IAM 자격 증명이다. 하지만 한 사람에게만 연관되지 않고 해당 역할이 필요한 사람이라면 제한없이 맡을 수 있어야 한다.

IAM은 IAM 사용자와 루트 사용자에게 장기 보안 인증 정보를 부여하고 IAM 역할에 임시 보안 인증 정보를 부여한다. AWS IAM Identity Center의 사용자, OIDC 및 SAML SSO 페더레이션 보안 주체는 AWS에 로그인이 필요할 때 IAM 역할을 수임하고, 임시 자격 증명이 부여된다.

## 1.1 IAM 사용자와 IAM Identity Center의 사용자의 차이점

### 1.1.1 IAM 사용자

별도의 계정이 아니라 계정 내의 존재하는 사용자로 각 사용자에게 AWS Managment Console에 액세스하기 위한 자체 암호가 있기에 계정의 리소스를 사용하기 위한 프로그래밍 방식의 요청을 할 수 있도록 각 사용자에 대한 개별 액세스 키를 생성할 수 있다.

### 1.1.2 IAM Identity Center

작업 인력 자격 증명은 수행하는 역할에 따라 권한 요구 사항이 다르고 조직 전체의 다양한 AWS 계정에서 작업할 수 있는 AWS IAM Identity Center의 사용자이다. 액세스 키가 필요한 사용 사례가 있는 경우 AWS IAM Identity Center의 사용자를 사용해 해당 사용 사례를 지원할 수 있고 액세스 포털을 통하여 로그인하는 사용자는 AWS 리소스에 대한 단기 자격 증명이 포함된 액세스 키를 얻을 수 있다.

중앙 액세스 관리를 위해 AWS IAM Identity Center를 사용하여 계정에 대한 액세스 권한과 해당 계정내 권한을 관리하는 것이 좋으므로 사용자 및 그룹을 추가하여 리소스에 대한 액세스 수준을 할당할 수 있는 기본 ID 소스로 IAM Identity Center 디렉터리를 활용해 자동으로 구성된다. 

## 2. 기존 자격 증명 소스의 사용자 페더레이션

<img width="50%" height="25%" alt="image" src="https://docs.aws.amazon.com/images/IAM/latest/UserGuide/images/iam-intro-federation.diagram.png" />

조직 내 사용자가 회사 네트워크에 접속할 때 이미 인증된 경우 해당 사용자를 위해 별도의 IAM 사용자 또는 IAM Identity Center 사용자를 생성할 필요가 없다. 다만 IAM이나 AWS IAM Identity Center 사용을 통해 이러한 사용자 자격 증명을 AWS에 페더레이션할 수 있다. OIDC 및 SAML SSO 페더레이션 보안 주체는 특정 리소스에 액세스할 수 있는 권한을 부여하는 역할을 수임하면 된다.

> [!NOTE]
> **페더레이션은 다음과 같은 경우에 유용하다.**
> * **사용자가 이미 기업 디렉터리에 존재하는 경우**
>   * 기업 디렉터리가 SAML 2.0과 호환되는 경우, 기업 디렉터리를 구성하여 사용자에게 AWS Management Console에 대해 SSO 액세스를 제공할 수 있다. 다만 SAML 2.0과 호환되지 않는 경우 ID 브로커 애플리케이션을 생성하여 사용자에게 SSO 액세스를 제공할 수 있다.
>   * 기억 디렉터리가 Microsoft AD인 경우 AWS IAM Identity Center를 사용하여 AD의 자체 관리형 디렉터리 또는 AWS Directory Service의 디렉토리에 연결해 기업 디렉토리와 AWS 간의 신뢰를 설정할 수 있다.
>   * Okta 또는 Microsoft Entra와 같은 외부 ID IdP를 사용하여 사용자를 관리하는 경우 AWS IAM Identity Center을 사용해 IdP와 AWS 계정 간에 신뢰를 설정할 수 있다.
> * **사용자가 이미 인터넷 자격 증명을 보유한 경우**
>   * 

---

## 참조

[AWS Account root user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#id_root-user-access-management)