package content
type canal struct {
	tableName 			string
	cacheKey			string
}
func (c *canal)GetTableName(){

}
func (c *canal)SetTableName(){

}

func (c *canal)GetCacheKey(){

}
func (c *canal)SetCacheKey(){

}

func (c *canal)Canal(tableName string,cacheKey string)  {
	c.cacheKey  = cacheKey
	c.tableName = tableName
}