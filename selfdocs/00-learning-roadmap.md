# 00. Learning Roadmap

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

GoClaw 학습 경로를 "실행 가능한 순서"로 제시합니다.

## 선행지식

- Docker / Docker Compose 기초(기본 전제: **Linux + Compose**)
- HTTP/WS API 개념
- 소스까지 수정할 때: Go 기본 문법, PostgreSQL 기초

## 단계별 로드맵

1. 제품 개요 파악: `README.md`, 공식 사이트, 공식 문서
2. 실행 환경 준비: Linux 에서 Docker + (선택) Make
3. 첫 실행: `prepare-env.sh` → `make up` 및 `/health` 확인 ([03](./03-initial-setup-and-first-run.md))
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
- selfdocs 기본 전제는 Linux + Docker 이므로, 다른 OS는 공식 `README.md` 와 병행
- DB 스키마 변경 시 PG/SQLite 동시 고려

## 다음 문서

[01-what-is-goclaw](./01-what-is-goclaw.md)
