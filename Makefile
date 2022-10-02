BINARY := terraform-provider-validation
LOCALPATH := ~/.terraform.d/plugins/github.com/tlkamp/validation/0.0.1/linux_amd64/
DOCS := docs/

.PHONY:
gendocs:
	@go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
	@go generate

.PHONY:
test:
	@go test -v ./...

.PHONY:
acctest:
	@TF_ACC=true go test -coverprofile=.coverage.txt -v ./...
	@go tool cover -html=.coverage.txt -o ./coverage.html

$(BINARY): gendocs test acctest
	@go build -o $(BINARY)

.PHONY:
testinstall: $(BINARY)
	@mkdir -p $(LOCALPATH)
	@rm -f terraform/.terraform.lock.hcl
	@mv $(BINARY) $(LOCALPATH)

.PHONY:
clean:
	@go mod tidy
	@go clean -testcache
	@rm -f $(LOCALPATH)/$(BINARY)
	@rm -f terraform/.terraform.lock.hcl
	@rm -rf terraform/.terraform
	@rm -f .coverage.txt
	@rm -f coverage.html
