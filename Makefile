# ■□ マイグレーション
#
migrate:
	goose up
migrate_down:
	goose down


run_server:
	docker-compose up


copy_proto_file:
	cp -r ../protocol-buffers/out/go/sample/ ./protoc/