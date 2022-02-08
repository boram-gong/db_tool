package es

import "testing"

func TestES(t *testing.T) {
	// 为已有的索引添加字段
	//_, err = client.PutMapping().Index("user").BodyString(mapping1).Do(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}

	// 查询
	//get1, err := client.Get().Index("user").Id("1").Do(ctx)
	//if err != nil{
	//	panic(err)
	//}
	//if get1.Found{
	//	fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	//	// Got document 1 in version 824633838776 from index user, type _doc
	//}

	// Flush to make sure the documents got written.将文档涮入磁盘
	//_, err = client.Flush().Index("user").Do(ctx)
	//if err != nil {
	//	panic(err)
	//}

	// 更新文档 update
	//update, err := client.Update().Index("user").Id("1").
	//	Script(elastic.NewScriptInline("ctx._source.age += params.num").Lang("painless").Param("num", 1)).
	//	//Upsert(map[string]interface{}{"created": "2020-06-17"}). // 插入未初始化的字段value
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
	//fmt.Printf("New version of user %q is now %d\n", update.Id, update.Version)
	// 更新方法2
	//update,err := client.Update().Index("user").Id("1").
	//	Script(elastic.NewScriptInline("ctx._source.created=params.date").Lang("painless").Param("date","2020-06-17")).
	//	Do(ctx)
	//termQuery := elastic.NewTermQuery("name", "bob")
	//update,err := client.UpdateByQuery("user").Query(termQuery).
	//	Script(elastic.NewScriptInline("ctx._source.age += params.num").Lang("painless").Param("num", 1)).
	//	Do(ctx)
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Printf("New version of user %q is now %d\n", update.Id, update.Version)
	//fmt.Println(update)
	// 删除文档

	//termQuery := elastic.NewTermQuery("name", "mike")
	//_, err = client.DeleteByQuery().Index("user"). // search in index "user"
	//	Query(termQuery). // specify the query
	//	Do(ctx)
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
}
