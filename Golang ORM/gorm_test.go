package Golang_ORM

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	sqlDb.SetConnMaxIdleTime(time.Minute * 10)

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteRawSQL(t *testing.T) {
	err := db.Exec("insert into samples(id, name) values(?, ?)", "1", "Gayuh").Error
	assert.Nil(t, err)

	err = db.Exec("insert into samples(id, name) values(?, ?)", "2", "Raharjo").Error
	assert.Nil(t, err)
}

type Samples struct {
	Id   string
	Name string
}

func TestSelectRawSQL(t *testing.T) {
	var sample Samples
	err := db.Raw("select id, name from samples where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "Gayuh", sample.Name)

	var samples []Samples
	err = db.Raw("select id, name from samples").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(samples))
	fmt.Println(samples)
}

func TestSQLRows(t *testing.T) {
	rows, err := db.Raw("select id, name from samples").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Samples
	for rows.Next() {
		var sample Samples
		err = rows.Scan(&sample.Id, &sample.Name)
		assert.Nil(t, err)

		samples = append(samples, sample)
	}

	assert.Equal(t, 2, len(samples))
	fmt.Println(samples)
}

func TestSQLScanRows(t *testing.T) {
	rows, err := db.Raw("select id, name from samples").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Samples
	for rows.Next() {
		err = db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}

	assert.Equal(t, 2, len(samples))
	fmt.Println(samples)
}

func TestCreateUser(t *testing.T) {
	user := User{
		ID:       "2",
		Password: "rahasia",
		Name: Name{
			FirstName:  "Ahmad",
			MiddleName: "Solikhin",
			LastName:   "Gayuh",
		},
		Information: "Test bang yak",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 0; i < 10; i++ {
		users = append(users, User{
			ID:       strconv.Itoa(i),
			Password: "rahasia",
			Name: Name{
				FirstName: "user " + strconv.Itoa(i),
			},
			Information: "Test bang",
		})
	}

	result := db.Create(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(10), result.RowsAffected)
}

func TestTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID:       "99",
			Password: "rahasia",
			Name: Name{
				FirstName:  "Ahmad",
				MiddleName: "Gayuh",
				LastName:   "Raharjo",
			},
			Information: "Test transaction",
		}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID:       "88",
			Password: "rahasia",
			Name: Name{
				FirstName:  "Ahmad",
				MiddleName: "Gayuh",
				LastName:   "Raharjo",
			},
			Information: "Test transaction",
		}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID:       "77",
			Password: "rahasia",
			Name: Name{
				FirstName:  "Ahmad",
				MiddleName: "Gayuh",
				LastName:   "Raharjo",
			},
			Information: "Test transaction",
		}).Error

		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionError(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID:       "21",
			Password: "rahasia",
			Name: Name{
				FirstName:  "Ahmad",
				MiddleName: "Gayuh",
				LastName:   "Raharjo",
			},
			Information: "Test transaction",
		}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID:       "88",
			Password: "rahasia",
			Name: Name{
				FirstName:  "Ahmad",
				MiddleName: "Gayuh",
				LastName:   "Raharjo",
			},
			Information: "Test transaction",
		}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID:       "77",
			Password: "rahasia",
			Name: Name{
				FirstName:  "Ahmad",
				MiddleName: "Gayuh",
				LastName:   "Raharjo",
			},
			Information: "Test transaction",
		}).Error

		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestManualTransaction(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{
		ID:       "15",
		Password: "rahasia",
		Name: Name{
			FirstName: "Ahmad",
			LastName:  "Raharjo",
		},
		Information: "Haiyaa",
	}).Error
	if err != nil {
		panic(err)
	}

	err = tx.Create(&User{
		ID:       "16",
		Password: "rahasia",
		Name: Name{
			FirstName: "Ahmad",
			LastName:  "Raharjo",
		},
		Information: "Haiyaa",
	}).Error
	if err != nil {
		panic(err)
	}

	tx.Commit()
}

func TestManualTransactionError(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{
		ID:       "17",
		Password: "rahasia",
		Name: Name{
			FirstName: "Ahmad",
			LastName:  "Raharjo",
		},
		Information: "Haiyaa",
	}).Error
	if err != nil {
		panic(err)
	}

	err = tx.Create(&User{
		ID:       "16",
		Password: "rahasia",
		Name: Name{
			FirstName: "Ahmad",
			LastName:  "Raharjo",
		},
		Information: "Haiyaa",
	}).Error
	if err != nil {
		panic(err)
	}

	tx.Commit()
}

