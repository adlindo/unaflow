package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/adlindo/gocom"
	"github.com/adlindo/gocom/config"
)

type InstanceData struct {
	InstanceId string `gorm:"primary_key:true"`
	DataName   string `gorm:"primary_key:true"`
	DataType   string
	DataValue  string
}

func (o *InstanceData) TableName() string {
	return "unaflow_instance_datas"
}

//-------------------------------------------------

type InstanceDataRepo struct {
	gocom.BaseRepo
}

var instanceDataRepo *InstanceDataRepo
var instanceDataRepoOnce sync.Once

func GetInstanceDataRepo() *InstanceDataRepo {

	instanceDataRepoOnce.Do(func() {

		instanceDataRepo = &InstanceDataRepo{
			BaseRepo: gocom.BaseRepo{
				ConnName: config.Get("unaflow.dbname"),
			},
		}

		instanceDataRepo.AutoMigrate(&InstanceData{})
	})
	return instanceDataRepo
}

func (o *InstanceDataRepo) SetData(instanceId, name string, value interface{}) error {

	dataType := ""
	dataValue := ""

	switch value.(type) {
	case int:
		dataType = "int"
		dataValue = fmt.Sprintf("%d", value)
	case float32:
	case float64:
		dataType = "float"
		dataValue = fmt.Sprintf("%f", value)
	case string:
		dataType = "string"
		dataValue = value.(string)
	case bool:
		dataType = "bool"
		dataValue = fmt.Sprintf("%t", value)
	case []interface{}:
		dataType = "array"

		bytea, err := json.Marshal(value)

		if err == nil {
			dataValue = string(bytea)
		}

	case map[string]interface{}:
		dataType = "map"

		bytea, err := json.Marshal(value)

		if err == nil {
			dataValue = string(bytea)
		}
	default:
		return errors.New("data type not supported")
	}

	o.Exec(`insert into unaflow_instance_datas 
	(instance_id, data_name, data_type, data_value)
	values 
	(?, ?, ?, ?)
	on conflict on constraint unaflow_instance_datas_pkey 
	do update set data_value = ?`,
		instanceId, name, dataType, dataValue, dataValue)

	return nil
}

func (o *InstanceDataRepo) ParseData(dataType, dataValue string) interface{} {

	switch dataType {
	case "int":
		val, err := strconv.Atoi(dataValue)
		if err != nil {
			return 0
		}
		return val
	case "float":
		val, err := strconv.ParseFloat(dataValue, 64)
		if err != nil {
			return 0
		}
		return val
	case "string":
		return dataValue
	case "bool":
		val, err := strconv.ParseBool(dataValue)
		if err != nil {
			return 0
		}
		return val
	case "array":
		val := []interface{}{}
		json.Unmarshal([]byte(dataValue), &val)
		return val
	case "map":
		val := map[string]interface{}{}
		json.Unmarshal([]byte(dataValue), &val)
		return val
	default:
		return nil
	}
}

func (o *InstanceDataRepo) GetData(instanceId, name string) interface{} {

	ret := &InstanceData{}

	if o.Model(*ret).
		Where("instance_id = ? and data_name = ?", instanceId, name).
		First(ret).Error != nil {

		return nil
	}

	return o.ParseData(ret.DataType, ret.DataValue)
}

func (o *InstanceDataRepo) ListData(instanceId string) map[string]interface{} {

	ret := map[string]interface{}{}
	retMdl := []InstanceData{}

	if o.Model(InstanceData{}).
		Where("instance_id = ?", instanceId).
		Find(retMdl).Error != nil {

		for _, item := range retMdl {

			ret[item.DataName] = o.ParseData(item.DataType, item.DataValue)
		}
	}

	return ret
}
