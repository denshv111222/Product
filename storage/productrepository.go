package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strconv"
	"strings"
)

type Productrepository struct {
	storage *Storage
}

var (
	tableProduct string = "products"
)

func (prRep *Productrepository) DeleteProducts(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d ", tableProduct, id)
	fmt.Println(query)
	if _, err := prRep.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (prRep *Productrepository) CreateProduct(p *models.Product) error {
	fmt.Println("Запрос")
	fmt.Println(p)
	querry := fmt.Sprintf("INSERT INTO %s (name, slug, sku, short_description, full_description,sort, brand_id) VALUES($1, $2, $3, $4, $5, $6,$7) RETURNING id", tableProduct)
	fmt.Println(querry)
	if err := prRep.storage.db.QueryRow(querry, p.Name, p.Slug, p.SKU, p.Short_description, p.Full_description, p.Sort, p.Brand.Id).Scan(&p.Id); err != nil {

		fmt.Println(err)
		return err
	}

	return nil
}
func (prRep *Productrepository) FilterAllProducts(fil *models.Filter) ([]*models.Product, error) { //срань которую надо переделать(а можно и не переделывать)
	fieldlist := make([]string, 0)
	sortList := make([]string, 0)

	if len(*fil.Fields) != 0 {
		for _, filters := range *fil.Fields {
			if filters.Value != "" {
				where = "where "
				if vaild.IsInt(filters.Value) == true {
					fieldlist = append(fieldlist, filters.Field+filters.Operations+filters.Value)
				} else {
					fieldlist = append(fieldlist, tableProduct+"."+filters.Field+" like "+"'%"+filters.Value+"%'")
				}
			} else {
				where = ""
			}
		}
	}
	request := strings.Join(fieldlist, " and ")
	where = where + request
	request = ""
	if len(*fil.Sorts) != 0 {
		sort = "order by "
		for _, sorts := range *fil.Sorts {
			if sorts.Sort != "" {
				sortList = append(sortList, tableProduct+"."+sorts.Sort+" "+sorts.SortView)
			} else {
				sort = ""
			}
		}
	}
	query := fmt.Sprintf("Select count(*) FROM %s %s %s", tableProduct, where, sort)
	fmt.Println(query)
	if err := prRep.storage.db.QueryRow(query).Scan(&fil.Pages.AllRecords); err != nil {
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
	query = fmt.Sprintf("Select %s.id,%s.name,%s.slug,%s.sku,%s.short_description,brands.name from brands join %s on %s.brand_id = brands.id %s %s LIMIT %d OFFSET %d", tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, where, sort, fil.Pages.CountsRecordOnPage, (fil.Pages.СurrentPage-1)*fil.Pages.CountsRecordOnPage)
	where, sort = "", ""

	fmt.Print(query)
	rows, err := prRep.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	product := make([]*models.Product, 0)
	log.Println(product)
	for rows.Next() {
		var (
			brand models.Brand
		)
		a := models.Product{
			Brand: &brand,
		}
		log.Println(rows)
		err := rows.Scan(&a.Id, &a.Name, &a.Slug, &a.SKU, &a.Short_description, &a.Brand.Name)
		if err != nil {
			log.Println(err)
			continue
		}
		product = append(product, &a)
	}
	return product, nil
}
func (prRep *Productrepository) UpdateProduct(id int, p *models.Product) error {
	fmt.Println(id)
	fieldlist := make([]string, 0)
	fmt.Println(fieldlist)
	if p.Name != "" {
		fieldlist = append(fieldlist, ("name='" + p.Name + "'"))
	}
	fmt.Println(fieldlist)
	if p.Slug != "" {
		fieldlist = append(fieldlist, ("slug='" + p.Slug + "'"))
	}
	fmt.Println(fieldlist)
	if p.Brand.Id != 0 {
		fieldlist = append(fieldlist, ("brand_id='" + strconv.Itoa(p.Brand.Id) + "'"))
	}
	fmt.Println(fieldlist)
	if p.SKU != "" {
		fieldlist = append(fieldlist, ("SKU='" + p.SKU + "'"))
	}
	fmt.Println(fieldlist)
	if p.Short_description != "" {
		fieldlist = append(fieldlist, ("Short_description='" + p.Short_description + "'"))
	}
	fmt.Println(fieldlist)
	if p.Full_description != "" {
		fieldlist = append(fieldlist, ("Full_description='" + p.Full_description + "'"))
	}
	fmt.Println(fieldlist)

	if p.Sort != 0 {
		fieldlist = append(fieldlist, ("sort='" + strconv.Itoa(p.Sort) + "'"))
	}
	fmt.Println(fieldlist)

	request := strings.Join(fieldlist, ", ")
	fmt.Println(request)

	querry := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d RETURNING *", tableProduct, request, id)
	fmt.Println(querry)
	if err := prRep.storage.db.QueryRow(querry).Scan(&p.Id, &p.Name, &p.Slug, &p.Brand.Id, &p.SKU, &p.Short_description, &p.Full_description, &p.Sort); err != nil {
		return err
	}
	return nil
}
func (prRep *Productrepository) GetProductById(id int) (*models.Product, error) {
	var brand models.Brand
	prod := models.Product{
		Brand: &brand,
	}
	query := fmt.Sprintf("Select %s.id,%s.name,%s.slug,%s.sku,%s.short_description,%s.full_description ,%s.sort,brands.id,brands.name,brands.slug from brands join %s on %s.brand_id = brands.id where products.id = %d", tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, id)
	if err := prRep.storage.db.QueryRow(query).Scan(&prod.Id, &prod.Name, &prod.Slug, &prod.SKU, &prod.Short_description, &prod.Full_description, &prod.Sort, &prod.Brand.Id, &prod.Brand.Name, &prod.Brand.Slug); err != nil {
		return nil, err
	}
	fmt.Println(query)
	return &prod, nil
}
