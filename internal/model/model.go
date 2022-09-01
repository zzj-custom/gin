package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type XTime struct {
	time.Time
}

type Model struct {
	ID        int   `gorm:"primary_key" json:"id"`
	CreatedAt XTime `gorm:"column:create_at;autoCreateTime;not null;comment:创建时间" json:"created_at" `
	UpdatedAt XTime `gorm:"column:update_at;autoUpdateTime;not null;comment:更新时间" json:"updated_at" `
}

// MarshalJSON 2. 为 Time 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// Value 3. 为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 4. 为 Time 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
