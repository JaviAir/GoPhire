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

# How To Use
* go get cloud.google.com/go/storage
* go get google.golang.org/api/option
* go get github.com/JaviAir/GoPhire

- Create a Firebase or Google Cloud project.
- Create a Service Account Key for that project.

* Open a Go project file. 
 > Ex. main.go
* import the GoPhire/Storage module
 > Ex.  import (
 > "github.com/JaviAir/GoPhire/Storage"
 > )
* Initialize a variable with your Service Account Key json data pasted as a string.
> Ex. serviceKey := `{ServiceAccountKey Json Data pasted from ServiceAccountKey.json}`
* Initialize the Storage module by calling the Init() method.
> * Pass in the serviceKey variable, and your default storage bucket name as a string WITHOUT the gs:// prefix.
> * Ex. Storage.Init(serviceKey, "bucketname")

# Initialize Storage
* This library will not work without running the Init() method first.
The Init() method takes in 2 string args.
* arg1: the service key json data
* arg2: the bucket name without the "gs://" prefix

> These are used to initialize default values required by all the other methods in the storage library.
> * You can change the bucket name later after you have chosen the default bucket.
> * Note: all buckets must be accessible under the same service key. 
> * * If trying to access a file from a bucket that needs another service key. 
> * * You will have to re-initialize the library. 
> * * Meaning get the new service key data and run the Storage.Init() method again with the new data.

* Once the Storage module has been initialized you can use the other methods.

# Download File
> * Ex. Storage.DownloadFile("../mypenguins.jpg", "myStorageDir/Penguins.jpg")
> * or
> * Ex. opt := Storage.OptionalParameters{ShowStatusPercentage: true}
> * Storage.DownloadFile("../mypenguins.jpg", "myStorageDir/Penguins.jpg", opt)
 * Arg1: string of destination path/name to save file as.
 * Arg2: string of path to source file on the storage bucket.
 * arg3: Optional parameters. (Optional, can be excluded completely.)
 * When referencing a storage path there must be no leading forward slash. So the root path would be "".
 
 # Upload File
> * Ex. Storage.UploadFile("myStorageDir/", "./mypenguins.jpg")
> * or
> * Ex. opt := Storage.OptionalParameters{ShowStatusPercentage: true}
> * Storage.UploadFile("", "./mypenguins.jpg", opt)
 - Arg1: string of destination path on the bucket to save file as.
 - Arg2: string of path to local source file.
 - arg3: Optional parameters. (Optional, can be excluded completely.)
 - When referencing a storage path there must be no leading forward slash. So the root path would be "".

# Optional Parameters
Methods DownloadFile(), UploadFile(), and GetFileAttributes() can take in optional paramters.
These optional parameters can be:
* CustomBucket         string       // default is bucket passed from Init() method
* ConnTimeout          int          // default is 50 seconds
* Permission           fs.FileMode  // default is 0644 or 0777, depending on next param
* ShowStatusPercentage bool         // default is false

- ShowStatusPercentage 
 - if true this prints, (fmt), the progress of the download or upload. In ONE line. (Multiple shown for example purposes.)
 > * Ex. Downloading... 2 % complete. (777835 bytes)
 > * Ex. Downloading... 69 % complete. (777835 bytes)
 > * Ex. Downloading... 100 % complete. (777835 bytes)
 > 

# Get File Attributes

