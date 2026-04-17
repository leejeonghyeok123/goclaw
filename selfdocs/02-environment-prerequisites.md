# 02. Environment Prerequisites

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

실행 전 필수 도구/버전을 명확히 맞춰 실패를 예방합니다. 이 저장소의 selfdocs 기본 전제는 **Linux 호스트 + Docker Compose** 입니다.

## 필수 구성요소 (Docker 경로)

- **Docker Engine** + **Docker Compose v2** (`docker compose`)
- (선택) **GNU Make** — `Makefile` 의 `make up` 등을 쓰려면 필요

컨테이너 안에서 빌드되므로, **호스트에 Go를 설치하지 않아도** Docker 만으로 백엔드를 띄울 수 있다.

## 소스 개발·테스트까지 할 때 (추가)

- **Go** 1.26+ (버전은 `README.md` / CI 와 맞출 것)
- **pnpm** — 웹 UI (`ui/web`) 개발 시
- **Wails CLI** — Desktop Lite 빌드/개발 시

## 코드/설정 위치

- 실행 지침: `README.md`
- Compose / `up`: `Makefile`, `docker-compose.yml`, `docker-compose.postgres.yml`
- 환경 예시: `.env.example`, `prepare-env.sh`
- CI 기준: `.github/workflows/ci.yaml`

## Linux에서 환경 확인 (bash)

```bash
docker --version
docker compose version
command -v make && make --version
# 소스 개발 시에만:
go version 2>/dev/null || true
pnpm --version 2>/dev/null || true
```

## 적용 예시

- **백엔드만 Docker 로 검증**: 위 Docker 항목만 갖춘 뒤 [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md) 진행
- **UI/데스크톱까지**: pnpm / Wails 추가 설치 후 저장소 루트의 `Makefile` 타깃 사용

## 활용 전략

- 팀 온보딩 문서 맨 위에 “`docker compose version` 한 줄” 검증 블록을 두면 환경 이슈를 빨리 걸러낸다.

## 체크리스트

- [ ] `docker` / `docker compose` 명령이 동작한다
- [ ] 지금 목표가 **Docker 만** 인지 **로컬 Go 빌드** 까지 인지 정했다

## 다음 문서

[03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md)
