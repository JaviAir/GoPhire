# GoPhire
Google-Cloud or Fire-Base storage api wrapper.

Simplified methods to help when working with Google Storage Buckets.

Currently, only Storage works. Additional modules will be added in the future and are a work in progress


Wtih this library you can
  * DownloadFile()
  * UploadFile()
  * GetFileAttributes()
  * Init() Initialize
  * SayHi()

Methods DownloadFile(), UploadFile(), and GetFileAttributes() can take in optional paramters.
These optional parameters can be:
  * CustomBucket         string
	* ConnTimeout          int
	* Permission           fs.FileMode
	* ShowStatusPercentage bool

# Download File
 Arg1: string of destination path/name to save file as.
 Arg2: string of path to source file on the storage bucket.
 arg3: Optional parameters. (Optional, can be excluded completely.)
 
 # Upload File
 Arg1: string of destination path on the bucket to save file as.
 Arg2: string of path to local source file.
 arg3: Optional parameters. (Optional, can be excluded completely.)