>	* returns // File name is: LatestStorageDir/Penguins.jpg , File Size is 777835
>	

	attrs, err := Storage.GetFileAttributes("LatestStorageDir/Penguins.jpg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("File name is: %s , File Size is %d \n", attrs.Name, attrs.Size)

* This method returns an attributes object that contains all the values a storage bucket creates when it recieves the file.

// ObjectAttrs represents the metadata for a Google Cloud Storage (GCS) object.

	// Bucket is the name of the bucket containing this GCS object.
	// This field is read-only.
	Bucket string

	// Name is the name of the object within the bucket.
	// This field is read-only.
	Name string

	// ContentType is the MIME type of the object's content.
	ContentType string

	// ContentLanguage is the content language of the object's content.
	ContentLanguage string

	// CacheControl is the Cache-Control header to be sent in the response
	// headers when serving the object data.
	CacheControl string

	// EventBasedHold specifies whether an object is under event-based hold. New
	// objects created in a bucket whose DefaultEventBasedHold is set will
	// default to that value.
	EventBasedHold bool

	// TemporaryHold specifies whether an object is under temporary hold. While
	// this flag is set to true, the object is protected against deletion and
	// overwrites.
	TemporaryHold bool

	// RetentionExpirationTime is a server-determined value that specifies the
	// earliest time that the object's retention period expires.
	// This is a read-only field.
	RetentionExpirationTime time.Time

	// ACL is the list of access control rules for the object.
	ACL []ACLRule

	// If not empty, applies a predefined set of access controls. It should be set
	// only when writing, copying or composing an object. When copying or composing,
	// it acts as the destinationPredefinedAcl parameter.
	// PredefinedACL is always empty for ObjectAttrs returned from the service.
	// See https://cloud.google.com/storage/docs/json_api/v1/objects/insert
	// for valid values.
	PredefinedACL string

	// Owner is the owner of the object. This field is read-only.
	//
	// If non-zero, it is in the form of "user-<userId>".
	Owner string

	// Size is the length of the object's content. This field is read-only.
	Size int64

	// ContentEncoding is the encoding of the object's content.
	ContentEncoding string

	// ContentDisposition is the optional Content-Disposition header of the object
	// sent in the response headers.
	ContentDisposition string

	// MD5 is the MD5 hash of the object's content. This field is read-only,
	// except when used from a Writer. If set on a Writer, the uploaded
	// data is rejected if its MD5 hash does not match this field.
	MD5 []byte

	// CRC32C is the CRC32 checksum of the object's content using the Castagnoli93
	// polynomial. This field is read-only, except when used from a Writer or
	// Composer. In those cases, if the SendCRC32C field in the Writer or Composer
	// is set to is true, the uploaded data is rejected if its CRC32C hash does
	// not match this field.
	CRC32C uint32

	// MediaLink is an URL to the object's content. This field is read-only.
	MediaLink string

	// Metadata represents user-provided metadata, in key/value pairs.
	// It can be nil if no metadata is provided.
	Metadata map[string]string

	// Generation is the generation number of the object's content.
	// This field is read-only.
	Generation int64

	// Metageneration is the version of the metadata for this
	// object at this generation. This field is used for preconditions
	// and for detecting changes in metadata. A metageneration number
	// is only meaningful in the context of a particular generation
	// of a particular object. This field is read-only.
	Metageneration int64

	// StorageClass is the storage class of the object. This defines
	// how objects are stored and determines the SLA and the cost of storage.
	// Typical values are "STANDARD", "NEARLINE", "COLDLINE" and "ARCHIVE".
	// Defaults to "STANDARD".
	// See https://cloud.google.com/storage/docs/storage-classes for all
	// valid values.
	StorageClass string

	// Created is the time the object was created. This field is read-only.
	Created time.Time

	// Deleted is the time the object was deleted.
	// If not deleted, it is the zero value. This field is read-only.
	Deleted time.Time

	// Updated is the creation or modification time of the object.
	// For buckets with versioning enabled, changing an object's
	// metadata does not change this property. This field is read-only.
	Updated time.Time

	// CustomerKeySHA256 is the base64-encoded SHA-256 hash of the
	// customer-supplied encryption key for the object. It is empty if there is
	// no customer-supplied encryption key.
	// See // https://cloud.google.com/storage/docs/encryption for more about
	// encryption in Google Cloud Storage.
	CustomerKeySHA256 string

	// Cloud KMS key name, in the form
	// projects/P/locations/L/keyRings/R/cryptoKeys/K, used to encrypt this object,
	// if the object is encrypted by such a key.
	//
	// Providing both a KMSKeyName and a customer-supplied encryption key (via
	// ObjectHandle.Key) will result in an error when writing an object.
	KMSKeyName string

	// Prefix is set only for ObjectAttrs which represent synthetic "directory
	// entries" when iterating over buckets using Query.Delimiter. See
	// ObjectIterator.Next. When set, no other fields in ObjectAttrs will be
	// populated.
	Prefix string

	// Etag is the HTTP/1.1 Entity tag for the object.
	// This field is read-only.
	Etag string

	// A user-specified timestamp which can be applied to an object. This is
	// typically set in order to use the CustomTimeBefore and DaysSinceCustomTime
	// LifecycleConditions to manage object lifecycles.
	//
	// CustomTime cannot be removed once set on an object. It can be updated to a
	// later value but not to an earlier one. For more information see
	// https://cloud.google.com/storage/docs/metadata#custom-time .
	CustomTime time.Time
