package entity

import (
	"auth"
	"convention/agendaerror"
	"convention/codec"
	log "util/logger"
)

// var logln = util.Log
// var logf = util.Logf

// Username represents username, a unique identifier, of User
// Identifier
type Username string

// Empty checks if Username empty
func (name Username) Empty() bool {
	return name == ""
}

// Valid checks if Username valid
func (name Username) Valid() bool {
	return !name.Empty() // NOTE: may not only !empty
}

func (name Username) String() string {
	return string(name)
}

type UserInfoPublic struct {
	Name Username

	Mail  string
	Phone string
}

type Auth = auth.Auth

// UserInfo represents the informations of a User
type UserInfo struct {
	UserInfoPublic

	Auth Auth
}

// UserInfoSerializable represents serializable UserInfo
type UserInfoSerializable = UserInfo

// User represents a User, which is the actor of the operations like sponsor/join/cancel a meeting, etc
type User struct {
	UserInfo
}

// NewUser creates a User object with given UserInfo
func NewUser(info UserInfo) *User {
	u := new(User)
	u.UserInfo = info
	return u
}

// LoadUser load a User into given container(u) from given decoder
func LoadUser(decoder codec.Decoder, u *User) {
	uInfoSerial := new(UserInfoSerializable)
	err := decoder.Decode(uInfoSerial)
	if err != nil {
		log.Fatal(err)
	}
	u.UserInfo = *uInfoSerial // omit the deserial
}

// LoadedUser returns loaded User from given decoder
func LoadedUser(decoder codec.Decoder) *User {
	u := new(User)
	LoadUser(decoder, u)
	return u
}

// Save saves User with given encoder
func (u *User) Save(encoder codec.Encoder) error {
	return encoder.Encode(u.UserInfo) // omit the serial
}

// omit the serializer/deserializer

// ................................................................

// UserList represents a list/group of User (of the form of pointers of Users)
type UserList struct {
	Users map[Username](*User)
}

// UserListRaw also represents a list of User, but it is more trivial and more simple, i.e. it basically is ONLY a list of User, besides this, nothing
// NOTE: these type may be modified/removed in future
type UserListRaw = []*User

// UserInfoSerializableList represents a list of serializable UserInfo
type UserInfoSerializableList []UserInfoSerializable

// Serialize just serializes from UserList to UserInfoSerializableList
func (ul *UserList) Serialize() UserInfoSerializableList {
	ret := make(UserInfoSerializableList, 0, ul.Size())

	ul.ForEach(func(u *User) error {
		if u == nil {
			log.Warning("A nil User is to be used. Just SKIP OVER it.")
			return nil
		}
		ret = append(ret, u.UserInfo) // omit the serial
		return nil
	})

	return ret
}

// Size just returns the size
func (ulSerial UserInfoSerializableList) Size() int {
	return len(ulSerial)
}

// Deserialize deserializes from serialized UserInfoList to UserList
// CHECK: Now no used (, to loading a UserList)
func (ulSerial UserInfoSerializableList) Deserialize() *UserList {
	ret := NewUserList()

	for _, uInfo := range ulSerial {
		if uInfo.Name.Empty() {
			log.Warning("A No-Name UserInfo is to be used. Just SKIP OVER it.")
			continue
		}
		u := NewUser(uInfo) // omit the deserial
		if err := ret.Add(u); err != nil {
			log.Error(err)
		}
	}
	return ret
}

// NewUserList creates a UserList object
func NewUserList() *UserList {
	ul := new(UserList)
	ul.Users = make(map[Username](*User))
	return ul
}

// CHECK: Need in-place load method ?

// LoadUserList loads a UserList into given container(ul) from given decoder
func LoadUserList(decoder codec.Decoder, ul *UserList) {
	// CHECK: Need clear ul ?

	ulSerial := new(UserInfoSerializableList)
	if err := decoder.Decode(ulSerial); err != nil {
		log.Fatal(err)
	}
	for _, uInfoSerial := range *ulSerial {
		u := NewUser(uInfoSerial)
		if err := ul.Add(u); err != nil {
			log.Error(err)
		}
	}
}

