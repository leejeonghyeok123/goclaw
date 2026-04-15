# 22. Tenant vs Account (초보자용 정리)

## 목적

"테넌트"와 "계정(사용자)" 개념을 헷갈리지 않게 구분하고, 실제 운영에서 어떻게 연결되는지 이해합니다.

## 한 줄 정의

- 테넌트(tenant): 데이터/권한이 격리되는 "조직 경계"
- 계정(account/user): 그 조직 안에서 활동하는 "개인 식별자"

## 쉬운 비유

- 테넌트 = 건물
- 계정 = 건물 안 입주자

건물이 다르면 문서/대화/에이전트/세션이 섞이면 안 됩니다.

## 코드 기준 위치

- 멀티테넌트 안내: `docs/23-multi-tenant-architecture.md`
- 컨텍스트 스코프: `internal/store/context.go`
- 테넌트 저장소: `internal/store/tenant_store.go`
- HTTP 인증/테넌트 해석: `internal/http/auth.go`, `internal/gateway/server.go`
- API 키 권한: `internal/gateway/methods/api_keys.go`

## 관계 정리

1. 사용자(계정)는 하나 이상의 테넌트에 속할 수 있음
2. 요청이 들어오면 인증정보(API 키/토큰)로 테넌트가 결정됨
3. 이후 DB 조회는 tenant scope로 제한됨
4. 같은 user_id여도 테넌트가 다르면 다른 데이터 공간으로 동작

## 왜 필요한가

- 고객사 A와 B 데이터가 절대 섞이지 않게 하기 위해
- 운영 권한을 조직 단위로 나누기 위해
- SaaS 환경에서 보안 사고를 방지하기 위해

## 실무 운영 모델

### 싱글 테넌트(개인/소규모)

- 사실상 하나의 master 테넌트만 사용
- 빠르게 시작 가능

### 멀티 테넌트(SaaS/고객사 분리)

- 고객사마다 tenant 생성
- API 키를 tenant-bound로 발급
- 요청 시 자동 tenant scope 적용

## 계정/테넌트/권한이 함께 작동하는 방식

- 인증: API 키 또는 게이트웨이 토큰
- 사용자 식별: `X-GoClaw-User-Id` 또는 WS connect user 정보
- 테넌트 결정: API 키 바인딩 또는 허용된 tenant hint
- 권한 결정: role/scope (`admin`, `operator`, `viewer` 등)

## 자주 묻는 질문

### Q1. 계정만 나누면 되지, 테넌트까지 왜 필요?

계정만 분리하면 같은 DB 공간 안에서 실수로 섞일 위험이 큽니다.  
테넌트는 경계 자체를 분리하는 안전장치입니다.

### Q2. 한 계정이 여러 테넌트를 볼 수 있나?

가능하지만 권한/멤버십이 있어야 합니다.  
운영자(owner/admin)와 일반 사용자의 범위는 다르게 제한됩니다.

### Q3. API 키는 어떻게 쓰는 게 안전한가?

- 브라우저가 아니라 서버 백엔드에만 저장
- 테넌트 바운드 키를 기본으로 사용
- 시스템 키(광범위 권한)는 최소한으로 사용

## 적용 예시

- SaaS 예시:
  - `tenant-acme`, `tenant-globex` 분리
  - 각 백엔드가 자기 tenant API 키로 GoClaw 호출
  - 같은 endpoint를 호출해도 결과는 tenant별로 다름

## 초보자 체크리스트

- 테넌트와 계정 차이를 설명할 수 있다
- 요청 하나가 어떤 tenant로 해석되는지 추적할 수 있다
- 시스템 키와 tenant-bound 키의 차이를 이해한다

## 다음 문서

[18-expert-checklist-and-next-steps](./18-expert-checklist-and-next-steps.md)
