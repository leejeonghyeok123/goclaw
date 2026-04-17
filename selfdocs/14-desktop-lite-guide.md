# 14. Desktop Lite Guide

> **Selfdocs 공통 전제:** 이 문서만 **Desktop Lite**(SQLite·Wails) 예외 경로다. **Linux 서버 + Docker** 표준 게이트웨이는 [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md), 테넌트·키는 [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

인프라 없이 로컬에서 GoClaw Lite를 실행/개발/배포하는 방법을 정리합니다.

## 핵심 위치

- 데스크탑 엔트리: `ui/desktop/main.go`
- 앱 브리지: `ui/desktop/app.go`
- 프론트엔드: `ui/desktop/frontend/`
- Edition 제한: `internal/edition/edition.go`

## 사용 상세 절차

1. 설치
   - Linux/macOS: `curl -fsSL https://raw.githubusercontent.com/nextlevelbuilder/goclaw/main/scripts/install-lite.sh | bash`
   - Windows PowerShell: `irm https://raw.githubusercontent.com/nextlevelbuilder/goclaw/main/scripts/install-lite.ps1 | iex`
2. 개발 실행
   - `make desktop-dev`
3. 빌드
   - `make desktop-build`

## 기능 제한 인지

- Lite는 SQLite 기반, 에이전트/팀 수 제한, 일부 엔터프라이즈 기능 미지원

## 적용 예시

- 개인 업무 자동화 도구로 오프라인/로컬 우선 운영

## 활용 전략

- 팀 표준 운영은 Standard 서버, 개인 생산성/POC는 Lite로 분리

## 체크리스트

- 앱 실행 후 로컬 헬스체크 성공
- 버전/업데이트 배너 동작 확인

## 다음 문서

[15-testing-debugging-and-quality-gates](./15-testing-debugging-and-quality-gates.md)
