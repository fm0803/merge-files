package main
/*
import "sync"

func main(){
	fd1 := Open(pathStr1)
	defer fd1.Close()
	c := make(chan string,100)
	//done := make(chan int)
	wg := sync.WaitGroup()
	wg.Add(len(pathArr))
	go func(){
		for{
			select str := <-c {
				if str != ""{
				fd1.WriteLine(str)
			}else{
				return
			}
				// case str := <-c :
				//  fd1.WriteLine(str)
				//case <-done:
				// return
			}
			}
		}()
		var helper func(string){}
		helper = func(path string){
			defer wg.Done()
			fd := Open(path)
			defer fd.Close()
			content := ""
			contentStr,err := fd.ReadLine()
			for err != EOF {
				c<-contentStr
				contentStr,err = fd.ReadLine()
			}

		}
		for _,path := range pathArr{
			go helper(path)
		}
		wg.Wait()
		close(c)
		//done<-1
	}*/