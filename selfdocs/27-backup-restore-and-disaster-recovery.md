# 27. Backup, Restore, Disaster Recovery

## 대상 독자

운영자/플랫폼 담당자

## 목적

장애/실수/데이터 손실 상황에서 GoClaw를 복구할 수 있도록 백업/복원 전략을 정리합니다.

## 코드 기준 위치

- 백업 코어: `internal/backup/*`
- HTTP 핸들러: `internal/http/backup_handler.go`, `internal/http/restore_handler.go`
- 테넌트 단위: `internal/http/tenant_backup_handler.go`, `internal/http/tenant_restore_handler.go`
- CLI 관련: `cmd/backup.go`, `cmd/restore.go`

## 백업 범위 설계

1. 시스템 전체 백업
2. tenant 단위 백업
3. 파일/문서 저장소 포함 여부

권장:

- 운영은 "정기 전체 + 수시 tenant 단위" 병행

## 복구 시나리오

### 시나리오 A: 운영 실수로 데이터 삭제

1. 영향 tenant 파악
2. 가장 최근 정상 백업 선택
3. tenant 단위 복구 우선 시도
4. 검증 후 서비스 재개

### 시나리오 B: DB 장애/손상

1. 서비스 보호 모드 전환(쓰기 제한)
2. 전체 백업 복원
3. 마이그레이션/버전 일치 확인
4. 핵심 시나리오 테스트 후 오픈

## 운영 체크포인트

- 백업 파일 무결성 검증
- 복구 리허설(월 1회 이상)
- 백업 보관 주기/암호화 정책
- 접근 권한 최소화

## 자주 하는 실수

- 백업만 하고 복구 테스트를 안 함
- 버전 불일치 상태로 복구 시도
- 보관 정책 없이 디스크만 증가

## 적용 예시

- 부서별 tenant 운영:
  - 일일 전체 백업 + 릴리즈 전 tenant 스냅샷
  - 장애 시 tenant 단위 우선 복구로 다운타임 최소화

## 초보자용 복구 최소 절차

1. "무엇이 망가졌는지" 먼저 분류(tenant? 전체?)
2. 복구 대상과 시점(언제 상태로 되돌릴지) 확정
3. staging/격리 환경에서 1차 복구 검증
4. 운영 반영 후 핵심 API/채팅/세션 검증

## 다음 문서

[15-testing-debugging-and-quality-gates](./15-testing-debugging-and-quality-gates.md)