// InitFrom loads UserList in-place from given []Username; Just like `init`
// CHECK: Not sure whether need/should return error
func (ul *UserList) InitFrom(li []Username) error {
	// clear ...
	ul.Users = NewUserList().Users

	for _, id := range li {
		u := id.RefInAllUsers()
		if err := ul.Add(u); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

// LoadFrom loads UserList in-place from given decoder; Just like `init`
func (ul *UserList) LoadFrom(decoder codec.Decoder) {
	LoadUserList(decoder, ul)
}

// LoadedUserList returns loaded UserList from given decoder
func LoadedUserList(decoder codec.Decoder) *UserList {
	ul := NewUserList()
	LoadUserList(decoder, ul)
	return ul
}

func (ul *UserList) Identifiers() []Username {
	ret := make([]Username, 0, ul.Size())
	for _, u := range ul.PublicInfos() {
		ret = append(ret, u.Name)
	}
	return ret
}
func (ul *UserList) PublicInfos() []UserInfoPublic {
	ret := make([]UserInfoPublic, 0, ul.Size())

	ul.ForEach(func(u *User) error {
		if u == nil {
			log.Warning("A nil User is to be used. Just SKIP OVER it.\n")
			return nil
		}
		ret = append(ret, u.UserInfoPublic)
		return nil
	})

	return ret
}

// Save use given encoder to Save UserList
func (ul *UserList) Save(encoder codec.Encoder) error {
	sl := ul.Serialize()
	return encoder.Encode(sl)
}

// Size just returns the number of User reference in UserList
func (ul *UserList) Size() int {
	return len(ul.Users)
}

// Ref just returns the reference of user with the given name
func (ul *UserList) Ref(name Username) *User {
	return ul.Users[name] // NOTE: if directly return accessed result from a map like this, would not get the (automatical) `ok`
}

// Contains just check if contains
func (ul *UserList) Contains(name Username) bool {
	u := ul.Ref(name)
	return u != nil
}

// Add just add
func (ul *UserList) Add(user *User) error {
	if user == nil {
		return agendaerror.ErrNilUser
	}
	name := user.Name
	if ul.Contains(name) {
		return agendaerror.ErrExistedUser
	}
	ul.Users[name] = user
	return nil
}

// Remove just remove
func (ul *UserList) Remove(user *User) error {
	if user == nil {
		return agendaerror.ErrNilUser
	}
	name := user.Name
	if ul.Contains(name) {
		delete(ul.Users, name) // NOTE: never error, according to 'go-maps-in-action'
		return nil
	}
	return agendaerror.ErrUserNotFound
}

// PickOut =~= Ref and then Remove
func (ul *UserList) PickOut(name Username) (*User, error) {
	if name.Empty() {
		return nil, agendaerror.ErrEmptyUsername
	}
	u := ul.Ref(name)
	if u == nil {
		return u, agendaerror.ErrUserNotFound
	}
	defer ul.Remove(u)
	return u, nil
}

// Slice returns a UserListRaw based on UserList ul
// NOTE: for the need of this simple agenda system, this seems somewhat needless
func (ul *UserList) Slice() UserListRaw {
	users := make(UserListRaw, 0, ul.Size())
	for _, u := range ul.Users {
		users = append(users, u) // CHECK: maybe better to use index in golang ?
	}
	return users
}

// ForEach used for all extension/concrete logic for whole UserList
func (ul *UserList) ForEach(fn func(*User) error) error {
	for _, v := range ul.Users {
		if err := fn(v); err != nil {
			// CHECK: Or, lazy error ?
			return err
		}
	}
	return nil
}

// Filter used for all extension/concrete select for whole UserList
func (ul *UserList) Filter(pred func(User) bool) *UserList {
	ret := NewUserList()
	for _, u := range ul.Users {
		if pred(*u) {
			ret.Add(u)
		}
	}
	return ret
}
