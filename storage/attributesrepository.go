package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strconv"
	"strings"
)

type Attributsrepository struct {
	storage *Storage
}

var (
	tableAttributes string = "atributes"
)

func (atRep *Attributsrepository) DeleteAttributes(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id_atribute = %d ", tableAttributes, id)
	fmt.Println(query)
	if _, err := atRep.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (atRep *Attributsrepository) CreateAttributes(attr *models.Attributes) error {
	fmt.Println("Запрос")
	querry := fmt.Sprintf("INSERT INTO %s (name,slug,unit_id) VALUES($1, $2, $3) RETURNING id_atribute", tableAttributes)
	fmt.Println(querry)
	if err := atRep.storage.db.QueryRow(querry, attr.Name, attr.Slug, attr.Units.Id).Scan(&attr.Id); err != nil {

		fmt.Println(err)
		return err
	}

	return nil
}
func (atRep *Attributsrepository) UpdateAttribute(id int, a *models.Attributes) error {

	fmt.Println(id)
	fmt.Println(a)
	fieldlist := make([]string, 0)
	fmt.Println(fieldlist)
	if a.Name != "" {
		fieldlist = append(fieldlist, ("name='" + a.Name + "'"))
	}
	fmt.Println(fieldlist)
	if a.Slug != "" {
		fieldlist = append(fieldlist, ("slug='" + a.Slug + "'"))
	}
	fmt.Println(fieldlist)
	if a.Units.Id != 0 {
		fieldlist = append(fieldlist, ("unit_id='" + strconv.Itoa(a.Units.Id) + "'"))
	}
	fmt.Println(fieldlist)
	request := strings.Join(fieldlist, ", ")
	fmt.Println(request)

	querry := fmt.Sprintf("UPDATE %s SET %s WHERE id_atribute = %d RETURNING *", tableAttributes, request, id)
	fmt.Println(querry)
	if err := atRep.storage.db.QueryRow(querry).Scan(&a.Id, &a.Name, &a.Slug, &a.Units.Id); err != nil {
		return err
	}
	return nil
}
func (atRep *Attributsrepository) FilterAllAttributes(fil *models.PageRequest) ([]*models.Attributes, error) { //срань которую надо переделать(а можно и не переделывать)
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
			fmt.Println(sorts.Order)
			sort = "Order by "
			if !sorts.Order {
				sortList = append(sortList, sorts.Name+" "+"DESC ")
			} else {
				sortList = append(sortList, sorts.Name+" "+"ASC ")

			}
		}
	}
	request := strings.Join(fieldlist, " and ")
	where = where + request
	request = ""

	request = strings.Join(sortList, ",")
	sort = sort + request
	fmt.Println(request)
	query := fmt.Sprintf("Select count(*) FROM %s %s %s", tableAttributes, where, sort)
	if err := atRep.storage.db.QueryRow(query).Scan(&fil.TotalRecords); err != nil {
		return nil, err
	}

	request = strings.Join(sortList, ",")
	sort = sort + request
	fmt.Println(request)
	query = fmt.Sprintf("Select %s.id_atribute,%s.name,%s.slug ,%s.id_unit,%s.name,%s.slug from %s inner join %s on  %s.unit_id = %s.id_unit %s %s LIMIT %d OFFSET %d", tableAttributes, tableAttributes, tableAttributes, tableunits, tableunits, tableunits, tableunits, tableAttributes, tableAttributes, tableunits, where, sort, fil.PageLength, (fil.PageNumber-1)*fil.PageLength)
	where, sort = "", ""

	fmt.Println(query)

	fmt.Print(query)
	rows, err := atRep.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	attributs := make([]*models.Attributes, 0)
	log.Println(attributs)
	for rows.Next() {
		var (
			units models.Units
		)
		a := models.Attributes{
			Units: &units,
		}
		log.Println(rows)
		err := rows.Scan(&a.Id, &a.Name, &a.Slug, &a.Units.Id, &a.Units.Name, &a.Units.Slug)
		if err != nil {
			log.Println(err)
			continue
		}
		attributs = append(attributs, &a)
	}
	return attributs, nil
}
