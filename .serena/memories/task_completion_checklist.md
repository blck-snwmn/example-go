# タスク完了時のチェックリスト

## 必須実行項目

### 1. リンターチェック
```bash
# 全ディレクトリでリント実行
./run_lints.sh

# または特定のディレクトリで
cd <directory> && golangci-lint run --enable=gosec
```

### 2. テスト実行
```bash
# 全テスト実行
./run_tests.sh

# または特定のディレクトリで
cd <directory> && go test ./...
```

### 3. フォーマット確認
```bash
# コードフォーマット
go fmt ./...
```

### 4. 依存関係の整理
```bash
# 各moduleディレクトリで実行
go mod tidy
```

## 新しいモジュール追加時の追加作業

### 1. モジュール作成
```bash
./genmod.sh <directory_name>
```

### 2. 必要な依存関係の追加
```bash
cd <directory_name>
go get <required_packages>
```

### 3. golangci-lintの追加（自動実行される）
```bash
go get --tool github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2
```

## チェック項目

- [ ] リンターエラーなし
- [ ] テスト全通過
- [ ] コードフォーマット適用済み
- [ ] 依存関係整理済み
- [ ] 新しいファイルにはappropriateなnolintコメント
- [ ] セキュリティ考慮事項確認済み（gosecチェック通過）
- [ ] 適切なエラーハンドリング実装済み

## セキュリティチェックポイント
- HTTP サーバーにReadHeaderTimeout設定
- 適切なnolintコメント使用
- gosecのwarning対応
- 機密情報のハードコーディング回避