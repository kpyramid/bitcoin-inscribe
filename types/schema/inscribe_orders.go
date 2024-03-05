package schema

import (
	"gorm.io/gorm"
)

const (
	OrderStatusPending    = "Pending"
	OrderStatusInscribing = "Inscribing"

	OrderStatusCommitBroadcast = "CommitBroadcast"
	OrderStatusRevealBroadcast = "RevealBroadcast"

	OrderStatusUnConfirmed = "UnConfirmed"
	OrderStatusSuccessful  = "Successful"

	OrderStatusFailure = "Failure"
	OrderStatusExpired = "Closed"
)

type InscribeOrder struct {
	gorm.Model
	OrderId              string `gorm:"column:order_id;type:varchar(255);not null;unique;" json:"order_id"`
	UserAddress          string `gorm:"column:user_address;type:varchar(255);not null;index;" json:"user_address"`
	ReceiptAddress       string `gorm:"column:receipt_address;type:varchar(255);not null;" json:"receipt_address"`
	ReceiptAddressNumber int64  `gorm:"column:receipt_address_number;type:bigint;not null;unique;" json:"receipt_address_number"`
	TotalAmount          int64  `gorm:"column:total_amount;type:bigint;not null;" json:"total_amount"`
	FeeRate              int64  `gorm:"column:fee_rate;type:bigint;not null;" json:"fee_rate"`
	NetWork              string `gorm:"column:net_work;type:varchar(255);not null;" json:"net_work"`
	TokenId              int64  `gorm:"column:token_id;type:bigint;not null;unique;" json:"token_id"`
	InscriptionId        string `gorm:"column:inscription_id;type:varchar(255);not null;" json:"inscription_id"`
	InscriptionContent   string `gorm:"column:inscription_content;type:text;not null;" json:"inscription_content"`
	CommitData           string `gorm:"column:commit_data;type:text;not null;" json:"commit_data"`
	CommitTxHash         string `gorm:"column:commit_tx_hash;type:varchar(255);not null;" json:"commit_tx_hash"`
	RevealData           string `gorm:"column:reveal_data;type:text;not null;" json:"reveal_data"`
	RevealTxHash         string `gorm:"column:reveal_tx_hash;type:varchar(255);not null;" json:"reveal_tx_hash"`
	Status               string `gorm:"column:status;type:varchar(255);index;not null;" json:"status"`
}
