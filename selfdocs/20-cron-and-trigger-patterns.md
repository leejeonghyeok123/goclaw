# 20. Cron and Trigger Patterns

## 대상 독자

중급 (입문자는 먼저 `06-http-ws-cli-interfaces.md`와 `99-glossary-for-beginners.md` 권장)

## 목적

GoClaw에서 에이전트를 자동으로 작동시키는 트리거 구조를 정확히 이해하고 운영에 적용합니다.

## 핵심 결론

"트리거가 별도로 없고 REST/WS 요청만 보내면 끝"은 절반만 맞습니다.

- 외부 트리거: 맞음. REST/WS 호출로 언제든 실행 가능
- 내부 트리거: 있음. `cron`과 `heartbeat`가 내장 스케줄 트리거 역할 수행
- 이벤트 트리거: 채널(텔레그램/슬랙 등) 인입 메시지도 실행 트리거

## 코드 기준 위치

- Cron 서비스: `internal/cron/service.go`, `internal/cron/service_execution.go`
- Cron RPC: `internal/gateway/methods/cron.go`
- Heartbeat: `internal/heartbeat/ticker.go`
- Heartbeat RPC: `internal/gateway/methods/heartbeat.go`
- 스케줄링 레인: `internal/scheduler/`
- 채널 인입: `internal/channels/manager.go` 및 채널별 구현

## 트리거 유형 정리

1. Request Trigger (수동/외부)
   - HTTP API 호출 (`/v1/chat/completions` 등)
   - WS 메서드 호출 (`chat.send` 등)
2. Time Trigger (내장 자동화)
   - cron job: `at`, `every`, `cron expr`
   - heartbeat: 에이전트 주기 점검/알림
3. Event Trigger (채널/시스템 이벤트)
   - 메신저 인입 메시지
   - 내부 이벤트 기반 후속 처리

## Cron 상세 사용법

1. 스케줄 정의
   - one-shot: `at`
   - 고정 주기: `every`
   - 크론식: `cron`
2. 대상 agent/session 지정
3. job enable/disable 관리
4. run logs로 성공/실패/재시도 확인

운영 팁:
- 사용자 채팅과 충돌하지 않게 cron lane 동시성 설정을 점검
- 실패 재시도(backoff) 로그를 보고 과도 재시도 방지

## Heartbeat 상세 사용법

1. 에이전트별 heartbeat 설정
2. `HEARTBEAT.md` 체크리스트 작성
3. active hours/채널/모델 override 설정
4. `HEARTBEAT_OK` 기반 suppress 동작 확인

운영 팁:
- heartbeat는 "헬스/점검 자동화"에 적합
- 일반 업무 예약은 cron이 더 적합

## 적용 예시

- 매일 09:00 리포트 생성
  - cron으로 리포트 프롬프트 실행
- 10분마다 운영 상태 점검
  - heartbeat로 체크리스트 수행 후 이상 시에만 알림
- 실시간 고객 문의 대응
  - Slack/Telegram 인입 메시지 자체를 트리거로 사용

## 질문에 대한 명확한 답

질문: "agent 작동 트리거는 별도 없고, 트리거 발생 시 REST/WS 요청이 끝 아닌가?"

정리:
- 외부 시스템에서 트리거를 걸 때는 보통 REST/WS 요청이 맞음
- 하지만 GoClaw 내부에도 cron/heartbeat라는 "별도 트리거 시스템"이 존재함
- 즉, 운영 관점에서는 "외부 트리거 + 내부 트리거"를 함께 설계해야 함

## 체크리스트

- 수동 실행(HTTP/WS)과 자동 실행(cron/heartbeat) 차이를 설명 가능
- cron/heartbeat 로그에서 실패 원인 추적 가능
- 업무 성격에 맞는 트리거 방식을 선택 가능
