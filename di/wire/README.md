# Google Wire DI Example

このディレクトリには、Google Wireフレームワークを使った依存関係注入（Dependency Injection）の簡単なサンプルが含まれています。

## 構造

```
di/wire/
├── main.go          # メイン実行ファイル
├── wire.go          # Wire設定ファイル（依存関係の定義）
├── wire_gen.go      # Wireによって自動生成されたDIコード
├── database.go      # データベース層のProvider
├── service.go       # サービス層
├── server.go        # HTTP サーバー層
├── go.mod          # Go モジュール設定
└── README.md       # このファイル
```

## 依存関係の構造

```
Server
├── UserService (インターフェース)
│   └── Database (インターフェース)
│       └── DSN (設定値)
└── Port (設定値)
```

## 実行方法

1. Wireコード生成（既に生成済み）:
```bash
go run github.com/google/wire/cmd/wire
```

2. アプリケーション実行:
```bash
go run .
```

3. ブラウザまたはcurlでテスト:
```bash
curl http://localhost:8080/users
```

## Wireの特徴

- **コンパイル時DI**: ランタイムではなくコンパイル時に依存関係を解決
- **型安全**: Go の型システムを活用した安全なDI
- **性能**: リフレクションを使わないため高速
- **explicit**: 依存関係が明示的で理解しやすい

## カスタマイズ

Provider関数を追加・変更する場合は：

1. `wire.go` の `wire.Build()` に新しいProviderを追加
2. Wireコード再生成: `go run github.com/google/wire/cmd/wire`
3. アプリケーション再ビルド: `go build`