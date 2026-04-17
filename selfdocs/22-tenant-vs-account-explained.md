# 22. Tenant vs Account — GoClaw 기준 구축·활용 가이드

> **Selfdocs 공통 전제:** **Linux** + **Docker Compose**(`make up` 등). 초기 기동: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 인증·스코프 심화: [24-auth-api-keys-rbac-operations](./24-auth-api-keys-rbac-operations.md).

## 목적

- **테넌트(tenant)** 와 **계정(user / `X-GoClaw-User-Id`)** 을 구분한다.
- Linux + Docker 로 띄운 GoClaw에서 **테넌트를 실제로 만들고**, **API 키와 요청 스코프까지** 연결해 쓰는 흐름을 따라 할 수 있게 한다.

## 한 줄 정의

| 용어 | 의미 (GoClaw) |
|------|----------------|
| **테넌트** | DB와 파일 워크스페이스에서 데이터가 격리되는 **조직 단위** (`tenants` 테이블 행, `slug`로 디렉터리도 잡힘). |
| **계정(사용자)** | 그 테넌트 안에서 활동하는 **외부 사용자 ID** 문자열. HTTP에서는 주로 `X-GoClaw-User-Id` 로 넘긴다. |

같은 `user_id`라도 **테넌트가 다르면** 세션·에이전트·메모리 등 tenant-scoped 데이터가 완전히 다르다.

## 마스터 테넌트(기본 공간)

처음 온보딩만 하고 별도 테넌트를 만들지 않으면, 대부분의 데이터는 **마스터 테넌트** 하나에 모인다.

- 고정 ID: `0193a5b0-7000-7000-8000-000000000001` (`internal/store/tenant_store.go` 의 `MasterTenantID`)
- 개인/소규모/단일 조직: **추가 테넌트 없이 마스터만** 써도 된다.
- SaaS·고객사별 격리: **고객마다 테넌트 행을 추가**하고, **테넌트에 묶인 API 키**로만 붙인다.

## 테넌트를 “구축”한다는 것 (GoClaw에서의 실체)

1. **DB에 `tenants` 레코드 생성** — 이름(`name`), URL/경로용 **`slug`** (소문자·숫자·하이픈, 하이픈으로 시작/끝 불가: `internal/http/validate.go` 의 `isValidSlug`).
2. **(서버가 워크스페이스를 쓰는 경우)** 디렉터리 생성: `{workspace}/tenants/{slug}` (`internal/http/tenants.go` 의 `handleCreate`).
3. **그 테넌트에 사용자 멤버십 부여** — `tenant_users` 에 `user_id` + 역할 (`owner`, `admin`, `operator`, `member`, `viewer` 등).
4. **해당 테넌트에 바인딩된 API 키 발급** — 키 레코드에 `tenant_id`가 저장되면, 이후 요청은 **자동으로 그 테넌트 스코프**가 된다.

## 요청마다 테넌트가 어떻게 정해지나 (핵심)

HTTP 인증은 `internal/http/auth.go` 의 `resolveAuth` → `enrichContext` 로 요청 `context` 에 `tenant_id`가 들어가고, 스토어 계층이 이를 따라 `WHERE tenant_id = …` 로 격리한다.

요약 표:

| 인증 방식 | 테넌트 결정 |
|-----------|-------------|
| **테넌트에 묶인 API 키** (`api_keys.tenant_id` 가 nil 이 아님) | **항상 그 키의 `tenant_id`**. 별도 헤더 불필요. **SaaS 연동에 권장.** |
| **시스템 레벨 API 키** (`tenant_id` 가 nil) | `X-GoClaw-Tenant-Id` 에 **UUID 또는 `slug`** 를 넣어 스코프 (없으면 마스터로 떨어질 수 있음). |
| **게이트웨이 토큰** (`Authorization: Bearer` + 설정된 게이트웨이 시크릿) + **Owner 사용자** | `X-GoClaw-Tenant-Id` 로 특정 테넌트를 가리킬 수 있음(교차 테넌트 관리). Owner 목록은 환경변수 `GOCLAW_OWNER_IDS` (쉼표 구분); 비어 있으면 기본 Owner 는 `"system"` (`internal/http/auth.go`, `InitOwnerIDs`). |
| **게이트웨이 토큰** + **Owner가 아닌 사용자** | 멤버십이 검증된 테넌트 힌트만 허용; 힌트 없으면 마스터 등 정책에 따름. |
| **인증 없음**(개발용으로 게이트웨이 토큰이 비어 있는 경우) | 마스터 테넌트·고정 사용자 등 **개발 모드 규칙** 적용 — 운영 환경에서는 사용하지 말 것. |

