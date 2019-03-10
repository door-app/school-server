# ■□ マイグレーション
#
migrate:
	goose up
migrate_down:
	goose down

# ■□ Protoファイルをコピーする
#
copy_proto_file:
	cp -r ../protocol-buffers/out/go/sample/ ./protoc/