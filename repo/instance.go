package repo

import (
	"sync"

	"github.com/adlindo/gocom"
	"github.com/adlindo/gocom/config"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Instance struct {
	gorm.Model
	Id     string
	FlowId string
	StepId string
}

func (o *Instance) BeforeCreate(db *gorm.DB) error {

	if o.Id == "" {

		o.Id = ulid.Make().String()
	}

	return nil
}

func (o *Instance) TableName() string {
	return "unaflow_instances"
}

//-----------------------------------------------------------------------------------

type InstanceRepo struct {
	gocom.BaseRepo
}

var instanceRepo *InstanceRepo
var instanceRepoOnce sync.Once

func GetInstanceRepo() *InstanceRepo {

	instanceRepoOnce.Do(func() {

		instanceRepo = &InstanceRepo{
			BaseRepo: gocom.BaseRepo{
				ConnName: config.Get("unaflow.dbname"),
			},
		}

		instanceRepo.AutoMigrate(&Instance{})
	})

	return instanceRepo
}

func (o *InstanceRepo) GetById(instanceId string) *Instance {

	ret := &Instance{}

	if o.Model(*ret).Where("id = ?", instanceId).First(ret).Error != nil {
		return nil
	}

	return ret
}

func (o *InstanceRepo) SetStepId(instanceId, stepIs string) {

	o.Exec("update ")
}
