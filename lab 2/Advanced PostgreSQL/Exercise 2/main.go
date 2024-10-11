package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func establishGORMConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=123 dbname=postgres2 port=5432 sslmode=disable TimeZone=Asia/Almaty"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	fmt.Println("Connected to PostgreSQL with GORM successfully!")
	return db, nil
}

func migrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		return fmt.Errorf("unable to migrate models: %w", err)
	}
	fmt.Println("Tables have been created successfully, including their associations!")
	return nil
}

func addUserWithProfile(db *gorm.DB, user User) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("error inserting user and profile: %w", err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("User and Profile added successfully!")
	return nil
}

func fetchUsersWithProfiles(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Preload("Profile").Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("error fetching users with profiles: %w", err)
	}

	fmt.Printf("Fetched %d users with profiles.\n", len(users))
	return users, nil
}

func modifyUserProfile(db *gorm.DB, userID uint, updatedName string, updatedAge int, updatedBio string, updatedProfilePic string) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&User{}).Where("id = ?", userID).Updates(User{Name: updatedName, Age: updatedAge}).Error; err != nil {
			return fmt.Errorf("error updating user: %w", err)
		}

		if err := tx.Model(&Profile{}).Where("user_id = ?", userID).Updates(Profile{Bio: updatedBio, ProfilePictureURL: updatedProfilePic}).Error; err != nil {
			return fmt.Errorf("error updating profile: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("User and Profile updated successfully!")
	return nil
}

func removeUserWithProfile(db *gorm.DB, userID uint) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&User{}, userID).Error; err != nil {
			return fmt.Errorf("error deleting user: %w", err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("User and related Profile removed successfully!")
	return nil
}

func main() {
	db, err := establishGORMConnection()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	err = migrateModels(db)
	if err != nil {
		log.Fatal("Model migration failed:", err)
	}

	user := User{
		Name: "Anuar",
		Age:  22,
		Profile: Profile{
			Bio:               "QA engineer from Kazakhstan",
			ProfilePictureURL: "https://i.pinimg.com/originals/34/63/34/346334640dd06e074b5b37c1e1263931.jpg",
		},
	}

	err = addUserWithProfile(db, user)
	if err != nil {
		log.Fatal("Failed to add user and profile:", err)
	}

	users, err := fetchUsersWithProfiles(db)
	if err != nil {
		log.Fatal("Error retrieving users:", err)
	}
	fmt.Println("Users with Profiles:", users)

	err = modifyUserProfile(db, users[0].ID, "Anuar Daniyar", 24, "Senior QA Engineer", "https://i.pinimg.com/736x/31/46/08/314608fdfdbeaadaafc7c5f5feadbbc9.jpg")
	if err != nil {
		log.Fatal("Error updating user profile:", err)
	}

}

type User struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `gorm:"size:50;not null"`
	Age     int     `gorm:"not null"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Profile struct {
	ID                uint   `gorm:"primaryKey"`
	UserID            uint   `gorm:"not null;unique"`
	Bio               string `gorm:"size:255"`
	ProfilePictureURL string `gorm:"size:255"`
}
