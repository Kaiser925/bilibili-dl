// Developed by Kaiser925 on 2023/5/24.
// Lasted modified 2023/5/24.
// Copyright (c) 2023.  All rights reserved
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package fetch

import (
	"fmt"
	"github.com/Kaiser925/requests4go"
	"github.com/antchfx/htmlquery"
	"log"
	"net/url"
	"path"
	"strings"
)

var headers = requests4go.Headers(map[string]string{
	"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
	"Host":            "www.bilibili.com",
	"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15",
	"Accept-Language": "zh-cn",
	"Accept-Encoding": "gzip, deflate, br",
})

// FetchCover fetches the cover of given video.
func FetchCover(u string, filename string) {
	URL, err := url.Parse(u)
	if err != nil {
		log.Fatalf("not valid url: %v\n", err)
	}
	vid := path.Base(URL.Path)
	if len(filename) == 0 {
		filename = fmt.Sprintf("%s.jpg", vid)
	}

	log.Println("get image src of", u)
	resp, err := requests4go.Get(u, headers)
	if err != nil {
		log.Fatalf("coudn't send request to %s: %v", u, err)
	}
	if !resp.Ok() {
		log.Fatalf("coudn't load response: got status %s\n", resp.Status)
	}
	txt, err := resp.Text()
	if err != nil {
		log.Fatalf("coudn't load response content: %v\n", err)
	}
	doc, err := htmlquery.Parse(strings.NewReader(txt))
	if err != nil {
		log.Fatalf("coudn't parse response html: %v\n", err)
	}
	meta := htmlquery.FindOne(doc, "//meta[@itemprop='image']")
	imgSrc := "https:" + htmlquery.SelectAttr(meta, "content")

	imgSrc, _, _ = strings.Cut(imgSrc, "@")
	log.Println("downloading image from", imgSrc)

	resp, err = requests4go.Get(imgSrc, headers)
	if err != nil {
		log.Fatalf("coudn't send request to %s: %v\n", u, err)
	}

	if err := resp.SaveContent(filename); err != nil {
		log.Fatalf("coudn't save image file: %v\n", err)
	}

	log.Printf("cover of %s saved\n", u)
}
