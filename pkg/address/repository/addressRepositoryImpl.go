package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/onizukazaza/tar-ecom-api/entities"
    "github.com/gofiber/fiber/v2/log"
	"fmt"
    _AddressException "github.com/onizukazaza/tar-ecom-api/pkg/address/exception"

)

type addressRepositoryImpl struct {
    db *sqlx.DB
}

func NewAddressRepositoryImpl(db *sqlx.DB) AddressRepository {
    return &addressRepositoryImpl{db: db}
}

func (r *addressRepositoryImpl) CreateAddress(address *entities.Address) error {
    query := `
        INSERT INTO addresses (
            id, user_id, recipient_name, province, district, subdistrict, postal, 
            address_line, contact, favorite, created_at, updated_at
        ) VALUES (
            :id, :user_id, :recipient_name, :province, :district, :subdistrict, :postal, 
            :address_line, :contact, :favorite, :created_at, :updated_at
        )
    `

    _, err := r.db.NamedExec(query, address)
    if err != nil {
        log.Errorf("Failed to create address for user %s: %v", address.UserID, err)
        return &_AddressException.FailedToCreateAddress{
            UserID: address.UserID.String(),
            Reason: err.Error(),
        }
    }

    return nil
}


func (r *addressRepositoryImpl) EditAddress(address *entities.Address) error {
    query := `
        UPDATE addresses 
        SET recipient_name = :recipient_name, province = :province, district = :district, 
            subdistrict = :subdistrict, postal = :postal, address_line = :address_line, 
            contact = :contact, favorite = :favorite, updated_at = NOW()
        WHERE id = :id AND user_id = :user_id
    `
    result, err := r.db.NamedExec(query, address)
    if err != nil {
        log.Errorf("Failed to edit address with ID %s: %v", address.ID, err)
        return &_AddressException.FailedToUpdateAddress{}
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get affected rows: %w", err)
    }
    if rowsAffected == 0 {
        return  &_AddressException.AddressNotFound{}
    }

    return nil
}

func (r *addressRepositoryImpl) ListAddresses(userID string) ([]*entities.Address, error) {
    addresses := make([]*entities.Address, 0)
    query := `
        SELECT 
            id, user_id, recipient_name, province, district, subdistrict, postal, 
            address_line, contact, favorite, created_at, updated_at
        FROM addresses 
        WHERE user_id = $1
    `

    err := r.db.Select(&addresses, query, userID)
    if err != nil {
        log.Errorf("Failed to list addresses for user %s: %v", userID, err)
        return nil, err
    }

    return addresses, nil
}



func (r *addressRepositoryImpl) FindAddressByID(id string, userID string) (*entities.Address, error) {
    var address entities.Address
    query := "SELECT * FROM addresses WHERE id = $1 AND user_id = $2"

    err := r.db.Get(&address, query, id, userID)
    if err != nil {
        log.Errorf("Failed to find address by ID: %v", err)
        return nil, &_AddressException.AddressNotFound{}
    }

    return &address, nil
}



func (r *addressRepositoryImpl) UpdateFavoriteAddress(id string, userID string, favorite bool) error {
    query := `
        UPDATE addresses
        SET favorite = $1, updated_at = NOW()
        WHERE id = $2 AND user_id = $3
    `

    result, err := r.db.Exec(query, favorite, id, userID)
    if err != nil {
        log.Errorf("Failed to update favorite address %s for user %s: %v", id, userID, err)
        return fmt.Errorf("failed to update favorite address: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get affected rows: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("address not found or not owned by user")
    }

    return nil
}


func (r *addressRepositoryImpl) ClearAllFavorites(userID string) error {
    query := `
        UPDATE addresses
        SET favorite = FALSE, updated_at = NOW()
        WHERE user_id = $1 AND favorite = TRUE
    `

    _, err := r.db.Exec(query, userID)
    if err != nil {
        log.Errorf("Failed to clear existing favorites for user %s: %v", userID, err)
        return &_AddressException.FailedToUpdateFavorite{}
    }

    return nil
}


func (r *addressRepositoryImpl) DeleteAddress(id string, userID string) error {
    query := `
        DELETE FROM addresses
        WHERE id = $1 AND user_id = $2
    `

    result, err := r.db.Exec(query, id, userID)
    if err != nil {
        log.Errorf("Failed to hard delete address %s for user %s: %v", id, userID, err)
        return fmt.Errorf("failed to delete address: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to check affected rows: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("address not found or not owned by user")
    }

    return nil
}
