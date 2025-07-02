mv unibee.go unibee.backup
mv configuration.go configuration.backup
rm -rf *.go
rm -rf docs/*.md
java -jar openapi-generator-cli.jar generate \
-i https://api.unibee.top/api.sdk.generator.json?hideSecurity=true \
-g go \
-o . \
--git-repo-id unibee-go-client \
--git-user-id UniBee-Billing \
-p packageName=unibee \
-c config.yaml
mv unibee.backup unibee.go
rm configuration.go
mv configuration.backup configuration.go
sed -i '' 's/map\[string\]map\[string\]interface/map\[string\]interface/g' *.go

# 删除所有 Go 文件中的 decoder.DisallowUnknownFields() 行
echo "正在删除 decoder.DisallowUnknownFields() 行..."
find . -name "*.go" -type f -exec sed -i '' '/decoder\.DisallowUnknownFields()/d' {} \;
echo "decoder.DisallowUnknownFields() 行删除完成"
