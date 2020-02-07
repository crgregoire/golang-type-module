package types

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Importing this for gorm to designate the db driver
	uuid "github.com/satori/go.uuid"
	"github.com/tespo/satya/v2/scoping"
)

//
// Regimen struct for tracking regimen regimen data
//
type Regimen struct {
	ID                            uuid.UUID  `json:"id" gorm:"type:char(36);primary_key" scope:"regimen.id"`
	AccountID                     uuid.UUID  `json:"account_id" gorm:"type:char(36);foreign_key" scope:"regimen.account_id"`
	UserID                        *uuid.UUID `json:"user_id" gorm:"type:char(36);foreign_key" scope:"regimen.user_id"`
	PodID                         *uuid.UUID `json:"pod_id" gorm:"type:char(36);foreign_key" scope:"regimen.pod_id"`
	LastReportedServingsRemaining uint       `json:"last_reported_servings_remaining" scope:"regimen.last_reported_servings_remaining"`
	Reminders                     Reminders  `json:"reminders" scope:"regimen.reminders"`
	Usages                        Usages     `json:"usages" scope:"regimen.usages"`
	Account                       Account    `json:"account" scope:"regimen.account"`
	User                          User       `json:"user" scope:"regimen.user"`
	Pod                           Pod        `json:"pod" scope:"regimen.pod"`
	Meta                          JSON       `json:"meta" gorm:"type:json" scope:"regimen.meta"`
	CreatedAt                     time.Time  `json:"created_at" scope:"regimen.created_at"`
	UpdatedAt                     time.Time  `json:"updated_at" scope:"regimen.updated_at"`
	DeletedAt                     *time.Time `json:"deleted_at"`
}

//
// BeforeCreate will set a UUID rather than a numeric ID
//
func (regimen *Regimen) BeforeCreate(scope *gorm.Scope) error {
	if regimen.UserID != nil {
		if regimen.UserID.String() == "00000000-0000-0000-0000-000000000000" {
			if err := scope.SetColumn("UserID", nil); err != nil {
				return err
			}
		}
	}
	if regimen.PodID != nil {
		if regimen.PodID.String() == "00000000-0000-0000-0000-000000000000" {
			if err := scope.SetColumn("PodID", nil); err != nil {
				return err
			}
		}
	}
	uuid := uuid.NewV4()
	if regimen.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return scope.SetColumn("ID", uuid)
	}
	return nil
}

//
// Regimens is a slice of Regimen
//
type Regimens []Regimen

//
// Get gets all regimens
//
func (regimens *Regimens) Get(db *gorm.DB) error {
	if err := db.Preload("Account").Preload("User").Preload("Pod").Preload("Reminders").Preload("Usages").Find(&regimens).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByID gets one regimen by id
//
func (regimen *Regimen) GetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Preload("Account").Preload("User").Preload("Pod").Preload("Reminders").Preload("Usages").Where("id = ?", id).Find(&regimen).Error; err != nil {
		return err
	}
	return nil
}

//
// GetByQuery gets one regimens by query
//
func (regimens *Regimens) GetByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Preload("Account").Preload("User").Preload("Pod").Preload("Reminders").Preload("Usages").Where(where, query...).Find(&regimens).Error; err != nil {
		return err
	}
	return nil
}

//
// GetOneByQuery gets one regimen by query
//
func (regimen *Regimen) GetOneByQuery(db *gorm.DB, where string, query ...interface{}) error {
	if err := db.Preload("Account").Preload("User").Preload("Pod").Preload("Reminders").Preload("Usages").Where(where, query...).First(&regimen).Error; err != nil {
		return err
	}
	return nil
}

//
// UnscopedGetByID gets one regimen by id
//
func (regimen *Regimen) UnscopedGetByID(db *gorm.DB, id uuid.UUID) error {
	if err := db.Preload("Account").Preload("User").Preload("Pod").Preload("Reminders").Preload("Usages").Unscoped().Where("id = ?", id).Find(&regimen).Error; err != nil {
		return err
	}
	return nil
}

