package repo

import (
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

func (o *InstanceDataRepo) SetData(instanceId, name string, value interface{}) {

	dataType := ""
	dataValue := ""

	o.Exec(`insert into unaflow_instance_datas 
	(instance_id, data_name, data_type, data_value)
	values 
	(?, ?, ?, ?)
	on conflict on constraint unaflow_instance_datas_pkey 
	do update set data_value = ?`,
		instanceId, name, dataType, dataValue, dataValue)
}

func (o *InstanceDataRepo) GetData(instanceId, name string) interface{} {

	return nil
}

func (o *InstanceDataRepo) ListData(instanceId string) map[string]interface{} {

	return nil
}
