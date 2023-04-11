package repo

import (
	"strings"
	"sync"
	"time"

	"github.com/adlindo/gocom"
	"github.com/adlindo/gocom/config"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Instance struct {
	gorm.Model
	Id         string
	FlowId     string
	Key        string
	StepId     string
	Status     int
	StopReason string
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

func (o *InstanceRepo) SetStepStatus(instanceId, stepId string, status int) {

	o.Exec("update unaflow_instances set updated_at = ?, step_id = ?, status = ? where id = ?", time.Now(), stepId, status, instanceId)
}

func (o *InstanceRepo) SetStepId(instanceId, stepId string) {

	o.Exec("update unaflow_instances set updated_at = ?, step_id = ? where id = ?", time.Now(), stepId, instanceId)
}

func (o *InstanceRepo) SetStatus(instanceId string, status int) {

	o.Exec("update unaflow_instances set updated_at = ?, status = ? where id = ?", time.Now(), status, instanceId)
}

func (o *InstanceRepo) SetStatusReason(instanceId string, status int, reason string) {

	o.Exec("update unaflow_instances set updated_at = ?, status = ?, stop_reason = ? where id = ?", time.Now(), status, reason, instanceId)
}

func (o *InstanceRepo) Search(flowId, stepId, filter string, pageNo, pageLength int) ([]*Instance, int64) {

	ret := []*Instance{}

	tx := o.Model(Instance{}).Debug()

	if flowId != "" {

		tx = tx.Where("flow_id = ?", flowId)
	}

	if stepId != "" {

		tx = tx.Where("step_id = ?", stepId)
	}

	if filter != "" {

		tx = tx.Where("upper(key) like ?", "%"+strings.ToUpper(filter)+"%")
	}

	if pageLength > 0 {
		tx = tx.Offset((pageNo - 1) * pageLength).Limit(pageLength)
	}

	var total int64 = 0

	tx.Count(&total)
	tx.Find(&ret)

	return ret, total
}
