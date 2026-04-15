# 24. Auth, API Keys, RBAC Operations

## 대상 독자

초급~중급 운영자 (부서 배포/권한 설계 담당)

## 목적

GoClaw 인증/권한 구조를 이해하고, 안전하게 운영할 수 있게 합니다.

## 먼저 알아야 할 개념

- 인증(Authentication): "누구인지 확인"
- 인가(Authorization): "무엇을 할 수 있는지 결정"
- RBAC: 역할 기반 권한 제어

## 코드 기준 위치

- 인증/미들웨어: `internal/http/auth.go`, `internal/gateway/server.go`
- API 키: `internal/http/api_keys.go`, `internal/gateway/methods/api_keys.go`
- 권한 정책: `internal/permissions/policy.go`
- 멀티테넌트 해석: `docs/23-multi-tenant-architecture.md`

## 인증 수단 정리

1. Gateway token
   - 운영/관리용, 노출 위험이 크므로 제한적으로 사용
2. Tenant-bound API key
   - 부서 서비스 연동 기본값(권장)
3. System-level key
   - 크로스테넌트 작업용, 최소한으로 사용

## HTTP/WS에서 어떻게 적용되나

### HTTP

- `Authorization: Bearer <token>`
- `X-GoClaw-User-Id`로 사용자 식별
- 필요 시 tenant hint를 사용하지만, 기본은 키 바인딩 tenant를 따름

### WebSocket

- 최초 `connect`에서 토큰 기반 인증
- 연결 후 요청/이벤트는 tenant scope가 유지됨

## RBAC/Scope 운영 전략

- admin: 플랫폼 운영자만
- operator: 부서 리더/자동화 담당
- viewer: 조회 중심 사용자

운영 원칙:

- 사람에게는 최소 권한
- 시스템 키는 백엔드 서버에서만 보관
- 키는 주기적으로 회전(rotate)

## 실무 체크리스트

1. 브라우저에 시스템 키를 절대 저장하지 않는다
2. 부서별 tenant-bound key를 분리 발급한다
3. 키별 scope를 문서화한다
4. 관리자 기능 API는 네트워크/IP로 이중 차단한다
5. 폐기/교체 절차(incident runbook)를 만든다

## 자주 발생하는 실수

- gateway token을 테스트 편의로 프론트에 직접 노출
- admin scope를 너무 많은 사용자에게 배포
- tenant-bound key 대신 system key를 상시 사용

## 적용 예시

- 사내 포털 연동:
  - 사용자 -> 사내 포털
  - 포털 백엔드 -> tenant-bound key로 GoClaw 호출
  - 사용자는 GoClaw 관리자 대시보드 직접 접근 안 함

## 다음 문서

[25-mcp-and-grants-operations](./25-mcp-and-grants-operations.md)
