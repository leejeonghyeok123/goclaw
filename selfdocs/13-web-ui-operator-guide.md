# 13. Web UI Operator Guide

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

운영자가 웹 대시보드에서 주요 기능을 실무적으로 다루는 방법을 익힙니다.

## 핵심 위치

- 앱 시작: `ui/web/src/main.tsx`
- 라우팅: `ui/web/src/routes.tsx`
- 페이지: `ui/web/src/pages/*`
- 상태관리: `ui/web/src/stores/*`
- i18n: `ui/web/src/i18n/locales/*`

## 주요 운영 화면

- Agents / Chat / Teams / Providers / Skills
- Channels / MCP / Vault / Memory / Traces / Usage
- Tenants / API Keys / Backup-Restore

## 사용 상세 절차

1. 관리자 로그인/토큰 확보
   - Docker 기본: 브라우저에서 `http://localhost:18790` (호스트 포트는 `GOCLAW_PORT`, `docker-compose.yml` 참고). 게이트웨이 토큰은 `.env` 의 `GOCLAW_GATEWAY_TOKEN` 등.
2. Providers 설정
3. Agent 생성 및 Tool 정책 확인
4. Team/Task/Delegation 흐름 점검
5. Trace/Usage로 품질/비용 모니터링

## 적용 예시

- 운영 대시보드에서 Provider 장애 감지 후 대체 모델로 전환

## 활용 전략

- 페이지별 "읽기 전용 역할"과 "변경 권한 역할"을 명확히 분리해 사고를 줄임

## 체크리스트

- 주요 페이지 진입/조회/수정 권한이 예상대로 작동
- 다국어 문자열 추가 시 en/vi/zh 동시 반영

## 다음 문서

[14-desktop-lite-guide](./14-desktop-lite-guide.md)
