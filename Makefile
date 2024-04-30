.PHONY: dao
dao:
	@echo "Moving files from internal/dao/internal/ to model/dao/internal/"
	@mv internal/dao/*.go dao/
	@mv internal/dao/internal/*.go dao/internal/
	@mv internal/model/do/*.go model/do/
	@mv internal/model/entity/*.go model/entity/
	@echo "Files moved successfully."
