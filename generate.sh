rm -rf *.go
rm -rf docs/*.md
java -jar openapi-generator-cli.jar generate \
-i http://api.unibee.top/api.sdk.generator.json \
-g go \
-o . \
--git-repo-id unibee-go-client \
--git-user-id UniB-e-e \
-c config.yaml