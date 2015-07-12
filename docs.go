package acousticid

/*
 Program flow:
 * Read input directory
  * For each file do the following
  	* Check if the file has id3 tag
  	* Is the id3 tag valid: Artist, Song Title, Album Name
  	* If not valid:
  	 * get the acoustic fingerprint?
  	 * Make the acoustic ID REST call
  	 * If valid result found:
  	  * Make ID3 tag struct
  	  * Write into the mp3 file
  	 * If no valid result found:
  	  * Check for details in file name
  	  * Create the id3 tag struct from file name
  	  * Write into the mp3 file
  	* If valid:
  	 * continue

 Files:
 acoustic-id.go

 id3.go

 id3-utils.go

 mp3-file-ops.go

 main.go
 */
