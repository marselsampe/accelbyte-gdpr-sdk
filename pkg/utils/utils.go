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

package utils

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func CreateZipFile(namespace, userID string, data map[string][]byte) ([]byte, error) {
	zipFile, err := os.Create(fmt.Sprintf("%s-%s.zip", namespace, userID))
	if err != nil {
		return nil, err
	}
	defer func() {
		zipFile.Close()
		_ = os.Remove(zipFile.Name())
	}()
	zipWriter := zip.NewWriter(zipFile)

	var includedModules []string
	for moduleID, moduleBytes := range data {
		if isEmptyJson(moduleBytes) {
			continue
		}
		writer, errCreate := zipWriter.Create(moduleID)
		if errCreate != nil {
			return nil, errCreate
		}
		reader := bytes.NewReader(moduleBytes)
		if _, errCopy := io.Copy(writer, reader); errCopy != nil {
			return nil, errCopy
		}
		includedModules = append(includedModules, moduleID)
	}

	if len(includedModules) == 0 {
		return nil, nil
	}

	// create summary file
	summaryBytes, errMarshal := json.Marshal(map[string]interface{}{
		"namespace": namespace,
		"userId":    userID,
		"modules":   includedModules,
	})
	if errMarshal != nil {
		return nil, err
	}
	writer, errCreate := zipWriter.Create("summary.json")
	if errCreate != nil {
		return nil, err
	}
	reader := bytes.NewReader(summaryBytes)
	if _, errCopy := io.Copy(writer, reader); errCopy != nil {
		return nil, errCopy
	}
	zipWriter.Close()

	// return zip file bytes
	var zipBytes []byte
	_, err = zipFile.Read(zipBytes)
	if err != nil {
		return nil, err
	}
	return zipBytes, nil
}

func UploadFile(ctx context.Context, uploadURL string, data []byte) error {
	reader := bytes.NewReader(data)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, uploadURL, reader)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/zip")
	req.ContentLength = reader.Size()

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		respBody, errRead := io.ReadAll(resp.Body)
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

func isEmptyJson(bytes []byte) bool {
	if bytes == nil {
		return true
	}
	str := string(bytes)
	if str == "" || str == "[]" || str == "{}" {
		return true
	}
	return false
}
