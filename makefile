include META


.PHONY: run
run: 
	@echo "Running the service..."
	@GOPRIVATE=github.com/MPay-Solutions GONOPROXY=github.com/MPay-Solutions GONOSUMDB=github.com/MPay-Solutions CGO_ENABLED=0 go run -mod=mod -installsuffix cgo .
	@echo "Done!"


.PHONY: push
push:
	@echo "[BUILD] Committing and pushing to remote repository"
	@echo " - Committing"
	@git add META
	@git commit -am "v$(VERSION)"
	@echo " - Tagging"
	@git tag v${VERSION}
	@echo " - Pushing"
	@git push --tags origin ${BRANCH}

