# コーディングスタイルと規約

## 一般的なGo規約
- パッケージ名は小文字
- 関数名はCamelCase
- プライベート関数は小文字から開始
- パブリック関数は大文字から開始

## プロジェクト固有の規約

### エラーハンドリング
- HTTPレスポンスのwriteエラーは通常無視（nolintコメント付き）
- 例: `//nolint:errcheck,gosec // HTTP response write errors aren't useful`

### テスト規約
- テスト関数名は `Test_` プレフィックス
- サブテストには `t.Run()` を使用
- 環境変数テストには `t.Setenv()` を使用

### HTTP サーバー設定
- ReadHeaderTimeoutを設定してセキュリティを向上
- デフォルトタイムアウト: 5秒

### リンター設定
- golangci-lintを使用
- gosecを有効化してセキュリティチェック

### nolintコメント使用パターン
- HTTP write エラーの無視: `//nolint:errcheck,gosec`
- 理由を明記: `// HTTP response write errors aren't useful`

## ファイル構造規約
- 各ディレクトリに独自のgo.modを配置
- READMEファイルで各例の説明を提供
- テストファイルは `_test.go` サフィックス