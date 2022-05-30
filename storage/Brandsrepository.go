package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strings"
)

type Brandrepository struct {
	storage *Storage
}

var (
	tableBrends string = "brands"
	fildFilter  string
	where       string
	sort        string
)

func (brRep *Brandrepository) CreateBrand(c *models.Brand) (*models.Brand, error) {
	fmt.Println(c.Slug)
	query := fmt.Sprintf("INSERT INTO %s (name,slug) VALUES ($1,$2) RETURNING id", tableBrends)
	if err := brRep.storage.db.QueryRow(query, c.Name, c.Slug).Scan(&c.Id); err != nil {
		return nil, err
	}
	return c, nil
}
func (brRep *Brandrepository) DeleteBrand(id int) error {
	fmt.Println(id)
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d ", tableBrends, id)
	fmt.Println(query)
	if _, err := brRep.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (brRep *Brandrepository) UpdateBrand(id int, c *models.Brand) error {
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
	request := strings.Join(fieldlist, ", ")
	fmt.Println(request)

	querry := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d RETURNING *", tableBrends, request, id)
	fmt.Println(querry)
	if err := brRep.storage.db.QueryRow(querry).Scan(&c.Id, &c.Name, &c.Slug); err != nil {
		return err
	}
	return nil
}
func (brRep *Brandrepository) FilterAllBrands(fil *models.PageRequest) ([]*models.Brand, error) { //срань которую надо переделать(а можно и не переделывать)
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
	query := fmt.Sprintf("Select count(*) FROM %s %s %s", tableBrends, where, sort)
	if err := brRep.storage.db.QueryRow(query).Scan(&fil.TotalRecords); err != nil {
		return nil, err
	}

	query = fmt.Sprintf("Select * FROM %s %s %s LIMIT %d OFFSET %d", tableBrends, where, sort, fil.PageLength, (fil.PageNumber-1)*fil.PageLength)
	where, sort = "", ""

	fmt.Println(query)

	rows, err := brRep.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	brand := make([]*models.Brand, 0)
	for rows.Next() {
		c := models.Brand{}
		err := rows.Scan(&c.Id, &c.Name, &c.Slug)
		if err != nil {
			log.Println(err)
			continue
		}
		brand = append(brand, &c)
	}
	return brand, nil
}
func (brRep *Brandrepository) GetBrandById(id int) (*models.Brand, error) {

	var brand models.Brand
	query := fmt.Sprintf("Select * FROM %s WHERE id = %d", tableBrends, id)
	if err := brRep.storage.db.QueryRow(query).Scan(&brand.Id, &brand.Name, &brand.Slug); err != nil {
		return nil, err
	}
	fmt.Println(query)
	return &brand, nil

}
