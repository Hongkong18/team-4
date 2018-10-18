package inmemory_store

import (
	"c4g_go/models"
	"strings"
)

type InstitutionService struct {
	InvIndex     models.InvertedIndex
	institutions map[int64]*models.Institution
	latestId     int64
}

func NewInstitutionService(invIndex models.InvertedIndex) *InstitutionService {
	ins := make(map[int64]*models.Institution)
	return &InstitutionService{InvIndex: invIndex, institutions: ins, latestId: 0}
}

//GetById(id int64) *Institution
//Insert() *Institution
//Update(id int64, ins *Institution)

func (i *InstitutionService) ListAll() []*models.Institution {
	rv := make([]*models.Institution, 0)
	for _, v := range i.institutions {
		rv = append(rv, v)
	}

	return rv
}

func (i *InstitutionService) GetById(id int64) *models.Institution {
	return i.institutions[id]
}

func (i *InstitutionService) Insert(ins *models.Institution) {
	i.latestId = i.latestId + 1
	i.institutions[i.latestId] = ins

	wordbag := make([]string, 0)
	wordbag = append(wordbag, strings.Split(ins.Name, " ")...)
	wordbag = append(wordbag, strings.Split(ins.Speciality, " ")...)
	wordbag = append(wordbag, strings.Split(ins.Location, " ")...)
	wordbag = append(wordbag, strings.Split(ins.Description, " ")...)

	for _, word := range wordbag {
		i.InvIndex.Insert(word, i.latestId)
	}
}
