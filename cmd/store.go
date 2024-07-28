package main

import (
	"context"
	"errors"
	"example/models"
	"github.com/MrSametBurgazoglu/enterprise/client"
	"log"
	"time"
)

func main() {
	dbUrl := "postgresql://testuser:54M3754M37@localhost:5433/testdb?search_path=public"
	db, err := models.NewDB(dbUrl)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	tes := models.NewTest(ctx, db)
	tes.SetName("name")
	tes.SetCreatedAt(time.Now())

	den := models.NewDeneme(ctx, db)
	den.SetCount(20)
	den.SetDenemeType(models.DenemeTypeTestType)
	den.SetTestIDValue(tes.GetID())

	acc := models.NewAccount(ctx, db)
	acc.SetName("name")
	acc.SetSurname("surname")
	acc.SetDenemeIDValue(den.GetID())

	acc2 := models.NewAccount(ctx, db)
	acc2.SetName("name")
	acc2.SetSurname("surname")
	acc2.SetDenemeIDValue(den.GetID())
	println(acc.GetID().String(), "id")
	println(acc2.GetID().String(), "id")

	err = tes.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = den.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = acc.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = acc2.Create()
	if err != nil {
		log.Fatal(err)
	}

	t := models.NewTest(ctx, db)
	t.Where(t.IsIDEqual(tes.GetID()))
	println("test", tes.GetID().String())
	t.WithDenemeList(func(denemeList *models.DenemeList) {
		denemeList.Where(
			denemeList.IsCountEqual(20),
		)
		denemeList.Order(models.DenemeIDField)
		denemeList.WithAccountList()
	})

	err, _ = t.Get()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range t.DenemeList.Items[0].AccountList.Items {
		println(item.GetID().String())
	}

	var count, maximum, minimum, sum int

	t.DenemeList.Where(
		t.DenemeList.IsIsActiveEqual(true),
	)
	scanner, err := t.DenemeList.Aggregate(func(aggregate *client.Aggregate) {
		aggregate.Count("*", &count)
		aggregate.Max(models.DenemeCountField, &maximum)
		aggregate.Min(models.DenemeCountField, &minimum)
		aggregate.Sum(models.DenemeCountField, &sum)
		aggregate.GroupBy(models.DenemeDenemeTypeField)
	})
	err = scanner()
	for err == nil {
		println(count, maximum, minimum, sum)
		err = scanner()
	}
	if err != nil && !errors.Is(err, client.ErrFinalRow) {
		log.Fatal(err)
	}

	den.SetDenemeType(models.DenemeTypeDenemeType)
	err = den.Update()
	if err != nil {
		log.Fatal(err)
	}
	err = den.Refresh()
	if err != nil {
		log.Fatal(err)
	}

	transaction, err := db.NewTransaction(ctx)
	if err != nil {
		log.Fatal(err)
	}

	test2 := models.NewTest(ctx, transaction)
	test2.SetName("transaction_name")
	test2.SetCreatedAt(time.Now())

	err = test2.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = transaction.Rollback(ctx)
	if err != nil {
		log.Fatal(err)
	}

	transaction2, err := db.NewTransaction(ctx)
	if err != nil {
		log.Fatal(err)
	}

	test3 := models.NewTest(ctx, transaction2)
	test3.SetName("transaction_name")
	test3.SetCreatedAt(time.Now())

	err = test3.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = transaction2.Commit(ctx)
	if err != nil {
		log.Fatal(err)
	}

	testList := models.NewTestList(ctx, db)
	testList.Where(
		testList.IsNameEqual("name"),
	)
	testList.Paging(0, 5)
	testList.WithDenemeList(func(denemeList *models.DenemeList) {
		denemeList.WithAccountList()
	})
	err, found := testList.List()
	if err == nil && found {
		for i, test := range testList.Items {
			println(i, test.GetID().String())
			for i2, deneme := range test.DenemeList.Items {
				println(i2, deneme.GetID().String())
				for i3, account := range deneme.AccountList.Items {
					println(i3, account.GetID().String())
				}
			}
			//println(item.Deneme.GetCount())
		}
	} else {
		println(err.Error())
	}

	acc3 := models.NewAccount(ctx, db)
	acc3.SetName("with_group")
	acc3.SetSurname("surname")

	err = acc3.Create()
	if err != nil {
		log.Fatal(err)
	}

	group := models.NewGroup(ctx, db)
	group.SetName("with_account")
	group.SetSurname("surname")

	err = group.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = acc3.AddIntoGroup(group)
	if err != nil {
		log.Fatal(err)
	}

	acc4 := models.NewAccount(ctx, db)
	acc4.Where(acc4.IsIDEqual(acc3.GetID()))
	acc4.WithGroupList()
	err, ok := acc4.Get()
	if err == nil && ok {
		println(acc4.GroupList.Items[0].GetID().String())
	}

}
