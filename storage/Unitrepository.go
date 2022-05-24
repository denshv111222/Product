package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strings"
)

type Unitrepository struct {
	storage *Storage
}

var (
	tableunits string = "units"
)

func (unRep *Unitrepository) CreateUnits(u *models.Units) (*models.Units, error) {

	fmt.Println(u.Slug)
	query := fmt.Sprintf("INSERT INTO %s (name,slug) VALUES ($1,$2) RETURNING id_unit", tableunits)
	if err := unRep.storage.db.QueryRow(query, u.Name, u.Slug).Scan(&u.Id); err != nil {
		fmt.Println(query)
		return nil, err
	}
	fmt.Println(query)
	return u, nil
}
func (unRep *Unitrepository) UpdateUnit(id int, u *models.Units) error {
	fmt.Println(id)
	fmt.Println(u)
	fieldlist := make([]string, 0)
	if u.Name != "" {
		fieldlist = append(fieldlist, ("name='" + u.Name + "'"))
	}
	if u.Slug != "" {
		fieldlist = append(fieldlist, ("slug='" + u.Slug + "'"))
	}
	fmt.Println(fieldlist)
	request := strings.Join(fieldlist, ", ")
	fmt.Println(request)

	querry := fmt.Sprintf("UPDATE %s SET %s WHERE id_unit = %d", tableunits, request, id)
	fmt.Println(querry)
	if err := unRep.storage.db.QueryRow(querry).Scan(&u.Id, &u.Name, &u.Slug); err != nil {
		return err
	}
	return nil
}
func (unRep *Unitrepository) DeleteUnit(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id_unit = %d ", tableunits, id)
	fmt.Println(query)
	if _, err := unRep.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (unRep *Unitrepository) FilterAllUnit(fil *models.Filter) ([]*models.Units, error) { //срань которую надо переделать(а можно и не переделывать)
	fieldlist := make([]string, 0)
	sortList := make([]string, 0)

	//
	if len(*fil.Fields) >= 1 {
		for _, filters := range *fil.Fields {
			if filters.Value != "" {
				where = "where "
				if vaild.IsInt(filters.Value) == true {
					fieldlist = append(fieldlist, filters.Field+filters.Operations+filters.Value)
				} else {
					fieldlist = append(fieldlist, filters.Field+" like "+"'%"+filters.Value+"%'")
				}
			} else {
				where = ""
			}
		}
	}
	request := strings.Join(fieldlist, " and ")
	where = where + request
	request = ""
	if len(*fil.Sorts) >= 1 {
		sort = "order by "
		for _, sorts := range *fil.Sorts {
			if sorts.Sort != "" {
				sortList = append(sortList, sorts.Sort+" "+sorts.SortView)
			} else {
				sort = ""
			}
		}
	}
	query := fmt.Sprintf("Select count(*) FROM %s,%s,%s", tableunits, where, sort)
	if err := unRep.storage.db.QueryRow(query).Scan(&fil.Pages.AllRecords); err != nil {
		return nil, err
	}
	fil.Pages.AllPages = allPage(fil.Pages.AllRecords, fil.Pages.CountsRecordOnPage)
	fil.Pages.RemainedRecords = fil.Pages.AllRecords - fil.Pages.CountsRecordOnPage*fil.Pages.СurrentPage
	//подумать над этим
	if fil.Pages.RemainedRecords < 0 {
		fil.Pages.RemainedRecords = 0
	}
	request = strings.Join(sortList, ",")
	sort = sort + request
	fmt.Println(request)
	query = fmt.Sprintf("Select * FROM %s %s %s LIMIT %d OFFSET %d", tableunits, where, sort, fil.Pages.CountsRecordOnPage, (fil.Pages.СurrentPage-1)*fil.Pages.CountsRecordOnPage)
	where, sort = "", ""

	fmt.Println(query)

	rows, err := unRep.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	units := make([]*models.Units, 0)
	for rows.Next() {
		c := models.Units{}
		err := rows.Scan(&c.Id, &c.Name, &c.Slug)
		if err != nil {
			log.Println(err)
			continue
		}
		units = append(units, &c)
	}
	return units, nil
}
