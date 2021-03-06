package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
	vaild "github.com/asaskevich/govalidator"
	"log"
	"strings"
)

type Videorepository struct {
	storage *Storage
}

var (
	tableVideo string = "videos"
)

func (vi *Videorepository) DeleteVideo(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id_unit = %d ", tableVideo, id)
	fmt.Println(query)
	if _, err := vi.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}
func (vi *Videorepository) CreateVideo(v *models.Videos) (*models.Videos, error) {
	query := fmt.Sprintf("INSERT INTO %s (name,storage,path) VALUES ($1,$2,$3) RETURNING id_video", tableVideo)
	fmt.Println(v.Storage)
	fmt.Println(v.Path)
	fmt.Println(v.Name)
	if err := vi.storage.db.QueryRow(query, v.Name, v.Storage, v.Path).Scan(&v.Id); err != nil {
		fmt.Println(query)
		return nil, err
	}
	fmt.Println(query)
	return v, nil
}
func (vi *Videorepository) Update(id int, v *models.Videos) error {
	fmt.Println(id)
	fmt.Println(v)
	fieldlist := make([]string, 0)
	if v.Name != "" {
		fieldlist = append(fieldlist, ("name='" + v.Name + "'"))
	}
	if v.Storage != "" {
		fieldlist = append(fieldlist, ("storage='" + v.Storage + "'"))
	}
	if v.Path != "" {
		fieldlist = append(fieldlist, ("path='" + v.Path + "'"))
	}
	fmt.Println(fieldlist)
	request := strings.Join(fieldlist, ", ")
	fmt.Println(request)

	querry := fmt.Sprintf("UPDATE %s SET %s WHERE id_video = %d", tableVideo, request, id)
	fmt.Println(querry)
	if err := vi.storage.db.QueryRow(querry).Scan(&v.Id, &v.Name, &v.Storage, &v.Path); err != nil {
		return err
	}
	return nil
}
func (vi *Videorepository) FilterAllVideos(fil *models.PageRequest) ([]*models.Videos, error) {
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
	query := fmt.Sprintf("Select count(*) FROM %s %s %s", tableVideo, where, sort)
	if err := vi.storage.db.QueryRow(query).Scan(&fil.TotalRecords); err != nil {
		return nil, err
	}
	request = strings.Join(sortList, ",")
	sort = sort + request
	fmt.Println(request)
	query = fmt.Sprintf("Select * FROM %s %s %s LIMIT %d OFFSET %d", tableVideo, where, sort, fil.PageLength, (fil.PageNumber-1)*fil.PageLength)
	where, sort = "", ""

	fmt.Println(query)

	rows, err := vi.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	videos := make([]*models.Videos, 0)
	for rows.Next() {
		i := models.Videos{}
		err := rows.Scan(&i.Id, &i.Name, &i.Storage, &i.Path)
		if err != nil {
			log.Println(err)
			continue
		}
		videos = append(videos, &i)
	}
	return videos, nil
}
