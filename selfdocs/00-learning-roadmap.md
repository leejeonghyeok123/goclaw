# 00. Learning Roadmap

## 목적

GoClaw 학습 경로를 "실행 가능한 순서"로 제시합니다.

## 선행지식

- Go 기본 문법
- Docker/PostgreSQL 기초
- HTTP/WS API 개념

## 단계별 로드맵

1. 제품 개요 파악: `README.md`, 공식 사이트, 공식 문서
2. 개발 환경 준비: Go/DB/pnpm/Wails 설치
3. 첫 실행: `goclaw onboard` 및 초기 연결 확인
4. 구조 파악: `cmd`, `internal`, `pkg`, `ui`, `migrations`, `tests`
5. 런타임 이해: `main.go` -> `cmd/root.go` -> `cmd/gateway.go`
6. 외부 인터페이스 이해: HTTP/WS/CLI
7. 데이터 계층/멀티테넌트 이해: store + migration + context scope
8. 에이전트 코어 이해: loop + 8-stage pipeline
9. 확장 지점 학습: tools/providers/channels/memory/vault
10. 운영 능력 확보: 테스트/배포/릴리즈/모니터링

## 실제 적용 예시

- 신규 팀원이 2일 안에 기능 개발 온보딩:
  - Day1: 00~08 완료
  - Day2: 09~18 + 작은 기능 하나 직접 변경

## 실수 방지 체크리스트

- 문서 순서를 건너뛰지 않기
- 실행 명령은 현재 OS에 맞춰 변환해서 사용하기
- DB 스키마 변경 시 PG/SQLite 동시 고려

## 다음 문서

[01-what-is-goclaw](./01-what-is-goclaw.md)
