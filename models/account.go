package models

import (
	"context"
	"github.com/MrSametBurgazoglu/enterprise/client"

	"github.com/google/uuid"
)

const AccountTableName = "account"

const (
	AccountIDField       string = "id"
	AccountNameField     string = "name"
	AccountSurnameField  string = "surname"
	AccountDenemeIDField string = "deneme_id"
)

func NewAccount(ctx context.Context, dc client.DatabaseClient) *Account {
	v := &Account{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.changedFields = make(map[string]any)
	v.result.Init()
	v.Default()
	return v
}

func NewRelationAccount(ctx context.Context, dc client.DatabaseClient) *Account {
	v := &Account{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.changedFields = make(map[string]any)
	v.result.Init()
	return v
}

type Account struct {
	id uuid.UUID

	name string

	surname string

	denemeid *uuid.UUID

	changedFields map[string]any
	serialFields  []*client.SelectedField

	ctx    context.Context
	client *client.Client
	AccountPredicate
	relations *client.RelationList

	Deneme    *Deneme
	GroupList *GroupList

	result AccountResult
}

func (t *Account) GetDBName() string {
	return AccountTableName
}

func (t *Account) GetSelector() *AccountResult {
	t.result.selectedFields = nil
	return &t.result
}

func (t *Account) GetRelationList() *client.RelationList {
	return t.relations
}

func (t *Account) IsExist() bool {
	var v uuid.UUID
	return t.id != v
}

func (t *Account) GetPrimaryKey() uuid.UUID {
	return t.id
}

func NewAccountList(ctx context.Context, dc client.DatabaseClient) *AccountList {
	v := &AccountList{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.result.Init()
	return v
}

func NewRelationAccountList(ctx context.Context, dc client.DatabaseClient) *AccountList {
	v := &AccountList{client: client.NewClient(dc), ctx: ctx}
	v.relations = new(client.RelationList)
	v.relations.RelationMap = make(map[string]*client.Relation)
	v.result.Init()
	return v
}

type AccountList struct {
	Items []*Account

	ctx    context.Context
	client *client.Client
	AccountPredicate
	order     []*client.Order
	paging    *client.Paging
	relations *client.RelationList
	result    AccountResult
}

func (t *AccountList) GetDBName() string {
	return AccountTableName
}

func (t *AccountList) GetRelationList() *client.RelationList {
	return t.relations
}

func (t *AccountList) IsExist() bool {
	return t.Items[len(t.Items)-1].IsExist()
}

func (t *Account) SetID(v uuid.UUID) {
	t.id = v
	t.SetIDField()
}
func (t *Account) SetName(v string) {
	t.name = v
	t.SetNameField()
}
func (t *Account) SetSurname(v string) {
	t.surname = v
	t.SetSurnameField()
}
func (t *Account) SetDenemeID(v *uuid.UUID) {
	t.denemeid = v
	t.SetDenemeIDField()
}

func (t *Account) SetDenemeIDValue(v uuid.UUID) {
	t.SetDenemeID(&v)
	t.SetDenemeIDField()
}

func (t *Account) SetIDNillable(v *uuid.UUID) {
	if v == nil {
		return
	}
	t.SetID(*v)
}
func (t *Account) SetNameNillable(v *string) {
	if v == nil {
		return
	}
	t.SetName(*v)
}
func (t *Account) SetSurnameNillable(v *string) {
	if v == nil {
		return
	}
	t.SetSurname(*v)
}

func (t *Account) ParseID(v string) error {
	parsedID, err := uuid.Parse(v)
	if err != nil {
		return err
	}
	t.id = parsedID
	t.SetIDField()
	return nil
}

func (t *Account) ParseDenemeID(v string) error {
	parsedID, err := uuid.Parse(v)
	if err != nil {
		return err
	}
	t.denemeid = &parsedID
	t.SetDenemeIDField()
	return nil
}

func (t *Account) IDIN(v ...uuid.UUID) bool {
	for _, x := range v {
		if t.id == x {
			return true
		}
	}
	return false
}

func (t *Account) NameIN(v ...string) bool {
	for _, x := range v {
		if t.name == x {
			return true
		}
	}
	return false
}

func (t *Account) SurnameIN(v ...string) bool {
	for _, x := range v {
		if t.surname == x {
			return true
		}
	}
	return false
}

func (t *Account) DenemeIDIN(v ...uuid.UUID) bool {
	if t.denemeid == nil {
		return false
	}
	for _, x := range v {
		if *t.denemeid == x {
			return true
		}
	}
	return false
}

func (t *Account) IDNotIN(v ...uuid.UUID) bool {
	for _, x := range v {
		if t.id == x {
			return false
		}
	}
	return true
}

func (t *Account) NameNotIN(v ...string) bool {
	for _, x := range v {
		if t.name == x {
			return false
		}
	}
	return true
}

func (t *Account) SurnameNotIN(v ...string) bool {
	for _, x := range v {
		if t.surname == x {
			return false
		}
	}
	return true
}

func (t *Account) DenemeIDNotIN(v ...uuid.UUID) bool {
	if t.denemeid == nil {
		return true
	}
	for _, x := range v {
		if *t.denemeid == x {
			return false
		}
	}
	return true
}

func (t *Account) GetID() uuid.UUID {
	return t.id
}
func (t *Account) GetName() string {
	return t.name
}
func (t *Account) GetSurname() string {
	return t.surname
}
func (t *Account) GetDenemeID() *uuid.UUID {
	return t.denemeid
}

func (t *Account) SetIDField() {
	t.changedFields[AccountIDField] = t.id
}
func (t *Account) SetNameField() {
	t.changedFields[AccountNameField] = t.name
}
func (t *Account) SetSurnameField() {
	t.changedFields[AccountSurnameField] = t.surname
}
func (t *Account) SetDenemeIDField() {
	t.changedFields[AccountDenemeIDField] = t.denemeid
}

func (t *Account) WithDeneme(opts ...func(*Deneme)) {
	t.Deneme = NewRelationDeneme(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(t.Deneme)
	}
	t.result.Deneme = new(DenemeResult)
	t.result.Deneme.Init()
	t.result.relations = append(t.result.relations, t.result.Deneme)
	t.result.relationsMap["deneme"] = t.result.Deneme
	for _, Relation := range t.Deneme.relations.Relations {
		t.result.Deneme.relations = append(t.result.Deneme.relations, Relation.RelationResult)
		t.result.Deneme.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:  t.Deneme,
			RelationTable:  "deneme",
			RelationResult: t.result.Deneme,
			Where:          t.Deneme.where,

			RelationWhere: &client.RelationCondition{
				RelationValue: "id",
				TableValue:    "deneme_id",
			},
		},
	)
	t.relations.RelationMap["deneme"] = t.relations.Relations[len(t.relations.Relations)-1]
}
func (t *Account) WithGroupList(opts ...func(*GroupList)) {
	t.GroupList = NewRelationGroupList(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(t.GroupList)
	}
	t.result.Group = new(GroupResult)
	t.result.Group.Init()
	t.result.relations = append(t.result.relations, t.result.Group)
	t.result.relationsMap["group"] = t.result.Group
	for _, Relation := range t.GroupList.relations.Relations {
		t.result.Group.relations = append(t.result.Group.relations, Relation.RelationResult)
		t.result.Group.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:   t.GroupList,
			RelationTable:   "group",
			RelationResult:  t.result.Group,
			Where:           t.GroupList.where,
			ManyToManyTable: "account_group",
			RelationWhere: &client.RelationCondition{
				RelationValue:      "account_id",
				TableValue:         "group_id",
				RelationTableValue: "id",
			},
		},
	)
	t.relations.RelationMap["group"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *AccountList) WithDeneme(opts ...func(*Deneme)) {
	//t.Deneme = NewRelationDeneme(t.ctx, t.client.Database)
	v := NewRelationDeneme(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(v)
	}
	t.result.Deneme = new(DenemeResult)
	t.result.Deneme.Init()
	t.result.relations = append(t.result.relations, t.result.Deneme)
	t.result.relationsMap["deneme"] = t.result.Deneme
	for _, Relation := range v.relations.Relations {
		t.result.Deneme.relations = append(t.result.Deneme.relations, Relation.RelationResult)
		t.result.Deneme.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:  v,
			RelationTable:  "deneme",
			RelationResult: t.result.Deneme,
			Where:          v.where,
			RelationWhere: &client.RelationCondition{
				RelationValue: "id",
				TableValue:    "deneme_id",
			},
		},
	)
	t.relations.RelationMap["deneme"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *AccountList) cleanDeneme() {
	Relation := t.Items[len(t.Items)-1].relations
	p := 0
	for i, v := range Relation.Relations {
		if v.RelationTable == "deneme" {
			p = i
		}
	}
	Relation.Relations = append(Relation.Relations[:p], Relation.Relations[p+1:]...)
}
func (t *AccountList) WithGroupList(opts ...func(*GroupList)) {
	//t.GroupList = NewRelationGroupList(t.ctx, t.client.Database)
	v := NewRelationGroupList(t.ctx, t.client.Database)
	for _, opt := range opts {
		opt(v)
	}
	t.result.Group = new(GroupResult)
	t.result.Group.Init()
	t.result.relations = append(t.result.relations, t.result.Group)
	t.result.relationsMap["group"] = t.result.Group
	for _, Relation := range v.relations.Relations {
		t.result.Group.relations = append(t.result.Group.relations, Relation.RelationResult)
		t.result.Group.relationsMap[Relation.RelationTable] = Relation.RelationResult
	}
	t.relations.Relations = append(t.relations.Relations,
		&client.Relation{
			RelationModel:  v,
			RelationTable:  "group",
			RelationResult: t.result.Group,
			Where:          v.where,
			RelationWhere: &client.RelationCondition{
				RelationValue: "account_id",
				TableValue:    "group_id",
			},
		},
	)
	t.relations.RelationMap["group"] = t.relations.Relations[len(t.relations.Relations)-1]
}

func (t *AccountList) cleanGroupList() {
	Relation := t.Items[len(t.Items)-1].relations
	p := 0
	for i, v := range Relation.Relations {
		if v.RelationTable == "group" {
			p = i
		}
	}
	Relation.Relations = append(Relation.Relations[:p], Relation.Relations[p+1:]...)
}

func (t *Account) Default() {
	t.id = uuid.New()
	t.changedFields[AccountIDField] = t.id

}

func (t *Account) SetResult(result client.Result) {
	if t == nil {
		v := NewRelationAccount(t.ctx, t.client.Database)
		*t = *v
	}
	t.result = *result.(*AccountResult)
}

func (t *AccountList) SetResult(result client.Result) {
	if t == nil {
		v := NewRelationAccountList(t.ctx, t.client.Database)
		*t = *v
	}
	t.result = *result.(*AccountResult)
}

func (t *Account) ScanResult() {
	t.id = t.result.id
	t.name = t.result.name
	t.surname = t.result.surname
	t.denemeid = t.result.denemeid

	if _, ok := t.relations.RelationMap["deneme"]; ok {
		if t.Deneme == nil {
			t.Deneme = NewRelationDeneme(t.ctx, t.client.Database)
		}
		t.Deneme.relations = t.relations.RelationMap["deneme"].RelationModel.GetRelationList()
		t.Deneme.SetResult(t.result.relationsMap["deneme"])
		t.Deneme.ScanResult()
	}
	if _, ok := t.relations.RelationMap["group"]; ok {
		if t.GroupList == nil {
			t.GroupList = NewRelationGroupList(t.ctx, t.client.Database)
		}
		t.GroupList.relations = t.relations.RelationMap["group"].RelationModel.GetRelationList()
		t.GroupList.SetResult(t.result.relationsMap["group"])
		t.GroupList.ScanResult()
	}
}

func (t *Account) CheckPrimaryKey(v uuid.UUID) bool {
	return t.id == v
}

func (t *AccountList) ScanResult() {
	var v *Account
	if len(t.Items) == 0 {
		v = NewRelationAccount(t.ctx, t.client.Database)
		t.Items = append(t.Items, v)
	} else {
		for _, item := range t.Items {
			if item.CheckPrimaryKey(t.result.id) {
				v = item
				break
			}
		}
	}
	if v == nil {
		v = NewRelationAccount(t.ctx, t.client.Database)
		t.Items = append(t.Items, v)
	}
	v.result = t.result
	v.relations = t.relations
	v.ScanResult()
}

func (t *Account) Get() (error, bool) {
	return t.client.Get(t.ctx, t.where, t, &t.result)
}

func (t *Account) Refresh() error {
	return t.client.Refresh(t.ctx, t, &t.result, AccountIDField, t.id)
}

func (t *Account) Create() error {
	return t.client.Create(t.ctx, AccountTableName, t.changedFields, t.serialFields)
}

func (t *Account) Update() error {
	return t.client.Update(t.ctx, AccountTableName, t.changedFields, AccountIDField, t.id)
}

func (t *Account) Delete() error {
	return t.client.Delete(t.ctx, AccountTableName, AccountIDField, t.id)
}

func (t *AccountList) List() (error, bool) {
	return t.client.List(t.ctx, t.where, t, &t.result, t.order, t.paging)
}

func (t *AccountList) Aggregate(f func(aggregate *client.Aggregate)) (func() error, error) {
	a := new(client.Aggregate)
	f(a)
	return t.client.Aggregate(t.ctx, t.where, t, a)
}

func (t *AccountList) Order(field string) *AccountList {
	t.order = append(t.order, &client.Order{Field: field})
	return t
}

func (t *AccountList) OrderDesc(field string) *AccountList {
	t.order = append(t.order, &client.Order{Field: field, Desc: true})
	return t
}

func (t *AccountList) Paging(skip, limit int) *AccountList {
	t.paging = &client.Paging{Skip: skip, Limit: limit}
	return t
}

type AccountResult struct {
	id       uuid.UUID
	name     string
	surname  string
	denemeid *uuid.UUID

	selectedFields []*client.SelectedField

	Deneme *DenemeResult
	Group  *GroupResult

	relations    []client.Result
	relationsMap map[string]client.Result
}

func (t *AccountResult) Init() {
	t.selectedFields = []*client.SelectedField{}
	t.relationsMap = make(map[string]client.Result)
	t.prepare()
	t.SelectAll()
}

func (t *AccountResult) GetSelectedFields() []*client.SelectedField {
	return t.selectedFields
}

func (t *AccountResult) GetRelations() []client.Result {
	return t.relations
}

func (t *AccountResult) prepare() {
	t.denemeid = &uuid.Nil

}

func (t *AccountResult) SelectID() {
	v := &client.SelectedField{Name: AccountIDField, Value: &t.id}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *AccountResult) SelectName() {
	v := &client.SelectedField{Name: AccountNameField, Value: &t.name}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *AccountResult) SelectSurname() {
	v := &client.SelectedField{Name: AccountSurnameField, Value: &t.surname}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *AccountResult) SelectDenemeID() {

	v := &client.SelectedField{Name: AccountDenemeIDField, Value: t.denemeid}
	t.selectedFields = append(t.selectedFields, v)
}

func (t *AccountResult) GetDBName() string {
	return AccountTableName
}

func (t *AccountResult) SelectAll() {
	t.SelectID()
	t.SelectName()
	t.SelectSurname()
	t.SelectDenemeID()

}

func (t *AccountResult) IsExist() bool {
	if t == nil {
		return false
	}
	var v uuid.UUID
	return t.id != v
}

func (t *Account) AddIntoGroup(relationship *Group) error {
	return t.client.AddManyToManyRelation(t.ctx, "account_group", "account_id", "group_id", t.id, relationship.GetPrimaryKey())
}

func (t *Account) RemoveFromGroup(relationship *Group) error {
	return t.client.DeleteManyToManyRelation(t.ctx, "account_group", "account_id", "group_id", t.id, relationship.GetPrimaryKey())
}

func (t *Account) IsInGroup(relationship *Group) (bool, error) {
	return t.client.ExistManyToManyRelation(t.ctx, "account_group", "account_id", "group_id", t.id, relationship.GetPrimaryKey())
}
