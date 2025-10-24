package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Validate   *validator.Validate
	Translator ut.Translator
)

func init() {
	// 初始化validator
	Validate = validator.New()

	// 初始化中文翻译器
	zhLocale := zh.New()
	uni := ut.New(zhLocale, zhLocale)
	Translator, _ = uni.GetTranslator("zh")

	// 注册中文翻译
	err := zh_translations.RegisterDefaultTranslations(Validate, Translator)
	if err != nil {
		panic(fmt.Sprintf("注册中文翻译失败: %v", err))
	}

	// 注册自定义翻译
	registerCustomTranslations()
}

// registerCustomTranslations 注册自定义验证规则的中文翻译
func registerCustomTranslations() {
	// 自定义验证规则翻译
	customTranslations := map[string]string{
		"required": "{0}是必填字段",
		"min":      "{0}长度不能少于{1}个字符",
		"max":      "{0}长度不能超过{1}个字符",
		"email":    "{0}必须是有效的邮箱地址",
		"len":      "{0}长度必须是{1}个字符",
		"gte":      "{0}必须大于或等于{1}",
		"lte":      "{0}必须小于或等于{1}",
		"gt":       "{0}必须大于{1}",
		"lt":       "{0}必须小于{1}",
		"oneof":    "{0}必须是以下值之一: {1}",
		"numeric":  "{0}必须是数字",
		"alpha":    "{0}只能包含字母",
		"alphanum": "{0}只能包含字母和数字",
		"url":      "{0}必须是有效的URL地址",
		"uuid":     "{0}必须是有效的UUID格式",
		"datetime": "{0}必须是有效的日期时间格式",
		"phone":    "{0}必须是有效的手机号码",
	}

	for tag, translation := range customTranslations {
		registerTranslation(tag, translation)
	}

	// 注册字段名翻译
	registerFieldTranslations()
}

// registerTranslation 注册单个验证规则的中文翻译
func registerTranslation(tag, translation string) {
	Validate.RegisterTranslation(tag, Translator, func(ut ut.Translator) error {
		return ut.Add(tag, translation, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field(), fe.Param())
		return t
	})
}

// TranslateError 将验证错误翻译为中文
func TranslateError(err error) string {
	if err == nil {
		return ""
	}

	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Translate(Translator))
	}

	return strings.Join(errors, "; ")
}

// registerFieldTranslations 注册字段名的中文翻译
func registerFieldTranslations() {
	// 注册字段名翻译函数
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// 优先使用label标签
		if label := fld.Tag.Get("label"); label != "" {
			return label
		}
		// 其次使用json标签
		if json := fld.Tag.Get("json"); json != "" {
			return json
		}
		// 最后使用字段名
		return fld.Name
	})
}

// TranslateFirstError 将验证错误翻译为中文，只返回第一个错误
func TranslateFirstError(err error) string {
	if err == nil {
		return ""
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	if len(validationErrors) > 0 {
		return validationErrors[0].Translate(Translator)
	}

	return ""
}

// ValidateStruct 验证结构体并返回中文错误信息
func ValidateStruct(s interface{}) error {
	err := Validate.Struct(s)
	if err != nil {
		return fmt.Errorf("%s", TranslateError(err))
	}
	return nil
}

// ValidateStructFirstError 验证结构体并返回第一个中文错误信息
func ValidateStructFirstError(s interface{}) error {
	err := Validate.Struct(s)
	if err != nil {
		return fmt.Errorf("%s", TranslateFirstError(err))
	}
	return nil
}
