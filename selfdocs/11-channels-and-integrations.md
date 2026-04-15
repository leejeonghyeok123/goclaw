# 11. Channels and Integrations

## 목적

외부 메신저/채널과 GoClaw 에이전트를 연결해 실제 사용자 접점을 운영합니다.

## 핵심 위치

- 채널 매니저: `internal/channels/manager.go`
- 인스턴스 로더: `internal/channels/instance_loader.go`
- 채널별 구현: `internal/channels/*`
- 관리 API: `internal/http/channel_instances.go`
- CLI: `cmd/channels_cmd.go`

## 지원 범주

- Telegram, Discord, Slack, Zalo, Feishu/Lark, WhatsApp 등

## 사용 상세 절차

1. 채널 자격증명 준비
   - 토큰/시크릿/웹훅 정보
2. 채널 인스턴스 생성
   - HTTP `/v1/channels/instances` 또는 UI/CLI
3. 에이전트와 바인딩
   - 대상 agent/team 연결
4. 정책 설정
   - DM/그룹 처리 규칙, 응답 포맷, 권한 범위 설정

## 적용 예시

- Slack 고객지원 봇:
  - Slack 채널 인스턴스 생성
  - 지원 전용 에이전트 연결
  - `team_tasks`와 연동해 티켓 전환

## 활용 전략

- 채널별 메시지 포맷 차이를 어댑터 계층에서 흡수하고, 도메인 로직은 공통 유지

## 체크리스트

- 인스턴스 연결 상태 정상
- 채널에서 실제 왕복 메시지 테스트 완료

## 다음 문서

[12-memory-kg-vault](./12-memory-kg-vault.md)
