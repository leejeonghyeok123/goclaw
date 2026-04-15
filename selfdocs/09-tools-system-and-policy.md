# 09. Tools System and Policy

## 목적

툴 등록/실행/보안 정책 구조를 이해하고, 안전하게 확장합니다.

## 핵심 위치

- 레지스트리: `internal/tools/registry.go`
- 정책: `internal/tools/policy.go`
- 타입/컨텍스트: `internal/tools/types.go`
- 실행 스크럽: `internal/tools/scrub.go`

## 기능 개요

- 파일시스템/실행/web/memory/media/teams/automation 등 다수 툴 제공
- 정책 엔진으로 허용/차단/승인 흐름 제어
- 민감정보 마스킹/스크러빙

## 사용 상세 절차

1. 새 툴 추가
   - 툴 구현 파일 생성
   - 레지스트리 등록
2. 정책 연결
   - allow/deny/rate limit/approval 규칙 반영
3. 인터페이스 노출
   - 에이전트 런타임과 필요 시 HTTP invoke 라우트 점검

## 적용 예시

- 사내 전용 API 조회 툴 추가:
  - HTTP 호출 로직 툴 구현
  - SSRF/도메인 allowlist 정책 적용
  - 팀 운영용 에이전트에만 권한 부여

## 활용 전략

- 툴 자체 기능보다 "정책/감사/승인"을 먼저 설계하면 운영 안정성이 높음

## 체크리스트

- 새 툴은 레지스트리 + 정책 + 테스트를 함께 반영
- 민감정보 로그 노출 없음

## 다음 문서

[10-providers-model-routing](./10-providers-model-routing.md)
