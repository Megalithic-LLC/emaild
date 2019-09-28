package dao

import (
	"github.com/Megalithic-LLC/on-prem-emaild/model"
	"github.com/asdine/genji"
)

func (self MailboxMessagesDAO) CreateTx(tx *genji.Tx, mailboxMessage *model.MailboxMessage) error {
	if table, err := tx.GetTable(model.MailboxMessageTable); err != nil {
		return err
	} else {
		_, err := table.Insert(mailboxMessage)
		return err
	}
}

func (self MailboxMessagesDAO) FindByIDsTx(tx *genji.Tx, mailboxID, messageID string) (*model.MailboxMessage, error) {
	messageTable, err := tx.GetTable(model.MailboxMessageTable)
	if err != nil {
		return nil, err
	}
	searchFor := &model.MailboxMessage{MailboxID: mailboxID, MessageID: messageID}
	mailboxMessagePK, err := searchFor.PrimaryKey()
	if err != nil {
		return nil, err
	}
	r, err := messageTable.GetRecord(mailboxMessagePK)
	if err != nil {
		return nil, err
	}
	var mailboxMessage model.MailboxMessage
	if err := mailboxMessage.ScanRecord(r); err != nil {
		return nil, err
	}
	return &mailboxMessage, nil
}
