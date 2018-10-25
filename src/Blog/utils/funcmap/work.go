package funcmap

//==自定义模板函数==
//注意，1.自定义模板函数必须在beego.Run()之前注册和申明。
//注意，2.模板函数没有对参数的取值范围做任何约束，读者可以依情况加以控制。

//向下取整算中值
func Middle(start int, end int) (mid int) {
	mid = (start + end) / 2
	return
}

//分页时派上用场，用于判断TotalPage-CurrentPage>=n是否成立
func GeN(totalPage int, currentPage int, n int) (bool) {
	return totalPage-currentPage >= n
}

//分页时派上用场，用于自减/加n，n为负数时做加操作
func SubN(page int, n int) (int) {
	return page - n
}

