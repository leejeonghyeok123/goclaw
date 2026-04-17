# 21. Sessions and Context (Beginner Guide)

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

세션과 컨텍스트를 처음 보는 사람도 "대화가 왜 이어지고, 왜 가끔 잊어버리는지" 이해하도록 정리합니다.

## 한 줄 정의

- 세션(session): 대화의 "대화방 ID"
- 컨텍스트(context): 에이전트가 현재 답변할 때 참고하는 "재료 묶음"

## 비유로 이해하기

- 세션 = 카카오톡 채팅방
- 컨텍스트 = 채팅방 최근 대화 + 메모 + 규칙 + 첨부자료

즉, 같은 에이전트라도 **세션이 다르면 대화 맥락이 달라질 수 있습니다.**

## 코드 기준 위치

- 세션 관리: `internal/sessions/manager.go`, `internal/gateway/methods/sessions.go`
- 채팅 처리: `internal/gateway/methods/chat.go`
- 파이프라인 컨텍스트 구성: `internal/pipeline/context_stage.go`, `internal/pipeline/pipeline.go`
- 프롬프트 조립: `internal/agent/prompt_builder_impl.go`

## 실제로 어떻게 동작하나 (단계)

1. 사용자가 메시지 전송
2. 서버가 `sessionKey`를 기준으로 기존 대화를 찾음
3. 파이프라인이 컨텍스트를 조립
   - 최근 히스토리
   - 시스템/에이전트 지시문
   - 도구 실행 결과 일부
   - 메모리/Vault/KG 검색 결과(설정/상황에 따라)
4. 모델 호출 후 응답 생성
5. 응답/상태를 다시 세션에 저장

## 왜 세션이 중요한가

- 같은 세션: "아까 말한 것"을 이어서 답하기 쉬움
- 새 세션: 초기화된 대화처럼 동작
- 잘못된 세션키 사용: 엉뚱한 맥락 섞임

## 컨텍스트가 항상 완전하지 않은 이유

- 토큰 제한 때문에 오래된 내용은 요약/제외될 수 있음
- 컨텍스트 pruning/compaction 정책이 동작할 수 있음
- 메모리/KG는 비동기 업데이트라 즉시 반영이 늦을 수 있음

## 실무 사용법 (HTTP/WS 기준)

### HTTP

- 같은 사용자가 같은 업무를 이어갈 때는 같은 `session` 관련 키를 유지
- 업무를 분리하고 싶으면 새 세션키 사용

### WebSocket

- `chat.send` 시 세션 식별자를 일관되게 사용
- 실시간 앱에서는 탭/업무 단위로 세션을 분리

## 적용 예시

- 고객지원:
  - 고객 1명당 1세션 유지 -> 맥락 유지 강함
- 운영자동화:
  - 작업유형별 세션 분리(장애대응/릴리즈/문서화)

## 자주 헷갈리는 포인트

- "에이전트가 같으면 기억도 완전 공유된다?" -> 아니오. 세션 단위 영향이 큼
- "메모리가 있으니 세션 없어도 된다?" -> 아니오. 세션은 즉시 대화 연속성의 핵심
- "컨텍스트는 DB 전체를 다 넣는다?" -> 아니오. 선택/요약/우선순위 기반으로 구성

## 초보자 체크리스트

- 세션과 컨텍스트 차이를 설명할 수 있다
- 같은 세션/새 세션 동작 차이를 테스트해봤다
- 내 앱에서 세션키 생성 규칙(사용자/업무 단위)을 정의했다

## 다음 문서

[22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md)
