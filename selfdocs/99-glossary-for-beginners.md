# 99. Glossary for Beginners

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

처음 보는 용어를 빠르게 이해할 수 있도록 핵심 개념을 짧게 정리합니다.

## 핵심 용어

- Agent: 사용자의 요청을 처리하는 AI 실행 단위
- Session: 대화를 이어가는 대화방 식별자
- Context: 현재 답변에 사용되는 입력 재료 묶음
- Tenant: 데이터/권한 격리 단위(조직 경계)
- User(Account): tenant 안에서 활동하는 사용자 식별자
- Provider: LLM 서비스 제공자(OpenAI, Anthropic 등)
- Model Routing: 요청별로 어떤 모델을 쓸지 결정하는 규칙
- Tool: 에이전트가 수행하는 실행 기능(파일, 웹, 실행, 팀작업 등)
- Policy: tool/권한 사용 규칙(허용, 차단, 승인)
- Pipeline: 요청을 단계별로 처리하는 내부 흐름(context->history->...->summarize)
- Working Memory: 현재/최근 대화 중심 기억
- Episodic Memory: 대화 요약 기반 기억
- Semantic Memory/KG: 엔티티-관계 중심 장기 지식 그래프
- Vault: 문서 저장/검색 레이어
- Cron: 시간 기반 자동 실행
- Heartbeat: 에이전트 상태 점검 자동 실행
- Delegate/Spawn: 다른 에이전트에게 하위 작업 위임/생성
- RBAC: 역할 기반 권한 제어(admin/operator/viewer)
- API Key Scope: API 키가 가진 기능 범위. 코드상 문자열은 `operator.read`, `operator.write`, `operator.admin` 등 (`internal/permissions/policy.go`, [24](./24-auth-api-keys-rbac-operations.md))
- Master Scope: 시스템 전역 관리 범위
- Tenant Scope: 특정 tenant 내부로 제한된 범위

## 초보자 읽기 순서 팁

1. 이 문서(용어집) 먼저 읽기
2. `21-sessions-and-context-for-beginners.md`
3. `22-tenant-vs-account-explained.md`
4. 나머지 심화 문서 순서대로 읽기
