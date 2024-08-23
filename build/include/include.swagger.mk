#
# This file provides all targets needed in work with swagger.
#

swagger/pull:
	@docker pull quay.io/goswagger/swagger

swagger/version:
	bash -e $(SCRIPTS_PATH)/swagger.sh version

swagger/generate:
	$(SCRIPTS_PATH)/swagger.sh generate spec -o /opt/work/swagger/swagger.json -w /opt/work

swagger/serve:
	$(SCRIPTS_PATH)/swagger.sh serve -F swagger /opt/work/swagger/swagger.json -p 8090 --no-open

swagger/help:
	@echo
	@echo '*** Swagger utility targets ***'
	@echo
	@echo 'Usage:'
	@echo '    make swagger/pull'
	@echo '    make swagger/version'
	@echo '    make swagger/generate'
	@echo '    make swagger/serve'
