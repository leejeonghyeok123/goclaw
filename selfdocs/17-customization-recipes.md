# 17. Customization Recipes

## 목적

실무에서 자주 요구되는 커스터마이징을 "어디를 어떻게 수정할지" 레시피로 제공합니다.

## 레시피 1: 새 HTTP API 추가

- 위치: `internal/http`, `cmd/gateway_http_wiring.go`
- 절차:
  1. 핸들러 추가
  2. 라우트 등록
  3. store 호출 연결
  4. contract 테스트 추가

## 레시피 2: 새 WS 메서드 추가

- 위치: `internal/gateway/methods`, `cmd/gateway_methods.go`, `pkg/protocol`
- 절차:
  1. 메서드 핸들러 구현
  2. 등록 표 반영
  3. WS contract 테스트 작성

## 레시피 3: 새 Tool 추가

- 위치: `internal/tools/*`, `registry.go`, `policy.go`
- 절차:
  1. 툴 구현
  2. 레지스트리 등록
  3. 허용 정책/승인 정책 반영
  4. 회귀 테스트

## 레시피 4: DB 스키마 변경

- 위치: `migrations/`, `internal/store/sqlitestore/schema*`, `internal/upgrade/version.go`
- 절차:
  1. PG migration 작성
  2. SQLite schema/migration 동기화
  3. store 양쪽 구현 업데이트
  4. integration 테스트

## 레시피 5: UI 기능 추가

- 위치: `ui/web/src/pages`, `ui/web/src/routes.tsx`, `ui/web/src/i18n/locales/*`
- 절차:
  1. 페이지/컴포넌트 생성
  2. 라우트/권한 가드 연결
  3. 다국어 키 3개 언어 동시 반영

## 적용 예시

- "Provider 연결 테스트 버튼 추가"
  - backend verify endpoint 재사용
  - UI Providers 페이지에 액션 연결

## 체크리스트

- 레시피 수행 후 관련 테스트 통과
- 문서/번역/권한 가드 누락 없음

## 다음 문서

[18-expert-checklist-and-next-steps](./18-expert-checklist-and-next-steps.md)
