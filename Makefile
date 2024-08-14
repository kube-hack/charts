.PHONY: index

# This would be a good place to have someone manually create a pull request, rather than trying to automate it.
index:
	@echo -n "What is the URL of the github repository? "
	@read line; echo $$line 