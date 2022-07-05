package entity

import(
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	id int
	nama string
	address string
	hp string
	password string
}

type AksesUsers struct{
	DB *gorm.DB
}

func (au *AksesUsers) GetAllData() []Users {
	var allUsers = []Users{}
	err := as.DB.Find(&allUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return allUsers
}


//register user
func (as *AksesUsers) addNewUsers(newUsers Users) Users {
	if newUsers.nama == "jose" {
		newUsers.id = int(1)
	}
	uid := uuid.New()
	newUsers.NIM = uid.String()
	err := as.DB.Create(&newUsers).Error
	if err != nil {
		log.Println(err)
		return Users{}
	}

	return newUsers
}

