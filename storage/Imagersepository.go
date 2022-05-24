package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strings"
)

type Imagerepository struct {
	storage *Storage
}

var (
	tableImages string = "images"
)

func (im *Imagerepository) DeleteImage(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id_image = %d ", tableImages, id)
	fmt.Println(query)
	if _, err := im.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (im *Imagerepository) CreateImage(i *models.Images) (*models.Images, error) {
	fmt.Println(i)
	query := fmt.Sprintf("INSERT INTO %s (name,storage,path) VALUES ($1,$2,$3) RETURNING id_image", tableImages)
	if err := im.storage.db.QueryRow(query, i.Name, i.Storage, i.Path).Scan(&i.Id); err != nil {
		return nil, err
	}
	return i, nil
}
func (im *Imagerepository) Update(id int, i *models.Images) error {
	fmt.Println(id)
	fmt.Println(i)
	fieldlist := make([]string, 0)
	fmt.Println(fieldlist)
	if i.Name != "" {
		fieldlist = append(fieldlist, ("name='" + i.Name + "'"))
	}
	fmt.Println(fieldlist)
	if i.Storage != "" {
		fieldlist = append(fieldlist, ("storage='" + i.Storage + "'"))
	}
	fmt.Println(fieldlist)
	if i.Path != "" {
		fieldlist = append(fieldlist, ("path='" + i.Path + "'"))
	}
	fmt.Println(fieldlist)
	request := strings.Join(fieldlist, ", ")
	fmt.Println(request)

	querry := fmt.Sprintf("UPDATE %s SET %s WHERE id_image = %d RETURNING *", tableImages, request, id)
	fmt.Println(querry)
	if err := im.storage.db.QueryRow(querry).Scan(&i.Id, &i.Name, &i.Storage, &i.Path); err != nil {
		return err
	}
	return nil
}
func (im *Imagerepository) FilterAllImages(fil *models.Filter) ([]*models.Images, error) { //срань которую надо переделать(а можно и не переделывать)
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
	query := fmt.Sprintf("Select count(*) FROM %s,%s,%s", tableImages, where, sort)
	if err := im.storage.db.QueryRow(query).Scan(&fil.Pages.AllRecords); err != nil {
		return nil, err
	}
	fil.Pages.AllPages = allPage(fil.Pages.AllRecords, fil.Pages.CountsRecordOnPage)
	fil.Pages.RemainedRecords = fil.Pages.AllRecords - fil.Pages.CountsRecordOnPage*fil.Pages.СurrentPage
	if fil.Pages.RemainedRecords < 0 {
		fil.Pages.RemainedRecords = 0
	}

	request = strings.Join(sortList, ",")
	sort = sort + request
	fmt.Println(request)
	query = fmt.Sprintf("Select * FROM %s %s %s LIMIT %d OFFSET %d", tableImages, where, sort, fil.Pages.CountsRecordOnPage, (fil.Pages.СurrentPage-1)*fil.Pages.CountsRecordOnPage)
	where, sort = "", ""

	fmt.Println(query)

	rows, err := im.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	images := make([]*models.Images, 0)
	for rows.Next() {
		i := models.Images{}
		err := rows.Scan(&i.Id, &i.Name, &i.Storage, &i.Path)
		if err != nil {
			log.Println(err)
			continue
		}
		images = append(images, &i)
	}
	return images, nil
}
