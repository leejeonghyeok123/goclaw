# 05. Runtime Entrypoints and Wiring

## 목적

GoClaw가 실제로 어떤 순서로 구성요소를 조립해 실행되는지 이해합니다.

## 핵심 파일

- `main.go`: 실행 시작점 (`cmd.Execute()`)
- `cmd/root.go`: Cobra 루트 및 서브커맨드
- `cmd/gateway.go`: config/store/provider/tool/channel/server 조립

## 런타임 흐름

1. `main.go` 진입
2. `cmd/root.go`에서 명령 해석
3. 기본 명령 `gateway` 실행
4. `cmd/gateway.go`에서 인프라 구성
5. `internal/gateway/server.go`가 HTTP/WS 노출

## 기능 사용을 위한 상세 포인트

- 새 전역 의존성(예: 새 서비스)은 보통 `cmd/gateway.go`에서 생성/주입
- HTTP 라우트 추가는 wiring 파일들과 `internal/http` 핸들러 동시 수정
- WS 메서드는 methods 등록 표에 빠짐없이 연결해야 함

## 적용 예시

- 새 "운영 메트릭 서비스"를 붙일 때:
  - 서비스 생성: `cmd/gateway.go`
  - API 노출: `internal/http/usage*.go` 유사 패턴 참고

## 활용 전략

- 기능 개발 전 `gateway.go`를 먼저 읽으면 아키텍처 의사결정이 쉬워짐

## 체크리스트

- 실행 순서를 그림 없이 말로 설명 가능
- 새 컴포넌트 추가 위치를 알고 있음

## 다음 문서

[06-http-ws-cli-interfaces](./06-http-ws-cli-interfaces.md)
