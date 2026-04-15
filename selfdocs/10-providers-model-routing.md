# 10. Providers and Model Routing

## 목적

다중 LLM 공급자 구성을 안정적으로 운영하고 모델 라우팅을 제어합니다.

## 핵심 위치

- Provider 구현: `internal/providers/`
- Provider wiring: `cmd/gateway_providers.go`
- API/관리: `internal/http/providers.go`, `cmd/providers_cmd.go`

## 기능 개요

- Anthropic/OpenAI-compat/Codex/DashScope/Claude CLI 등 다수 지원
- 공통 어댑터 인터페이스로 통합
- 검증, 모델 목록, 설정 저장, 키 암호화 처리

## 사용 상세 절차

1. 프로바이더 등록
   - UI 또는 HTTP `/v1/providers` 또는 CLI
2. 연결 검증
   - `/v1/providers/{id}/verify`
3. 모델 선택/기본값 설정
   - provider settings 및 agent 설정 연동

## 적용 예시

- 운영비 절감:
  - 일반 질의는 저비용 모델, 고난도는 고성능 모델로 분리
  - 팀/에이전트 단위로 provider profile 분리

## 활용 전략

- 모델 장애 대비를 위해 최소 2개 provider를 준비해 페일오버 전략을 수립

## 체크리스트

- provider verify 통과
- 에이전트별 모델 정책이 문서화됨

## 다음 문서

[11-channels-and-integrations](./11-channels-and-integrations.md)
