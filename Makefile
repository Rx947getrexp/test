dao:
	@echo "Moving files from internal/dao/internal/ to model/dao/internal/"
	@mv internal/dao/internal/*.go dao/internal/
	@echo "Files moved successfully."

clean:
	@mv cmd/go-admin/config.yaml ~/src/backup/config.yaml.admin
	@mv cmd/go-api/config.yaml ~/src/backup/config.yaml.api
	@mv cmd/go-executor/config.yaml ~/src/backup/config.yaml.executor
	@mv cmd/go-job/config.yaml ~/src/backup/config.yaml.job
	@mv cmd/go-upload/config.yaml ~/src/backup/config.yaml.upload