**계정(`X-GoClaw-User-Id`)** 은 세션·사용자별 데이터를 나누는 데 쓰이고, **테넌트**는 조직 경계이다. SaaS에서는 보통 **백엔드마다 하나의 tenant-bound 키** + **앱의 end-user id 를 `X-GoClaw-User-Id` 로 전달**하는 패턴이 된다.

## Linux + Docker: 서버를 띄운 뒤 테넌트까지 만드는 예시

전제: 저장소 루트에서 `.env` 준비 후 `make up` 으로 기동했다고 가정한다. 상세는 [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md) 을 본다.

환경에서 **게이트웨이 토큰**을 확인한다 (예: `.env` 의 `GOCLAW_GATEWAY_TOKEN`). API 기본 포트는 compose 기준으로 보통 **18790** (`docker-compose.yml` 주석).

아래는 **Owner** 로만 가능한 **테넌트 HTTP API** 이다 (`internal/http/tenants.go`: `IsOwnerRole` 검사).  
즉 `Authorization` 에 게이트웨이 토큰을 넣고, `X-GoClaw-User-Id` 는 `GOCLAW_OWNER_IDS` 에 포함된 값(미설정 시 `system`)을 사용한다.

### 1) 테넌트 생성

`POST /v1/tenants` — 바디: `name`, `slug`.

```bash
BASE=http://localhost:18790
TOKEN='여기에_GOCLAW_GATEWAY_TOKEN'
OWNER_USER=system

curl -sS -X POST "$BASE/v1/tenants" \
  -H "Authorization: Bearer $TOKEN" \
  -H "X-GoClaw-User-Id: $OWNER_USER" \
  -H "Content-Type: application/json" \
  -d '{"name":"Acme Corp","slug":"acme"}'
```

응답 JSON 안의 `id` 가 새 테넌트 UUID다. 이후 `TENANT_ID` 변수로 쓴다.

### 2) 테넌트에 사용자(멤버십) 추가

대시보드/게이트웨이 사용자가 특정 테넌트에 들어가려면 **멤버십**이 필요할 수 있다(`resolveTenantHint` 등). 운영자가 넣어 줄 `user_id` 는 **당신의 SaaS가 쓰는 사용자 ID 문자열**과 맞추면 된다.

`POST /v1/tenants/{id}/users` — 바디: `user_id`, `role` (예: `admin`, `operator`, `viewer`).

```bash
TENANT_ID='위에서_받은_uuid'

curl -sS -X POST "$BASE/v1/tenants/$TENANT_ID/users" \
  -H "Authorization: Bearer $TOKEN" \
  -H "X-GoClaw-User-Id: $OWNER_USER" \
  -H "Content-Type: application/json" \
  -d '{"user_id":"saas-user-123","role":"operator"}'
```

### 3) 그 테넌트 전용 API 키 발급

`POST /v1/api-keys` (`internal/http/api_keys.go`). Owner 는 바디에 `tenant_id` 를 넣어 **해당 테넌트에 귀속된 키**를 만들 수 있다.

스코프 문자열은 `internal/permissions/policy.go` 의 `operator.*` 계열이 유효하다. 예: 채팅 호출 위주라면 `operator.read`, `operator.write` 조합이면 **Operator 역할**이 된다(`RoleFromScopes`).  
**새 API 키를 만들려면** HTTP 라우트가 `RoleAdmin` 이상을 요구하므로, 키 발급용으로는 **`operator.admin` 포함**이 필요하다.

