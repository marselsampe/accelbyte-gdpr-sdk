/*
 * Copyright (c) 2023 AccelByte Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 *
 */

package service

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func CreateTempFile(fileName string, bytes []byte) (*os.File, error) {
	logrus.Debugf("[CreateTempFile] Creating file [%s]...", fileName)
	tempFile, err := os.CreateTemp("", fileName)
	if err != nil {
		return nil, err
	}

	_, err = tempFile.Write(bytes)
	return tempFile, err
}

func UploadFile(ctx context.Context, uploadURL, filePath string) error {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer file.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, uploadURL, file)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/zip")

	fi, err := file.Stat()
	if err != nil {
		return err
	}
	req.ContentLength = fi.Size()

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		respBody, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			return errRead
		}
		defer resp.Body.Close()

		errMsg := fmt.Sprintf("response code: %v, response body: %v", resp.Status, string(respBody))
		logrus.Errorf("Fail upload file: %s", errMsg)
		return errors.New("response code: " + resp.Status)
	}

	return nil
}
