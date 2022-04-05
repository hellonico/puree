# puree 

<img src="index.png" width="64px" height="64px"/>

By default, mongoimport does not add useful metadata like date or filename of the file to be imported.
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

<img src="index.png" width="64px" height="64px"/>

Developed with [Jetbrains Goland](https://www.jetbrains.com/go/promo/)

