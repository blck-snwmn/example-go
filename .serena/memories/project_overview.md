# Example-Go プロジェクト概要

## プロジェクトの目的
example-goは、Goプログラミング言語の様々な機能や技術を実例で学習するためのサンプルコード集です。以下の分野をカバーしています：

- **標準ライブラリ**: HTTP、slog、iter、unique、genericsなど
- **テスト**: 標準テスト、gotestsum、ginkgo、カバレッジ測定
- **OpenAPI**: ogen、oapi-codegenを使ったAPI開発
- **データベース**: sqlx、sqlcを使ったデータベース操作
- **CLI**: コマンドラインツール開発
- **ツール**: コード生成、テンプレート処理

## 技術スタック
- **言語**: Go
- **テストフレームワーク**: 
  - 標準testing
  - gotestsum
  - Ginkgo
- **OpenAPI**: ogen、oapi-codegen
- **データベース**: PostgreSQL、MySQL（sqlx、sqlc使用）
- **リンター**: golangci-lint（gosec含む）
- **依存関係管理**: Go modules、go.work

## プロジェクト構造
プロジェクトは複数のGo modulesから構成されており、各ディレクトリが独立したmoduleとして動作します：

- `standard/`: 標準ライブラリの使用例
- `test/`: テスト手法の実例
- `openapi/`: OpenAPI関連の実装例
- `db/`: データベース操作の実例
- `cli/`: CLI開発の実例
- `tools/`: 開発ツールの実例