package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/timex"
	log "github.com/sirupsen/logrus"
	"go-api/internal/util"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

type DefaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultValidator{}

var timeMark = []string{"today", "yesterday", "last_week", "last_month"}

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *DefaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
		_ = v.validate.RegisterValidation("PosIntSplitBy", posIntSplitBy)
		_ = v.validate.RegisterValidation("EndAbleDate", endAbleDate)
		_ = v.validate.RegisterValidation("LteTime", lteTime)
		_ = v.validate.RegisterValidation("LtTime", ltTime)
		_ = v.validate.RegisterValidation("GteTime", gteTime)
		_ = v.validate.RegisterValidation("GtTime", gtTime)
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

// 验证以separator为分隔符的正整数集合
// 当分隔的部分使用strconv.Atoi转换失败时，验证不通过
// 当分隔的部分数值<=0时，验证不通过
// Example:
// binding:"posIntSplitBy" 默认使用","作为分隔符
// binding:"posIntSplitBy=0x2C"
func posIntSplitBy(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	param := fl.Param()
	if param == "" {
		param = ","
	}
	splits := strings.Split(fl.Field().String(), param)
	for _, str := range splits {
		atoi, err := strconv.Atoi(str)
		if err != nil {
			return false
		}
		if atoi <= 0 {
			return false
		}
	}
	return true
}

func endAbleDate(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		if date.After(time.Now()) {
			log.WithField(fl.FieldName(), date).Error("请求参数校验失败：无法提供今日之后的数据")
			return false
		}

		startDate := fl.Parent().FieldByName("Start").Interface().(time.Time)
		mapInterface := make(map[string]interface{})
		mapInterface["start"] = timex.FromTime(startDate)
		mapInterface["end"] = timex.FromTime(date)

		if !date.After(startDate) {
			log.WithFields(mapInterface).Error("请求参数校验失败：开始时间不能大于结束时间")
			return false
		}
		if date.Add(-1 * 24 * 31 * time.Hour).After(startDate) {
			log.WithFields(mapInterface).Error("请求参数校验失败：开始日期与结束日期差值超过31天")
			return false
		}
	}
	return true
}

func timeParams(fl validator.FieldLevel) (time.Time, string) {
	validateParams := fl.Param()
	if validateParams == "" {
		validateParams = "today"
	}

	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return time.Time{}, validateParams
	}
	return date, validateParams
}

func lteTime(fl validator.FieldLevel) bool {
	date, p := timeParams(fl)
	if date.IsZero() {
		return false
	}
	if arrutil.Contains(timeMark, p) {
		tm := util.MarkToTime(p)
		if date.Before(tm) || date.Equal(tm) {
			return true
		}
	}

	tf := util.TimeStringToTime(p)
	if !tf.IsZero() && (date.Before(tf) || date.Equal(tf)) {
		return true
	}
	return false
}

func ltTime(fl validator.FieldLevel) bool {
	date, p := timeParams(fl)
	if date.IsZero() {
		return false
	}
	if arrutil.Contains(timeMark, p) {
		tm := util.MarkToTime(p)
		if date.Before(tm) {
			return true
		}
	}

	tf := util.TimeStringToTime(p)
	if !tf.IsZero() && date.Before(tf) {
		return true
	}
	return false
}

func gteTime(fl validator.FieldLevel) bool {
	date, p := timeParams(fl)
	if date.IsZero() {
		return false
	}

	if arrutil.Contains(timeMark, p) {
		tm := util.MarkToTime(p)
		if date.After(tm) || date.Equal(tm) {
			return true
		}
	}

	tf := util.TimeStringToTime(p)
	if !tf.IsZero() && (date.After(tf) || date.Equal(tf)) {
		return true
	}
	return false
}

func gtTime(fl validator.FieldLevel) bool {
	date, p := timeParams(fl)
	if date.IsZero() {
		return false
	}

	if arrutil.Contains(timeMark, p) {
		tm := util.MarkToTime(p)
		if date.After(tm) {
			return true
		}
	}

	tf := util.TimeStringToTime(p)
	if !tf.IsZero() && date.After(tf) {
		return true
	}
	return false
}
