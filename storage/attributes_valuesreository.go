package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strconv"
	"strings"
)

type Attributs_valuesrepository struct {
	storage *Storage
}

var (
	tableAttributes_values string = "atributes_values"
)

func (atvalRep *Attributs_valuesrepository) DeleteAttributes_values(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d ", tableAttributes, id)
	fmt.Println(query)
	if _, err := atvalRep.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (atvalRep *Attributs_valuesrepository) CreateAttributes(attr *models.Attributes_values) error {
	fmt.Println("Запрос")
	fmt.Println(attr.Attributes.Id)
	querry := fmt.Sprintf("INSERT INTO %s (name,atribute_id) VALUES($1, $2) RETURNING id", tableAttributes_values)
	fmt.Println(querry)
	if err := atvalRep.storage.db.QueryRow(querry, attr.Name, attr.Attributes.Id).Scan(&attr.Id); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
func (atvalRep *Attributs_valuesrepository) UpdateAttribute(id int, a *models.Attributes_values) error {

	fmt.Println(id)
	fmt.Println(a)
	fieldlist := make([]string, 0)
	fmt.Println(fieldlist)
	if a.Name != "" {
		fieldlist = append(fieldlist, ("name='" + a.Name + "'"))
	}
	fmt.Println(fieldlist)
	if a.Attributes.Id != 0 {
		fieldlist = append(fieldlist, ("atribute_id='" + strconv.Itoa(a.Attributes.Id) + "'"))
	}
	fmt.Println(fieldlist)
	request := strings.Join(fieldlist, ",")
	fmt.Println(request)

	querry := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d RETURNING *", tableAttributes_values, request, id)
	fmt.Println(querry)
	if err := atvalRep.storage.db.QueryRow(querry).Scan(&a.Id, &a.Name, &a.Attributes.Id); err != nil {
		return err
	}
	return nil
}
func (atvalRep *Attributs_valuesrepository) FilterAllatributes_values(fil *models.PageRequest) ([]*models.Attributes_values, error) { //срань которую надо переделать(а можно и не переделывать)
	fieldlist := make([]string, 0)
	sortList := make([]string, 0)

	//
	if len(*fil.Fields) != 0 {
		for _, filters := range *fil.Fields {
			if filters.Value != "" {
				where = "where "
				if vaild.IsInt(filters.Value) == true {
					fieldlist = append(fieldlist, filters.Name+filters.Operation+filters.Value)
				} else {
					fieldlist = append(fieldlist, filters.Name+" like "+"'%"+filters.Value+"%'")
				}
			} else {
				where = ""
			}
		}
		for _, sorts := range *fil.Fields {
			if sorts.Order == false {
				sortList = append(sortList, sorts.Name+" "+"DESC")
			} else {
				sort = ""
			}
		}
	}
	request := strings.Join(fieldlist, " and ")
	where = where + request
	request = ""

	request = strings.Join(sortList, ",")
	sort = sort + request
	fmt.Println(request)
	query := fmt.Sprintf("Select count(*) FROM %s %s %s", tableAttributes_values, where, sort)

	if err := atvalRep.storage.db.QueryRow(query).Scan(&fil.TotalRecords); err != nil {
		return nil, err
	}

	request = strings.Join(sortList, ",")
	sort = sort + request
	fmt.Println(request)
	query = fmt.Sprintf("select %s.id,%s.name,%s.id_atribute,%s.name,%s.slug,%s.id_unit,%s.name,%s.slug from %s inner join %s  on %s.atribute_id = %s.id_atribute inner join %s on %s.unit_id =%s.id_unit %s %s LIMIT %d OFFSET %d", tableAttributes_values, tableAttributes_values, tableAttributes, tableAttributes, tableAttributes, tableunits, tableunits, tableunits, tableAttributes, tableAttributes_values, tableAttributes_values, tableAttributes, tableunits, tableAttributes, tableunits, where, sort, fil.PageLength, (fil.PageNumber-1)*fil.PageLength)
	where, sort = "", ""

	fmt.Print(query)
	rows, err := atvalRep.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	attribute := make([]*models.Attributes_values, 0)
	log.Println(attribute)
	for rows.Next() {
		unit := models.Units{}
		attr := models.Attributes{
			Units: &unit,
		}

		a := models.Attributes_values{
			Attributes: &attr,
		}

		log.Println(rows)
		err := rows.Scan(&a.Id, &a.Name, &a.Attributes.Id, &a.Attributes.Name, &a.Attributes.Slug, &a.Attributes.Units.Id, &a.Attributes.Units.Name, &a.Attributes.Units.Slug)
		if err != nil {
			log.Println(err)
			continue
		}
		attribute = append(attribute, &a)
	}
	return attribute, nil
}
