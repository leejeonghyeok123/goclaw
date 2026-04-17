# 25. MCP and Grants Operations

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 대상 독자

중급 운영자/개발자 (외부 도구 연결 담당)

## 목적

MCP 서버 등록, 권한 부여(grants), 사용자 자격증명 정책을 안전하게 운영합니다.

## MCP를 쉽게 설명하면

MCP는 "외부 도구를 표준 방식으로 에이전트에 연결하는 브릿지"입니다.

예:

- 사내 검색 API
- 티켓 시스템
- 사내 지식베이스

## 코드 기준 위치

- MCP 코어: `internal/mcp/*`
- MCP HTTP: `internal/http/mcp*.go`
- MCP 도구 연동: `internal/tools/mcp_*`

## 운영 절차

1. MCP 서버 정의 등록
2. 연결 테스트
3. 접근 권한(grant) 정책 설정
4. 필요 시 사용자별 credential 요구 설정
5. 에이전트에서 실제 호출 테스트

## grant가 중요한 이유

- 연결만 해두면 누구나 쓰는 구조가 되면 위험
- 어떤 agent/user가 어떤 MCP를 쓸지 분리해야 보안 유지

## require_user_credentials 사용 시점

- 사용자별 API 키/권한이 달라야 하는 외부 시스템
- 감사 추적에서 "누가 호출했는지" 명확히 필요할 때

## 장애/실패 시 점검 포인트

1. 네트워크 도달성
2. 서버 설정값(transport, endpoint)
3. grants 누락 여부
4. 사용자 credential 누락 여부
5. tool policy 차단 여부

## 적용 예시

- Jira MCP 연결:
  - 운영팀이 서버 등록
  - 개발팀 agent에만 grant
  - 개인 토큰 필요 정책 적용(require_user_credentials=true)

## 체크리스트

- MCP 등록/테스트/grant/user credential 경로를 모두 검증
- 권한 없는 agent가 호출 시 차단되는지 확인

## 다음 문서

[26-hooks-and-automation-guardrails](./26-hooks-and-automation-guardrails.md)
