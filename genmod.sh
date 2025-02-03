#!/bin/bash

# 現在の作業ディレクトリを記録
ORIG_DIR=$(pwd)

# 引数が与えられていない場合はエラーメッセージを表示して終了
if [ "$#" -ne 1 ]; then
    echo "使用法: $0 <ディレクトリ名>"
    exit 1
fi

# 引数で指定されたディレクトリ名を変数に代入
DIR_NAME=$1

# ディレクトリを作成
mkdir -p "$DIR_NAME"

# 作成したディレクトリに移動
cd "$DIR_NAME"

# 現在のディレクトリパスを取得し、'github.com'までのパスを抽出
CURRENT_PATH=$(pwd)
GITHUB_PATH=$(echo $CURRENT_PATH | sed -n 's/.*\(github\.com.*\)/\1/p')

# go mod init コマンドを実行
go mod init $GITHUB_PATH

go work use .

# リポジトリルートを取得
REPO_ROOT=$(git rev-parse --show-toplevel)
WORK_FILE="$REPO_ROOT/go.work"
DEPENDABOT_FILE="$REPO_ROOT/.github/dependabot.yml"

# go.work の use ブロックから最新ディレクトリの一覧を取得
NEW_DIRS=$(awk '/use \(/,/\)/{
    if ($1 ~ /^\.\//) print $1
}' "$WORK_FILE" | sed 's/^/      - "/; s/$/"/')

# dependabot.yml の gomod ブロック内の directories リストを更新
awk -v new_dirs="$NEW_DIRS" '
    BEGIN { inGomod=0; inDirs=0 }
    /- package-ecosystem: "gomod"/ { inGomod=1 }
    {
      if(inGomod && $0 ~ /^[[:space:]]*directories:/) {
         print "    directories:";
         print new_dirs;
         inDirs=1;
         inGomod=0; next
      }
      if(inDirs && $0 ~ /^[[:space:]]*[^-[:space:]]/) { inDirs=0 }
      if(!inDirs) { print }
    }
' "$DEPENDABOT_FILE" > "$DEPENDABOT_FILE.tmp" && mv "$DEPENDABOT_FILE.tmp" "$DEPENDABOT_FILE"

# 元の作業ディレクトリへ戻る
cd "$ORIG_DIR"
