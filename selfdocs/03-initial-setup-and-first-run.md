# 03. Initial Setup and First Run

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

최초 실행을 성공시키고, 운영 가능한 기본 상태를 확보합니다. **기본 절차는 Linux + Docker Compose** 입니다.

## 코드 기준 위치

- 온보딩(CLI): `cmd/onboard.go`
- 마이그레이션: `cmd/migrate.go`
- 게이트웨이 기동: `cmd/root.go`, `cmd/gateway.go`
- Compose 기동: `Makefile` 의 `up`, `docker-compose.yml`

## 방법 A (권장): Docker Compose (Linux)

저장소 루트에서:

1. **환경 파일 준비**

   ```bash
   chmod +x prepare-env.sh
   ./prepare-env.sh
   ```

   - 루트에 `.env` 가 생기거나 보강된다.
   - `GOCLAW_GATEWAY_TOKEN`, `GOCLAW_ENCRYPTION_KEY` 등이 비어 있으면 자동 생성될 수 있다.
   - PostgreSQL 비밀번호 등은 `.env` / `.env.example` 설명에 따라 채운다.

2. **LLM 제공자 키**

   최소 한 개 이상의 제공자 API 키는 **대시보드 설정** 등을 통해 구성하는 흐름이 일반적이다(`.env.example` 주석 참고).

3. **기동**

   ```bash
   make up
   ```

   - 이미지 빌드·Postgres·GoClaw 컨테이너 기동과 마이그레이션용 `upgrade` 잡이 포함된다(`Makefile` 의 `up` 타깃).

4. **동작 확인**

   - API + 내장 웹 UI: `http://localhost:18790` (호스트 포트는 환경변수 `GOCLAW_PORT` 로 바꿀 수 있음, `docker-compose.yml` 참고)
   - 헬스: `curl -sS http://localhost:18790/health`

5. **대시보드 로그인**

   - 게이트웨이 토큰: `.env` 의 `GOCLAW_GATEWAY_TOKEN`
   - Owner 사용자 ID: 환경에 `GOCLAW_OWNER_IDS` 가 없으면 기본적으로 `"system"` 이 Owner 로 취급된다(`internal/http/auth.go` 와 [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md)).

## 방법 B: 호스트에서 바이너리 직접 빌드 (개발자용)

Go 가 설치된 머신에서:

```bash
make build
./goclaw onboard
set -a && source .env.local && set +a && ./goclaw
```

(`onboard` 가 `.env.local` 에 DSN·시크릿을 쓰는 흐름은 `cmd/onboard.go` 참고.)

## 실패 시

- DB 마이그레이션 오류: `./goclaw migrate up` (호스트 실행 시) 또는 compose 로그 `make logs` 로 원인 확인
- 포트 충돌: `.env` 의 `GOCLAW_PORT` 조정

## 체크리스트

- [ ] `make up` 후 `/health` 가 성공한다
- [ ] 브라우저에서 대시보드 URL 이 열린다
- [ ] `GOCLAW_GATEWAY_TOKEN` 을 안전한 곳에 보관했다

## 다음 문서

[04-project-structure-complete-map](./04-project-structure-complete-map.md)  
멀티테넌트 운영: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md)
