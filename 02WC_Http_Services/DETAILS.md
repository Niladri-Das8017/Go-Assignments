**WC_httpServics:**	Small service that accepts as input a body of text, such as that from a book,
and returns the top ten most-used words along with how many times they occur in the text. 

**View Result:**     Send text with POST method to port:
                                http://localhost:4040/wordcount


**routers** 

        Handles routes

**controller**

      i.    ServerHome: Homepage

      ii.   WordCount:  1. Fetch text request
                        2. Check error
                        3. Check empty String
                        4. Call helpers.Wordcount
                        5. Call helpers.SortWc
                        6. Send the sorted slice of WordCount as response
**helper**

        i.  WordCount: Counts frequency and occourance of words & return slice of Word-Count
        ii. SortWc: Sort the Slice by Count
    