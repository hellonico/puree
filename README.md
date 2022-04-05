# puree <img src="index.jpg" width="32px" height="32px"/>

By default, mongoimport does not add the date nor the filename.

*puree* is a very small utility to add metadata to a file to be imported.

Usage:

```bash
⋊> ~/G/puree on master ⨯ ./puree
  -in string
    	Input folder (default ".")
  -out string
    	Output folder (default "out")
    	
From folder < ., to folder > out 
```

Each csv file from the input folder will be added two columns, with the date, and the filename.
The resulting file will be copied into the specified out folder.


Developed with [Jetbrains Goland](https://www.jetbrains.com/go/promo/)

