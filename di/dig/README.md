# Dependency Injection with Uber Dig

このサンプルは、[Uber Dig](https://github.com/uber-go/dig)を使用した依存関係注入（DI）の実装例です。

## 概要

Digは、Uberが開発したGoのための依存関係注入フレームワークです。リフレクションベースでランタイムに依存関係を解決する特徴があります。

## 特徴

- **ランタイム依存関係解決**: リフレクションを使用してランタイムに依存関係を解決
- **シンプルなAPI**: `container.Provide()`で依存関係を登録、`container.Invoke()`で実行
- **循環依存の検出**: ランタイムに循環依存を検出してエラーを出力
- **コード生成不要**: Wireとは異なり、コード生成が不要

## プロジェクト構造

```
di/dig/
├── main.go        # エントリーポイント
├── container.go   # DIコンテナの設定
├── server.go      # HTTPサーバーの実装
├── service.go     # ビジネスロジック層
├── database.go    # データアクセス層
├── go.mod
└── README.md
```

## 依存関係グラフ

```
DSN → Database → UserService → Server → App
```

## 実行方法

```bash
# 依存関係をダウンロード
go mod tidy

# サーバーを起動
go run .
```

## API使用例

サーバーが起動したら、以下のAPIを使用できます：

```bash
# ユーザー一覧を取得
curl http://localhost:8080/users

# 新しいユーザーを作成
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"新しいユーザー"}'
```

## WireとDigの違い

| 項目 | Wire | Dig |
|------|------|-----|
| 依存関係解決 | コンパイル時 | ランタイム |
| コード生成 | 必要 | 不要 |
| パフォーマンス | 高速 | わずかに遅い |
| エラー検出 | コンパイル時 | ランタイム |
| 学習コスト | 高い | 低い |

## 実装のポイント

### 1. プロバイダーの登録

```go
container := dig.New()
container.Provide(ProvideDSN)
container.Provide(NewDatabase)
container.Provide(NewUserService)
```

### 2. 依存関係の注入

```go
err := container.Invoke(func(app *App) error {
    return app.Start()
})
```

### 3. コンストラクタ関数

```go
func NewUserService(db Database) UserService {
    return &userService{db: db}
}
```

## 参考

- [Uber Dig公式ドキュメント](https://github.com/uber-go/dig)
- [Go DIパターンの比較](https://blog.golang.org/wire)