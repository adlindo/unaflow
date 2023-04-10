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

func (o *InstanceRepo) Create(mdl *Instance) error {

	err := o.BaseRepo.Create(mdl).Error

	if err != nil {
		return err
	}

	return nil
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

func (o *InstanceRepo) Search(flowId, stepId string, pageNo, pageLength int) ([]*Instance, int64) {

	ret := []*Instance{}

	tx := o.Model(Instance{})

	if flowId != "" {

		tx = tx.Where("flow_id = ?", flowId)
	}

	if stepId != "" {

		tx = tx.Where("step_id = ?", flowId)
	}

	if pageLength > 0 {
		tx = tx.Offset((pageNo - 1) * pageLength).Limit(pageLength)
	}

	var total int64 = 0

	tx.Count(&total)
	tx.Find(ret)

	return ret, total
}
