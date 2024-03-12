mv unibee.go unibee.backup
mv configuration.go configuration.backup
rm -rf *.go
rm -rf docs/*.md
java -jar openapi-generator-cli.jar generate \
-i http://api.unibee.top/api.sdk.generator.json \
-g go \
-o . \
--git-repo-id unibee-go-client \
--git-user-id UniB-e-e \
-p packageName=unibee \
-c config.yaml
mv unibee.backup unibee.go
rm configuration.go
mv configuration.backup configuration.go