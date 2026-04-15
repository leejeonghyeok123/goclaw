# 04. Project Structure Complete Map

## 목적

전체 폴더의 책임을 빠르게 파악해, 변경 시 "어디를 고쳐야 하는지" 즉시 판단합니다.

## 최상위 디렉토리 역할

- `cmd/`: CLI와 gateway 조립
- `internal/`: 핵심 비즈니스/런타임
- `pkg/`: 재사용 가능 프로토콜/브라우저 패키지
- `ui/web/`: 웹 대시보드
- `ui/desktop/`: Wails 데스크탑 앱
- `migrations/`: PostgreSQL 스키마 변경
- `tests/`: invariant/contract/scenario/integration
- `docs/`: 설계 문서
- `.github/workflows/`: CI/CD

## internal 하위 핵심

- `agent/`, `pipeline/`: 에이전트 루프
- `tools/`: 툴 레지스트리와 정책
- `providers/`: LLM 공급자 어댑터
- `store/pg`, `store/sqlitestore`: DB 구현
- `gateway/`, `http/`: WS/HTTP 인터페이스
- `channels/`: Telegram/Slack/Discord 등
- `memory/`, `knowledgegraph/`, `vault/`: 지식 계층

## 어디에서 무엇을 해야 하나

- 새 HTTP API 추가: `internal/http` + `cmd/gateway_http_wiring.go`
- WS 메서드 추가: `internal/gateway/methods` + `cmd/gateway_methods.go`
- DB 컬럼 추가: `migrations` + `internal/store/sqlitestore/schema*`
- UI 페이지 추가: `ui/web/src/pages` + `ui/web/src/routes.tsx`

## 적용 예시

- "에이전트 설정 화면에 옵션 추가" 작업 시:
  - API: `internal/http/agents*.go`
  - 저장: `internal/store/pg/agents*.go` + sqlite 대응
  - UI: `ui/web/src/pages/agents*`

## 체크리스트

- 신규 기능을 넣을 디렉토리를 1분 안에 판단 가능
- 변경 영향 범위(backend + UI + tests)를 같이 생각함

## 다음 문서

[05-runtime-entrypoints-and-wiring](./05-runtime-entrypoints-and-wiring.md)
