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

# 修复 openapi-generator 生成的命名冲突：
# hasDiscountCode 字段会与 DiscountCode 的 HasDiscountCode() 方法冲突。
# 仅重命名 Go 字段标识符，保留 json tag: `hasDiscountCode` 不变。
echo "正在修复 hasDiscountCode 命名冲突..."
for f in model_unibee_api_merchant_invoice_list_req.go model_unibee_api_merchant_payment_time_line_list_req.go; do
  if [ -f "$f" ]; then
    perl -0777 -i -pe 's/\bHasDiscountCode\s+\*bool\s+`json:"hasDiscountCode,omitempty"`/HasDiscountCodeFilter *bool `json:"hasDiscountCode,omitempty"`/g; s/\bo\.HasDiscountCode\b/o.HasDiscountCodeFilter/g' "$f"
  fi
done
echo "hasDiscountCode 命名冲突修复完成"
