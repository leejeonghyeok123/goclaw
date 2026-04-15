# 03. Initial Setup and First Run

## 목적

최초 실행을 성공시키고, 운영 가능한 기본 상태를 확보합니다.

## 코드 기준 위치

- 온보딩 로직: `cmd/onboard.go`
- 마이그레이션 명령: `cmd/migrate.go`
- 앱 시작: `cmd/root.go`, `cmd/gateway.go`

## 방법 A: 소스 실행

1. 빌드
   - `make build`
2. 온보딩
   - `./goclaw onboard`
3. 환경 로드 후 실행
   - Linux/macOS: `source .env.local && ./goclaw`
   - PowerShell: `. .\.env.local; .\goclaw.exe`

## 방법 B: Docker 실행

1. 환경 파일 생성
   - Linux/macOS: `chmod +x prepare-env.sh && ./prepare-env.sh`
2. API 키 최소 1개 입력 (`GOCLAW_*_API_KEY`)
3. 기동
   - `make up`

## 기능 사용을 위한 상세 포인트

- `onboard`는 DB 연결/마이그레이션/기본 시드 설정까지 유도
- 실패 시 `./goclaw migrate up` 수동 수행 후 재시도
- 서버 실행 후 `http://localhost:18790/health` 확인

## 적용 예시

- 신규 환경에서 15분 안에 "헬스체크 + 기본 채팅"까지 도달:
  - `onboard` -> 실행 -> Health -> Web 대시보드 접속

## 활용 전략

- CI와 동일한 기본 루틴(`build -> migrate -> run`)을 로컬에서도 유지

## 체크리스트

- 앱 기동 성공
- DB 마이그레이션 버전 정상
- `/health` 응답 정상

## 다음 문서

[04-project-structure-complete-map](./04-project-structure-complete-map.md)
