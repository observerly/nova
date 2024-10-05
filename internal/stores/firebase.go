/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/birpc
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
