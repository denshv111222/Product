package storage

import (
	"GitHab/Standart_Server_API/internal/app/models"
	"fmt"
)

type Attributs_values_productsrepository struct {
	storage *Storage
}

var (
	tableAttributes_values_products string = "atributes_values_products"
)

func (atvalprodRep *Attributs_values_productsrepository) DeleteAttributes_values_productsById(attrvalprod *models.Attributes_values_products) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE products_id = %d and atributes_values_id = %d", tableAttributes_values_products, attrvalprod.Produkt.Id, attrvalprod.Attributes_values.Id)
	if _, err := atvalprodRep.storage.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (atvalprodRep *Attributs_values_productsrepository) CreateAttributes_values_products(attrvalprod *models.Attributes_values_products) error {
	fmt.Println("Запрос")
	querry := fmt.Sprintf("INSERT INTO %s VALUES($1, $2,$3)", tableAttributes_values_products)
	fmt.Println(querry)
	atvalprodRep.storage.db.QueryRow(querry, attrvalprod.Produkt.Id, attrvalprod.Attributes_values.Id, attrvalprod.Sort)
	return nil
}

func (atvalprodRep *Attributs_values_productsrepository) FilterAllAttributes_values_products(fil *models.PageRequest) ([]*models.Attributes_values_products, error) { //срань которую надо переделать(а можно и не переделывать)
	/*	fieldlist := make([]string, 0)
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
		query = fmt.Sprintf("select %s.id,%s.name,%s.id_atribute,%s.name,%s.slug ,%s.id_unit,%s.name,%s.slug,%s.id,%s.name,%s.slug,%s.sku,%s.short_description,%s.full_description,%s.sort,%s.id,%s.name,%s.slug, %s.sort from %s inner join %s on %s.atribute_id = %s.id_atribute inner join %s on %s.unit_id = %s.id_unit inner join %s on %s.atributes_values_id = %s.id inner join %s on %s.products_id = %s.id inner join %s on %s.id = %s.id %s %s LIMIT %d OFFSET %d", tableAttributes_values, tableAttributes_values, tableAttributes, tableAttributes, tableAttributes, tableunits, tableunits, tableunits, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableProduct, tableBrends, tableBrends, tableBrends, tableAttributes_values_products, tableAttributes_values, tableAttributes, tableAttributes_values, tableAttributes, tableunits, tableAttributes, tableunits, tableAttributes_values_products, tableAttributes_values_products, tableAttributes_values, tableProduct, tableAttributes_values_products, tableProduct, tableBrends, tableBrends, tableProduct, where, sort, fil.Pages.CountsRecordOnPage, (fil.Pages.СurrentPage-1)*fil.Pages.CountsRecordOnPage)
		where, sort = "", ""

		fmt.Print(query)
		rows, err := atvalprodRep.storage.db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		unit := models.Units{}
		attr := models.Attributes{
			Units: &unit,
		}
		attr_val := models.Attributes_values{
			Attributes: &attr,
		}
		brand := models.Brand{}
		product := models.Product{
			Brand: &brand,
		}
		a := models.Attributes_values_products{
			Produkt:           &product,
			Attributes_values: &attr_val,
		}
		attribute := make([]*models.Attributes_values_products, 0)
		log.Println(attribute)
		for rows.Next() {
			log.Println(rows)
			err := rows.Scan(&a.Attributes_values.Id, &a.Attributes_values.Name, &a.Attributes_values.Attributes.Id, &a.Attributes_values.Attributes.Name, &a.Attributes_values.Attributes.Slug, &a.Attributes_values.Attributes.Units.Id, &a.Attributes_values.Attributes.Units.Name, &a.Attributes_values.Attributes.Units.Slug, &a.Produkt.Id, &a.Produkt.Name, &a.Produkt.Slug, &a.Produkt.SKU, &a.Produkt.Short_description, &a.Produkt.Full_description, &a.Produkt.Sort, &a.Produkt.Brand.Id, &a.Produkt.Brand.Name, &a.Produkt.Brand.Slug, &a.Sort)
			if err != nil {
				log.Println(err)
				continue
			}
			attribute = append(attribute, &a)
		}*/
	return nil, nil
}