//
// Create makes an Regimen
//
func (regimen *Regimen) Create(db *gorm.DB) error {
	if err := db.Create(&regimen).Scan(&regimen).Error; err != nil {
		return err
	}
	return nil
}

//
// Update makes an Regimen
//
func (regimen *Regimen) Update(db *gorm.DB) error {
	if err := db.Save(&regimen).Error; err != nil {
		return err
	}
	return nil
}

//
// Delete will soft delete an regimen
//
func (regimen *Regimen) Delete(db *gorm.DB, id uuid.UUID) error {
	if err := db.Where("id = ?", id).Delete(&regimen).Error; err != nil {
		return err
	}
	return nil
}

//
// UpdateAccountRegimen updates a Regimen
//
func (regimen *Regimen) UpdateAccountRegimen(db *gorm.DB) error {
	if err := db.Save(&regimen).Error; err != nil {
		return err
	}
	return nil
}

//
// GetAccountRegimens returns the account regimens
//
func (regimens *Regimens) GetAccountRegimens(db *gorm.DB, accountID uuid.UUID) error {
	if err := db.Preload("Account").Preload("User").Preload("Pod").Preload("Reminders").Preload("Usages").Where("account_id = ?", accountID).Find(&regimens).Error; err != nil {
		return err
	}
	return nil
}

//
// GetAccountRegimenByID returns an account regimen by ID
//
func (regimen *Regimen) GetAccountRegimenByID(db *gorm.DB, regimenID, accountID uuid.UUID) error {
	if err := db.Preload("Account").Preload("User").Preload("Pod").Preload("Reminders").Preload("Usages").Find(&regimen, "id = ? and account_id = ?", regimenID, accountID).Error; err != nil {
		return err
	}
	return nil
}

//
// GetUserRegimens returns the user regimens
//
func (regimens *Regimens) GetUserRegimens(db *gorm.DB, userID uuid.UUID) error {
	if err := db.Preload("Account").Preload("User").Preload("Pod").Preload("Reminders").Preload("Usages").Find(&regimens, "user_id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}

//
// GetUserRegimenByID returns a user regimens by ID
//
func (regimen *Regimen) GetUserRegimenByID(db *gorm.DB, regimenID, userID uuid.UUID) error {
	if err := db.Preload("Account").Preload("User").Preload("Pod").Preload("Reminders").Preload("Usages").Find(&regimen, "id = ? and userID_id = ?", regimenID, userID).Error; err != nil {
		return err
	}
	return nil
}

//
// DeleteAccountRegimenByID removes a regimen on an account
//
func (regimen *Regimen) DeleteAccountRegimenByID(db *gorm.DB, regimenID, accountID uuid.UUID) error {
	if err := db.Delete(&regimen, "account_id = ? and id = ?", accountID, regimenID).Error; err != nil {
		return err
	}
	return nil
}

//
// GetReminders returns the reminds associated with a regimen
//
func (regimen *Regimen) GetReminders(db *gorm.DB) error {
	if err := db.Model(&regimen).Association("Reminders").Find(&regimen.Reminders).Error; err != nil {
		return err
	}
	return nil
}

//
// CreateReminder creates a reminder on a regimen
//
func (regimen *Regimen) CreateReminder(db *gorm.DB, reminder Reminder) error {
	if err := db.Model(&regimen).Association("Reminders").Append(reminder).Error; err != nil {
		return err
	}
	return nil
}

//
// UpdateReminder updates a reminder on a regimen
//
func (regimen *Regimen) UpdateReminder(db *gorm.DB, reminder Reminder) error {
	if err := db.Model(&regimen).Association("Reminders").Append(&reminder).Error; err != nil {
		return err
	}
	return nil
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (regimen *Regimen) Scope(scopes []string) {
	*regimen = scoping.FilterByScopes(scopes, *regimen).(Regimen)
}

//
// Scope limits the fields being returned based on the
// passed in scopes
//
func (regimens *Regimens) Scope(scopes []string) {
	regimenSlice := *regimens
	for i, regimen := range regimenSlice {
		regimenSlice[i] = scoping.FilterByScopes(scopes, regimen).(Regimen)
	}
	*regimens = regimenSlice
}
