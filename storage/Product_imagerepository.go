package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strings"
)

type Products_imagesrepository struct {
	storage *Storage
}

var (
	Products_imagesTable string = "images_products"
)

func (prim *Products_imagesrepository) DeleteProduct_imageById(prodimage *models.Products_images) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE products_id = %d and images_id = %d", Products_imagesTable, prodimage.Product.Id, prodimage.Image.Id)
	if _, err := prim.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (prim *Products_imagesrepository) CreateProduct_image(prodimage *models.Products_images) error {
	fmt.Println("Запрос")
	querry := fmt.Sprintf("INSERT INTO %s VALUES($1, $2)", Products_imagesTable)
	fmt.Println(querry)
	prim.storage.db.QueryRow(querry, prodimage.Product.Id, prodimage.Image.Id)
	return nil
}
func (prim *Products_imagesrepository) FilterAllProducts_images(fil *models.Filter) ([]*models.Products_images, error) { //срань которую надо переделать(а можно и не переделывать)
	fieldlist := make([]string, 0)
	sortList := make([]string, 0)

	//
	if len(*fil.Fields) != 0 {
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
	if len(*fil.Sorts) != 0 {
		sort = "order by "
		for _, sorts := range *fil.Sorts {
			if sorts.Sort != "" {
				sortList = append(sortList, sorts.Sort+" "+sorts.SortView)
			} else {
				sort = ""
			}
		}
	}
	query := fmt.Sprintf("Select count(*) FROM %s,%s,%s", Products_imagesTable, where, sort)
	if err := prim.storage.db.QueryRow(query).Scan(&fil.Pages.AllRecords); err != nil {
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
	query = fmt.Sprintf("Select %s.id,%s.name,%s.slug,%s.sku,%s.short_description,%s.full_description ,%s.sort,%s.id,%s.name,%s.slug,%s.* from images_products inner join %s on %s.id = images_products.products_id inner join %s on %s.id_image = images_products.images_id inner join %s on %s.id = products.brand_id %s %s LIMIT %d OFFSET %d", tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableBrends, tableBrends, tableBrends, tableImages, tableProduct, tableProduct, tableImages, tableImages, tableBrends, tableBrends, where, sort, fil.Pages.CountsRecordOnPage, (fil.Pages.СurrentPage-1)*fil.Pages.CountsRecordOnPage)
	where, sort = "", ""

	fmt.Print(query)
	rows, err := prim.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	product_image := make([]*models.Products_images, 0)
	log.Println(product_image)
	for rows.Next() {
		brand := models.Brand{}
		product := models.Product{
			Brand: &brand,
		}

		var (
			image models.Images
		)
		a := models.Products_images{
			Product: &product,
			Image:   &image,
		}
		log.Println(rows)
		err := rows.Scan(&a.Product.Id, &a.Product.Name, &a.Product.Slug, &a.Product.SKU, &a.Product.Short_description, &a.Product.Full_description, &a.Product.Sort, &a.Product.Brand.Id, &a.Product.Brand.Name, &a.Product.Brand.Slug, &a.Image.Id, &a.Image.Name, &a.Image.Storage, &a.Image.Path)
		if err != nil {
			log.Println(err)
			continue
		}
		product_image = append(product_image, &a)
	}
	return product_image, nil
}