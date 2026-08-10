# AGENTS.md

本文件適用於整個 `gitops-demo-backend` repository。工作前請先閱讀本文件與 `README.md`。

## 專案定位

- 本 repo 維護 Go 後端 API、Dockerfile 與後端映像建置 workflow。
- Kubernetes manifests、Kustomize overlays 與 ArgoCD ApplicationSet 由 `gitops-demo-apps` 管理。
- Cluster、ArgoCD 平台與 GitOps bootstrap 分別由 `gitops-demo-cluster`、`gitops-demo-infra` 管理。
- 除非使用者明確要求，不得在本 repo 加入或修改其他 repository 的部署責任。

## 註解與術語規範

- 人工維護的 Go、Dockerfile、GitHub Actions、設定檔與腳本註解必須使用繁體中文。
- 專有名詞、產品名稱、API、資源種類、欄位名稱、命令、路徑與識別字可保留英文，但英文專有名詞必須放在中文敘述中，不得以完整英文句子撰寫註解。
- 本 repo 的技術名稱統一寫作 `Go`、`HTTP`、`API`、`Dockerfile`、`Docker Hub` 與 `GitHub Actions`，不得改用非官方大小寫或自行翻譯。
- Go package、module、function、environment variable、HTTP endpoint 與其他程式識別字必須保留原始拼法，例如 `main`、`COMMIT`、`/version`、`/health` 與 `/meta`。
- README 與 docs 使用繁體中文敘述，並遵守相同的專有名詞大小寫。
- Workflow／job／step、composite action 的 `name` 與 `description` 必須使用英文。
- 程式碼內的文字必須使用英文，包括 CLI／UI 文字、log、error、warning、summary 與其他執行訊息；但等待／重試迴圈中即時印給人類觀察進度的狀態訊息（例如第幾次嘗試、剩餘秒數、失敗原因、逾時後的診斷輸出）例外，使用繁體中文。
- 產品名稱的唯一允許拼法為 `ArgoCD`。
- 自動生成檔案的生成器註解、shebang、lint directive 與被註解掉的程式碼不需翻譯或改寫。

## Validation

- 依全域「最小必要 Validation」規範，先根據實際執行路徑與風險，再從 repository 既有的 formatting check、targeted Go test／build、Docker build 與 `actionlint` 中選擇最小子集。
- `go test ./...`、完整 Go build、Docker build 與完整 CI workflow 屬較大範圍驗證；只有 shared package／application entrypoint、image runtime contract、CI contract 或其他直接 consumers 確實受影響時才執行。
- 每項 validation 必須說明對應的變更與風險；不得只因工具存在或過去慣例而全部執行。
- 缺少工具或無法安全執行時，標示 `BLOCKED` 或 `NOT RUN`，並說明未取得的信心。
- 不得回復使用者既有未提交變更。
