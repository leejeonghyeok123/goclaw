# 06. HTTP / WS / CLI Interfaces

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

GoClaw의 3개 인터페이스를 목적별로 올바르게 사용하는 방법을 익힙니다.

## 코드 위치

- HTTP: `internal/http/`
- WS: `internal/gateway/methods/`, `pkg/protocol/`
- CLI: `cmd/*.go`

## 인터페이스 선택 기준

- HTTP: 외부 시스템 연동, REST 기반 자동화
- WS: 실시간 스트리밍/이벤트 중심 대화
- CLI: 운영자 로컬 작업/관리

## 상세 사용 절차

1. HTTP
   - 기본 베이스 URL(Docker Compose): `http://localhost:18790` (호스트 포트는 `GOCLAW_PORT` 로 변경 가능, `docker-compose.yml` 참고)
   - `/v1/openapi.json`으로 계약 확인
   - 예: `/v1/chat/completions`, `/v1/providers`, `/v1/skills`
2. WS
   - 첫 프레임은 반드시 `connect`
   - 그 다음 `chat.send` 등 메서드 요청
3. CLI
   - `goclaw providers ...`, `goclaw skills ...`, `goclaw migrate ...`

## 적용 예시

- 사내 포털에서 에이전트 채팅 호출:
  - HTTP `POST /v1/chat/completions`
- 실시간 중간 토큰 표시:
  - WS `chat.send` + 이벤트 스트림 소비

## 활용 전략

- B2B 연동은 HTTP 중심, 오퍼레이터 콘솔은 CLI 중심, 실시간 UX는 WS 중심으로 분리 설계

## 체크리스트

- connect-first WS 규칙을 알고 있음
- HTTP/WS/CLI 중 상황별 선택이 가능

## 다음 문서

[07-store-migrations-and-tenancy](./07-store-migrations-and-tenancy.md)
