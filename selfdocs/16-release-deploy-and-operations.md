# 16. Release, Deploy, and Operations

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

개발 결과를 안정적으로 배포하고 운영하는 실제 절차를 정리합니다.

## 핵심 위치

- 릴리즈 워크플로우: `.github/workflows/release.yaml`, `release-beta.yaml`, `release-desktop.yaml`
- 배포 구성: `docker-compose.yml`, `compose.d/`, `Dockerfile*`
- 운영 명령: `Makefile`

## 릴리즈 모델

- 안정 릴리즈: `vX.Y.Z`
- 베타/RC: `vX.Y.Z-beta.N` / `vX.Y.Z-rc.N`
- 데스크탑 Lite: `lite-vX.Y.Z`

## 사용 상세 절차

1. 브랜치/테스트 기준 충족
2. 태그 발행
3. GitHub Actions에서 아티팩트/이미지 생성 확인
4. 배포 환경 업데이트
   - 저장소가 정의한 파일과 동일하게: `docker compose pull && docker compose up -d` (로컬 개발 기준은 [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md) 의 `make up` 흐름)

## 운영 포인트

- 헬스/로그/트레이스 모니터링
- API 키/시크릿 교체 정책
- 실패 시 롤백 전략(이전 이미지/태그)

## 적용 예시

- 베타를 내부 테넌트에 먼저 배포한 뒤 안정 태그 전환

## 활용 전략

- "릴리즈 태그 정책"과 "환경별 배포 정책"을 분리하면 리스크를 낮출 수 있음

## 체크리스트

- 태그 규칙 준수
- 배포 후 헬스체크/핵심 시나리오 통과

## 다음 문서

[17-customization-recipes](./17-customization-recipes.md)
