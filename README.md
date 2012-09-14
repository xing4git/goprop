This is a golang package.
Package prop attempts to provider a similar usage with Properties in java.


Usage, consume you have a my.conf. then, you invoke:


myinfo, err := prop.Load(my.conf)
if err != nil {
  // ...
}
for k, v := range myinfo {
	fmt.Println(key + "=" + value)
}


Get:
go get github.com/xing4git/goprop