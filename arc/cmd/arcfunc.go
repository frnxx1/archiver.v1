package cmd

import (
	/* "archive/zip"
	"bytes"
	"fmt"
	"io" */
	"archive/zip"
	
	"log"
	"os"
)
 



func osCreate(archiv,archfile string)(string,error){
      a := archfile
	filezip,err := os.Create(archiv)
	
	  if err != nil{
		log.Fatal("Вероятно такой файл уже создан")
	  }
	  defer filezip.Close()

	  wr := zip.NewWriter(filezip)
      defer wr.Close()
	  f, err := wr.Create(a)
	  
  if err != nil {
    log.Fatal(err,f)
  }
	  return "", err


	
}


