# 26. Hooks and Automation Guardrails

## 대상 독자

중급~고급 운영자/개발자 (자동화/품질 게이트 설계 담당)

## 목적

Hook 기반 자동화를 안전하게 적용하고, 실패/오작동 리스크를 관리합니다.

## Hook를 쉽게 설명하면

Hook는 "특정 이벤트가 발생했을 때 자동으로 실행되는 규칙/동작"입니다.

예:

- 작업 완료 시 슬랙 알림
- 특정 조건에서 승인 요청 생성
- 보안 정책 위반 시 실행 차단

## 코드 기준 위치

- Hook 타입/설정: `internal/hooks/types.go`, `internal/hooks/config.go`
- 감사/기록: `internal/hooks/audit.go`
- 연동 지점: gateway/agent/tool 실행 경로 전반

## 운영 전략

1. 처음에는 읽기/알림 Hook부터 시작
2. 그 다음 비차단형 자동화 적용
3. 마지막에 차단형(게이트) Hook 적용

이유:

- 초기에 차단형을 바로 넣으면 업무 중단 위험이 큼

## 권장 설계 원칙

- 실패 기본값 명확화(fail-open vs fail-closed)
- timeout과 재시도 정책 설정
- 중요한 Hook는 감사 로그를 반드시 남김
- Hook 변경은 코드리뷰/승인 프로세스 포함

## 적용 예시

- 코드 실행 도구(`exec`) 사용 시:
  - 고위험 명령 감지 Hook
  - 관리자 승인 없으면 보류/차단

## 자주 하는 실수

- 조건식이 너무 넓어 정상 트래픽까지 차단
- timeout 설정 누락으로 전체 응답 지연
- 감사 로그 없는 자동화로 사후 추적 불가

## 체크리스트

- Hook별 목적/조건/실패정책이 문서화됨
- staging에서 오탐률 확인 후 운영 반영
- 장애시 임시 우회 스위치가 준비됨

## 다음 문서

[27-backup-restore-and-disaster-recovery](./27-backup-restore-and-disaster-recovery.md)
