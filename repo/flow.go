package repo

import (
	"strings"
	"sync"

	"github.com/adlindo/gocom"
	"github.com/adlindo/gocom/config"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Flow struct {
	gorm.Model
	Id       string
	Code     string
	Name     string
	Script   string
	IsActive bool
}

func (o *Flow) BeforeCreate(db *gorm.DB) error {

	if o.Id == "" {

		o.Id = ulid.Make().String()
	}

	return nil
}

func (o *Flow) TableName() string {
	return "unaflow_flows"
}

// -------------------

type FlowRepo struct {
	gocom.BaseRepo
}

var flowRepo *FlowRepo
var flowRepoOnce sync.Once

func GetFlowRepo() *FlowRepo {

	flowRepoOnce.Do(func() {

		flowRepo = &FlowRepo{
			BaseRepo: gocom.BaseRepo{
				ConnName: config.Get("unaflow.dbname"),
			},
		}

		flowRepo.AutoMigrate(&Flow{})
	})

	return flowRepo
}

func (o *FlowRepo) Create(mdl *Flow) {

	o.BaseRepo.Create(mdl)
}

func (o *FlowRepo) GetById(id string) *Flow {

	ret := &Flow{}

	if o.Model(*ret).Where("id = ?", id).First(ret).Error != nil {

		return nil
	}

	return ret
}

func (o *FlowRepo) GetByCode(code string) *Flow {

	ret := &Flow{}

	if o.Model(*ret).Where("code = ?", code).First(ret).Error != nil {

		return nil
	}

	return ret
}

func (o *FlowRepo) Search(filter string, pageNo, pageLength int) ([]*Flow, int64) {

	ret := []*Flow{}

	tx := o.Model(Flow{})

	if filter != "" {

		filter = "%" + strings.ToUpper(filter) + "%"
		tx = tx.Where("upper(name) like ? or upper(code) like ?", filter, filter)
	}

	if pageLength > 0 {
		tx = tx.Offset((pageNo - 1) * pageLength).Limit(pageLength)
	}

	var total int64 = 0

	tx.Count(&total)
	tx.Find(&ret)

	return ret, total
}
