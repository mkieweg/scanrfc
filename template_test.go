package scanrfc

import (
	"reflect"
	"testing"
)

func TestJsonDate_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		jd      *JsonDate
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.jd.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("JsonDate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJsonDate_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		j       JsonDate
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonDate.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonDate.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateEntry(t *testing.T) {
	type args struct {
		n int
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *RfcEntry
		wantErr bool
	}{
		{
			name: "rfc8000",
			args: args{
				n: 8000,
				b: []byte(`
{
	"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
	"rev": "11",
	"pages": 17,
	"time": "2016-11-30 15:53:01",
	"group": {
		"name": "Network File System Version 4",
		"type": "WG",
		"acronym": "nfsv4"
	},
	"expires": "2017-03-05 11:39:38",
	"title": "Requirements for NFSv4 Multi-Domain Namespace Deployment",
	"abstract": "This document presents requirements for the deployment of the NFSv4 protocols for the construction of an NFSv4 file namespace in environments with multiple NFSv4 Domains.  To participate in an NFSv4 multi-domain file namespace, the server must offer a multi-domain-capable file system and support RPCSEC_GSS for user authentication.  In most instances, the server must also support identity-mapping services.",
	"aliases": [
		"draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rfc8000"
	],
	"state": "RFC",
	"intended_std_level": "Proposed Standard",
	"std_level": "Proposed Standard",
	"authors": [
		{
		"name": "Andy Adamson",
		"email": "andros@netapp.com",
		"affiliation": "NetApp"
		},
		{
		"name": "Nicol\u00e1s Williams",
		"email": "nico@cryptonector.com",
		"affiliation": "Cryptonector"
		}
	],
	"shepherd": "Spencer Shepler <spencer.shepler@gmail.com>",
	"ad": "Spencer Dawkins <spencerdawkins.ietf@gmail.com>",
	"rev_history": [
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "00",
		"published": "2014-10-25T08:33:57",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/00/",
		"pages": 13
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "01",
		"published": "2015-01-23T09:12:05",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/01/",
		"pages": 13
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "02",
		"published": "2015-07-22T05:34:03",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/02/",
		"pages": 13
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "03",
		"published": "2015-08-07T10:51:24",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/03/",
		"pages": 14
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "04",
		"published": "2015-08-14T09:25:19",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/04/",
		"pages": 15
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "05",
		"published": "2015-08-21T07:03:24",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/05/",
		"pages": 15
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "06",
		"published": "2015-10-01T10:38:45",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/06/",
		"pages": 15
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "07",
		"published": "2016-05-06T07:24:06",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/07/",
		"pages": 15
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "08",
		"published": "2016-06-22T12:23:02",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/08/",
		"pages": 16
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "09",
		"published": "2016-06-29T09:48:31",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/09/",
		"pages": 16
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "10",
		"published": "2016-08-29T16:03:51",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/10/",
		"pages": 16
		},
		{
		"name": "draft-ietf-nfsv4-multi-domain-fs-reqs",
		"rev": "11",
		"published": "2016-09-01T11:39:38",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/11/",
		"pages": 16
		},
		{
		"name": "rfc8000",
		"rev": "rfc8000",
		"published": "2016-11-30T15:53:01",
		"url": "/doc/draft-ietf-nfsv4-multi-domain-fs-reqs/"
		}
	],
	"iesg_state": "RFC Published",
	"rfceditor_state": null,
	"iana_review_state": "Version Changed - Review Needed",
	"iana_action_state": "No IANA Actions",
	"consensus": true,
	"stream": "IETF"
	}

			`),
			},
			want: &RfcEntry{
				Number: 8000,
				Authors: []Author{
					{
						Name:        "Andy Adamson",
						Email:       "andros@netapp.com",
						Affiliation: "NetApp",
					},
					{
						Name:        "Nicol\u00e1s Williams",
						Email:       "nico@cryptonector.com",
						Affiliation: "Cryptonector",
					},
				},
				Title:    "Requirements for NFSv4 Multi-Domain Namespace Deployment",
				Pages:    17,
				Year:     2016,
				Month:    "November",
				Abstract: "This document presents requirements for the deployment of the NFSv4 protocols for the construction of an NFSv4 file namespace in environments with multiple NFSv4 Domains.  To participate in an NFSv4 multi-domain file namespace, the server must offer a multi-domain-capable file system and support RPCSEC_GSS for user authentication.  In most instances, the server must also support identity-mapping services.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateEntry(tt.args.n, tt.args.b)
			got.Date = JsonDate{}
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}
