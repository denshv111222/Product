package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strings"
)

type Products_videosrepository struct {
	storage *Storage
}

var (
	Products_videosTable string = "videos_products"
)

func (prvi *Products_videosrepository) DeleteProduct_VideoById(prodvideo *models.Products_videos) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE products_id = %d and videos_id = %d", Products_videosTable, prodvideo.Product.Id, prodvideo.Videos.Id)
	if _, err := prvi.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (prvi *Products_videosrepository) CreateProduct_video(prodvideo *models.Products_videos) error {
	fmt.Println("Запрос")
	querry := fmt.Sprintf("INSERT INTO %s VALUES($1, $2)", Products_videosTable)
	fmt.Println(querry)
	prvi.storage.db.QueryRow(querry, prodvideo.Product.Id, prodvideo.Videos.Id)
	return nil
}

func (prvi *Products_videosrepository) FilterAllProducts_video(fil *models.Filter) ([]*models.Products_videos, error) { //срань которую надо переделать(а можно и не переделывать)
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
	query := fmt.Sprintf("Select count(*) FROM %s,%s,%s", Products_videosTable, where, sort)
	if err := prvi.storage.db.QueryRow(query).Scan(&fil.Pages.AllRecords); err != nil {
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
	query = fmt.Sprintf("Select %s.id,%s.name,%s.slug,%s.sku,%s.short_description,%s.full_description ,%s.sort,%s.id,%s.name,%s.slug,%s.* from videos_products inner join %s on %s.id = videos_products.products_id inner join %s on %s.id_video = videos_products.videos_id inner join %s on %s.id = products.brand_id %s %s LIMIT %d OFFSET %d", tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableBrends, tableBrends, tableBrends, tableVideo, tableProduct, tableProduct, tableVideo, tableVideo, tableBrends, tableBrends, where, sort, fil.Pages.CountsRecordOnPage, (fil.Pages.СurrentPage-1)*fil.Pages.CountsRecordOnPage)
	where, sort = "", ""

	fmt.Print(query)
	rows, err := prvi.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	product_videos := make([]*models.Products_videos, 0)
	log.Println(product_videos)
	for rows.Next() {
		brand := models.Brand{}
		product := models.Product{
			Brand: &brand,
		}

		var (
			video models.Videos
		)
		a := models.Products_videos{
			Product: &product,
			Videos:  &video,
		}
		log.Println(rows)
		err := rows.Scan(&a.Product.Id, &a.Product.Name, &a.Product.Slug, &a.Product.SKU, &a.Product.Short_description, &a.Product.Full_description, &a.Product.Sort, &a.Product.Brand.Id, &a.Product.Brand.Name, &a.Product.Brand.Slug, &a.Videos.Id, &a.Videos.Name, &a.Videos.Storage, &a.Videos.Path)
		if err != nil {
			log.Println(err)
			continue
		}
		product_videos = append(product_videos, &a)
	}
	return product_videos, nil
}