```bash
curl -sS -X POST "$BASE/v1/api-keys" \
  -H "Authorization: Bearer $TOKEN" \
  -H "X-GoClaw-User-Id: $OWNER_USER" \
  -H "Content-Type: application/json" \
  -d "{
    \"name\": \"acme-backend-prod\",
    \"tenant_id\": \"$TENANT_ID\",
    \"scopes\": [\"operator.read\", \"operator.write\"]
  }"
```

응답에 **한 번만 노출되는 `key`** 를 고객사 백엔드 시크릿에 저장한다.

### 4) 고객사 백엔드에서 실제 호출 (테넌트는 자동)

테넌트 바운드 키로 호출할 때는 **`X-GoClaw-Tenant-Id` 없이** 동작한다. 사용자만 구분하면 된다.

```bash
ACME_KEY='goclaw_sk_....'

curl -sS -X POST "$BASE/v1/chat/completions" \
  -H "Authorization: Bearer $ACME_KEY" \
  -H "X-GoClaw-User-Id: end-user-999" \
  -H "Content-Type: application/json" \
  -d '{"model":"agent:default","messages":[{"role":"user","content":"hello"}]}'
```

Acme 키로는 Acme 테넌트 데이터만 보인다. 다른 테넌트 키로는 같은 URL을 쳐도 데이터가 섞이지 않는다.

## 테넌트를 “어떻게 활용”하나 (운영 패턴)

| 목적 | 권장 |
|------|------|
| 고객사 / 부서 / 환경(prod vs staging) 분리 | 테넌트를 나누고 **키를 테넌트별로** 발급 |
| B2B SaaS | 고객사 = 테넌트 1개 + 백엔드에 tenant-bound 키 1개(또는 로테이션용 여러 개) |
| 플랫폼 관리자만 교차 조회 | 게이트웨이 토큰 + Owner `user_id` + 필요 시 `X-GoClaw-Tenant-Id` |
| 단일 팀 자가 호스팅 | 추가 테넌트 없이 **마스터만** 사용해도 됨 |

## 코드·공식 문서 위치

- HTTP 테넌트 CRUD·멤버: `internal/http/tenants.go` (`GET/POST /v1/tenants`, `.../users`, …)
- HTTP 인증·테넌트 해석: `internal/http/auth.go`
- API 키 생성(테넌트 지정): `internal/http/api_keys.go`
- 컨텍스트에 심는 값: `internal/store/context.go`
- 아키텍처·연결 유형 표: [docs/23-multi-tenant-architecture.md](../docs/23-multi-tenant-architecture.md)

## 자주 묻는 질문

### Q. 계정만 나누면 되지 않나?

`X-GoClaw-User-Id` 만으로는 **같은 테넌트 안**에서의 구분이다. **조직 단위 격리**는 `tenant_id` (키 또는 헤더로 결정)가 담당한다.

### Q. `slug` 를 잘못 쓰면?

`isValidSlug` 에 맞지 않으면 `400` 이 된다. 예: `Acme` (대문자), `-acme` (하이픈 시작), `acme_` (언더스코어 불가 패턴).

### Q. 테넌트 생성은 Operator API 키로 되나?

HTTP `/v1/tenants` 는 핸들러에서 **Owner 전용**이다. 일반 운영 키로는 생성하지 말고, **플랫폼 관리 파이프라인(게이트웨이 토큰 + Owner user)** 으로 provisioning 하면 된다.

## 체크리스트

- [ ] 마스터 테넌트와 “추가 테넌트”를 언제 쓸지 결정했다.
- [ ] 테넌트 생성 → 멤버 추가 → tenant-bound API 키 발급 순서를 실행해 봤다.
- [ ] SaaS 백엔드에는 tenant-bound 키만 넣고, `X-GoClaw-User-Id` 로 최종 사용자를 식별한다.
- [ ] 시스템 레벨 키 + `X-GoClaw-Tenant-Id` 가 필요한 예외 케이스를 구분한다.

## 다음 문서

- [24-auth-api-keys-rbac-operations](./24-auth-api-keys-rbac-operations.md)
- 심화 아키텍처: [docs/23-multi-tenant-architecture.md](../docs/23-multi-tenant-architecture.md)
