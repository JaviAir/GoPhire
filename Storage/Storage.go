package Storage

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var serviceAccountKey string
var defaultBucket string

func Init(userServiceAccountKey string, userDefaultBucket string) {
	serviceAccountKey = userServiceAccountKey
	defaultBucket = userDefaultBucket
}

func SayHi() {
	fmt.Println("Hi there!")
}

type OptionalParameters struct {
	CustomBucket         string
	ConnTimeout          int
	Permission           fs.FileMode
	ShowStatusPercentage bool
}

type percentWriter struct {
	Current int
}

var totalBytes int
var loadStatus string = "Loading"

func (pw *percentWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Current += n

	percentage := (float64(pw.Current) / float64(totalBytes)) * 100

	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\r%s... %s %% complete. (%d bytes)", loadStatus, strconv.Itoa(int(percentage)), totalBytes)
	if int(percentage) == 100 {
		fmt.Println("\n ")
	}
	return n, nil
}

func createClient() (*storage.Client, error) { // internal, do not edit unless you know what you're doing
	ctx := context.Background()
	opt := option.WithCredentialsJSON([]byte(serviceAccountKey))
	// opt := option.WithCredentialsFile(serviceAccountKey)
	newClient, err := storage.NewClient(ctx, opt)

	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	// defer newClient.Close()

	return newClient, nil
}

func DownloadFile(dstFileName string, srcFileName string, opts ...OptionalParameters) ([]int, error) {

	ctx := context.Background()

	// default option param values
	conectionTimeout := 50
	bucket := defaultBucket
	perm := fs.FileMode(0777)
	showPercentage := false

	if len(opts) > 0 {
		if opts[0].ConnTimeout != 0 {
			conectionTimeout = opts[0].ConnTimeout
		}
		if opts[0].CustomBucket != "" {
			bucket = opts[0].CustomBucket
		}
		if opts[0].Permission != fs.FileMode(0000) {
			perm = opts[0].Permission
		}
		if opts[0].ShowStatusPercentage {
			showPercentage = true
		}
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(conectionTimeout))
	defer cancel()

	newClient, err := createClient()
	defer newClient.Close()

	rc, err := newClient.Bucket(bucket).Object(srcFileName).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("srcFileName(%q).NewReader: %v", srcFileName, err)
	}
	defer rc.Close()

	perc := &percentWriter{Current: 0}
	if showPercentage {
		// Create the file
		out, err := os.Create(dstFileName)
		if err != nil {
			return nil, err
		}
		defer out.Close()

		//set total file size & load status: upload/download
		totalBytes = int(rc.Attrs.Size)
		loadStatus = "Downloading"

		if _, err = io.Copy(out, io.TeeReader(rc, perc)); err != nil {
			out.Close()
			return nil, fmt.Errorf("io.Copy: %v", err)
		}
	} else {
		data, err := io.ReadAll(rc)
		if err != nil {
			return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
		}
		// fmt.Println(len(data))

		err = os.WriteFile(dstFileName, data, perm)
		if err != nil {
			return nil, fmt.Errorf("os.WriteFile error: %v", err)
		}
	}

	return nil, nil
}

func GetFileAttributes(srcFileName string, opts ...OptionalParameters) (*storage.ObjectAttrs, error) {
	ctx := context.Background()
	bucket := defaultBucket

	if len(opts) > 0 {
		if opts[0].CustomBucket != "" {
			bucket = opts[0].CustomBucket
		}
	}

	newClient, err := createClient()
	defer newClient.Close()

	attrs, err := newClient.Bucket(bucket).Object(srcFileName).Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("attrs error: %v", err)
	}
	return attrs, nil
}

func UploadFile(dstPath string, srcFilePath string, opts ...OptionalParameters) error {

	ctx := context.Background()

	newClient, err := createClient()
	defer newClient.Close()

	// Open local file.
	f, err := os.Open(srcFilePath)
	if err != nil {
		return fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	// default option param values
	conectionTimeout := 50
	bucket := defaultBucket
	showPercentage := false

	if len(opts) > 0 {
		if opts[0].ConnTimeout != 0 {
			conectionTimeout = opts[0].ConnTimeout
		}
		if opts[0].CustomBucket != "" {
			bucket = opts[0].CustomBucket
		}
		if opts[0].ShowStatusPercentage {
			showPercentage = true
		}
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(conectionTimeout))
	defer cancel()

	srcFileName := filepath.Base(srcFilePath)
	fullPath := dstPath + srcFileName

	if showPercentage {

		// Upload an object with storage.Writer.
		wc := newClient.Bucket(bucket).Object(fullPath).NewWriter(ctx)

		perc := &percentWriter{Current: 0}

		//get total file size
		fi, err := f.Stat()
		if err != nil {
			return fmt.Errorf("f.stat() error: %v", err)
		}
		totalBytes = int(fi.Size())
		loadStatus = "Uploading"

		if _, err = io.Copy(wc, io.TeeReader(f, perc)); err != nil {
			return fmt.Errorf("io.Copy: %v", err)
		}

		if err := wc.Close(); err != nil {
			return fmt.Errorf("Writer.Close: %v", err)
		}
	} else {

		// Upload an object with storage.Writer.
		wc := newClient.Bucket(bucket).Object(fullPath).NewWriter(ctx)
		if _, err = io.Copy(wc, f); err != nil {
			return fmt.Errorf("io.Copy: %v", err)
		}
		if err := wc.Close(); err != nil {
			return fmt.Errorf("Writer.Close: %v", err)
		}
	}
	// fmt.Fprintf(w, "Blob %v uploaded.\n", object)
	return nil
}
