. ./script/ent/config

# 引数：スキーマ（テーブル）名、パスカルケース
go run entgo.io/ent/cmd/ent init --target $ENT_SRC_PATH/schema $@
