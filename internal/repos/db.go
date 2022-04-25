package repos

import (
	"database/sql"
	"github.com/iamkirillnb/avtodom/internal"
	"github.com/iamkirillnb/avtodom/internal/entities"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)


type Postgres struct {
	*sqlx.DB

	Config *internal.DbConfig
}

func NewPostgres(c *internal.DbConfig) *Postgres {
	db, err := sqlx.Connect("postgres", c.DbUrlConnection())
	if err != nil {
		log.Println("connetction to posgtres failed")
		log.Fatal(err)
	}

	return &Postgres{
		DB:     db,
		Config: c,
	}
}

type DbRepo struct {
	*Postgres
}

func NewDbRepo(p *Postgres) *DbRepo {
	return &DbRepo{p}
}


func (d *DbRepo) GetByInnerUrl(innerUrl string) string {
	const qry = `SELECT incomming_url, out_url, code FROM url_redirect WHERE incomming_url=$1;`

	data := &entities.Url{}
	err := d.Get(data, qry, innerUrl)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("get data from DB failed")
			log.Fatal(err)
		} else {
			return innerUrl
		}
	}
	return data.OutUrl
}






