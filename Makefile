_run:
	@$(MAKE) --warn-undefined-variables -f tools/make/common.mk $(MAKECMDGOALS)
.PHONY: _run
$(if $(MAKECMDGOALS),$(MAKECMDGOALS): %: _run)