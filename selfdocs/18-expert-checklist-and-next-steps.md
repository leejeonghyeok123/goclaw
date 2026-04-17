# 18. Expert Checklist and Next Steps

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

GoClaw 전문가 수준 도달 여부를 자가 진단하고, 다음 성장 경로를 정의합니다.

## 전문가 체크리스트

- 런타임 흐름을 코드 기준으로 설명 가능 (`main` -> `cmd` -> `gateway` -> `http/ws`)
- 인터페이스 선택을 상황별로 설계 가능 (HTTP/WS/CLI)
- store/migration 변경 시 PG + SQLite 동시 처리 가능
- tenant scope와 보안 정책 누락 없이 구현 가능
- tool/provider/channel/memory/vault 연동 구조를 실무적으로 설명 가능
- 웹/데스크탑 배포 경로와 제한사항 이해
- 테스트 레이어(P0/P1/P2) 기준으로 품질 게이트 설계 가능

## 실전 과제 (권장)

1. 새 tool 1개 추가 + 정책 적용 + 테스트
2. provider fallback 전략 문서화 및 적용
3. vault ingest 파이프라인 튜닝
4. ws 이벤트 기반 운영 대시보드 기능 1개 추가

## 운영 활용 전략

- 팀 내 역할 분리:
  - Platform: gateway/store/security
  - AI Runtime: agent/pipeline/tools/providers
  - Product UI: web/desktop
- 기능 추가 시 항상 "보안/테스트/운영성" 3축으로 검토

## 다음 단계

- `docs/` 하위 심화 문서와 연결해 특정 도메인(예: Vault, Teams, Channels) 전문화
- 실제 PR 단위로 개선 기록을 `selfdocs`에 주기적으로 축적

## 완료 기준

- 위 체크리스트를 모두 충족하고, 최소 1개 실무 기능을 end-to-end로 구현/배포 가능
