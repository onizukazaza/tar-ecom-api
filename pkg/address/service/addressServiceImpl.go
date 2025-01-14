package service

import (
    "fmt"
    "time"

    "github.com/google/uuid"
    "github.com/onizukazaza/tar-ecom-api/entities"
    _userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
    _addressRepository "github.com/onizukazaza/tar-ecom-api/pkg/address/repository"
    _addressModel "github.com/onizukazaza/tar-ecom-api/pkg/address/model"
)

type addressServiceImpl struct {
    addressRepository _addressRepository.AddressRepository
    userRepository    _userRepository.UserRepository
}

func NewAddressServiceImpl(
	addressRepository _addressRepository.AddressRepository, 
	userRepository _userRepository.UserRepository ,
	) AddressService {
    return &addressServiceImpl{
        addressRepository,
        userRepository,
    }
}

func (s *addressServiceImpl) CreateAddress(req *_addressModel.CreateAddressReq) (*_addressModel.Address, error) {
    user, err := s.userRepository.FindUserByID(req.UserID)
    if err != nil || user == nil {
        return nil, fmt.Errorf("user not found")
    }

    address := &entities.Address{
        ID:            uuid.New(),
        UserID:        uuid.MustParse(req.UserID),
        RecipientName: req.RecipientName,
        Province:      req.Province,
        District:      req.District,
        SubDistrict:   req.SubDistrict,
        Postal:        req.Postal,
        AddressLine:   req.AddressLine,
        Contact:       req.Contact,
        Favorite:      req.Favorite,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }

    if err := s.addressRepository.CreateAddress(address); err != nil {
        return nil, fmt.Errorf("failed to create address: %w", err)
    }

    return address.ToModel(), nil
}

func (s *addressServiceImpl) ListAddresses(userID string) ([]*_addressModel.Address, error) {
    addresses, err := s.addressRepository.ListAddresses(userID)
    if err != nil {
        return nil, fmt.Errorf("failed to list addresses: %w", err)
    }

    if len(addresses) == 0 {
        return nil, fmt.Errorf("no addresses found for user %s", userID)
    }

    models := make([]*_addressModel.Address, len(addresses))
    for i, address := range addresses {
        models[i] = address.ToModel()
    }

    return models, nil
}


func (s *addressServiceImpl) FindAddressByID(id string, userID string) (*_addressModel.Address, error) {
    address, err := s.addressRepository.FindAddressByID(id, userID)
    if err != nil {
        return nil, fmt.Errorf("address not found: %w", err)
    }

    return address.ToModel(), nil
}

func (s *addressServiceImpl) UpdateFavoriteAddress(id string, userID string, favorite bool) error {
    // ตรวจสอบว่าที่อยู่เป็นของผู้ใช้งาน
    address, err := s.addressRepository.FindAddressByID(id, userID)
    if err != nil || address == nil {
        return fmt.Errorf("address not found or not owned by user")
    }

    if favorite {
        // หากกำหนดให้ Favorite เป็น True, ต้องปิด Favorite อื่นทั้งหมดก่อน
        if err := s.addressRepository.ClearAllFavorites(userID); err != nil {
            return fmt.Errorf("failed to clear existing favorites: %w", err)
        }
    }

    // อัปเดตสถานะ Favorite
    if err := s.addressRepository.UpdateFavoriteAddress(id, userID, favorite); err != nil {
        return fmt.Errorf("failed to update favorite address: %w", err)
    }

    return nil
}
func (s *addressServiceImpl) DeleteAddress(id string, userID string) error {
    // ตรวจสอบว่าที่อยู่เป็นของผู้ใช้งาน
    address, err := s.addressRepository.FindAddressByID(id, userID)
    if err != nil || address == nil {
        return fmt.Errorf("address not found or not owned by user")
    }

    // ลบที่อยู่
    if err := s.addressRepository.DeleteAddress(id, userID); err != nil {
        return fmt.Errorf("failed to delete address: %w", err)
    }

    return nil
}

func (s *addressServiceImpl) EditAddress(req *_addressModel.EditAddressReq) error {
    // ตรวจสอบว่า Address เป็นของ User
    address, err := s.addressRepository.FindAddressByID(req.ID, req.UserID)
    if err != nil || address == nil {
        return fmt.Errorf("address not found or not owned by user")
    }

    // ตรวจสอบฟิลด์ที่ส่งมาและอัปเดตเฉพาะฟิลด์ที่ไม่เป็นค่าเริ่มต้น
    if req.RecipientName != "" {
        address.RecipientName = req.RecipientName
    }
    if req.Province != "" {
        address.Province = req.Province
    }
    if req.District != "" {
        address.District = req.District
    }
    if req.SubDistrict != "" {
        address.SubDistrict = req.SubDistrict
    }
    if req.Postal != "" {
        address.Postal = req.Postal
    }
    if req.AddressLine != "" {
        address.AddressLine = req.AddressLine
    }
    if req.Contact != "" {
        address.Contact = req.Contact
    }
    if req.Favorite {
        address.Favorite = req.Favorite
    }
    address.UpdatedAt = time.Now()

    // อัปเดตข้อมูลใน Repository
    if err := s.addressRepository.EditAddress(address); err != nil {
        return fmt.Errorf("failed to edit address: %w", err)
    }

    return nil
}

