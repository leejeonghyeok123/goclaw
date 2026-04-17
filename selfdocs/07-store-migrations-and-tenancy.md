# 07. Store, Migrations, and Tenancy

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

데이터 계층 변경 시 장애/격리 이슈를 방지합니다.

## 핵심 위치

- 인터페이스: `internal/store/`
- PostgreSQL 구현: `internal/store/pg/`
- SQLite 구현: `internal/store/sqlitestore/`
- PG 마이그레이션: `migrations/`

## 중요 원칙

- PG와 SQLite는 별도 마이그레이션 체계
- 멀티테넌트 스코프를 항상 컨텍스트와 SQL 양쪽에서 보장 (테넌트 행·키·HTTP 헤더 흐름은 [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md))
- 쿼리는 파라미터 바인딩만 사용

## 상세 작업 절차

1. 스키마 변경 필요 시
   - PG: `migrations/*.up.sql` + 버전 상수 갱신
   - SQLite: `schema.sql` + incremental migration map + 버전 상수 갱신
2. Store 변경
   - 인터페이스 -> PG 구현 -> SQLite 구현 순으로 반영
3. 격리 검증
   - tenant scope 조건 누락 여부 점검
4. 테스트
   - 관련 invariant/contract/integration 실행

## 적용 예시

- `agents` 테이블에 컬럼 추가:
  - PG migration 추가
  - SQLite schema 및 migration 동기화
  - `internal/store/pg/agents*.go`, `internal/store/sqlitestore/agents*.go` 반영

## 활용 전략

- DB 변경 PR 템플릿에 "PG/SQLite 동시 반영 체크"를 강제 항목으로 둠

## 체크리스트

- 두 저장소 구현이 모두 업데이트됨
- tenant scope 조건 누락 없음

## 다음 문서

[08-agent-loop-and-8-stage-pipeline](./08-agent-loop-and-8-stage-pipeline.md)
