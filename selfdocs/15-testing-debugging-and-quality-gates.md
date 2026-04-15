# 15. Testing, Debugging, and Quality Gates

## 목적

변경사항을 안전하게 검증하고 회귀를 막는 표준 절차를 확보합니다.

## 핵심 위치

- 테스트 코드: `tests/`
- CI 파이프라인: `.github/workflows/ci.yaml`
- 가이드: `CONTRIBUTING.md`, `Makefile`

## 테스트 레이어

- Invariants (P0): 테넌트 격리/권한
- Contracts (P1): API 계약
- Scenarios (P2): 사용자 여정
- Integration: 실 DB/실행환경 기반

## 사용 상세 절차

1. 빠른 로컬 점검
   - `go test ./...`
2. 레이어별 실행
   - `make test-invariants`
   - `make test-contracts`
   - `make test-scenarios`
3. 변경 영향 큰 경우 integration 실행

## 디버깅 포인트

- WS 문제: `pkg/protocol` + `internal/gateway/methods`
- API 문제: `internal/http`
- 데이터 문제: `internal/store/*` + migration
- 에이전트 품질: `internal/agent`, `internal/pipeline`, `internal/tools`

## 적용 예시

- API 스키마 수정 후 contract 테스트 실패를 통해 브레이킹 체인지 조기 발견

## 활용 전략

- PR 전 최소 P0/P1 자동 실행을 로컬 훅 또는 CI 필수로 유지

## 체크리스트

- 변경 범위에 맞는 테스트 레이어를 모두 실행
- 실패 로그를 파일/함수 단위로 추적 완료

## 다음 문서

[16-release-deploy-and-operations](./16-release-deploy-and-operations.md)
