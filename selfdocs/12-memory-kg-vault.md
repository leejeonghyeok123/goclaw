# 12. Memory, Knowledge Graph, Vault

> **Selfdocs 공통 전제:** **Linux** 호스트·**Docker Compose**(`make up` 등)로 게이트웨이를 띄우는 것을 기본으로 한다. 초기 기동·포트·헬스: [03-initial-setup-and-first-run](./03-initial-setup-and-first-run.md). 테넌트·API 키·요청 스코프: [22-tenant-vs-account-explained](./22-tenant-vs-account-explained.md).

## 대상 독자

중급(입문자는 먼저 `99-glossary-for-beginners.md`, `21-sessions-and-context-for-beginners.md` 권장)

## 목적

대화 내용을 지식 자산으로 축적하고 재활용하는 구조를 운영합니다.

## 핵심 위치

- 메모리: `internal/memory/`
- 통합/승격 워커: `internal/consolidation/`
- 지식그래프: `internal/knowledgegraph/`
- Vault: `internal/vault/`
- 관련 API: `internal/http/memory*.go`, `knowledge_graph*.go`, `vault*.go`

## 기능 개요

- Working/Episodic/Semantic 다층 메모리
- KG 추출과 그래프 탐색
- Vault 문서 레지스트리 + 하이브리드 검색

## 가장 많이 묻는 질문 (핵심 답)

질문: "내가 agent랑 대화하면 메모리랑 지식그래프가 자동 생성되나요?"

정답:

- 메모리(특히 working/요약 기반 episodic)는 **대체로 자동**으로 쌓입니다.
- 지식그래프(KG)는 **기능 활성화 + 백그라운드 파이프라인 + 데이터 품질** 조건이 맞아야 안정적으로 쌓입니다.
- 즉, "완전 무조건 자동"이 아니라, 운영 설정이 맞을 때 자동화가 제대로 작동합니다.

## 자동 생성 vs 설정 필요 (구분표)

| 영역 | 대화만으로 기본 동작 | 추가 설정 필요 | 비고 |
|---|---|---|---|
| Working memory (대화 히스토리) | 예 | 낮음 | 세션 기반으로 기본 축적 |
| Episodic summary | 보통 예 | 중간 | 파이프라인/요약 품질 영향 |
| Semantic/KG | 부분적 | 높음 | consolidation, extractor, scope, 데이터 품질 영향 |
| Vault 검색 | 아니오(문서 없으면) | 높음 | 문서 ingest/index 선행 필요 |

## 대화 후 내부에서 일어나는 흐름 (실제 운영 관점)

1. 사용자가 대화 요청
   - 엔진: `internal/agent/*`, `internal/pipeline/*`
2. 실행 종료 시 memory/summarize 단계 반영
   - finalize 경로에서 후속 이벤트가 준비됨 (`internal/pipeline/finalize_stage.go`, `internal/agent/loop_finalize.go`)
3. domain event 발생
   - `session.completed` -> `episodic.created` 계열 이벤트로 consolidation 파이프라인 유입 (`internal/eventbus/event_types.go`)
4. consolidation worker 비동기 처리
   - episodic -> semantic(KG) -> dreaming 순으로 승격/정제 (`internal/consolidation/workers.go`)
5. 검색 시점
   - memory/KG/vault를 합쳐 검색하거나 개별 검색 (`internal/vault/search.go`, `internal/tools/knowledge_graph.go`)

## 왜 "대화했는데 KG가 비어보일 수 있나"

- 백그라운드 worker가 아직 처리 중
- KG 추출할 만한 엔티티/관계가 대화에 부족
- 스코프가 다름(사용자별 KG vs shared KG)
- Lite/환경 설정에서 의미 검색 품질이 낮거나 파이프라인 의존 기능이 제한

## 사용 상세 절차 (권장 운영 루틴)

1. 메모리/벡터 관련 설정 확인
2. 문서 ingest 및 인덱싱
3. 검색 API/툴 호출로 회수 테스트
4. 품질 검증(정확도/누락/노이즈) 후 조정

## 어디에서 뭘 확인하면 되는가 (점검 포인트)

1. 메모리 설정
   - `internal/config/config.go`
   - `internal/config/config_system.go`
   - `internal/memory/auto_injector.go` (`auto_inject`, `consolidation_enabled`)
2. 파이프라인 이벤트 연결
   - `internal/pipeline/finalize_stage.go`
   - `internal/agent/loop_finalize.go`
   - `internal/consolidation/workers.go`
3. 저장소/검색
   - episodic: `internal/store/*/episodic*.go`
   - KG: `internal/store/*/knowledge_graph*.go`, `internal/tools/knowledge_graph.go`
   - vault fan-out: `internal/vault/search.go`

## 적용 예시

- 사내 SOP 문서를 Vault에 올려 에이전트 답변 근거 강화
- 고객 대화 히스토리를 episodic summary로 축약해 컨텍스트 비용 절감
- 운영 알림 대화 로그에서 사람/시스템/이슈 관계를 KG로 추출해 원인 추적 속도 향상

## 활용 전략

- "요약 품질"과 "검색 품질"을 분리 측정하여 튜닝 포인트를 명확히 함
- "즉시성"이 중요한 질의는 working memory 우선, "관계 추론"은 KG, "근거문서 회수"는 Vault 중심으로 라우팅

## 체크리스트

- 검색 결과에 근거 문서가 포함됨
- 메모리 승격이 비동기적으로 정상 수행됨
- 대화 직후와 몇 분 후 KG 결과 차이를 비교해 파이프라인 지연을 확인함
- 같은 질문을 `memory_search`, `knowledge_graph_search`, `vault_search`로 각각 실행해 결과 특성 차이를 파악함

## 다음 문서

[13-web-ui-operator-guide](./13-web-ui-operator-guide.md)
