package supastorage

import (
	"bytes"
	"io"

	storage_go "github.com/supabase-community/storage-go"
)

/*
supaClient struct to define an object that implements io.ReadCloser and io.WriterCloser interfaces
for writing and uploading to the supabase storage.
*/
type supaClient struct {
	bucket      string
	storagePath string
	objectData  *bytes.Buffer
	*storage_go.Client
}

/*
NewSupaClient to create an instance of supaClient
*/
func NewSupaClient(api_url, api_key, bucket string) *supaClient {

	client := storage_go.NewClient(api_url, api_key, nil)

	return &supaClient{
		bucket,
		"",
		new(bytes.Buffer),
		client,
	}
}

/*
supaClient.Write, will write to supabase storage
*/
func (supa *supaClient) Write(data []byte) (n int, err error) {
	supa.objectData.Write(data)
	return
}

/*
supaClient.close, will close and start uploading to supabase
*/
func (supa *supaClient) Close() (err error) {
	if supa.objectData.Len() > 0 {
		_, err = supa.UploadFile(supa.bucket, supa.storagePath, supa.objectData)
		if err != nil {
			return err
		}
		supa.objectData.Reset()
	}

	return nil
}

/*
supaClient.Read, will read data from supabase
*/
func (supa *supaClient) Read(p []byte) (n int, err error) {
	data, err := supa.DownloadFile(supa.bucket, supa.storagePath)
	if err != nil {
		return 0, err
	}

	n = copy(p, data)

	return n, nil
}

/*
supaClient.Reader will implement io.ReadCloser
*/
func (supa *supaClient) Reader(storagePath string) (resp io.ReadCloser) {
	supa.storagePath = storagePath
	return supa
}

/*
supaClient.Writer will implement io.WriteCloser
*/
func (supa *supaClient) Writer(storagePath string) (resp io.WriteCloser) {
	supa.storagePath = storagePath
	return supa
}
