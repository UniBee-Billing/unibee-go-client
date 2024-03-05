rm -rf *.go
rm -rf test/*.go
rm -rf docs/*.md
jdk17
openapi-generator generate \
-i http://api.unibee.top/api.sdk.generator.json \
-g go \
-o . \
--git-repo-id unibee-go-client \
--git-user-id UniB-e-e