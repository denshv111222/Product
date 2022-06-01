package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strings"
)

type Category_productrepository struct {
	storage *Storage
}

var (
	tableCategoties_products string = "categories_products"
)

func (catproRep *Category_productrepository) DeleteCategories_productsById(catprod *models.Categories_products) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE products_id = %d and categories_id = %d", tableCategoties_products, catprod.Product.Id, catprod.Categories.Id)
	if _, err := catproRep.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (catproRep *Category_productrepository) CreateCategories_products(catprod *models.Categories_products) error {
	querry := fmt.Sprintf("INSERT INTO %s VALUES($1, $2,$3)", tableCategoties_products)
	fmt.Println(querry)
	catproRep.storage.db.QueryRow(querry, catprod.Product.Id, catprod.Categories.Id, catprod.Sort)
	return nil
}
func (catproRep *Category_productrepository) FilterAllCategories_products(fil *models.PageRequest) ([]*models.Categories_products, error) { //срань которую надо переделать(а можно и не переделывать)
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
	query := fmt.Sprintf("Select count(*) FROM %s %s %s", tableCategoties_products, where, sort)
	if err := catproRep.storage.db.QueryRow(query).Scan(&fil.TotalRecords); err != nil {
		return nil, err
	}

	request = strings.Join(sortList, ",")
	sort = sort + request
	fmt.Println(request)
	query = fmt.Sprintf("select %s.id_categories,%s.name,%s.slug,%s.parent_id,%s.id,%s.name,%s.slug,%s.sku,%s.short_description,%s.full_description,%s.sort,%s.id,%s.name,%s.slug,%s.sort from %s inner join %s on  %s.categories_id = %s.id_categories inner join %s on %s.products_id = %s.id inner join %s on %s.id = %s.id %s %s LIMIT %d OFFSET %d", tableCategories, tableCategories, tableCategories, tableCategories, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableBrends, tableBrends, tableBrends, tableCategoties_products, tableCategoties_products, tableCategories, tableCategoties_products, tableCategories, tableProduct, tableCategoties_products, tableProduct, tableBrends, tableBrends, tableProduct, where, sort, fil.PageLength, (fil.PageNumber-1)*fil.PageLength)
	where, sort = "", ""

	fmt.Println(query)

	fmt.Print(query)
	rows, err := catproRep.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories_products := make([]*models.Categories_products, 0)
	log.Println(categories_products)
	for rows.Next() {
		brand := models.Brand{}
		product := models.Product{
			Brand: &brand,
		}
		categ := models.Categories{}
		a := models.Categories_products{
			Product:    &product,
			Categories: &categ,
		}
		log.Println(rows)
		err := rows.Scan(&a.Categories.Id, &a.Categories.Name, &a.Categories.Slug, &a.Categories.Parent_id, &a.Product.Id, &a.Product.Name, &a.Product.Slug, &a.Product.SKU, &a.Product.Short_description, &a.Product.Full_description, &a.Product.Sort, &a.Product.Brand.Id, &a.Product.Brand.Name, &a.Product.Brand.Slug, &a.Sort)
		if err != nil {
			log.Println(err)
			continue
		}
		categories_products = append(categories_products, &a)
	}
	return categories_products, nil
}
