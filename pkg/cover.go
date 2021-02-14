// Developed by Kaiser925 on 2021/2/14.
// Lasted modified 2021/2/14.
// Copyright (c) 2021.  All rights reserved
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkg

import (
	"errors"
	"fmt"
	"github.com/Kaiser925/requests4go"
	"github.com/antchfx/htmlquery"
	"strings"
)

var headers = requests4go.Headers(map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Host":            "www.bilibili.com",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15",
	"Accept-Language": "zh-cn",
	"Accept-Encoding": "gzip, deflate, br",
})

func GetCover(bvNum string, filename string) error {
	if !isValidName(bvNum) {
		return errors.New("not valid bv number: " + bvNum)
	}

	if len(filename) == 0 {
		filename = fmt.Sprintf("%s.jpg", bvNum)
	}

	return getCover(bvNum, filename)
}

func getCover(bvNum string, filename string) error {
	image, err := getCoverSrc(bvNum)
	if err != nil {
		return err
	}

	resp, err := requests4go.Get(image, headers)
	if err != nil {
		return err
	}

	if err := resp.SaveContent(filename); err != nil {
		return err
	}

	return nil
}

func getCoverSrc(bvNum string) (string, error) {
	url := fmt.Sprintf("https://www.bilibili.com/video/%s/", bvNum)
	resp, err := requests4go.Get(url, headers)
	if err != nil {
		return "", err
	}

	if !resp.Ok() {
		return "", errors.New("get video page failed: " + resp.Status)
	}
	txt, err := resp.Text()
	if err != nil {
		return "", err
	}
	doc, err := htmlquery.Parse(strings.NewReader(txt))
	if err != nil {
		return "", err
	}
	meta := htmlquery.FindOne(doc, "//meta[@itemprop='image']")
	return htmlquery.SelectAttr(meta, "content"), nil
}

func isValidName(bvNum string) bool {
	return strings.HasPrefix(bvNum, "BV") || strings.HasPrefix(bvNum, "bv")
}