func TestQuerySingleObject(t *testing.T) {
	user := User{}
	result := db.First(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, user.ID, "0")

	user = User{}
	result = db.Last(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, user.ID, "99")
}

func TestQueryInlineCondition(t *testing.T) {
	user := User{}
	result := db.Take(&user, "id = ?", "5")
	assert.Nil(t, result.Error)
	assert.Equal(t, user.ID, "5")
}

func TestQueryAllObjects(t *testing.T) {
	var users []User
	result := db.Find(&users, "id in ?", []string{"0", "1", "2"})
	assert.Nil(t, result.Error)
	assert.Equal(t, len(users), 3)
}

func TestWhereQuery(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%Ahmad%").Where("password = ?", "rahasia").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, len(users), 7)
}

func TestWhereOrdQuery(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%Ahmad%").Or("password = ?", "rahasia").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, len(users), 17)
}

func TestWhereNotQuery(t *testing.T) {
	var users []User
	result := db.Not("first_name like ?", "%Ahmad%").Or("password = ?", "rahasia").Find(&users)
	assert.Nil(t, result.Error)
}

func TestSelectField(t *testing.T) {
	var users []User
	result := db.Select("id", "first_name").Where("id = ?", "0").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, len(users), 1)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
		assert.Equal(t, "", user.Name.LastName)
	}
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "Ahmad",
		},
		Password: "rahasia",
	}

	var users []User

	result := db.Where(userCondition).Find(&users)
	assert.Nil(t, result.Error)
}

func TestMapCondition(t *testing.T) {
	userCondition := map[string]interface{}{
		"last_name": "",
	}

	var users []User
	result := db.Where(userCondition).Find(&users)
	assert.Nil(t, result.Error)
}

func TestOrderLimitOffset(t *testing.T) {
	var users []User
	result := db.Order("id asc, first_name desc").Limit(5).Offset(5).Find(&users)
	assert.Nil(t, result.Error)
}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	result := db.Model(&User{}).Where("id = ?", "1").Find(&users)
	assert.Nil(t, result.Error)
}

func TestDataUpdt(t *testing.T) {
	user := User{}
	result := db.Where("id = ?", "1").Take(&user)
	assert.Nil(t, result.Error)

	user.Name.MiddleName = "Huwalahumba"
	user.Name.LastName = "Haiyaaa"
	user.Password = "rahasia123"

	result = db.Save(&user)
	assert.Nil(t, result.Error)
}

func TestSelectedColumns(t *testing.T) {
	result := db.Model(&User{}).Where("id = ?", "0").Updates(map[string]interface{}{
		"last_name": "Haiyaaa",
	})
	assert.Nil(t, result.Error)

	result = db.Model(&User{}).Where("id = ?", "0").Update("password", "ada deh")
	assert.Nil(t, result.Error)

	result = db.Model(&User{}).Where("id = ?", "0").Updates(User{
		Name: Name{
			MiddleName: "Bejir",
		},
	})
	assert.Nil(t, result.Error)
}

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserId: strconv.Itoa(i),
			Action: "New",
		}
		result := db.Create(&userLog)
		assert.Nil(t, result.Error)

		assert.NotEqual(t, userLog.ID, 0)

		fmt.Println(userLog)
	}
}

func TestUpsert(t *testing.T) {
	userLog := UserLog{
		UserId: "1",
		Action: "New",
	}

	result := db.Save(&userLog)
	assert.Nil(t, result.Error)

	userLog.UserId = "2"
	result = db.Save(&userLog)
	assert.Nil(t, result.Error)
}

func TestUpsertIdNotAutoIncrement(t *testing.T) {
	user := User{
		ID: "999",
		Name: Name{
			FirstName: "Haiyaa Looo",
		},
	}

	result := db.Save(&user)
	assert.Nil(t, result.Error)

	user.Name.MiddleName = "Huwalahumba"
	result = db.Save(&user)
	assert.Nil(t, result.Error)
}

func TestConflict(t *testing.T) {
	user := User{
		ID: "999",
		Name: Name{
			FirstName: "Gayuh",
		},
	}

	result := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user)
	assert.Nil(t, result.Error)
}

