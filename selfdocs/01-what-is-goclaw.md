# 01. What is GoClaw

## 목적

GoClaw의 정체, 문제 해결 범위, 핵심 가치를 정확히 이해합니다.

## 기능 개요

GoClaw는 Go 기반 멀티테넌트 AI Agent Gateway/Platform입니다.

- 다수 LLM Provider 통합
- HTTP + WebSocket + CLI + Web UI 제공
- Team/Task/Delegation 오케스트레이션
- Tool 실행 정책, 보안, 멀티테넌트 격리
- Memory/KG/Vault 기반 장기 지식화
- Desktop Lite(로컬 SQLite) 지원

## 코드 기준 위치

- 엔트리포인트: `main.go`
- 앱 조립: `cmd/gateway.go`
- 프로토콜 타입: `pkg/protocol/`
- API 엔드포인트: `internal/http/`
- 에이전트 루프: `internal/agent/`, `internal/pipeline/`

## 어디에서 무엇을 해야 하나

1. 제품 정체 확인
   - `README.md`의 Core Features 섹션 확인
   - 공식 사이트: <https://goclaw.sh/>
2. 공식 문서 범위 확인
   - <https://docs.goclaw.sh/>
3. 로컬 코드와 매핑
   - README 기능 목록을 `internal/*` 패키지와 일대일 매칭

## 적용 예시

- "우리는 Slack + OpenAI + 사내 문서 검색 에이전트가 필요" 같은 요구를 GoClaw로 빠르게 조합 가능:
  - 채널: `internal/channels/slack`
  - 프로바이더: `internal/providers/openai_*.go`
  - 문서 검색: `internal/vault/`

## 활용 전략

- 제품 소개 시 "기능 나열"보다 "운영 흐름(요청->루프->툴->메모리->응답)"으로 설명하면 설득력이 높음

## 체크리스트

- GoClaw가 단일 챗봇이 아니라 "에이전트 운영 플랫폼"임을 설명 가능
- Standard와 Lite 차이를 설명 가능

## 다음 문서

[02-environment-prerequisites](./02-environment-prerequisites.md)
