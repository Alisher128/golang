package useCase

import (
	"Projects/day/services/contact/internal/domain"
	"Projects/day/services/contact/internal/repository"
	"fmt"
)

type ContactManager struct {
	Contacts []domain.Contact
	Groups   []domain.Group
}

type UserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{}
}

func (cm *ContactManager) CreateContact(contact domain.Contact) {
	cm.Contacts = append(cm.Contacts, contact)
}

func (cm *ContactManager) ReadContact(contactID int) (domain.Contact, error) {
	for _, contact := range cm.Contacts {
		if contact.ID == contactID {
			return contact, nil
		}
	}
	return domain.Contact{}, fmt.Errorf("Contact with ID %d is not found", contactID)
}

func (cm *ContactManager) UpdateContact(updatedContact domain.Contact) error {
	for i, contact := range cm.Contacts {
		if contact.ID == updatedContact.ID {
			cm.Contacts[i] = updatedContact
			return nil
		}
	}
	return fmt.Errorf("Contact with ID %d is not found", updatedContact.ID)
}

func (cm *ContactManager) DeleteContact(contactID int) error {
	for i, contact := range cm.Contacts {
		if contact.ID == contactID {
			cm.Contacts = append(cm.Contacts[:i], cm.Contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Contact with ID %d is not found", contactID)
}

func (cm *ContactManager) CreateGroup(group domain.Group) error {
	cm.Groups = append(cm.Groups, group)
	return nil
}

func (cm *ContactManager) ReadGroup(groupID int) (domain.Group, error) {
	for _, group := range cm.Groups {
		if group.ID == groupID {
			return group, nil
		}
	}
	return domain.Group{}, fmt.Errorf("Group with ID %d is not found", groupID)
}

func (cm *ContactManager) AddContactToGroup(contactID, groupID int) error {
	contact, err := cm.ReadContact(contactID)
	if err != nil {
		return err
	}

	for i, group := range cm.Groups {
		if group.ID == groupID {
			cm.Groups[i].Contacts = append(group.Contacts, contact)
			return nil
		}
	}
	return fmt.Errorf("Group with ID %d is not found", groupID)
}
