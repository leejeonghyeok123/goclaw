# 19. Agent Teams Deep Dive

## 대상 독자

중급~고급 (입문자는 먼저 `21-sessions-and-context-for-beginners.md`, `22-tenant-vs-account-explained.md` 권장)

## 목적

Agent Team의 개념과 실제 운영 방법(생성, 멤버링, 작업 분배, 위임, 모니터링)을 실무 기준으로 정리합니다.

## Team 개념 정리

GoClaw의 Team은 여러 에이전트를 묶어 협업시키는 운영 단위입니다.

- 단일 에이전트: 한 에이전트가 모든 업무 수행
- 팀 기반: 역할별 에이전트로 분리(예: PM, 리서처, 코더, 리뷰어)
- 오케스트레이션: task board + delegation + team messaging

## 코드 기준 위치

- WS 팀 메서드: `internal/gateway/methods/teams.go`, `teams_crud.go`, `teams_tasks*.go`
- 팀 툴 계층: `internal/tools/team_tasks_tool.go`, `team_tool_dispatch.go`, `delegate_tool.go`
- 팀 저장소:
  - PostgreSQL: `internal/store/pg/teams*.go`, `teams_tasks*.go`
  - SQLite: `internal/store/sqlitestore/teams*.go`, `teams_tasks*.go`
- 오케스트레이션 모드: `internal/agent/orchestration_mode.go`

## Team으로 작동시키는 상세 절차

1. 팀 생성
   - UI Teams 페이지 또는 WS `teams.*` 계열 메서드 사용
2. 팀 멤버(에이전트) 연결
   - 역할 기반으로 멤버 구성(기획/실행/검수)
3. 작업 생성
   - `team_tasks`(create)로 작업 등록
4. 담당자 배정
   - assignee를 지정하거나 팀 정책으로 자동 분배
5. 위임 실행
   - `delegate`/`spawn`을 통해 하위 작업 분기
6. 상태 전이 관리
   - TODO -> IN_PROGRESS -> DONE 등 라이프사이클 관리
7. 결과 통합
   - 팀 메시지/태스크 결과를 기반으로 최종 응답 산출

## 운영 예시

- 시나리오: "신규 기능 배포"
  - PM 에이전트: 요구사항 정리 + task 분할
  - 코더 에이전트: 구현
  - 리뷰어 에이전트: 검토/수정요청
  - 릴리즈 에이전트: 체크리스트 검증 후 배포 안내

## 자주 하는 실수

- 팀은 만들었지만 실제 위임 규칙(`delegate`, `team_tasks`)을 연결하지 않음
- 역할이 겹치는 멤버를 과도하게 넣어 작업 충돌 발생
- 상태 전이 규칙 없이 task만 생성해 보드가 금방 무의미해짐

## 활용 전략

- 작은 팀(2~3 에이전트)으로 시작해 역할 분리 효과를 먼저 검증
- 팀 작업에 SLA(응답시간/완료조건)를 붙여 운영 품질을 계량화
- 실패 태스크를 재시도하기보다 원인 분류(권한/툴/컨텍스트) 후 재할당

## 체크리스트

- 팀 생성/멤버링/태스크 생성/위임/완료 흐름이 end-to-end로 작동
- `team_tasks`, `delegate`, `spawn`의 목적 차이를 설명 가능
- 팀 운영 로그를 기반으로 병목 구간을 찾을 수 있음

## 다음 문서

[20-cron-and-trigger-patterns](./20-cron-and-trigger-patterns.md)
