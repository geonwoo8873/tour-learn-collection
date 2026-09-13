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


**페더레이션은 다음과 같은 경우에 유용하다.**
* **사용자가 이미 기업 디렉터리에 존재하는 경우**
  * 기업 디렉터리가 SAML 2.0과 호환되는 경우, 기업 디렉터리를 구성하여 사용자에게 AWS Management Console에 대해 SSO 액세스를 제공할 수 있다. 다만 SAML 2.0과 호환되지 않는 경우 ID 브로커 애플리케이션을 생성하여 사용자에게 SSO 액세스를 제공할 수 있다.
  * 기억 디렉터리가 Microsoft AD인 경우 AWS IAM Identity Center를 사용하여 AD의 자체 관리형 디렉터리 또는 AWS Directory Service의 디렉토리에 연결해 기업 디렉토리와 AWS 간의 신뢰를 설정할 수 있다.
  * Okta 또는 Microsoft Entra와 같은 외부 ID IdP를 사용하여 사용자를 관리하는 경우 AWS IAM Identity Center을 사용해 IdP와 AWS 계정 간에 신뢰를 설정할 수 있다.
* **사용자가 이미 인터넷 자격 증명을 보유한 경우**
  * 사용자가 Platform 또는 OIDC 호환 자격 증명 공급자 등의 인터넷 자격 증명 공급자를 통하여 사용자를 식별할 수 있도록 모바일 웹 기반 앱을 만들면, 해당 앱에서 연동을 통해 AWS에 액세스할 수 있다.

## 1.2 사용자 액세스를 제공하는 방법

* **사용자 액세스 유형**
  * IAM Identity Center를 사용하여 AWS 리소스에 액세스 하기 위한 SSO 액세스
    * IAM Identity Center는 사용자의 관리와 AWS 계정 및 클라우드 애플리케이션에 대한 액세스를 통합하는 중앙 위치를 제공
    * IAM Idnetity Center 내에서 ID 스토어를 설정하거나 기존 IdP와의 페더레이션을 구성할 수 있으며 보안상 가장 좋은 방법은 사용자에게 AWS 리소스에 대한 제한된 보안 인증 정보를 부여하는 것
    * 사용자는 보다 쉽게 로그인할 수 있어 단일 시스템에서 리소스에 대한 액세스를 제어할 수 있고, IAM Identity Center는 추가 계정 보안을 위해 MFA를 지원
  * IAM ID IdP를 사용하여 AWS 서비스에 액세스하기 위한 페더레이션 액세스
    * IAM은 OIDC 또는 SAML 2.0과 호환되는 IdP를 지원한다. IAM IdP를 생성 후 페더레이션 보안 주체에게 동적으로 할당할 수 있는 IAM 역할을 하나 이상 생성
  * AWS 계정 간 크로스 계정 액세스
    * 일부 AWS 리소스에 대한 액세스를 AWS 계정 사용자와 공유하려고 할 경우에 유용하며 역할은 교차 계정 액세스를 부여하는 기본적인 방법이지만 일부 AWS 서비스는 정책을 리소스에 직접 연결할 수 있는 리소스 기반 정책을 지원한다.
  * AWS 계정의 지정된 IAM 사용자를 위한 장기 보안 인증
    * AWS에서 IAM 사용자의 장기 보안 인증이 필요한 특정 사용 사례가 있으나 IAM을 사용하여 AWS 계정에서 이러한 사용자를 생성하고 IAM을 통해 해당 권한을 관리한다.

## 1.3 프로그래밍 방식의 사용자 액세스 지원

사용자가 AWS Management Console 외부에서 AWS와 상호 작용하려면 프로그램 ㅇ방식의 액세스 권한이 필요해 액세스를 부여하는 방법은 AWS에 액세스 요청하는 사용자 유형에 따라 다르다.

* IAM Identity Center에서 ID를 관리하는 경우 AWS API에는 프로필이 필요하고 AWS Command Line Interface에는 프로필이나 환경 변수가 필요하다.
* IAM 사용자가 있는 경우 AWS API 및 AWS Command Line Interface에는 액세스 키가 별도로 필요하며, 액세스 키 ID, 비밀 액세스 키 및 보안 인증 정보가 만료되는 시간을 나타내는 보안 토큰으로 구성된 임시 보안 인증 정보를 만든다.

* **프로그래밍 방식 액세스가 필요한 사용자**
  * 작업 인력 ID
    * 단기 자격 증명을 사용하여 AWS CLI 또는 AWS API에 대한 프로그래밍 방식 요청에 직접 또는 AWS SDK를 사용하여 서명해야 한다.
  * IAM 사용자
    * 단기 또는 장기 자격 증명을 사용하여 AWS CLI 또는 AWS API에 대한 프로그래밍 방식 요청에 직접 또는 AWS SDK를 사용하여 서명해야 한다.
  * 페더레이션 보안 주체
    * AWS STS API 작업을 사용하여 액세스 키 페어 및 세션 토큰을 포함하는 임시 보안 인증 정보로 새 세션을 생성한다.

# 2. 권한 및 정책이 액세스 관리 제공 방법

AWS Identity and Access Management의 액세스 관리를 통하여 계정에서 보안 주체 엔터티에 허용된 권한을 정의할 수 있다. 보안 주체 엔터티는 IAM 엔터티를 사용하여 인증된 사람 또는 애플리케이션에 한하여 액세스 관리를 통칭적으로 권한 부여라고 뜻한다. 정책을 생성하고 IAM ID 또는 AWS 리소스에 연결하여 자체적으로 액세스를 제어 및 관리를 통해 정책에 대한 평가를 진행하고 권한의 요청이 허용되거나 거부되는 지를 결정한다. 또한 대부분의 정책은 AWS 내부에 JSON 문서 형태로 저장된다.

## 2.1 정책 및 계정

AWS에서 하나의 계정을 관리하려면 정책을 사용한 해당 계정 내 권한을 정의한다. 여러 계정의 전체적인 권한을 관리하고자 한다면 IAM 사용자에 대한 권한을 관리하기가 어렵기 때문에 크로스 계정 권한에 대해 IAM 역할, 리소스 기반 정책 혹은 액세스 제어 목록 (`ACL`)을 사용할 수 있다.

## 2.2 정책 및 사용자



```json
{
  "Version":"2012-10-17",
  "Statement": {
    "Effect": "Allow",
    "Action": "dynamodb:*",
    "Resource": "arn:aws:dynamodb:us-east-2:123456789012:table/Books"
  }
}
```

---

## 참조

[AWS Account root user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#id_root-user-access-management)
[AWS OIDC](https://docs.aws.amazon.com/ko_kr/IAM/latest/UserGuide/id_roles_providers_oidc.html)