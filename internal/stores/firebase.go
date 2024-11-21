/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/nova
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package stores

/*****************************************************************************************************************/

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"firebase.google.com/go/v4/storage"
	"github.com/google/uuid"
)

/*****************************************************************************************************************/

type FirebaseStore struct {
	Store
	Client *storage.Client
}

/*****************************************************************************************************************/

func NewFirebaseStorageClient(client *storage.Client) *FirebaseStore {
	return &FirebaseStore{
		Store:  Store{},
		Client: client,
	}
}

/*****************************************************************************************************************/

func (f *FirebaseStore) RetriveBuffer(
	ctx context.Context,
	bucketName string,
	location string,
) (*bytes.Buffer, error) {
	// Create a new buffer to store the object data:
	buff := new(bytes.Buffer)

	bucket, err := f.Client.Bucket(bucketName)

	// If there is an error setting up the Firebase Storage bucket, return an error 500 response:
	if err != nil {
		return nil, err
	}

	// Get the object from the bucket:
	obj := bucket.Object(location)

	rc, err := obj.NewReader(ctx)

	if err != nil {
		return nil, fmt.Errorf("object(%v).NewReader: %w", obj, err)
	}

	defer rc.Close()

	_, err = io.Copy(buff, rc)

	if err != nil {
		return nil, fmt.Errorf("failed to read object data: %w", err)
	}

	return buff, nil
}

/*****************************************************************************************************************/

type StoreBufferParams struct {
	ContentType string
	Owner       string
}

/*****************************************************************************************************************/

func (f *FirebaseStore) StoreBuffer(
	ctx context.Context,
	buff *bytes.Buffer,
	bucketName string,
	location string,
	params StoreBufferParams,
) error {
	id := uuid.New()

	bucket, err := f.Client.Bucket(bucketName)

	// If there is an error setting up the Firebase Storage bucket, return an error 500 response:
	if err != nil {
		return err
	}

	// Create a new object in the bucket:
	obj := bucket.Object(location)

	// Create a new writer for the object:
	w := obj.NewWriter(ctx)

	// Set the ContentType of the object:
	w.ContentType = params.ContentType
	// Set the Cache-Control header for the object:
	w.ObjectAttrs.CacheControl = "public, max-age=31536000"
	// Set the ContentType metadata for the object:
	w.ObjectAttrs.ContentType = params.ContentType
	// Set the Owner metadata for the object:
	w.ObjectAttrs.Owner = fmt.Sprintf("user-%s", params.Owner)
	// Set the firebaseStorageDownloadTokens Metadata of the object:
	w.ObjectAttrs.Metadata = map[string]string{"firebaseStorageDownloadTokens": id.String()}

	// Write the buffer to the object:
	_, err = w.Write(buff.Bytes())

	if err != nil {
		return err
	}

	// Close the writer:
	err = w.Close()

	if err != nil {
		return err
	}

	// Reset the buffer:
	buff.Reset()

	// By default we return a nil error:
	return nil
}

/*****************************************************************************************************************/