func TestDelete(t *testing.T) {
	var user User

	result := db.First(&user, "id = ?", "1")
	assert.Nil(t, result.Error)
	result = db.Delete(&user)
	assert.Nil(t, result.Error)
	fmt.Println(user)

	result = db.Delete(&User{}, "id = ?", "999")
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", "99").Delete(&User{})
	assert.Nil(t, result.Error)
}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserId:      "1",
		Title:       "Test",
		Description: "Test",
	}

	result := db.Create(&todo)
	assert.Nil(t, result.Error)

	result = db.Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	result = db.Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, len(todos), 0)
}

func TestUnscoped(t *testing.T) {
	var todos []Todo
	result := db.Unscoped().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, len(todos), 1)

	result = db.Unscoped().Delete(&todos)
	assert.Nil(t, result.Error)
}

func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "2")
		if result.Error != nil {
			return result.Error
		}

		user.Name.FirstName = "Solikhin"
		result = tx.Save(&user)
		return result.Error
	})
	assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserId:  "2",
		Balance: 1000000,
	}

	tx := db.Create(&wallet)
	assert.Nil(t, tx.Error)
}

func TestRetrieveEagerRelation(t *testing.T) {
	var users []User
	tx := db.Model(&User{}).Preload("Wallet").Find(&users, "id in ?", []string{"13", "2"})

	assert.Nil(t, tx.Error)
}

func TestRetrieveJoinRelation(t *testing.T) {
	var user User
	tx := db.Model(&user).Joins("Wallet").Take(&user, "users.id = ?", "2")
	assert.Nil(t, tx.Error)
	assert.NotNil(t, user.Wallet)
}

func TestAutoUpsert(t *testing.T) {
	user := User{
		ID:       "100",
		Password: "123456",
		Name: Name{
			FirstName: "Ahmad Solikhin",
		},
		Wallet: Wallet{
			ID:      "3",
			Balance: 10000,
		},
	}

	result := db.Create(&user)
	assert.Nil(t, result.Error)
}

func TestSkipAutoUpsert(t *testing.T) {
	user := User{
		ID:       "101",
		Password: "123456",
		Name: Name{
			FirstName: "Ahmad Solikhin",
		},
		Wallet: Wallet{
			ID:      "3",
			Balance: 10000,
		},
	}

	result := db.Omit(clause.Associations).Create(&user)
	assert.Nil(t, result.Error)
}

func TestUserAndAddress(t *testing.T) {
	user := User{
		ID:       "103",
		Password: "123456",
		Name: Name{
			FirstName: "Ahmad Solikhin Gayuh",
		},
		Addresses: []Address{
			{
				Address: "Cibitung",
			},
			{
				Address: "Bekasi",
			},
		},
	}

	tx := db.Create(&user)
	assert.Nil(t, tx.Error)
}

func TestJoinOneToMany(t *testing.T) {
	user := User{}
	tx := db.Model(&user).Preload("Addresses").Take(&user, "users.id = ?", "102")
	assert.Nil(t, tx.Error)

	fmt.Println(user)
}

func TestBelongsTo(t *testing.T) {
	var addresses []Address
	tx := db.Model(&Address{}).Preload("User").Find(&addresses)
	assert.Nil(t, tx.Error)
	for _, address := range addresses {
		fmt.Println(address)
	}

	addresses = []Address{}
	tx = db.Model(&Address{}).Joins("User").Find(&addresses)
	for _, address := range addresses {
		fmt.Println(address)
	}
}

func TestCreateManyToMany(t *testing.T) {
	product := Product{
		ID:    "P001",
		Name:  "Product 1",
		Price: 1000000,
	}

	tx := db.Save(&product)
	assert.Nil(t, tx.Error)

	tx = db.Table("user_like_product").Create(map[string]interface{}{
		"product_id": product.ID,
		"user_id":    "2",
	})
	assert.Nil(t, tx.Error)
}

func TestPreloadManyToMany(t *testing.T) {
	product := Product{}
	tx := db.Preload("LikedByUsers").Take(&product, "products.id = ?", "P001")
	assert.Nil(t, tx.Error)
	assert.Equal(t, len(product.LikedByUsers), 1)
}

func TestPreloadManyToManyUser(t *testing.T) {
	user := User{}
	tx := db.Preload("LikedProducts").Take(&user, "users.id = ?", "2")
	assert.Nil(t, tx.Error)
	assert.Equal(t, len(user.LikedProducts), 1)
}

