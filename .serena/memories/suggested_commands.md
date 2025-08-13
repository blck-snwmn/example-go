# 推奨コマンド一覧

## 開発用コマンド

### テスト実行
```bash
# 全てのテストを実行
./run_tests.sh

# 特定のディレクトリでテスト実行
cd <directory> && go test ./...

# Ginkgoテスト実行（test/ginkoディレクトリ用）
cd test/ginkgo && go tool ginkgo -p
```

### リント実行
```bash
# 全ディレクトリでリント実行
./run_lints.sh

# 特定のディレクトリでリント実行
cd <directory> && golangci-lint run --enable=gosec
```

### 新しいモジュール作成
```bash
# 新しいGo moduleを作成し、dependabot.ymlを自動更新
./genmod.sh <directory_name>
```

### Goの基本コマンド
```bash
# 依存関係のインストール
go mod tidy

# モジュールの初期化
go mod init <module_path>

# ビルド
go build ./...

# フォーマット
go fmt ./...
```

## システムコマンド（Darwin用）
- `find` - ファイル検索
- `grep` - テキスト検索
- `git` - バージョン管理
- `ls` - ディレクトリ内容表示
- `cd` - ディレクトリ変更