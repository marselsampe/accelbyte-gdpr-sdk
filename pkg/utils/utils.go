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
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func IsEmptyJson(data []byte) bool {
	str := string(data)
	if str == "" || str == "[]" || str == "{}" {
		return true
	}
	return false
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
