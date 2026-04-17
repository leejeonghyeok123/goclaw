# 08. Agent Loop and 8-Stage Pipeline

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 목적

GoClaw의 핵심 실행 엔진(think-act-observe)을 코드 흐름 수준으로 이해합니다.

## 핵심 위치

- 루프: `internal/agent/loop*.go`
- 파이프라인: `internal/pipeline/*.go`
- 입력 가드/정제: `internal/agent/input_guard.go`, `sanitize.go`

## 8단계

1. context
2. history
3. prompt
4. think
5. act
6. observe
7. memory
8. summarize

## 기능 사용 상세

- 대화 요청이 들어오면 pipeline이 단계별 컨텍스트를 조립
- 모델 호출 결과에 따라 툴 실행(`act`)과 결과 반영(`observe`) 반복
- 종료 시 메모리/요약 저장 단계 수행

## 어디를 수정하면 되는가

- 프롬프트 구성 변경: prompt stage 관련 파일
- 반복/중단 정책 변경: loop control 로직
- 메모리 적재 정책 변경: memory/summarize stage

## 적용 예시

- "툴 호출 전에 안전 검사 강화" 요구:
  - input guard + tool policy + act stage 연결 흐름 점검

## 활용 전략

- 디버깅 시 "현재 어느 stage인지"부터 고정하면 원인 파악 속도가 급격히 올라감

## 체크리스트

- 요청 1건의 stage 이동을 설명 가능
- stage별 주요 파일을 알고 있음

## 다음 문서

[09-tools-system-and-policy](./09-tools-system-and-policy.md)
