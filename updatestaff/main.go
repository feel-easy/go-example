package main

import (
	"context"
	"flag"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	offdsn = ""
	ssodsn = ""
)

func init() {
	flag.StringVar(&offdsn, "off", offdsn, "")
	flag.StringVar(&ssodsn, "sso", ssodsn, "")
	flag.Parse()
}

type account struct {
	gorm.Model
	ResourceName string `gorm:"unique"`
	ResourceID   string `gorm:"index"`
}

type staff struct {
	UserID    uint
	AccountID string
}

func main() {
	h := NewHandler()
	if err := h.do(); err != nil {
		panic(err)
	}
}

type handler struct {
	offDB *gorm.DB
	ssoDB *gorm.DB
}

func NewHandler() *handler {
	ctx := context.Background()
	offDB, err := Open(ctx, offdsn)
	if err != nil {
		panic(err)
	}
	ssoDB, err := Open(ctx, ssodsn)
	if err != nil {
		panic(err)
	}
	return &handler{
		offDB: offDB,
		ssoDB: ssoDB,
	}
}

func (h *handler) do() error {
	stffs := make([]*staff, 0)
	if err := h.offDB.Find(&stffs).Error; err != nil {
		return err
	}
	fmt.Printf("len stffs: %d \n", len(stffs))
	userIDs := make([]uint, 0, len(stffs))
	for _, stff := range stffs {
		userIDs = append(userIDs, stff.UserID)
	}
	accounts := make([]*account, 0)
	if err := h.ssoDB.Where("id in ?", userIDs).Find(&accounts).Error; err != nil {
		return err
	}
	fmt.Printf("len accounts: %d\n", len(accounts))
	actMap := make(map[uint]*account, len(accounts))
	for _, act := range accounts {
		actMap[act.ID] = act
	}
	for _, stff := range stffs {
		if act, ok := actMap[stff.UserID]; ok {
			if err := h.offDB.Model(stff).Where("user_id= ?", stff.UserID).Update("account_id", act.ResourceID).Error; err != nil {
				return err
			}
			fmt.Printf("user_id: %d, account_id: %s \n", stff.UserID, act.ResourceID)
		} else {
			fmt.Printf("UserID:%d not exists \n", stff.UserID)
		}

	}
	return nil
}

func Open(ctx context.Context, dataSourceName string) (*gorm.DB, error) {
	d, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return d, nil
}
