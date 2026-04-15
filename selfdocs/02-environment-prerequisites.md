# 02. Environment Prerequisites

## 목적

실행 전 필수 도구/버전을 명확히 맞춰 실패를 예방합니다.

## 필수 구성요소

- Go 1.26+
- PostgreSQL 18 + pgvector (Standard 서버 모드)
- Docker / Docker Compose (권장)
- pnpm (웹/데스크탑 프론트엔드)
- Wails CLI (Desktop Lite 개발/빌드)

## 코드/설정 위치

- 실행 지침: `README.md`
- 자동화 명령: `Makefile`
- CI 기준 버전: `.github/workflows/ci.yaml`
- 데스크탑 빌드: `.github/workflows/release-desktop.yaml`

## 상세 준비 절차 (Windows PowerShell 우선)

1. Go 설치 확인
   - `go version`
2. Docker 설치 확인
   - `docker --version`
   - `docker compose version`
3. pnpm 설치 확인
   - `pnpm --version`
4. (옵션) Wails 설치
   - `go install github.com/wailsapp/wails/v2/cmd/wails@latest`

## 적용 예시

- 백엔드만 먼저 검증:
  - Go + PostgreSQL만 준비 후 `make build`, `./goclaw onboard`
- 데스크탑만 검증:
  - Wails + pnpm 준비 후 `make desktop-dev`

## 활용 전략

- 팀 공용 온보딩 문서에 "명령 1줄 검증" 블록을 넣어 환경 이슈를 초기에 차단

## 체크리스트

- `go`, `docker`, `pnpm` 버전 명령이 모두 동작
- Standard/Lite 중 현재 목표 모드가 결정됨

## 다음 문서

[03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md)
