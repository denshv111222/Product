package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strconv"
	"strings"
)

type Categoryrepository struct {
	storage *Storage
}

var (
	tableCategories string = "categories"
)

func (catRep *Categoryrepository) CreateCategory(c *models.Categories) (*models.Categories, error) {
	query := fmt.Sprintf("INSERT INTO %s (name,slug,parent_id) VALUES ($1,$2,$3) RETURNING id_categories", tableCategories)
	if err := catRep.storage.db.QueryRow(query, c.Name, c.Slug, c.Parent_id).Scan(&c.Id); err != nil {
		return nil, err
	}
	return c, nil
}
func (catRep *Categoryrepository) DeleteCategory(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id_categories = %d ", tableCategories, id)
	fmt.Println(query)
	if _, err := catRep.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (catRep *Categoryrepository) UpdateCategory(id int, c *models.Categories) error {
	fmt.Println(id)
	fmt.Println(c)
	fieldlist := make([]string, 0)
	fmt.Println(fieldlist)
	if c.Name != "" {
		fieldlist = append(fieldlist, ("name='" + c.Name + "'"))
	}
	fmt.Println(fieldlist)
	if c.Slug != "" {
		fieldlist = append(fieldlist, ("slug='" + c.Slug + "'"))
	}
	fmt.Println(fieldlist)
	if c.Parent_id != 0 {
		fieldlist = append(fieldlist, ("parent_id=" + strconv.Itoa(c.Parent_id) + ""))
	}
	fmt.Println(fieldlist)
	request := strings.Join(fieldlist, ", ")
	fmt.Println(request)

	querry := fmt.Sprintf("UPDATE %s SET %s WHERE id_categories = %d RETURNING *", tableCategories, request, id)
	fmt.Println(querry)
	if err := catRep.storage.db.QueryRow(querry).Scan(&c.Id, &c.Name, &c.Slug, &c.Parent_id); err != nil {
		return err
	}
	return nil
}
func (catRep *Categoryrepository) FilterAllCategories(fil *models.PageRequest) ([]*models.Categories, error) { //срань которую надо переделать(а можно и не переделывать)
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
	query := fmt.Sprintf("Select count(*) FROM %s %s %s", tableCategories, where, sort)
	if err := catRep.storage.db.QueryRow(query).Scan(&fil.TotalRecords); err != nil {
		return nil, err
	}

	request = strings.Join(sortList, ",")
	sort = sort + request
	fmt.Println(request)
	query = fmt.Sprintf("Select * FROM %s %s %s LIMIT %d OFFSET %d", tableCategories, where, sort, fil.PageLength, (fil.PageNumber-1)*fil.PageLength)
	where, sort = "", ""

	fmt.Println(query)

	rows, err := catRep.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := make([]*models.Categories, 0)
	for rows.Next() {
		c := models.Categories{}
		err := rows.Scan(&c.Id, &c.Name, &c.Slug, &c.Parent_id)
		if err != nil {
			log.Println(err)
			continue
		}
		categories = append(categories, &c)
	}
	return categories, nil
}
func (catRep *Categoryrepository) GetCategoryById(id int) (*models.Categories, error) {

	fmt.Println(id)
	var cat models.Categories
	query := fmt.Sprintf("Select * FROM %s WHERE id_categories = %d", tableCategories, id)
	fmt.Println(query)
	if err := catRep.storage.db.QueryRow(query).Scan(&cat.Id, &cat.Name, &cat.Slug, &cat.Parent_id); err != nil {
		return nil, err
	}
	fmt.Println(query)
	return &cat, nil

}