func TestAssociationFind(t *testing.T) {
	product := Product{}
	tx := db.Take(&product, "products.id = ?", "P001")
	assert.Nil(t, tx.Error)

	var users []User
	err := db.Model(&product).Where("first_name like ?", "Solikhin%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
}

func TestAssociationAppend(t *testing.T) {
	product := Product{}
	tx := db.Take(&product, "products.id = ?", "P001")
	assert.Nil(t, tx.Error)

	user := User{}
	tx = db.Take(&user, "users.id = ?", "3")
	assert.Nil(t, tx.Error)

	err := db.Model(&user).Association("LikedProducts").Append(&product)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		user := User{}
		result := tx.Take(&user, "users.id = ?", "100")
		if result.Error != nil {
			return result.Error
		}

		wallet := Wallet{
			ID:      "4",
			Balance: 10000,
		}

		return tx.Model(&user).Association("Wallet").Replace(&wallet)
	})

	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	product := Product{}
	tx := db.Take(&product, "products.id = ?", "P001")
	assert.Nil(t, tx.Error)

	user := User{}
	tx = db.Take(&user, "users.id = ?", "3")
	assert.Nil(t, tx.Error)

	err := db.Model(&user).Association("LikedProducts").Delete(&product)
	assert.Nil(t, err)
}

func TestAssociationClear(t *testing.T) {
	product := Product{}
	tx := db.Take(&product, "products.id = ?", "P001")
	assert.Nil(t, tx.Error)

	err := db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}

func TestPreloadWithCondition(t *testing.T) {
	user := User{}
	tx := db.Preload("Wallet", "balance > ?", 100000).Take(&user, "users.id = ?", "2")
	assert.Nil(t, tx.Error)
}

func TestNestedPreload(t *testing.T) {
	wallet := Wallet{}
	tx := db.Preload("User.Addresses").Take(&wallet, "wallets.id = ?", "1")
	assert.Nil(t, tx.Error)

	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}

func TestPreloadAll(t *testing.T) {
	user := User{}
	tx := db.Preload(clause.Associations).Take(&user, "users.id = ?", "2")
	assert.Nil(t, tx.Error)
}

func TestJoinQuery(t *testing.T) {
	var users []User
	tx := db.Joins("join wallets on wallets.user_id = users.id and wallets.balance > 100000").Find(&users)
	assert.Nil(t, tx.Error)

	users = []User{}
	tx = db.Joins("Wallet").Find(&users)
	assert.Nil(t, tx.Error)
}

func TestCount(t *testing.T) {
	var count int64
	tx := db.Model(&User{}).Count(&count)
	assert.Nil(t, tx.Error)
	fmt.Println("Count: ", count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	var result AggregationResult
	tx := db.Model(&Wallet{}).Select("sum(balance) as total_balance, min(balance) as min_balance, MAX(balance) as max_balance, avg(balance) as avg_balance").Take(&result)
	assert.Nil(t, tx.Error)
	fmt.Println(result)
}

func TestGroupByAndHaving(t *testing.T) {
	var results []AggregationResult
	tx := db.Model(&Wallet{}).Select("sum(balance) as total_balance, min(balance) as min_balance, MAX(balance) as max_balance, avg(balance) as avg_balance").
		Joins("User").Group("User.id").Having("sum(balance) > ?", 100000).
		Find(&results)
	assert.Nil(t, tx.Error)
	fmt.Println(results)
}

func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 100000)
}

func TestScopes(t *testing.T) {
	var wallets []Wallet
	tx := db.Scopes(BrokeWalletBalance).Find(&wallets)
	assert.Nil(t, tx.Error)

	wallets = []Wallet{}
	tx = db.Scopes(SultanWalletBalance).Find(&wallets)
	assert.Nil(t, tx.Error)

	wallets = []Wallet{}
	tx = db.Scopes(BrokeWalletBalance, SultanWalletBalance).Find(&wallets)
	assert.Nil(t, tx.Error)
}

func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)
}

func TestHook(t *testing.T) {
	user := User{
		Password: "123456",
		Name: Name{
			FirstName: "User Hook",
		},
	}

	tx := db.Create(&user)
	assert.Nil(t, tx.Error)
	assert.NotEqual(t, user.ID, "")
}
