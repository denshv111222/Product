package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // для того, чотбы отработала функция init()
)

//Instance of storage
type Storage struct {
	config                              *Config
	db                                  *sql.DB
	imageRepository                     *Imagerepository
	videoRepository                     *Videorepository
	unitRepository                      *Unitrepository
	categoryRepository                  *Categoryrepository
	brandRepository                     *Brandrepository
	productRepository                   *Productrepository
	products_imagesRepository           *Products_imagesrepository
	products_videosRepository           *Products_videosrepository
	attributsRepository                 *Attributsrepository
	attributs_valuesRepository          *Attributs_valuesrepository
	attributs_values_productsRepository *Attributs_values_productsrepository
	category_productRepository          *Category_productrepository
}

//Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

//Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfully!")
	return nil
}

//Close connection
func (storage *Storage) Close() {
	storage.db.Close()
}

func (s *Storage) Image() *Imagerepository {
	if s.imageRepository != nil {
		return s.imageRepository
	}
	s.imageRepository = &Imagerepository{
		storage: s,
	}
	return s.imageRepository
}

func (s *Storage) Video() *Videorepository {
	if s.videoRepository != nil {
		return s.videoRepository
	}
	s.videoRepository = &Videorepository{
		storage: s,
	}
	return s.videoRepository
}

func (s *Storage) Unit() *Unitrepository {
	if s.unitRepository != nil {
		return s.unitRepository
	}
	s.unitRepository = &Unitrepository{
		storage: s,
	}
	return s.unitRepository
}

func (s *Storage) Category() *Categoryrepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}
	s.categoryRepository = &Categoryrepository{
		storage: s,
	}
	return s.categoryRepository
}

func (s *Storage) Brand() *Brandrepository {
	if s.brandRepository != nil {
		return s.brandRepository
	}
	s.brandRepository = &Brandrepository{
		storage: s,
	}
	return s.brandRepository
}

func (s *Storage) Product() *Productrepository {
	if s.productRepository != nil {
		return s.productRepository
	}
	s.productRepository = &Productrepository{
		storage: s,
	}
	return s.productRepository
}

func (s *Storage) Products_images() *Products_imagesrepository {
	if s.unitRepository != nil {
		return s.products_imagesRepository
	}
	s.products_imagesRepository = &Products_imagesrepository{
		storage: s,
	}
	return s.products_imagesRepository
}

func (s *Storage) Products_videos() *Products_videosrepository {
	if s.products_videosRepository != nil {
		return s.products_videosRepository
	}
	s.products_videosRepository = &Products_videosrepository{
		storage: s,
	}
	return s.products_videosRepository
}

func (s *Storage) Attributes() *Attributsrepository {
	if s.attributsRepository != nil {
		return s.attributsRepository
	}
	s.attributsRepository = &Attributsrepository{
		storage: s,
	}
	return s.attributsRepository
}
func (s *Storage) Attributes_values() *Attributs_valuesrepository {
	if s.attributs_valuesRepository != nil {
		return s.attributs_valuesRepository
	}
	s.attributs_valuesRepository = &Attributs_valuesrepository{
		storage: s,
	}
	return s.attributs_valuesRepository
}

func (s *Storage) Attributes_values_products() *Attributs_values_productsrepository {
	if s.attributs_values_productsRepository != nil {
		return s.attributs_values_productsRepository
	}
	s.attributs_values_productsRepository = &Attributs_values_productsrepository{
		storage: s,
	}
	return s.attributs_values_productsRepository
}

func (s *Storage) Categories_products() *Category_productrepository {
	if s.category_productRepository != nil {
		return s.category_productRepository
	}
	s.category_productRepository = &Category_productrepository{
		storage: s,
	}
	return s.category_productRepository
}
