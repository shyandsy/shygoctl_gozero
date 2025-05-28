package template

import _ "embed"

var (
	//go:embed testdata/test_api_template.api
	TestApiTemplate string
	//go:embed testdata/test_multi_service_template.api
	TestMultiServiceTemplate string
	//go:embed testdata/ap_ino_info.api
	ApiNoInfo string
	//go:embed testdata/invalid_api_file.api
	InvalidApiFile string
	//go:embed testdata/anonymous_annotation.api
	AnonymousAnnotation string
	//go:embed testdata/api_has_middleware.api
	ApiHasMiddleware string
	//go:embed testdata/api_jwt.api
	ApiJwt string
	//go:embed testdata/api_jwt_with_middleware.api
	ApiJwtWithMiddleware string
	//go:embed testdata/api_has_no_request.api
	ApiHasNoRequest string
	//go:embed testdata/api_route_test.api
	ApiRouteTest string
	//go:embed testdata/has_comment_api_test.api
	HasCommentApiTest string
	//go:embed testdata/has_inline_no_exist_test.api
	HasInlineNoExistTest string
	//go:embed testdata/import_api.api
	ImportApi string
	//go:embed testdata/has_import_api.api
	HasImportApi string
	//go:embed testdata/no_struct_tag_api.api
	NoStructTagApi string
	//go:embed testdata/nest_type_api.api
	NestTypeApi string
	//go:embed testdata/import_twice.api
	ImportTwiceApi string
	//go:embed testdata/another_import_api.api
	AnotherImportApi string
	//go:embed testdata/example.api
	ExampleApi string
)

//go:embed config.tpl
var ConfigTemplate string

//go:embed etc.tpl
var EtcTemplate string

//go:embed handler.tpl
var HandlerTemplate string

//go:embed handler_test.tpl
var HandlerTestTemplate string

//go:embed logic.tpl
var LogicTemplate string

//go:embed logic_test.tpl
var LogicTestTemplate string

//go:embed main.tpl
var MainTemplate string

//go:embed middleware.tpl
var MiddlewareImplementCode string

//go:embed svc.tpl
var ContextTemplate string

//go:embed types.tpl
var TypesTemplate string
