package models

import "github.com/MrSametBurgazoglu/enterprise/client"

import "github.com/google/uuid"

type AccountPredicate struct {
	where []*client.WhereList
}

func (t *AccountPredicate) Where(w ...*client.Where) {
	t.where = nil
	wl := &client.WhereList{}
	wl.Items = append(wl.Items, w...)
	t.where = append(t.where, wl)
}

func (t *AccountPredicate) ORWhere(w ...*client.Where) {
	wl := &client.WhereList{}
	wl.Items = append(wl.Items, w...)
	t.where = append(t.where, wl)
}

func (t *AccountPredicate) IsIDEqual(v uuid.UUID) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     AccountIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsNameEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     AccountNameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsSurnameEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     AccountSurnameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsDenemeIDEqual(v uuid.UUID) *client.Where {
	return &client.Where{
		Type:     client.EQ,
		Name:     AccountDenemeIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsIDNotEqual(v uuid.UUID) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     AccountIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsNameNotEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     AccountNameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsSurnameNotEqual(v string) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     AccountSurnameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsDenemeIDNotEqual(v uuid.UUID) *client.Where {
	return &client.Where{
		Type:     client.NEQ,
		Name:     AccountDenemeIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsIDIN(v ...uuid.UUID) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     AccountIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsNameIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     AccountNameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsSurnameIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     AccountSurnameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsDenemeIDIN(v ...uuid.UUID) *client.Where {
	return &client.Where{
		Type:     client.ANY,
		Name:     AccountDenemeIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsIDNotIN(v ...uuid.UUID) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     AccountIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsNameNotIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     AccountNameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsSurnameNotIN(v ...string) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     AccountSurnameField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsDenemeIDNotIN(v ...uuid.UUID) *client.Where {
	return &client.Where{
		Type:     client.NANY,
		Name:     AccountDenemeIDField,
		HasValue: true,
		Value:    v,
	}
}

func (t *AccountPredicate) IsDenemeIDNil() *client.Where {
	return &client.Where{
		Type: client.NIL,
		Name: AccountDenemeIDField,
	}
}

func (t *AccountPredicate) IsDenemeIDNotNil() *client.Where {
	return &client.Where{
		Type: client.NNIL,
		Name: AccountDenemeIDField,
	}
}
