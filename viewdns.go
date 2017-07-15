/*
 * Copyleft 2017, Simone Margaritelli <evilsocket at protonmail dot com>
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *   * Redistributions of source code must retain the above copyright notice,
 *     this list of conditions and the following disclaimer.
 *   * Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *   * Neither the name of ARM Inject nor the names of its contributors may be used
 *     to endorse or promote products derived from this software without
 *     specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */
package xray

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ViewDNS struct {
	apikey string
}

type Query struct {
	tool   string
	domain string
}

type Response struct {
	Records []HistoryEntry `json:"records"`
}

type Result struct {
	Query    Query    `json:"query"`
	Response Response `json:"response"`
}

func NewViewDNS(apikey string) *ViewDNS {
	return &ViewDNS{apikey: apikey}
}

func (d *ViewDNS) GetHistory(domain string) []HistoryEntry {
	url := fmt.Sprintf("https://api.viewdns.info/iphistory/?domain=%s&apikey=%s&output=json", domain, d.apikey)
	history := make([]HistoryEntry, 0)

	if d.apikey != "" {
		if res, err := http.Get(url); err == nil {
			/*empijei: error not handled,
			suggestion to either handle it or throw it away properly, i.e.:
			defer func(){
				_ = res.Body.Close()
			}
			*/
			defer res.Body.Close()

			decoder := json.NewDecoder(res.Body)
			r := Result{}

			if err = decoder.Decode(&r); err == nil {
				history = r.Response.Records
			}
		}
	}

	return history
}